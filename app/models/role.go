package models

import "clockwerk/pkg/model"

type Role struct {
	model.BaseModel
	RoleName string    `json:"roleName" gorm:"size:128;"` // 角色名称
	RoleKey  string    `json:"roleKey" gorm:"size:128;"`  //角色权限字符串
	RoleSort int       `json:"roleSort" gorm:""`          //显示顺序
	Status   uint      `json:"status" gorm:"size:4;"`     // 状态
	Remark   string    `json:"remark" gorm:"size:255;"`   //备注
	SysMenu  []SysMenu `json:"sysMenu" gorm:"-"`
}

func (Role) TableName() string {
	return "sys_role"
}

func (e *Role) Generate() model.Builder {
	o := *e
	return &o
}

func (e *Role) GetId() interface{} {
	return e.Id
}
