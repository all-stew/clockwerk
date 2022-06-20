package routes

import (
	sysV1 "clockwerk/app/api/sys/v1"
	middleware "clockwerk/app/middlewares"
	"clockwerk/app/service/impl"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// UserRouters 用户相关路由
func UserRouters(r *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {

	u := sysV1.NewUserAPIHandler(impl.GetUserServiceImpl())

	rg := r.Group("user").Use(auth.MiddlewareFunc()).Use(middleware.Casbin) // 当前的路由都是需要登录的

	rg.POST("/change-password")
	// 获取用户列表
	rg.GET("/list")
	// 获取user详情
	rg.GET("/:userId")
	// 新增用户
	rg.POST("/create", u.Create)
	return rg
}
