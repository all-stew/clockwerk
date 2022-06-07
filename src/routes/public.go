package routes

import (
	"clockwerk/src/api"

	"github.com/gin-gonic/gin"
)

/*
	说明：公开的路由
*/

// 测试接口
func PublicRouters(r *gin.RouterGroup) gin.IRoutes {
	r.GET("/ping", api.PingHandler)
	return r
}
