package model

import "time"

type Procdef struct {
	Id           string `gorm:"primary_key"`
	Company_id   string
	Company_name string
	Name         string
	Ver          string
	Resource     string
	Create_id    string
	Update_id    string
	Create_time  *time.Time
	Update_time  *time.Time
	Status       string
}

type Procinst struct {
	Id               string `gorm:"primary_key"`
	Name             string
	Procdef_id       string
	Create_name      string
	Create_dept_id   string
	Create_dept_name string
	Create_email     string
	Create_phone     string
	Current_flow     string
	Current_name     string
	Start_id         string
	Create_id        string
	Version          int
	Update_id        string
	Create_time      *time.Time
	Update_time      *time.Time
	Status           string
}

type Procflow struct {
	Id           string `gorm:"primary_key"`
	Procdef_id   string
	Procinst_id  string
	Prev_flow_id string
	Node_id      string
	Node_name    string
	Node_type    string
	Prev_id      string
	Update_id    string
	Create_id    string
	Create_time  *time.Time
	Update_time  *time.Time
	Status       string
}

type Approve struct {
	Id             string `gorm:"primary_key"`
	Procinst_id    string
	Procflow_id    string
	Form_id        string
	User_id        string
	User_name      string
	User_dept_id   string
	User_dept_name string
	User_email     string
	User_phone     string
	Node_id        string
	Node_type      string
	Data           string
	Deal_time      *time.Time
	Deal_result    string
	Update_id      string
	Create_id      string
	Create_time    *time.Time
	Update_time    *time.Time
	Status         string
}

type Formdata struct {
	Id          string `gorm:"primary_key"`
	Form_id     string
	Procinst_id string
	Task_id     string
	Form_name   string
	Form_data   string
	Create_id   string
	Update_id   string
	Create_time *time.Time
	Update_time *time.Time
	Status      string
}

type Form struct {
	Id          string `gorm:"primary_key"`
	Form_name   string
	Form        string
	Form_desc   string
	Create_id   string
	Update_id   string
	Create_time *time.Time
	Update_time *time.Time
	Status      string
}

type ApproveList struct {
	Id          string `gorm:"primary_key"`
	Procinst_id   string
	Procflow_id   string     
	Form_id        string
	Name   string
	NodeId string
	Procdef_id string
}
