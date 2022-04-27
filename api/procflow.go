package api

import (
	"encoding/json"
	"fmt"
	"goflow/dao"
	"goflow/model"
	"goflow/util"
	"log"
	"time"

	"github.com/robertkrimen/otto"
)

func MoveToNextNode(user *model.User, ok, procDefId, procInstId, procFlowId , nodeId,approveId string, approver []*model.User, datamap map[string]any) (err error) {
	// get process
	process, err := FindProcessById(procDefId)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if process == nil {
		err = fmt.Errorf("no process found. procDefId= " + procDefId)
		return
	}
	inst,err:=dao.FindProcinstByID(procInstId)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if inst == nil {
		err = fmt.Errorf("no process instance found. procInstId= " + procInstId)
		return
	}
	
	procflows := make([]model.Procflow, 0)
	addApprove := make([]model.Approve, 0)
	updateApprove := make([]model.Approve, 0)

	// deal flow
	err = dealNextFlow(user, &procflows, &addApprove, &updateApprove, process, ok, procDefId, procInstId, procFlowId, procFlowId, nodeId, nodeId,approveId, approver, datamap,false)
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = dao.DealProcflow(procFlowId, inst,&procflows, &addApprove, &updateApprove)

	return
}

func dealNextFlow(user *model.User, flows *[]model.Procflow, addApprove *[]model.Approve, updateApprove *[]model.Approve, process *model.Process, ok, procDefId, procInstId, prevFlowId, procFlowId, prevNodeId, nodeId , approveId string, approver []*model.User, datamap map[string]any, isAuto bool) (err error) {
	//get next node by id
	defer func() {
		if e := recover(); e != nil {
			log.Println(err.Error())
		}
	}()
	
	//get node
	node, nodetype := GetNodeById(process, nodeId)
	if node == nil {
		return fmt.Errorf("no node found. nodeId= " + nodeId)
	}
	if nodetype == "end" {
		// process end
		endnode := node.(*model.End)
		flow := newFlow(procDefId, procInstId, prevFlowId, prevNodeId, nodeId, endnode.Name, "end", "1", datamap)
		*flows = append(*flows, *flow)
		return nil
	}
	if nodetype == "start" {
		// start node
		startnode := node.(*model.Start)
		return dealNextFlow(user, flows, addApprove, updateApprove, process, ok, procDefId, procInstId, prevFlowId, procFlowId, nodeId, startnode.NextId,approveId, approver, datamap,true)
	}

	if nodetype == "fork" {
		forknode := node.(*model.Fork)

		vm := otto.New()
		for _, p := range forknode.Params {
			vm.Set(p.Name, datamap[p.Name])
		}

		value, err := vm.Eval(forknode.Conds)
		vm = nil
		if err != nil {
			log.Println(err)
			return err
		}
		var forknext string
		if value.IsNumber() {
			ivalue, err := value.ToInteger()
			if err != nil {
				log.Println(err)
				return err
			}

			forknext = forknode.NextId[ivalue]
		} else {
			// default goto 1st node
			forknext = forknode.NextId[0]
		}
		flow := newFlow(procDefId, procInstId, prevFlowId, prevNodeId, nodeId,forknode.Name, "fork", "0", datamap)
		*flows = append(*flows, *flow)

		// set reslut to datamap
		for _, fresult := range forknode.Results {
			datamap[fresult.Name] = fresult.Value
		}

		prevFlowId = flow.Id
		err = dealNextFlow(user, flows, addApprove, updateApprove, process, ok, procDefId, procInstId, prevFlowId, procFlowId, nodeId, forknext,approveId, approver, datamap,true)
		if err != nil {
			return err
		}

		return nil
	}

	// if node is a task
	if nodetype == "task" {
		tasknode := node.(*model.Task)

		// need approve
		if ok == util.DEAL_RESULT_OK {
			// if ok means this task going approve
			if tasknode.ActType == util.ACT_TYPE_AND {
				// and sign
				oknum, err := dao.GetOkNum(procInstId, procFlowId)
				if err != nil {
					return err
				}

				if oknum+1 >= int64(tasknode.ApproveNum) {
					// all member approve, update approve status
					approve := setApprove(user.Id, user.Name, procInstId, procFlowId,tasknode.FormId, nodeId, ok,approveId, datamap, false)
					*updateApprove = append(*updateApprove, *approve)

					flow := newFlow(procDefId, procInstId, prevFlowId, prevNodeId, nodeId,tasknode.Name, "task", "0", datamap)
					*flows = append(*flows, *flow)
					prevFlowId = flow.Id
					return dealNextFlow(user, flows, addApprove, updateApprove, process, ok, procDefId, procInstId, prevFlowId, procFlowId, nodeId, tasknode.NextId,approveId, approver, datamap,true)
				} else {
					// not all member approve, return
					approve := setApprove(user.Id, user.Name, procInstId, procFlowId,tasknode.FormId, nodeId, ok,approveId, datamap, false)
					*updateApprove = append(*updateApprove, *approve)
					return nil
				}
			} else {
				// or sign
				approve := setApprove(user.Id, user.Name, procInstId, procFlowId,tasknode.FormId, nodeId, ok,approveId, datamap, false)
				*updateApprove = append(*updateApprove, *approve)

				flow := newFlow(procDefId, procInstId, prevFlowId, prevNodeId, nodeId,tasknode.Name, "task", "0", datamap)
				*flows = append(*flows, *flow)
				prevFlowId = flow.Id
				return dealNextFlow(user, flows, addApprove, updateApprove, process, ok, procDefId, procInstId, prevFlowId, procFlowId, nodeId, tasknode.NextId,approveId, approver, datamap,true)
			}
		} else if ok == util.DEAL_RESULT_NG {
			// if ok is ng, means this task going not approve
			if tasknode.ActType == util.ACT_TYPE_AND {
				// and sign
				ngnum, err := dao.GetNgNum(procInstId, procFlowId)
				if err != nil {
					return err
				}

				if ngnum+1 >= int64(tasknode.RejectNum) {
					// all member reject, update reject status
					approve := setApprove(user.Id, user.Name, procInstId, procFlowId,tasknode.FormId, nodeId, ok,approveId, datamap, false)
					*updateApprove = append(*updateApprove, *approve)

					flow := newFlow(procDefId, procInstId, prevFlowId, prevNodeId, nodeId,tasknode.Name, "task", "0", datamap)
					*flows = append(*flows, *flow)
					prevFlowId = flow.Id
					return dealNextFlow(user, flows, addApprove, updateApprove, process, ok, procDefId, procInstId, prevFlowId, procFlowId, nodeId, tasknode.NextId,approveId, approver, datamap,true)
				} else {
					// not all member reject, return
					approve := setApprove(user.Id, user.Name, procInstId, procFlowId,tasknode.FormId, nodeId, ok,approveId, datamap, false)
					*updateApprove = append(*updateApprove, *approve)
					return nil
				}
			} else {
				// or sign
				approve := setApprove(user.Id, user.Name, procInstId, procFlowId,tasknode.FormId, nodeId, ok,approveId, datamap, false)
				*updateApprove = append(*updateApprove, *approve)

				flow := newFlow(procDefId, procInstId, prevFlowId, prevNodeId, nodeId,tasknode.Name, "task", "0", datamap)
				*flows = append(*flows, *flow)
				prevFlowId = flow.Id
				return dealNextFlow(user, flows, addApprove, updateApprove, process, ok, procDefId, procInstId, prevFlowId, procFlowId, nodeId, tasknode.NextId,approveId, approver, datamap,true)
			}
		} else {
			// if ok is empty, means this task going set approve
			// APPROVE_TYPE_LIST = 1
			// APPROVE_TYPE_MANUAL = 2
			// APPROVE_TYPE_MANAGER = 4
			appr := util.GetIntValue(tasknode.ApproveType)
			var approvemap map[string]string = make(map[string]string)
			for _, apprint := range appr {
				if apprint == util.APPROVE_TYPE_MANAGER {
					// manager approve
					approvemap[user.ManagerId] = user.ManagerName
				} else if apprint == util.APPROVE_TYPE_MANUAL {
					for _, approv := range approver {
						approvemap[approv.Id] = approv.Name
					}
				} else if apprint == util.APPROVE_TYPE_LIST {
					for _, approv := range tasknode.Approvers {
						approvemap[approv.Id] = approv.Name
					}
				}
			}
			flow := newFlow(procDefId, procInstId, prevFlowId, prevNodeId, nodeId,tasknode.Name, "task", "1", datamap)
			*flows = append(*flows, *flow)
			for key, name := range approvemap {
				approve := setApprove(key, name, procInstId, flow.Id,tasknode.FormId, nodeId, ok,approveId, datamap, true)
				*addApprove = append(*addApprove, *approve)
			}
		}
	}

	return nil
}

func newFlow(procDefId, procInstId, prevFlowId, prevNodeId, nodeId, nodeName,nodeType, status string, data map[string]any) *model.Procflow {
	flow := &model.Procflow{}
	flow.Id = util.NextId()
	flow.Procdef_id = procDefId
	flow.Procinst_id = procInstId
	flow.Prev_flow_id = prevFlowId
	flow.Node_id = nodeId
	flow.Node_name = nodeName
	flow.Node_type = nodeType
	flow.Prev_id = prevNodeId
	flow.Status = status

	return flow
}

func setApprove(userId, userName, procInstId, procFlowId,formId, nodeId, dealResult,approveId string, data map[string]any, isNew bool) *model.Approve {
	now := time.Now()
	approve := &model.Approve{}
	approve.Procinst_id = procInstId
	approve.Procflow_id = procFlowId
	approve.Form_id = formId
	approve.Node_id = nodeId
	approve.User_id = userId
	approve.User_name = userName
	
	if isNew {
		approve.Id = util.NextId()
		approve.Create_time = &now
	} else {
		approve.Deal_time = &now
		approve.Deal_result = dealResult
		approve.Id = approveId
		if data != nil {
			v, err := json.Marshal(data)
			if err != nil {
				log.Println(err.Error())
			}
			approve.Data = string(v)
		}
	}

	return approve
}
