package models

import "clockwerk/pkg/model"

type UserRoleRelation struct {
	model.BaseModel
	UserId uint64 `json:"user_id" gorm:"not null"`
	RoleId uint64 `json:"role_id" gorm:"not null"`
}

func (ur UserRoleRelation) TableName() string {
	return "sys_user_role_relation"
}

func (ur *UserRoleRelation) Generate() model.Builder {
	o := *ur
	return &o
}

func (ur *UserRoleRelation) GetId() interface{} {
	return ur.Id
}
