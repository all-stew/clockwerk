package models

import (
	"clockwerk/pkg/model"
)

type Preset struct {
	model.BaseModel
	PresetName string `json:"preset_name" gorm:"size:100;comment:策略集名称;column:preset_name"`
	Policies   string `json:"policies" gorm:"text;comment:策略;column:policies"`
	Remark     string `json:"remark" gorm:"size:100;comment:备注;column:remark"`
}

func (Preset) TableName() string {
	return "sys_preset"
}

func (e *Preset) Generate() model.Builder {
	o := *e
	return &o
}

func (e *Preset) GetId() interface{} {
	return e.Id
}
