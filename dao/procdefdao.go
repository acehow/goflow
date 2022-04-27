package dao

import (
	"goflow/model"
	"goflow/util"
)

func FindProcdefByID(id string) (procdef *model.Procdef, err error) {
	procdef = &model.Procdef{}
	err = util.Db.First(procdef, "id = ?", id).Error
	return
}

func FindAllProcdef() (procdefs []*model.Procdef, err error) {
	procdefs = []*model.Procdef{}
	err = util.Db.Find(&procdefs, "status='0'").Error
	return
}
