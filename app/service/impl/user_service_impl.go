package impl

import (
	"clockwerk/app/global"
	"clockwerk/app/models"
	. "clockwerk/app/service"
	"context"
)

type UserServiceImpl struct {
}

var _ UserService = (*UserServiceImpl)(nil)

func (service *UserServiceImpl) Login(context context.Context, username string, password string) (models.User, error) {
	user, err := global.UserStore.FindByUsernameAndPassword(context, username, password)
	if err != nil {
		return models.User{}, err
	}
	// 组装用户 role
	user.Roles, err = RoleServiceImpl{}.FindRoleByUserId(context, user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}
