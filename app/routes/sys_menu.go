package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// MenuRouters 角色相关路由
func MenuRouters(r *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {

	rg := r.Group("menu").Use(auth.MiddlewareFunc())
	//.Use(middleware.Casbin) // 当前的路由都是需要登录的

	// 获取角色列表
	rg.GET("/list")
	// 获取角色详情
	rg.GET("/:menuId")
	// 新增角色
	rg.POST("/")
	// 修改角色
	rg.PUT("/")
	// 删除角色
	rg.DELETE("/:menuId")
	return rg
}
