package dao

import (
	"goflow/model"
	"goflow/util"
)

func GetOkNum(procInstId, procFlowId string) (oknum int64, err error) {
	var approve []model.Approve
	result := util.Db.Where("procinst_id = ? and procflow_id=? and deal_result='ok'", procInstId, procFlowId).Find(&approve).Count(&oknum)
	if result.Error != nil {
		return 0, result.Error
	}

	return
}

func GetNgNum(procInstId, procFlowId string) (ngnum int64, err error) {
	var approve []model.Approve
	result := util.Db.Where("procinst_id = ? and procflow_id=? and deal_result='ng'", procInstId, procFlowId).Find(&approve).Count(&ngnum)
	if result.Error != nil {
		return 0, result.Error
	}

	return
}
