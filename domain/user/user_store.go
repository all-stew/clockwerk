package user

import (
	. "clockwerk/config/mysql"
	"clockwerk/pkg/logger"
	"clockwerk/pkg/util/password_util"
)

type Store struct {
}

var _ Repository = (*Store)(nil)

func (*Store) Create(username string, nickname string, password string, email string, phone string, createdBy uint64) bool {

	user := User{Username: username, Nickname: nickname, Phone: phone, Email: email, CreatedBy: createdBy}

	// 生成盐和密码
	salt := password_util.GetSalt()
	passwordHashed := password_util.GetPassword(password, salt)

	user.Password = passwordHashed
	user.Salt = salt

	result := GetDb().Create(&user)

	if result.Error != nil {
		logger.Logf("create error %s\n", result.Error.Error())
		return false
	}
	return result.RowsAffected == 1
}

func (*Store) Update(id uint64, nickname string, email string, phone string, updatedBy uint64) bool {
	user := User{}
	GetDb().Where("id = ?", id).Take(&user)
	user.Nickname = nickname
	user.Email = email
	user.Phone = phone
	user.UpdatedBy = updatedBy

	result := GetDb().Save(&user)
	if result.Error != nil {
		logger.Logf("update error %s\n", result.Error.Error())
		return false
	}

	return result.RowsAffected == 1
}

func (*Store) Login() {

}

func validPassword(password string, salt string, passwordInDb string) bool {
	return passwordInDb == password_util.GetPassword(password, salt)
}
