package user

import "time"

type User struct {
	ID        uint64 `gorm:"primary_key;autoIncrement"`
	Username  string `gorm:"column:username"`
	Nickname  string `gorm:"column:nickname"`
	Email     string `gorm:"column:email"`
	Phone     string `gorm:"column:phone"`
	Password  string `gorm:"column:password_util"`
	Salt      string `gorm:"column:salt"`
	Status    Status `gorm:"column:status"`
	CreatedAt *time.Time
	CreatedBy uint64 `gorm:"type:bigint"`
	UpdatedAt *time.Time
	UpdatedBy uint64 `gorm:"type:bigint"`
	DeletedAt *time.Time
}

type Status int8

const (
	USER_STATUS_DISABLE Status = 0
	USER_STATUS_ENABLE  Status = 1
)
