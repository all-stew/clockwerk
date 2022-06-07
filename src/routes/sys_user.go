package routes

import (
	"clockwerk/src/api/v1/sys"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

/*
   说明：用户相关路由
*/
func UserRouters(r *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {

	rg := r.Group("user").Use(auth.MiddlewareFunc())
	//.Use(middleware.Casbin) // 当前的路由都是需要登录的
	r.POST("/create", sys.Create)

	return rg
}
