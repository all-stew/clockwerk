package service

import (
	"clockwerk/app/models"
	"context"
)

type UserService interface {
	Login(ctx context.Context, username string, password string) (models.User, error)
}
