package dao

import (
	"goflow/model"
	"goflow/util"
)

func FindMyList(uid string) (procinst []*model.Procinst, err error) {
	procinst = []*model.Procinst{}
	err = util.Db.Find(&procinst, "create_id = ?", uid).Order("create_time desc").Error
	return
}

func FindApproveList(uid string) (list []*model.ApproveList, err error) {
//	SELECT a.procinst_id , a.procflow_id , a.id ,p.name 
//FROM goflow.approve a join goflow.procinst p on a.procinst_id =p.id and a.procflow_id =p.current_flow where a.deal_result ="" and a.user_id ="2"
	err = util.Db.Table("approve").Select("approve.id ,approve.procinst_id , approve.procflow_id,approve.form_id, procinst.name,approve.node_id,procinst.procdef_id").Joins("join procinst on approve.procinst_id =procinst.id and approve.procflow_id =procinst.current_flow").Where("approve.deal_result ='' and approve.user_id= ?", uid).Order("approve.create_time desc").Scan(&list).Error
	return
}
