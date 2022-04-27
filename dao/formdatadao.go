package dao

import (
	"goflow/model"
	"goflow/util"
)

func CreateFormData(formdata *model.Formdata) error {
	result := util.Db.Create(formdata)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindFormData(formid string,procinstid string) (formdata string, err error) {
	form := &model.Formdata{}
	err = util.Db.First(form, "form_id = ? and procinst_id=?", formid,procinstid).Error
	
	if err != nil {
		return "", err
	}
	return form.Form_data, nil
}
