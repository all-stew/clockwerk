package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

/*
   说明：登录登出相关路由
*/

// 基础路由
func BaseRouters(r *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	r.POST("/login", auth.LoginHandler)          // 用户登录
	r.POST("/logout", auth.LogoutHandler)        // 用户登出
	r.POST("/refreshToken", auth.RefreshHandler) // 刷新Token
	return r
}
