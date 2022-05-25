package main

import (
	"clockwerk/config"
	"clockwerk/config/mysql"
	"clockwerk/router"
	"clockwerk/task"
	"fmt"
)

func main() {
	go func() {
		task.Schedule()
	}()
	if err := config.Init("./config.yaml"); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(config.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}

	// 注册路由
	r := router.SetupRouter(config.Config.Mode)
	err := r.Run(fmt.Sprintf(":%d", config.Config.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
