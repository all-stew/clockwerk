package casbin

import (
	"testing"
)

type testModel struct {
	name     string
	appId    string
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
			Actions: []string{"com.test.com/DeviceAPI.GetDevice"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "deny",
		},
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"com.test.com/DeviceAPI.*"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "2",
			// 动作
			Actions: []string{"com.test.com/ReportAPI.GetReport", "com.test.com/ReportAPI.SubmitReport"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "3",
			// 动作
			Actions: []string{"com.test.com/UserAPI*"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "3",
			// 动作
			Actions: []string{"com.test.com/UserAPI.GetUser", "com.test.com/UserAPI.DeleteUser"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "deny",
		},
	}
	s := []*PolicyGroup{
		{
			// app_id
			AppId: "hmy",
			// 策略集ID
			PresetId: "1",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "1",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "2",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "3",
		},
	}
	e, _ := NewCasbinPermission(p, s)

	tests := []testModel{
		{
			name:     "test1",
			appId:    "hmy",
			resource: "*",
			action:   "com.test.com/DeviceAPI.DeleteDevice",
			want:     true,
		},
		{
			name:     "test1",
			appId:    "hmy",
			resource: "*",
			action:   "com.test.com/DeviceAPI.GetDevice",
			want:     false,
		},
		{
			name:     "test1",
			appId:    "hmy",
			resource: "asodkasokdos",
			action:   "com.test.com/DeviceAPI.GetDevice",
			want:     false,
		},
		{
			name:     "test1",
			appId:    "hmy",
			resource: "asodkasokdos",
			action:   "com.test.com/DeviceAPI.DeleteDevice",
			want:     true,
		},
		{
			name:     "test1",
			appId:    "hmy",
			resource: "",
			action:   "com.test.com/DeviceAPI.DeleteDevice",
			want:     true,
		},
		{
			name:     "test1",
			appId:    "hmy",
			resource: "",
			action:   "com.test.com/DeviceAPI.GetDevice",
			want:     false,
		},
		{
			name:     "test1",
			appId:    "admin",
			resource: "*",
			action:   "com.test.com/UserAPI.GetUser",
			want:     false,
		},
		{
			name:     "test1",
			appId:    "admin",
			resource: "*",
			action:   "com.test.com/UserAPI.SignIn",
			want:     true,
		},
	}

	for _, v := range tests {
		// 验证
		t.Run(v.name, func(t *testing.T) {
			if ok, _ := e.CheckPermission(v.appId, v.action, v.resource); ok != v.want {
				t.Errorf("check permission failed,%s %s %s", v.appId, v.resource, v.action)
			}
		})
	}
}

func TestRemovePermission(t *testing.T) {
	p := []*PermissionPreset{
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"com.test.com/DeviceAPI.GetDevice"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "deny",
		},
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"com.test.com/DeviceAPI.*"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "2",
			// 动作
			Actions: []string{"com.test.com/ReportAPI.GetReport", "com.test.com/ReportAPI.SubmitReport"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "3",
			// 动作
			Actions: []string{"com.test.com/UserAPI*"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "3",
			// 动作
			Actions: []string{"com.test.com/UserAPI.GetUser", "com.test.com/UserAPI.DeleteUser"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "deny",
		},
	}
	s := []*PolicyGroup{
		{
			// app_id
			AppId: "hmy",
			// 策略集ID
			PresetId: "1",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "1",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "2",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "3",
		},
	}

	e, _ := NewCasbinPermission(p, s)

	err := e.RemovePermission("1")
	if err != nil {
		return
	}

	tests := []testModel{
		{
			name:     "test1",
			appId:    "hmy",
			resource: "*",
			action:   "com.test.com/DeviceAPI.DeleteDevice",
			want:     false,
		},
	}

	for _, v := range tests {
		// 验证
		t.Run(v.name, func(t *testing.T) {
			if ok, _ := e.CheckPermission(v.appId, v.action, v.resource); ok != v.want {
				t.Errorf("check permission failed,%s %s %s", v.appId, v.resource, v.action)
			}
		})
	}
}

func TestRemoveAppPermissionGroup(t *testing.T) {
	p := []*PermissionPreset{
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"com.test.com/DeviceAPI.GetDevice"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "deny",
		},
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"com.test.com/DeviceAPI.*"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "2",
			// 动作
			Actions: []string{"com.test.com/ReportAPI.GetReport", "com.test.com/ReportAPI.SubmitReport"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "3",
			// 动作
			Actions: []string{"com.test.com/UserAPI*"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "3",
			// 动作
			Actions: []string{"com.test.com/UserAPI.GetUser", "com.test.com/UserAPI.DeleteUser"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "deny",
		},
	}
	s := []*PolicyGroup{
		{
			// app_id
			AppId: "hmy",
			// 策略集ID
			PresetId: "1",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "1",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "2",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "3",
		},
	}

	e, _ := NewCasbinPermission(p, s)

	err := e.DeleteAppPermissionGroup("admin")
	if err != nil {
		return
	}

	tests := []testModel{
		{
			name:     "test1",
			appId:    "admin",
			resource: "*",
			action:   "com.test.com/DeviceAPI.DeleteDevice",
			want:     false,
		},
	}

	for _, v := range tests {
		// 验证
		t.Run(v.name, func(t *testing.T) {
			if ok, _ := e.CheckPermission(v.appId, v.action, v.resource); ok != v.want {
				t.Errorf("check permission failed,%s %s %s", v.appId, v.resource, v.action)
			}
		})
	}
}

func TestRemovePermissionGroup(t *testing.T) {
	p := []*PermissionPreset{
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"com.test.com/DeviceAPI.GetDevice"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "deny",
		},
		{
			// 策略集ID
			PresetId: "1",
			// 动作
			Actions: []string{"com.test.com/DeviceAPI.*"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "2",
			// 动作
			Actions: []string{"com.test.com/ReportAPI.GetReport", "com.test.com/ReportAPI.SubmitReport"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "3",
			// 动作
			Actions: []string{"com.test.com/UserAPI*"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "allow",
		},
		{
			// 策略集ID
			PresetId: "3",
			// 动作
			Actions: []string{"com.test.com/UserAPI.GetUser", "com.test.com/UserAPI.DeleteUser"},
			// * 客体
			// action作用的对象
			Resources: []string{"*"},
			// Allow/Deny
			Effect: "deny",
		},
	}
	s := []*PolicyGroup{
		{
			// app_id
			AppId: "hmy",
			// 策略集ID
			PresetId: "1",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "1",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "2",
		},
		{
			// app_id
			AppId: "admin",
			// 策略集ID
			PresetId: "3",
		},
	}

	e, _ := NewCasbinPermission(p, s)

	err := e.DeletePermissionGroup("admin", "1")
	if err != nil {
		return
	}

	tests := []testModel{
		{
			name:     "test1",
			appId:    "admin",
			resource: "*",
			action:   "com.test.com/DeviceAPI.DeleteDevice",
			want:     false,
		},
	}

	for _, v := range tests {
		// 验证
		t.Run(v.name, func(t *testing.T) {
			if ok, _ := e.CheckPermission(v.appId, v.action, v.resource); ok != v.want {
				t.Errorf("check permission failed,%s %s %s", v.appId, v.resource, v.action)
			}
		})
	}
}
