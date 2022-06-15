package models

import (
	"clockwerk/pkg/model"
)

type SysRolePreset struct {
	model.BaseModel
	RoleId   uint64 `json:"role_id" gorm:"comment:角色Id;column:role_id"`
	PresetId uint64 `json:"preset_id" gorm:"comment:策略Id;column:preset_id"`
}

func (SysRolePreset) TableName() string {
	return "sys_role_preset"
}

func (e *SysRolePreset) Generate() model.Builder {
	o := *e
	return &o
}

func (e *SysRolePreset) GetId() interface{} {
	return e.Id
}
