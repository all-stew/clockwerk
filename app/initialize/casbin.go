package initialize

import (
	"clockwerk/app/global"
	"clockwerk/app/models"
	"clockwerk/pkg/permission/casbin"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

/*
   说明：初始化 Casbin
*/
func Casbin() {
	// 从数据库读取数据
	// 加载preset策略集
	var presetsModels []models.SysPreset
	err := global.DB.Model(&models.SysPreset{}).Find(&presetsModels).Error
	if err != nil {
		panic(fmt.Sprintf("Casbin加载数据失败：%s", err.Error()))
	}

	// 加载角色与preset的关系
	var rolePresetsModels []models.SysRolePreset
	err = global.DB.Model(&models.SysRolePreset{}).Find(&rolePresetsModels).Error
	if err != nil {
		panic(fmt.Sprintf("Casbin加载数据失败：%s", err.Error()))
	}

	// 具体格式参照clockwerk/pkg/permission/casbin/casbin_policy.go
	// global.DB.FindXXXXXXXX
	// Sample
	// presets := []*casbin.PermissionPreset{
	// 	// Policy为Resources和Actions的笛卡尔积
	// 	// /system/user/list Post Allow
	// 	// /system/user/list Get Allow
	// 	// /system/user/list Put Allow
	// 	// /system/user/search Post Allow
	// 	// /system/user/search Get Allow
	// 	// /system/user/search Put Allow
	// 	{
	// 		// 策略集ID
	// 		PresetId: "1",
	// 		// 动作
	// 		Actions: []string{"Post", "Get", "Put"},
	// 		// 客体
	// 		// action作用的对象
	// 		Resources: []string{"/system/user/list", "/system/user/search"},
	// 		// Allow/Deny
	// 		Effect: "Allow",
	// 	},
	// 	// /system/user/get * Allow
	// 	// 可以匹配所有/system/user/get的Rest请求
	// 	{
	// 		// 策略集ID
	// 		PresetId: "1",
	// 		// 动作
	// 		Actions: []string{"*"},
	// 		// 客体
	// 		// action作用的对象
	// 		Resources: []string{"/system/user/get"},
	// 		// Allow/Deny
	// 		Effect: "Allow",
	// 	},
	// 	// /system/user/* * Allow
	// 	// 可以匹配所有/system/user/*的Rest请求
	// 	{
	// 		// 策略集ID
	// 		PresetId: "1",
	// 		// 动作
	// 		Actions: []string{"*"},
	// 		// 客体
	// 		// action作用的对象
	// 		Resources: []string{"/system/user/*"},
	// 		// Allow/Deny
	// 		Effect: "Allow",
	// 	},
	// }
	// group := []*casbin.PolicyGroup{
	// 	{
	// 		RoleId:   "admin",
	// 		PresetId: "1",
	// 	},
	// 	{
	// 		RoleId:   "test",
	// 		PresetId: "2",
	// 	},
	// }

	group, presets, err := convertPresetPolicies(rolePresetsModels, presetsModels)
	if err != nil {
		panic(fmt.Sprintf("Casbin数据转化失败：%s", err.Error()))
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

func convetIdToString(id uint64) string {
	return strconv.Itoa(int(id))
}

// convertPresetPolicies
func convertPresetPolicies(ap []models.SysRolePreset, ps []models.SysPreset) ([]*casbin.PolicyGroup, []*casbin.PermissionPreset, error) {
	gp := make([]*casbin.PolicyGroup, len(ap))
	for k, v := range ap {
		gp[k] = &casbin.PolicyGroup{
			// app_id
			RoleId: convetIdToString(v.RoleId),
			// 策略集ID
			PresetId: convetIdToString(v.PresetId),
		}
	}

	var pp []*casbin.PermissionPreset
	for k, v := range ps {
		var policies []*casbin.PermissionPreset
		err := json.Unmarshal([]byte(ps[k].Policies), &policies)
		if err != nil {
			return nil, nil, err
		}
		for _, j := range policies {
			j.PresetId = convetIdToString(v.Id)
		}
		pp = append(pp, policies...)
	}
	return gp, pp, nil
}
