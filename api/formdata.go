package api

import (
	"encoding/json"
	"goflow/dao"
	"goflow/model"
	"goflow/util"
	"time"
)

func SaveFormData(user *model.User, formid, procinstid, taskid string, formData map[string]any) error {
	//create a new process instance
	now := time.Now()
	formdata := &model.Formdata{}
	formdata.Id = util.NextId()
	formdata.Form_id = formid
	formdata.Procinst_id = procinstid
	formdata.Task_id = taskid
	v, err := json.Marshal(formData)
	if err != nil {
		return err
	}
	formdata.Form_data = string(v)
	formdata.Create_id = user.Id
	formdata.Update_id = user.Id
	formdata.Create_time = &now
	formdata.Update_time = &now
	formdata.Status = "0"
	return dao.CreateFormData(formdata)
}

func FindFormData(formid string,procinstid string) (formdata string, err error) {
	return dao.FindFormData(formid,procinstid)
}
