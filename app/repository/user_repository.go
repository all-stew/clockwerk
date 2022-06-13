package repository

import (
	"clockwerk/app/models"
	"clockwerk/pkg/dbutils"
	"context"
)

type SysUserRepository interface {
	dbutils.Tx
	// Create 创建user
	Create(ctx context.Context, username string, nickname string, phone string, email string, gender models.SYS_USER_GENDER) (models.SysUser, error)
	// FindByUsernameAndPassword 根据username和明文密码查询SysUser
	FindByUsernameAndPassword(ctx context.Context, username string, password string) (models.SysUser, error)
	// FindById 根据id查询用户
	FindById(ctx context.Context, id uint64) (models.SysUser, error)
}
