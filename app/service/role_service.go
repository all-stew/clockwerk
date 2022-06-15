package service

import (
	"clockwerk/app/models"
	"context"
)

type RoleService interface {
	FindRoleByUserId(ctx context.Context, userId uint64) ([]models.Role, error)
}
