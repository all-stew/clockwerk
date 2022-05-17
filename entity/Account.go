package entity

type Account struct {
	Model
	Account      string `gorm:"type:varchar(128);default:'';not null"`
	AccountType  uint8  `gorm:"type:tinyint;default:0;not null"`
	ParentId     uint64 `gorm:"type:bigint;default:0;not null"`
	AccountParam string `gorm:"type:text;default:''"`
	Status       uint8  `gorm:"type:tinyint;default:0;not null"`
}
