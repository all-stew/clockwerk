package casbin

import (
	"testing"
)

type testModel struct {
	name     string
	roleId   string
	action   string
	resource string
	want     bool
}

func TestMatch(t *testing.T) {
	p := []*PermissionPreset{
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"Get"},
			// * 客体
			// action作用的对象
			Resources: []string{"/system/user/list"},
			// Allow/Deny
			Effect: "deny",
		},
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"Post", "Put", "Delete"},
			// * 客体
			// action作用的对象
			Resources: []string{"/system/user/list"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "2",
			// 动作
			Actions: []string{"*"},
			// * 客体
			// action作用的对象
			Resources: []string{"/system/user/cast"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "3",
			// 动作
			Actions: []string{"Post"},
			// * 客体
			// action作用的对象
			Resources: []string{"/system/user/userinfo"},
			// Allow/Deny
			Effect: "allow",
		},
	}
	s := []*PolicyGroup{
		{
			// app_id
			RoleId: "admin",
			// 策略集ID
			PresetId: "1",
		},
		{
			// app_id
			RoleId: "work",
			// 策略集ID
			PresetId: "2",
		},
		{
			// app_id
			RoleId: "test",
			// 策略集ID
			PresetId: "3",
		},
	}
	e, _ := NewCasbinPermission(p, s)

	tests := []testModel{
		{
			name:     "test1",
			roleId:   "admin",
			resource: "/system/user/list",
			action:   "Get",
			want:     false,
		},
		{
			name:     "test1",
			roleId:   "admin",
			resource: "/system/user/list",
			action:   "Post",
			want:     true,
		},
		{
			name:     "test1",
			roleId:   "admin",
			resource: "/system/user/list",
			action:   "Delete",
			want:     true,
		},
		{
			name:     "test1",
			roleId:   "admin",
			resource: "/system/user/list",
			action:   "*",
			want:     false,
		},
		{
			name:     "test1",
			roleId:   "test",
			resource: "/system/user/userinfo",
			action:   "Post",
			want:     true,
		},
	}

	for _, v := range tests {
		// 验证
		t.Run(v.name, func(t *testing.T) {
			if ok, _ := e.CheckPermission(v.roleId, v.resource, v.action); ok != v.want {
				t.Errorf("check permission failed,%s %s %s", v.roleId, v.resource, v.action)
			}
		})
	}
}
