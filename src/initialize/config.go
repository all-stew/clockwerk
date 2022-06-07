package initialize

import (
	"bytes"
	"clockwerk/src/global"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/viper"
)

/*
	说明：配置文件初始化
*/

// 定义配置文件相关信息
const (
	envName       = "RUN_ENV"               // 环境变量名称，用于判定运行环境，从而适配配置
	configBoxName = "gin-config-box"        // packr 需要定义一个名称空间用于存放配置
	configType    = "yml"                   // 配置文件类型，viper 会根据它解析
	configPath    = "../config"             // 配置文件相对目录
	devConfig     = "application.dev.yaml"  // 开发环境配置文件名称
	testConfig    = "application.test.yaml" // 测试环境配置文件名称
	prodConfig    = "application.prod.yaml" // 生产环境配置文件名称
)

// 读取配置文件函数
func ReadConfig(v *viper.Viper, file string) {
	// 属性设置
	v.SetConfigType(configType)

	// 判断配置是否存在
	config, err := global.ConfBox.Find(file)
	if err != nil {
		panic(fmt.Sprintf("配置文件不存在：%s", err.Error()))
	}

	// 加载配置
	err = v.ReadConfig(bytes.NewReader(config))
	if err != nil {
		panic(fmt.Sprintf("配置文件初始化失败：%s", err.Error()))
	}
}

// 初始化配置文件函数
func Config() {
	// 设置打包配置文件
	global.ConfBox = packr.New(configBoxName, configPath)

	/*
	   读取配置，默认先加载开发配置
	   这意味着开发配置文件最好是包含全部配置项，其它配置文件只需要保存变更的配置即可
	*/
	v := viper.New()
	ReadConfig(v, devConfig)

	// 设置配置
	settings := v.AllSettings()
	for index, setting := range settings {
		v.SetDefault(index, setting)
	}

	// 读取环境变量(RUN_ENV)，确定运行环境
	runEnv := strings.ToLower(os.Getenv(envName))

	// 判断不同环境再次加载不同配置
	configName := ""
	if runEnv == "test" {
		configName = testConfig
	} else if runEnv == "prod" {
		configName = prodConfig
	}

	// 如果指定了环境，则加载指定环境
	if configName != "" {
		ReadConfig(v, configName)
	}

	// 将配置转换成结构体
	err := v.Unmarshal(&global.Conf)
	if err != nil {
		panic(fmt.Sprintf("配置文件解析失败：%s", err.Error()))
	}
	log.Println("配置初始化完成")
}
