package models

import (
	"fmt"
)

/*
   说明：Casbin 模型优化
*/
type CasbinRole struct {
	Id    uint   `gorm:"primaryKey;autoIncrement"`
	Ptype string `gorm:"size:100;uniqueIndex:uk_index;comment:策略类型"` // 多个字段联合唯一
	V0    string `gorm:"size:100;uniqueIndex:uk_index;comment:角色关键字"`
	V1    string `gorm:"size:100;uniqueIndex:uk_index;comment:资源名称"`
	V2    string `gorm:"size:100;uniqueIndex:uk_index;comment:请求类型"`
	V3    string `gorm:"size:100;uniqueIndex:uk_index"`
	V4    string `gorm:"size:100;uniqueIndex:uk_index"`
	V5    string `gorm:"size:100;uniqueIndex:uk_index"`
}

// 表名
func (c CasbinRole) TableName() string {
	return fmt.Sprintf("%s", "casbin_rule")
}

// 角色权限结构体
type CasbinRoleData struct {
	Keyword string `json:"keyword"` // 角色关键字
	Path    string `json:"path"`    // 请求地址
	Method  string `json:"method"`  // 访问路径
}
