package main

import (
	"clockwerk/config"
	"clockwerk/router"
	"clockwerk/util/logger"
	"fmt"
)

func main() {
	if err := config.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(config.Config.LogConfig, config.Config.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	// todo mysql 关闭 暂时不用
	//if err := mysql.init(config.Conf.MySQLConfig); err != nil {
	//	fmt.Printf("init mysql failed, err:%v\n", err)
	//	return
	//}
	//// 程序退出关闭数据库连接
	//defer mysql.Close()

	// 注册路由
	r := router.SetupRouter(config.Config.Mode)
	err := r.Run(fmt.Sprintf(":%d", config.Config.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
