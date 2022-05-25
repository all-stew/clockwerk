package account

import "time"

type Account struct {
	ID           uint64      `gorm:"primary_key;autoIncrement"`
	Account      string      `gorm:"type:varchar(128);default:'';not null"`
	AccountType  AccountType `gorm:"type:tinyint;default:0;not null"`
	ParentId     uint64      `gorm:"type:bigint;default:0;not null"`
	AccountParam string      `gorm:"type:text;default:''"`
	UserId       uint64      `gorm:"column:user_id"`
	Status       Status      `gorm:"type:tinyint;default:0;not null"`
	CreatedAt    *time.Time
	CreatedBy    uint64 `gorm:"type:bigint"`
	UpdatedAt    *time.Time
	UpdatedBy    uint64 `gorm:"type:bigint"`
	DeletedAt    *time.Time
}

type Status int8

const (
	ACCOUNT_ENABLE  Status = 0
	ACCOUNT_DISABLE Status = 1
	ACCOUNT_VALID   Status = 2
	ACCOUNT_INVALID Status = 3
	ACCOUNT_RUNNING Status = 4
	// mihoyo
	ACCOUNT_MIHOYO_ERROR               Status = 5
	ACCOUNT_MIHOYO_DO_NOT_HAVE_GENSHIN Status = 6
	// genshin
	ACCOUNT_GENSHIN_VALID   = 7
	ACCOUNT_GENSHIN_INVALID = 8
)

type AccountType int8

const (
	ACCOUNT_TYPE_MIHOYO_BBS AccountType = 0
	ACCOUNT_TYPE_GENSHIN    AccountType = 1
)
