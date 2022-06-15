package store

import (
	"clockwerk/app/models"
	. "clockwerk/app/repository"
	"clockwerk/pkg/dbutils"
	"context"
)

type RoleStore struct {
	// 数据库连接
	*dbutils.Connection
}

func NewRoleStore(db *dbutils.Connection) *RoleStore {
	return &RoleStore{
		Connection: db,
	}
}

var _ RoleRepository = (*RoleStore)(nil)

func (rs RoleStore) FindRoleByUserId(ctx context.Context, userId uint64) ([]models.Role, error) {
	//TODO implement me
	var roles []models.Role
	rs.GetConnection(ctx).Table("sys_role r").Select("*").Joins("left join `sys_user_role_relation` ur on ur.role_id = r.id").Where("ur.user_id = ?", userId).Scan(&roles)

	return roles, nil
}
