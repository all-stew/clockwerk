package models

import "clockwerk/pkg/model"

type SysRoleMenuRelation struct {
	model.BaseModel
	RoleId uint64 `json:"role_id" gorm:"not null"`
	MenuId uint64 `json:"menu_id" gorm:"not null"`
}
