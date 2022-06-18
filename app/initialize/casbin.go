package initialize

import (
	"clockwerk/app/global"
	"clockwerk/app/models"
	"clockwerk/pkg/permission/casbin"
	"encoding/json"
	"fmt"
	"strconv"
)

/*
   说明：初始化 Casbin
*/
func Casbin() {
	// 从数据库读取数据
	// 加载preset策略集
	var presetsModels []models.Preset
	err := global.DB.Model(&models.Preset{}).Find(&presetsModels).Error
	if err != nil {
		panic(fmt.Sprintf("Casbin加载数据失败：%s", err.Error()))
	}

	// 加载角色与preset的关系
	var rolePresetsModels []models.RolePreset
	err = global.DB.Model(&models.RolePreset{}).Find(&rolePresetsModels).Error
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

	global.Log.Info("Casbin初始化完成")
}

func convetIdToString(id uint64) string {
	return strconv.Itoa(int(id))
}

// convertPresetPolicies
func convertPresetPolicies(ap []models.RolePreset, ps []models.Preset) ([]*casbin.PolicyGroup, []*casbin.PermissionPreset, error) {
	gp := make([]*casbin.PolicyGroup, len(ap))
	for k, v := range ap {
		gp[k] = &casbin.PolicyGroup{
			// app_id
			RoleId: v.RoleId,
			// 策略集ID
			PresetId: v.PresetId,
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
