package dao

import (
	"errors"
	"goflow/model"
	"goflow/util"
)

func CreateProcflow(procflow *[]model.Procflow) error {
	result := util.Db.Create(procflow)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DealProcflow(procFlowId string, procinst *model.Procinst, procflow *[]model.Procflow, addApprove *[]model.Approve, updateApprove *[]model.Approve) error {
	tx := util.Db.Begin()
	if procflow != nil && len(*procflow) > 0 {
		result := tx.Create(procflow)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	if addApprove != nil && len(*addApprove) > 0 {
		result := tx.Create(addApprove)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}
	if updateApprove != nil && len(*updateApprove) > 0 {
		for _, approve := range *updateApprove {
			result := tx.Model(&model.Approve{}).Where("id=?", approve.Id).Updates(model.Approve{Deal_result: approve.Deal_result, Deal_time: approve.Deal_time, Data: approve.Data})
			if result.Error != nil {
				tx.Rollback()
				return result.Error
			}
		}
	}

	replace := false
	current := ""
	currentName := ""
	for _, flow := range *procflow {
		if flow.Node_type == "end" {
			result := tx.Model(procinst).Where("version = ?", procinst.Version).Updates(map[string]interface{}{"current_flow": "end", "current_name": flow.Node_name})
			if result.Error != nil {
				tx.Rollback()
				return result.Error
			}
			if result.RowsAffected == 0 {
				tx.Rollback()
				return errors.New("procinst update overtime")
			}
			tx.Commit()
			return nil
		}
		if flow.Status == "1" {
			replace = true
			current = flow.Id
			currentName = flow.Node_name
		}
	}
	if replace {
		err := tx.Model(procinst).Updates(map[string]interface{}{"current_flow": current, "current_name": currentName}).Error
		if err != nil {
			tx.Rollback()
			return err
		}

	}
	// update proc inst version
	result := tx.Model(procinst).Where("version = ?", procinst.Version).Update("version", procinst.Version+1)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("procinst update overtime")
	}
	tx.Commit()
	return nil
}
