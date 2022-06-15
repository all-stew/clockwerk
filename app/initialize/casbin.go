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

	// 转化为casbin的结构
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
