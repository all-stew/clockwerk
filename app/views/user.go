package views

import (
	"clockwerk/app/models"
	"clockwerk/pkg/validator"
)

type UserCreateRequestView struct {
	Nickname string                 `json:"nickname" form:"nickname" binding:"required"`
	Phone    string                 `json:"phone" form:"phone" binding:"required"`
	Email    string                 `json:"email" form:"email" binding:"required"`
	Gender   models.SYS_USER_GENDER `json:"gender" form:"gender" binding:"required"`
}

func (request UserCreateRequestView) GetMessages() validator.Messages {
	return validator.Messages{
		"NickName.required": "昵称不能为空",
		"Phone.required":    "手机号不能为空",
		"Email.required":    "邮箱不能为空",
		"Gender.required":   "性别不能为空",
	}
}

type LoginView struct {
	Username  string `json:"username" form:"username" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
	Captcha   string `json:"captcha" form:"captcha" binding:"required"`
	CaptchaId string `json:"captcha_id" form:"captcha_id" binding:"required"`
}

func (LoginView LoginView) GetMessages() validator.Messages {
	return validator.Messages{
		"Username.required":  "用户名不能为空",
		"Password.required":  "密码不能为空",
		"Captcha.required":   "验证码不能为空",
		"CaptchaId.required": "登陆异常，请重新登陆",
	}
}
