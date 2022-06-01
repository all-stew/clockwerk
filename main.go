package main

import (
	"clockwerk/global"
	"clockwerk/initialize"
	"clockwerk/middlewares"
	"fmt"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"time"
)

func main() {
	// 1. 初始化yaml配置
	color.Cyan("1. initialize config...")
	initialize.InitConfig()
	// 2. 初始化routers
	color.Cyan("2. initialize routers...")
	Router := initialize.Routers()

	// 3. 初始化日志信息
	color.Cyan("3. initialize logger...")
	initialize.InitLogger()

	color.Cyan("4. initialize translation...")
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	// 5.初始化mysql
	initialize.InitMysqlDB()

	//6. 初始化redis
	initialize.InitRedis()
	global.Redis.Set("test", "testValue", time.Second)
	value := global.Redis.Get("test")
	color.Blue(value.Val())

	color.Cyan("4. use gin logger and recovery...")
	// 4. use ginLogger
	Router.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))

	// 5. 启动
	color.Cyan("5. gin is running...")
	err := Router.Run(fmt.Sprintf(":%d", global.ServerSetting.Port)) // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		zap.L().Info("error when starting")
	}

}
