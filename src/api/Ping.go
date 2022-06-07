package api

import (
	"clockwerk/pkg/response"
	"clockwerk/src/global"
	"github.com/gin-gonic/gin"
)

/*
	说明：测试处理函数
*/
func PingHandler(ctx *gin.Context) {
	response.Success(ctx, 200, "成功", map[string]interface{}{
		"name":    global.Conf.System.Name,
		"message": "pong",
	})
}
