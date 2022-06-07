package views

import "clockwerk/src/models"

type UserCreateRequestView struct {
	Username string                 `json:"username" form:"username"`
	Phone    string                 `json:"phone" form:"phone"`
	Email    string                 `json:"email" form:"email"`
	Gender   models.SYS_USER_GENDER `json:"gender" form:"gender"`
}

type LoginView struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
