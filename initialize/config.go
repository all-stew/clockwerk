package initialize

import (
	"clockwerk/config"
	"clockwerk/global"
	"github.com/fatih/color"
	"github.com/spf13/viper"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile("./settings-dev.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	//给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递给全局变量
	global.ServerSetting = serverConfig
	color.Blue("test", global.ServerSetting.Name)
}
