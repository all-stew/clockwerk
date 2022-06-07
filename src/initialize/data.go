package initialize

import (
	"clockwerk/pkg/model"
	"clockwerk/pkg/utils"
	"clockwerk/src/global"
	"clockwerk/src/models"
	"errors"
	"gorm.io/gorm"
)

/*
   说明：初始化基础数据
*/

// 基础数据
var (
	// 密码：123456
	password = "$2a$10$CrHfdigkqfsh94kDgHc7SOKUa3dLM6VHvHDuq0EAH8rq0ybjMbpbK"
)

// 创建用户
func UserInitialize() {
	var users = []models.SysUser{
		{
			BaseModel: model.BaseModel{
				Id: 1,
			},
			Username: "test",
			Password: utils.CryptoPassword("123456"),
		},
	}

	for _, user := range users {
		u := models.SysUser{}
		err := global.DB.Where("id = ?", user.Id).First(&u).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.DB.Create(&user)
		}
	}
}

// 初始化数据
func InitData() {
	UserInitialize() // 用户初始化
}
