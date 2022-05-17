package entity

type UserNotification struct {
	Model
	NotificationType  int8   `gorm:"type:tinyint; not null; default:0"`
	NotificationKey   string `gorm:"type:varchar(256); not null"`
	NotificationParam string `gorm:"type:text; not null; default:0"`
	UserId            uint64 `gorm:"type:bigint; not null; default:0"`
}
