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
	// 用户登录
	// http POST :8080/api/v1/login username="test" password="123456"
	r.POST("/login", auth.LoginHandler)
	// 用户登出  http POST :8080/api/v1/logout Authorization="Bearer xxx"
	r.POST("/logout", auth.LogoutHandler)
	// 刷新Token  http POST :8080/api/v1/refreshToken Authorization:"Bearer xxx"
	r.POST("/refreshToken", auth.RefreshHandler)
	return r
}
