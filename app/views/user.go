package views

import (
	"clockwerk/app/models"
	. "clockwerk/pkg/validator"
)

type UserCreateRequestView struct {
	Username string                 `json:"username" form:"username"`
	Phone    string                 `json:"phone" form:"phone"`
	Email    string                 `json:"email" form:"email"`
	Gender   models.SYS_USER_GENDER `json:"gender" form:"gender"`
}

type LoginView struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (LoginView LoginView) GetMessages() Messages {
	return Messages{
		"Username.required": "用户名不能为空",
		"Password.required": "密码不能为空",
	}
}
