package dao

import (
	"goflow/model"
	"goflow/util"
)

func CreateProcinst(procinst *model.Procinst,flow *model.Procflow) error {
	tx := util.Db.Begin()
	result := tx.Create(procinst)
	if result.Error != nil {
		tx.Rollback()
		return  result.Error
	}
	result = tx.Create(flow)
	if result.Error != nil {
		tx.Rollback()
		return  result.Error
	}
	tx.Commit()
	return nil
}

func FindProcinstByID(id string) (procinst *model.Procinst, err error) {
	procinst = &model.Procinst{}
	err = util.Db.First(procinst, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return
}