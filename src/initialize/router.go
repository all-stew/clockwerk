package initialize

import (
	"clockwerk/src/global"
	"clockwerk/src/middlewares"
	"clockwerk/src/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

/*
	说明：路由初始化文件
*/

// 路由初始化
func Router() *gin.Engine {
	// 设置运行模式
	gin.SetMode(global.Conf.System.Mode)

	// 创建不带中间件的路由
	r := gin.New()

	// 中间件
	r.Use(middleware.AccessLog) // 访问中间件
	r.Use(middleware.Cors)      // 跨域中间件
	//r.Use(middleware.Exception)        // 异常中间件
	auth, err := middleware.AuthInit() // JWT中间件
	if err != nil {
		panic(fmt.Sprintf("JWT中间件初始化失败：%v", err))
	}

	// 添加 api 前缀
	apiGroup := r.Group(global.Conf.System.ApiPrefix)
	// 添加 v1 前缀
	vGroup := apiGroup.Group(global.Conf.System.ApiVersion)
	{
		routes.PublicRouters(vGroup)     // 开放接口
		routes.BaseRouters(vGroup, auth) // 登录登出接口
		// 系统模块路由入口
		systemGroup := vGroup.Group("system")
		{
			routes.UserRouters(systemGroup, auth) // 用户相关接口
		}
	}
	log.Println("路由初始化完成")
	return r
}
