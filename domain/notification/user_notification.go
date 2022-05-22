package notification

import (
	"time"
)

type UserNotification struct {
	ID                uint64           `gorm:"primary_key;autoIncrement"`
	NotificationType  NotificationType `gorm:"type:tinyint; not null; default:0"`
	NotificationKey   string           `gorm:"type:varchar(256); not null"`
	NotificationParam string           `gorm:"type:text; not null; default:0"`
	UserId            uint64           `gorm:"type:bigint; not null; default:0"`
	CreatedAt         *time.Time
	CreatedBy         uint64 `gorm:"type:bigint"`
	UpdatedAt         *time.Time
	UpdatedBy         uint64 `gorm:"type:bigint"`
	DeletedAt         *time.Time
}

type NotificationType int8

const (
	NOTIFICATION_TYPE_NONE      NotificationType = 0
	NOTIFICATION_TYPE_PUSH_PLUS NotificationType = 1
)
