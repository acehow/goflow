package api

import (
	"errors"
	"goflow/dao"
	"goflow/model"

	"log"
)

func FindFormById(formId string) (*model.Form, error) {
	form, err := dao.FindFormByID(formId)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if form == nil || form.Form == "" {
		return nil, errors.New("form is not found")
	}

	return form, nil
}
