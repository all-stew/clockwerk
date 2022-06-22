package routes

import (
	sysV1 "clockwerk/app/api/sys/v1"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

/*
   说明：登录登出相关路由
*/

// 基础路由
func BaseRouters(r *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	u := sysV1.NewCaptchaAPIHandler()

	// http GET :8080/api/v1/captcha
	r.GET("/captcha", u.GetCaptcha)
	// http POST :8080/api/v1/login username="test" password="123456"
	r.POST("/login", auth.LoginHandler)
	// 用户登出  http POST :8080/api/v1/logout Authorization="Bearer xxx"
	r.POST("/logout", auth.LogoutHandler)
	// 刷新Token  http POST :8080/api/v1/refreshToken Authorization:"Bearer xxx"
	r.POST("/refreshToken", auth.RefreshHandler)
	return r
}
