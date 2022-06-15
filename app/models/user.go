package models

import (
	"clockwerk/pkg/model"
	"time"
)

type User struct {
	model.BaseModel
	Username  string          `json:"username" gorm:"size:64;comment:用户名;column:username"`
	Nickname  string          `json:"nickname" gorm:"size:128;comment:昵称;column:nickname"`
	UserType  int             `json:"user_type" gorm:"comment:用户类型;column:user_type"`
	Email     string          `json:"email" gorm:"size:128;comment:邮箱;column:email"`
	Phone     string          `json:"phone" gorm:"size:32;comment:手机号;column:phone"`
	Gender    SYS_USER_GENDER `json:"gender" gorm:"comment:性别;column:gender"`
	Avatar    string          `json:"avatar" gorm:"size:255;comment:头像;column:avatar"`
	Password  string          `json:"-" gorm:"size:128;comment:密码;column:password"`
	Status    SYS_USER_STATUS `json:"status" gorm:"comment:状态;column:status"`
	LoginIp   string          `json:"login_ip" gorm:"size:128;comment:最后登陆ip"`
	LoginDate *time.Time      `json:"login_date" gorm:"comment:最后登陆ip"`
	Remark    string          `json:"remark" gorm:"size:255;comment:备注;column:remark"`
	Roles     []Role          `json:"roles" gorm:"-"`
}

type SYS_USER_GENDER int

const (
	SYS_USER_GENDER_MALE    SYS_USER_GENDER = 0
	SYS_USER_GENDER_FEMALE  SYS_USER_GENDER = 1
	SYS_USER_GENDER_UNKNOWN SYS_USER_GENDER = 2
)

type SYS_USER_STATUS int

const (
	SYS_USER_STATUS_ENABLE  SYS_USER_STATUS = 0
	SYS_USER_STATUS_DISABLE SYS_USER_STATUS = 1
)

func (User) TableName() string {
	return "sys_user"
}

func (e *User) Generate() model.Builder {
	o := *e
	return &o
}

func (e *User) GetId() interface{} {
	return e.Id
}
