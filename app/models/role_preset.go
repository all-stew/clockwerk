package models

import (
	"clockwerk/pkg/model"
)

type RolePreset struct {
	model.BaseModel
	RoleId   string `json:"role_id" gorm:"comment:角色Id;column:role_id"`
	PresetId string `json:"preset_id" gorm:"comment:策略Id;column:preset_id"`
}

func (RolePreset) TableName() string {
	return "sys_role_preset"
}

func (e *RolePreset) Generate() model.Builder {
	o := *e
	return &o
}

func (e *RolePreset) GetId() interface{} {
	return e.Id
}
