package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// RoleRouters 角色相关路由
func RoleRouters(r *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {

	rg := r.Group("role").Use(auth.MiddlewareFunc())
	//.Use(middleware.Casbin) // 当前的路由都是需要登录的

	// 获取角色列表
	rg.GET("/list")
	// 获取角色详情
	rg.GET("/:roleId")
	// 新增角色
	rg.POST("/")
	// 修改角色
	rg.PUT("/")
	// 删除角色
	rg.DELETE("/:roleId")
	// 获取用户授权角色
	rg.GET("/:roleId/authRole")
	// 获取用户未授权角色
	rg.GET("/:roleId/unAuthRole")
	// 给用户授权角色
	rg.POST("/authRole")
	// 给用户授权角色
	rg.DELETE("/authRole")
	return rg
}
