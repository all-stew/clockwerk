package entity

import (
	"time"
)

type Model struct {
	ID        uint64 `gorm:"primary_key;autoIncrement"`
	CreatedAt *time.Time
	CreatedBy uint64 `gorm:"type:bigint"`
	UpdatedAt *time.Time
	UpdatedBy uint64 `gorm:"type:bigint"`
	DeletedAt *time.Time
}
