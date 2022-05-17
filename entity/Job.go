package entity

type Job struct {
	Model
	AccountId          uint64 `gorm:"type:bigint; not null; default:0"`
	UserNotificationId uint64 `gorm:"type:bigint; not null; default:0"`
	UserId             uint64 `gorm:"type:bigint; not null; default:0"`
	TimeWindow         string `gorm:"type:int not null; default:0"`
	Status             int8   `gorm:"type:tinyint; not null; default:0"`
}
