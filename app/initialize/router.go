package initialize

import (
	"clockwerk/app/global"
	"clockwerk/app/middlewares"
	"clockwerk/app/routes"
	"clockwerk/pkg/response"
	"fmt"
	"log"
	"runtime/debug"

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

	r.NoRoute(HandleNotFound)
	r.NoMethod(HandleNotFound)
	r.Use(Recover)
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
		routes.TestRouters(vGroup)       // 测试接口接口
		routes.BaseRouters(vGroup, auth) // 登录登出接口
		// 系统模块路由入口
		systemGroup := vGroup.Group("system")
		{
			routes.UserRouters(systemGroup, auth) // 用户相关接口
			routes.RoleRouters(systemGroup, auth) // 角色相关接口
			routes.MenuRouters(systemGroup, auth) // 菜单相关接口
		}
	}
	log.Println("路由初始化完成")
	return r
}

func HandleNotFound(c *gin.Context) {
	global.Log.Errorf("handle not found: %v", c.Request.RequestURI)
	response.NewResult(c).Fail(404, "资源未找到")
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			//log.Printf("panic: %v\n", r)
			global.Log.Errorf("panic: %v", r)
			//log stack
			global.Log.Errorf("stack: %v", string(debug.Stack()))
			//print stack
			debug.PrintStack()
			//return
			response.NewResult(c).Fail(500, "服务器内部错误")
		}
	}()
	//继续后续接口调用
	c.Next()
}
