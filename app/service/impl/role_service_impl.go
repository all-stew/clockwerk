package impl

import (
	"clockwerk/app/global"
	"clockwerk/app/models"
	. "clockwerk/app/service"
	"context"
)

type RoleServiceImpl struct {
}

var _ RoleService = (*RoleServiceImpl)(nil)

func (rs RoleServiceImpl) FindRoleByUserId(ctx context.Context, userId uint64) ([]models.Role, error) {
	return global.RoleStore.FindRoleByUserId(ctx, userId)
}
