package api

import (
	"fmt"
	"goflow/dao"
	"goflow/model"
	"goflow/util"
	"log"
	"time"
)

func StartProcessInstance(procDefId string, user *model.User) (instId, flowId, taskId, formId string, err error) {
	process, err := FindProcessById(procDefId)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if process == nil {
		err = fmt.Errorf("no process found. procDefId=" + procDefId)
		return
	}

	//get start node
	taskId = process.Start.Id
	formId = process.Start.FormId
	now := time.Now()
	flowId = util.NextId()
	instId = util.NextId()
	//create a new process instance
	procinst := &model.Procinst{}
	procinst.Id = instId
	procinst.Name = user.Name + " " + process.Name
	procinst.Procdef_id = procDefId
	procinst.Create_name = user.Name
	procinst.Create_dept_id = user.DeptId
	procinst.Create_dept_name = user.DeptName
	procinst.Create_email = user.Email
	procinst.Create_phone = user.Phone
	procinst.Current_flow = flowId
	procinst.Start_id = flowId
	procinst.Create_id = user.Id
	procinst.Update_id = user.Id
	procinst.Create_time = &now
	procinst.Update_time = &now
	procinst.Status = "0"

	flow := &model.Procflow{}
	flow.Id = flowId
	flow.Procdef_id = procDefId
	flow.Procinst_id = instId
	flow.Node_id = taskId
	flow.Node_type = "start"
	flow.Prev_id = ""

	err = dao.CreateProcinst(procinst, flow)

	return
}
