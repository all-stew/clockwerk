package repository

import (
	"clockwerk/app/models"
	"clockwerk/pkg/dbutils"
	"context"
)

type RoleRepository interface {
	dbutils.Tx
	// FindRoleByUserId 根据id查询用户
	FindRoleByUserId(ctx context.Context, userId uint64) ([]models.Role, error)
}
