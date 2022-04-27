package api

import (
	"goflow/dao"
	"goflow/model"
)

func GetNodeById(process *model.Process, id string) (node any, nodeType string) {
	if process.Start.Id == id {
		return process.Start, "start"
	}
	
	for _, end := range process.End {
		if end.Id == id {
			return end, "end"
		}
	}
	
	for _, fork := range process.Forks {
		if fork.Id == id {
			return fork, "fork"
		}
	}
	for _, task := range process.Tasks {
		if task.Id == id {
			return task, "task"
		}
	}
	return nil, ""

}

func FindMyStartList(uid string) (procinst []*model.Procinst, err error) {
	return dao.FindMyList(uid)
}

func FindApproveList(uid string) (list []*model.ApproveList, err error) {
	return dao.FindApproveList(uid)
}