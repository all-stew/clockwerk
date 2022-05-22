package job

import "time"

type Job struct {
	ID                 uint64 `gorm:"primary_key;autoIncrement"`
	AccountId          uint64 `gorm:"type:bigint; not null; default:0"`
	UserNotificationId uint64 `gorm:"type:bigint; not null; default:0"`
	UserId             uint64 `gorm:"type:bigint; not null; default:0"`
	TimeWindow         int    `gorm:"type:int not null; default:0"`
	Status             Status `gorm:"type:tinyint; not null; default:0"`
	CreatedAt          *time.Time
	CreatedBy          uint64 `gorm:"type:bigint"`
	UpdatedAt          *time.Time
	UpdatedBy          uint64 `gorm:"type:bigint"`
	DeletedAt          *time.Time
}

type Status int8

const (
	JOB_DISABLE Status = 0
	JOB_ENABLE  Status = 1
	JOB_BAN     Status = 2
)
