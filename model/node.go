package model

// node type: start, end, task, fork, join
type Process struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Start *Start  `json:"start"`
	End   []*End  `json:"end"`
	Tasks []*Task `json:"tasks"`
	Forks []*Fork `json:"forks"`
}

type Node struct {
	Id   string
	Type string
}

type End struct {
	Name   string `json:"name,omitempty"`
	Id     string `json:"id"`
	PrevId string `json:"prevId,omitempty"`
}

type Start struct {
	Name   string `json:"name,omitempty"`
	Id     string `json:"id"`
	FormId string `json:"formId,omitempty"`
	NextId string `json:"nextId,omitempty"`
}

type Fork struct {
	Name   string   `json:"name,omitempty"`
	Id     string   `json:"id"`
	NextId []string `json:"nextId,omitempty"`
	Conds  string   `json:"conds,omitempty"`
	Params []*Param `json:"params,omitempty"`

	//result can set nodeid as name, stop as value to stop process
	Results []*Result `json:"results,omitempty"`
}

type Param struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type Result struct {
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type Task struct {
	Name   string `json:"name,omitempty"`
	Id     string `json:"id"`
	PrevId string `json:"prevId,omitempty"`
	NextId string `json:"nextId,omitempty"`
	// 0 noneed: noneed approve, direct to next node
	// 1 list: select approvers from approver list
	// 2 manual: manual select approvers
	// 4 manager: select user direct manager as approver
	// Support combine with add: i.e. 5=4+1=manager+list
	ApproveType int     `json:"approveType,omitempty"`
	Approvers   []*User `json:"approvers,omitempty"`

	// and sign need number, defaule need all sign
	ApproveNum int `json:"approveNum,omitempty"`
	RejectNum  int `json:"rejectNum,omitempty"`
	// and sign , or sign : default or sign,
	ActType  string         `json:"actType,omitempty"`
	Url      string         `json:"url,omitempty"`
	FormId   string         `json:"formId,omitempty"`
	ValueMap map[string]any `json:"valueMap,omitempty"`
}

type User struct {
	Name        string `json:"name,omitempty"`
	Id          string `json:"id,omitempty"`
	DeptId      string `json:"deptId,omitempty"`
	DeptName    string `json:"deptName,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Email       string `json:"email,omitempty"`
	ManagerId   string `json:"managerId,omitempty"`
	ManagerName string `json:"managerName,omitempty"`
}
