package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type BaseModel struct {
	Id        uint64         `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	CreatedBy uint64         `json:"createBy" gorm:"index;comment:创建者"`
	UpdatedBy uint64         `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

// SetCreatedBy 设置创建人id
func (e *BaseModel) SetCreatedBy(createdBy uint64) {
	e.CreatedBy = createdBy
}

// SetUpdatedBy 设置修改人id
func (e *BaseModel) SetUpdatedBy(updatedBy uint64) {
	e.UpdatedBy = updatedBy
}

// Builder 模型builder
type Builder interface {
	schema.Tabler
	SetCreatedBy(createBy uint64)
	SetUpdatedBy(updateBy uint64)
	Generate() Builder
	GetId() interface{}
}
