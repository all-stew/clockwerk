package global

import (
	"clockwerk/app/common"
	"clockwerk/app/repository/store"
	"clockwerk/pkg/permission/casbin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/gobuffalo/packr/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/*
	说明：该文件定义了提供给全局使用的变量和常量
*/

var (
	Conf           common.Configuration     // 配置信息
	ConfBox        *packr.Box               // 配置打包
	Log            *zap.SugaredLogger       // 日志输出
	DB             *gorm.DB                 // 数据库
	Validate       *validator.Validate      // validation.v10 校验器
	Translator     ut.Translator            // validation.v10 翻译器
	CasbinEnforcer *casbin.CasbinPermission // cabin实例
	Redis          *redis.Client            // Redis
	UserStore      *store.UserStore         // userStore
	RoleStore      *store.RoleStore         // roleStore
)

// 时间格式化常量
const (
	MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
	SecLocalTimeFormat  = "2006-01-02 15:04:05"
	DateLocalTimeFormat = "2006-01-02"
)
