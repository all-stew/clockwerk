package initialize

import (
	"clockwerk/app/global"
	"clockwerk/pkg/permission/casbin"
	"fmt"
	"log"
)

/*
   说明：初始化 Casbin
*/
func Casbin() {
	// 初始化数据库适配器，并自定义表名，默认 casbin_rule
	// a, err := gormadapter.NewAdapterByDBUseTableName(global.DB, global.Conf.Mysql.TablePrefix, "casbin_rule")
	// if err != nil {
	// 	panic(fmt.Sprintf("Casbin数据库连接失败：%s", err.Error()))
	// }

	// 读取配置文件，需要避免多次读取
	// config, err := global.ConfBox.Find(global.Conf.Casbin.Model)
	// cabinModel := model.NewModel()

	// 从字符串中加载配置
	// err = cabinModel.LoadModelFromText(string(config))
	// if err != nil {
	// 	panic(fmt.Sprintf("Casbin配置读取失败：%s", err.Error()))
	// }

	// e, err := casbin.NewEnforcer(cabinModel, a)
	// if err != nil {
	// 	panic(fmt.Sprintf("Casbin配置加载失败：%s", err.Error()))
	// }

	// 加载策略
	// err = e.LoadPolicy()
	// if err != nil {
	// 	panic(fmt.Sprintf("Casbin策略加载失败：%s", err.Error()))
	// }

	// TODO: 从数据库读取角色和preset的对应关系和所有的preset
	// 具体格式参照clockwerk/pkg/permission/casbin/casbin_policy.go
	// global.DB.FindXXXXXXXX
	// Sample
	presets := []*casbin.PermissionPreset{
		// Policy为Resources和Actions的笛卡尔积
		// /system/user/list Post Allow
		// /system/user/list Get Allow
		// /system/user/list Put Allow
		// /system/user/search Post Allow
		// /system/user/search Get Allow
		// /system/user/search Put Allow
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"Post", "Get", "Put"},
			// 客体
			// action作用的对象
			Resources: []string{"/system/user/list", "/system/user/search"},
			// Allow/Deny
			Effect: "Allow",
		},
		// /system/user/get * Allow
		// 可以匹配所有/system/user/get的Rest请求
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"*"},
			// 客体
			// action作用的对象
			Resources: []string{"/system/user/get"},
			// Allow/Deny
			Effect: "Allow",
		},
		// /system/user/* * Allow
		// 可以匹配所有/system/user/*的Rest请求
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"*"},
			// 客体
			// action作用的对象
			Resources: []string{"/system/user/*"},
			// Allow/Deny
			Effect: "Allow",
		},
	}
	group := []*casbin.PolicyGroup{
		{
			RoleId:   "admin",
			PresetId: "1",
		},
		{
			RoleId:   "test",
			PresetId: "2",
		},
	}

	e, err := casbin.NewCasbinPermission(presets, group)
	if err != nil {
		panic(fmt.Sprintf("Casbin策略加载失败：%s", err.Error()))
	}

	global.CasbinEnforcer = e

	// 全局设置
	//global.CasbinEnforcer = e
	log.Println("Casbin初始化完成")
}
