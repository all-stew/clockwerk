package initialize

import (
	"clockwerk/app/global"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	说明：日志初始化
*/

// 定义日志时间格式
func ZapLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.MsecLocalTimeFormat))
}

// 初始化日志打印
func Logger() {
	// 定义基础变量
	// 当前时间
	now := time.Now()
	// 日志保存路径
	filename := fmt.Sprintf("%s/%s.%04d-%02d-%02d.log", global.Conf.Logs.Path, global.Conf.System.Name, now.Year(), now.Month(), now.Day())
	hook := &lumberjack.Logger{
		Filename:   filename,                    // 日志保存路径
		MaxSize:    global.Conf.Logs.MaxSize,    // 文件最大
		MaxAge:     global.Conf.Logs.MaxAge,     // 文件保留天数
		MaxBackups: global.Conf.Logs.MaxBackups, // 文件备份个数
		Compress:   global.Conf.Logs.Compress,   // 是否压缩
	}
	defer hook.Close()

	// 简易配置 zap
	enConfig := zap.NewProductionEncoderConfig()
	enConfig.EncodeTime = ZapLocalTimeEncoder

	// 日志等级颜色输出处理
	if global.Conf.Logs.Colorful {
		enConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		enConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	// 日志输出位置配置(控制台/文件输出)
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(enConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)),
		global.Conf.Logs.Level,
	)

	// 处理日志输出中打印当前文件的问题
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// 配置全局答应变量
	global.Log = logger.Sugar()

	// 打印输出日志
	log.Println("日志初始化完成")
}
