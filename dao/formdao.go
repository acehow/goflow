package dao

import (
	"goflow/model"
	"goflow/util"
)

func FindFormByID(id string) (form *model.Form, err error) {
	form = &model.Form{}
	err = util.Db.First(form, "id = ?", id).Error
	return
}

