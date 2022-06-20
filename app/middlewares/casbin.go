package middleware

import (
	"clockwerk/app/global"
	"clockwerk/pkg/response"
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

	// 获取当前用户的相关信息
	userData := ctx.MustGet("user").(interface{})
	_, ok := userData.(map[string]interface{})
	if !ok {
		response.NewResult(ctx).Fail(403, "Unauth")
		ctx.Abort()
	}

	// 获取用户角色额
	rolesData := ctx.MustGet("roles").([]interface{})
	roles := make([]string, len(rolesData))
	for i, v := range rolesData {
		roles[i] = v.(string)
	}

	// 白名单不需要检测
	if whiteList[obj] == act {
		ctx.Next()
	}

	// 权限检查
	for _, sub := range roles {
		pass, _ := global.CasbinEnforcer.CheckPermission(sub, obj, act)
		// 只要有一个角色能够支持访问该接口即可
		if pass {
			// 继续处理请求
			ctx.Next()
			return
		}
	}

	// 没有一个角色能够满足该条件直接返回error
	ctx.Abort()
	fmt.Println("casbin out")
	response.NewResult(ctx).Fail(403, "Unauth")
}
