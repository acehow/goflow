package api

import (
	"encoding/json"
	"goflow/dao"
	"goflow/model"
	"goflow/util"

	"log"
	"strings"
)

func LoadAllProcess() error {
	procdefs, err := dao.FindAllProcdef()
	if err != nil {
		return err
	}
	for _, procdef := range procdefs {
		if procdef.Resource == "" {
			continue
		}
		process := &model.Process{}
		decoder1 := json.NewDecoder(strings.NewReader(procdef.Resource))
		decoder1.Decode(process)
		process.Name = procdef.Name
		util.PutProcess(procdef.Id, process)
	}

	return nil
}

func FindProcessById(procDefId string) (*model.Process, error) {
	process:=util.GetProcess(procDefId)
	if process!=nil {
		return process,nil
	}
	procdef, err := dao.FindProcdefByID(procDefId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if procdef == nil || procdef.Resource == "" {
		return nil, nil
	}

	process = &model.Process{}
	decoder := json.NewDecoder(strings.NewReader(procdef.Resource))
	decoder.Decode(process)
	process.Name = procdef.Name
	util.PutProcess(procdef.Id, process)
	return process, nil
}
