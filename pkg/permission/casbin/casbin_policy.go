package casbin

// 权限策略集合
type PermissionPreset struct {
	// 策略集ID
	PresetId string `json:"preset_id"`
	// 版本
	Version string `json:"version"`
	// 动作
	Actions []string `json:"actions"`
	// 客体
	// action作用的对象
	Resources []string `json:"resources"`
	// Allow/Deny
	Effect string `json:"effect"`
	// 备注
	Remark string `json:"remark"`
}

// 权限组
type PolicyGroup struct {
	// role_id
	RoleId string `json:"role_id"`
	// 策略集ID
	PresetId string `json:"preset_id"`
}
