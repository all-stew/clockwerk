package impl

import (
	"clockwerk/app/global"
	"clockwerk/app/models"
	. "clockwerk/app/service"
	"clockwerk/pkg/uuid"
	"context"
)

type UserServiceImpl struct {
}

func GetUserServiceImpl() *UserServiceImpl {
	return &UserServiceImpl{}
}

var _ UserService = (*UserServiceImpl)(nil)

func (service *UserServiceImpl) Login(username string, password string) (models.User, error) {
	ctx := global.UserStore.BeginTx(context.Background())

	user, err := global.UserStore.FindByUsernameAndPassword(ctx, username, password)
	if err != nil {
		return models.User{}, err
	}
	// 组装用户 role
	user.Roles, err = RoleServiceImpl{}.FindRoleByUserId(ctx, user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *UserServiceImpl) Create(nickname string, phone string, email string, gender models.SYS_USER_GENDER) (models.User, error) {

	ctx := global.UserStore.BeginTx(context.Background())
	username := uuid.GenerateUUIDv4WithNoMinus()
	// todo check

	createUser, err := global.UserStore.Create(ctx, username, nickname, phone, email, gender)
	if err != nil {
		return models.User{}, err
	}
	return createUser, nil
}
