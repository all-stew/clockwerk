package middleware

/*
   说明：Casbin 权限中间件
*/
//func Casbin(ctx *gin.Context) {
//	// 获取当前用户
//	user, _ := system.GETCurrentUserInfo(ctx)
//	// 获取当前用户的角色关键字作为 sub
//	sub := user.Role.Keyword
//	// 请求 URL 路径作为访问资源 obj（需要清除前缀）
//	urlPrefix := "/" + global.Conf.System.ApiPrefix + "/" + global.Conf.System.ApiVersion
//	obj := strings.Replace(ctx.Request.URL.Path, urlPrefix, "", 1)
//	// 请求方式作为请求的 act
//	act := ctx.Request.Method
//
//	// 权限检查
//	pass, _ := global.CasbinEnforcer.Enforce(sub, obj, act)
//	if !pass {
//		response2.FailWithCode(response2.Forbidden)
//		return
//	}
//
//	// 继续处理请求
//	ctx.Next()
//}
