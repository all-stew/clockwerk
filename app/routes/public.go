package routes

import (
	"clockwerk/app/api"

	"github.com/gin-gonic/gin"
)

// TestRouters 测试接口
func TestRouters(r *gin.RouterGroup) gin.IRoutes {
	// http :8080/api/v1/ping
	r.GET("/ping", api.PingHandler)
	// http :8080/api/v1/ping/id?id=20
	r.GET("/ping/id", api.PingValidatorHandler)
	// http POST :8080/api/v1/ping/id id:=20
	r.POST("/ping/id", api.PingValidatorFormHandler)
	return r
}
