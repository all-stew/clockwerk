package middleware

import (
	"clockwerk/src/global"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

/*
	说明：访问日志中间件
*/
func AccessLog(ctx *gin.Context) {
	// 开始时间
	requestStartTime := time.Now()
	// 处理请求
	ctx.Next()
	// 结束时间
	requestEndTime := time.Now()
	// 处理耗时
	requestExecTime := requestEndTime.Sub(requestStartTime)
	// 请求方式
	requestMethod := ctx.Request.Method
	// 请求路由
	requestUri := ctx.Request.RequestURI
	// 状态码
	requestCode := ctx.Writer.Status()
	// 请求 IP
	requestIP := ctx.ClientIP()

	// 判断请求方式，OPTIONS 使用 DEBUG 输出
	if requestMethod == "OPTIONS" {
		global.Log.Debug(fmt.Sprintf("%s\t%s\t%d\t%s\t%s", requestMethod, requestUri, requestCode, requestExecTime.String(), requestIP))
	} else {
		global.Log.Info(fmt.Sprintf("%s\t%s\t%d\t%s\t%s", requestMethod, requestUri, requestCode, requestExecTime.String(), requestIP))
	}
}
