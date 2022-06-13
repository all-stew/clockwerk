package models

import "clockwerk/pkg/model"

type SysUserRoleRelation struct {
	model.BaseModel
	UserId uint64 `json:"user_id" gorm:"not null"`
	RoleId uint64 `json:"role_id" gorm:"not null"`
}
