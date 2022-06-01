package initialize

import (
	"clockwerk/global"
	"clockwerk/utils"
	"fmt"
	"go.uber.org/zap"
	"time"
)

// InitLogger 初始化Logger
func InitLogger() {
	// 实例化zap 配置
	cfg := zap.NewDevelopmentConfig()
	// 注意global.Settings.LogsAddress是在settings-dev.yaml配置过的
	// 配置日志的输出地址
	cfg.OutputPaths = []string{
		fmt.Sprintf("./logs/%s_log_%s.log", global.ServerSetting.Name, utils.GetFormatTime("%02d-%02d-%02d", time.Now())), //
		"stdout",
	}
	// 创建logger实例
	logg, _ := cfg.Build()
	// 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(logg)
	// 注册到全局变量中
	global.Lg = logg
}
