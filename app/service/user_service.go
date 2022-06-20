package service

import (
	"clockwerk/app/models"
)

type UserService interface {
	Login(username string, password string) (models.User, error)
	Create(nickname string, phone string, email string, gender models.SYS_USER_GENDER) (models.User, error)
}
