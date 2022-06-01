package router

import (
	"clockwerk/controller"
	"clockwerk/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), controller.GetUserList)
		UserRouter.POST("login", controller.PasswordLogin)
	}
}
