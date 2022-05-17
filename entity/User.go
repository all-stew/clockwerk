package entity

type User struct {
	Model
	Username string `gorm:"type:varchar(64); not null;default:''"`
	Nickname string `gorm:"type:varchar(32); not null "`
	Email    string `gorm:"type:varchar(128); default:'' "`
	Phone    string `gorm:"type:varchar(64); default:''"`
	Status   int8   `gorm:"type:tinyint; not null default:0 "`
}
