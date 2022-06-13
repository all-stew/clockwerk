package store

import (
	"clockwerk/app/models"
	. "clockwerk/app/repository"
	"clockwerk/pkg/dbutils"
	"clockwerk/pkg/utils"
	"context"
	"errors"
)

type UserStore struct {
	// 数据库连接
	*dbutils.Connection
}

func NewUserStore(db *dbutils.Connection) *UserStore {
	return &UserStore{
		Connection: db,
	}
}

var _ SysUserRepository = (*UserStore)(nil)

func (us *UserStore) FindById(ctx context.Context, id uint64) (models.SysUser, error) {
	var user models.SysUser
	user.Id = id

	// 查询数据库第一条id为 {id}的user数据
	result := us.Connection.GetConnection(ctx).Model(user).First(&user)

	if result.Error != nil {
		return models.SysUser{}, result.Error
	}

	return user, nil
}

func (us *UserStore) Create(ctx context.Context, username string, nickname string, phone string, email string, gender models.SYS_USER_GENDER) (models.SysUser, error) {
	user := models.SysUser{
		Username: username,
		Password: "",
		Nickname: nickname,
		Phone:    phone,
		Avatar:   "",
		Gender:   gender,
		Email:    email,
		Remark:   "",
		Status:   models.SYS_USER_STATUS_ENABLE,
	}
	// 默认密码
	result := us.Connection.GetConnection(ctx).Create(&user)

	if result.Error != nil {
		return models.SysUser{}, result.Error
	}

	return user, nil
}

func (us *UserStore) FindByUsernameAndPassword(ctx context.Context, username string, password string) (models.SysUser, error) {
	var user models.SysUser
	result := us.Connection.GetConnection(ctx).Where("username = ?", username).First(&user)

	if result.Error != nil {
		return models.SysUser{}, result.Error
	}

	if !utils.ComparePassword(user.Password, password) {
		return models.SysUser{}, errors.New("wrong password")
	}

	return user, nil
}
