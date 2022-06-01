package initialize

import (
	"clockwerk/middlewares"
	"clockwerk/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/api/v1/")
	router.UserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	return Router
}
