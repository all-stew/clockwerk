package repository

import (
	. "clockwerk/config/mysql"
	. "clockwerk/entity"
	"log"
)

func Create(username string, nickname string, email string, phone string) bool {

	user := User{Username: username, Nickname: nickname, Phone: phone, Email: email}

	result := GetDb().Create(&user)

	if result.Error != nil {
		log.Printf("新增异常 %s\n", result.Error.Error())
		return false
	}
	return result.RowsAffected == 1
}

func Update(id uint64, username string, nickname string, email string, phone string) bool {
	user := User{}
	GetDb().Where("id = ?", id).Take(&user)
	user.Username = username
	user.Nickname = nickname
	user.Email = email
	user.Phone = phone

	result := GetDb().Save(&user)
	if result.Error != nil {
		log.Printf("更新异常 %s\n", result.Error.Error())
		return false
	}

	return result.RowsAffected == 1
}
