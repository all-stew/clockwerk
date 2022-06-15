package routes

import (
	v1 "clockwerk/app/api/sys/v1"
	middleware "clockwerk/app/middlewares"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// UserRouters 用户相关路由
func UserRouters(r *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {

	rg := r.Group("user").Use(auth.MiddlewareFunc()).Use(middleware.Casbin) // 当前的路由都是需要登录的

	// 获取用户列表
	rg.GET("/list")
	// 获取user详情
	rg.GET("/:userId")
	// 新增用户
	rg.POST("/create", v1.Create)
	return rg
}
