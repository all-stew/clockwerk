package middleware

import (
	"clockwerk/app/global"
	"fmt"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

// 白名单
var (
	whiteList = map[string]string{
		"/system/user/xxx": "Get",
	}
)

/*
   说明：Casbin 权限中间件
*/
func Casbin(ctx *gin.Context) {
	// 获取当前请求的信息
	urlPrefix := path.Join("/", global.Conf.System.ApiPrefix, "/", global.Conf.System.ApiVersion)
	obj := strings.Replace(ctx.Request.URL.Path, urlPrefix, "", 1)
	// 请求方式作为请求的 act
	act := ctx.Request.Method

	fmt.Println(obj)
	fmt.Println(act)

	fmt.Println("casbin in")
	// 获取当前用户的信息
	clamins, ok := ctx.MustGet("JWT_PAYLOAD").(map[string]interface{})
	fmt.Println("claims", clamins)
	if !ok {
		return
	}
	// 获取当前用户的相关信息
	user, ok := clamins["user"].(map[string]interface{})
	if !ok {
		return
	}
	fmt.Println("user", user)

	// 获取用户角色额
	roles, ok := user["role"].([]string)
	if !ok {
		return
	}

	fmt.Println("roles", roles)

	// 白名单不需要检测
	if whiteList[obj] == act {
		ctx.Next()
	}

	// 权限检查
	for _, sub := range roles {
		pass, _ := global.CasbinEnforcer.CheckPermission(sub, act, obj)
		// 只要有一个角色能够支持访问该接口即可
		if pass {
			// 继续处理请求
			ctx.Next()
		}
	}

	// 没有一个角色能够满足该条件直接返回error
	return
}
