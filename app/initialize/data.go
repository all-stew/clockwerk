package initialize

import (
	"clockwerk/app/global"
	"clockwerk/app/models"
	"clockwerk/pkg/model"
	"clockwerk/pkg/permission/casbin"
	"clockwerk/pkg/utils"
	"encoding/json"
	"errors"
	"fmt"
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

func Table() {
	err := global.DB.AutoMigrate(
		new(models.User),                // 用户 数据表
		new(models.UserRoleRelation),    // 用户-角色 数据表
		new(models.Role),                // 角色 数据表
		new(models.SysRoleMenuRelation), // 角色-菜单 数据表
		new(models.SysMenu),             // 菜单 数据表
		new(models.Preset),              // 策略集 数据表
		new(models.RolePreset),          // 角色-策略集 数据表
	)
	if err != nil {
		panic(fmt.Sprintf("数据库同步异常：%s", err.Error()))
	}
}

func Data() {
	UserInitialize()
	RoleInitialize()
	UserRoleRelationInitialize()
	PresetInitialize()
	RolePresetInitialize()
}

// UserInitialize 用户初始化数据
func UserInitialize() {
	var users = []models.User{
		{
			BaseModel: model.BaseModel{
				Id:        1,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Username: "admin",
			Nickname: "admin",
			Password: utils.CryptoPassword("admin123"),
		},
		{
			BaseModel: model.BaseModel{
				Id:        2,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			Username: "test",
			Nickname: "test",
			Password: utils.CryptoPassword("test123"),
		},
	}

	for _, user := range users {
		u := models.User{}
		err := global.DB.Where("id = ?", user.Id).First(&u).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.DB.Create(&user)
		}
	}
}

// RoleInitialize 角色初始化数据
func RoleInitialize() {
	var roles = []models.Role{
		{
			BaseModel: model.BaseModel{
				Id:        1,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			RoleName: "admin",
			RoleKey:  "admin",
			RoleSort: 1,
			Status:   0,
		},
	}

	for _, role := range roles {
		u := models.Role{}
		err := global.DB.Where("id = ?", role.Id).First(&u).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.DB.Create(&role)
		}
	}

}

func UserRoleRelationInitialize() {
	var userRoleRelations = []models.UserRoleRelation{
		{
			BaseModel: model.BaseModel{
				Id:        1,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			UserId: 1,
			RoleId: 1,
		},
	}
	for _, userRoleRelation := range userRoleRelations {
		r := models.UserRoleRelation{}
		err := global.DB.Where("id = ?", userRoleRelation.Id).First(&r).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.DB.Create(&userRoleRelation)
		}
	}
}

func PresetInitialize() {
	marshal, err := json.Marshal(&[]casbin.PermissionPreset{
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"GET"},
			// * 客体
			// action作用的对象
			Resources: []string{"/system/user/list"},
			// Allow/Deny
			Effect: "allow",
		},
	})
	if err != nil {
		return
	}
	var presets = []models.Preset{
		{
			BaseModel: model.BaseModel{
				Id:        1,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			PresetName: "admin",
			Policies:   string(marshal),
		},
	}
	for _, preset := range presets {
		r := models.Preset{}
		err := global.DB.Where("id = ?", preset.Id).First(&r).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.DB.Create(&preset)
		}
	}
}

func RolePresetInitialize() {
	var rolePresets = []models.RolePreset{
		{
			BaseModel: model.BaseModel{
				Id:        1,
				CreatedBy: 1,
				UpdatedBy: 1,
			},
			RoleId:   "admin",
			PresetId: "1",
		},
	}
	for _, rolePreset := range rolePresets {
		r := models.RolePreset{}
		err := global.DB.Where("id = ?", rolePreset.Id).First(&r).Error
		// 如果记录不存在则添加数据
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.DB.Create(&rolePreset)
		}
	}
}
