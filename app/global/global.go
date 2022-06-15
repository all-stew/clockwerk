package global

import (
	"clockwerk/app/common"
	"clockwerk/app/repository/store"
	"clockwerk/pkg/permission/casbin"
	"errors"
	"strings"

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
)

// 时间格式化常量
const (
	MsecLocalTimeFormat = "2006-01-02 15:04:05.000"
	SecLocalTimeFormat  = "2006-01-02 15:04:05"
	DateLocalTimeFormat = "2006-01-02"
)

// NewValidatorError 方法说明：对参数校验的错误进行重写参数：1. 原本的错误信息2. 字段翻译 map3. 错误翻译 map
func NewValidatorError(err error, fieldTrans map[string]string, fieldError map[string]string) error {
	if err == nil {
		return nil
	}

	// 获取校验错误
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		tranStr := e.Translate(Translator)

		// 字段名称
		field := e.Field()

		// 自定义错误：先判断错误是否被重写
		v, ok := fieldError[field]
		if ok {
			// 返回自定义错误
			return errors.New(v)
		}

		// 系统错误：为重写错误信息的字段
		v, ok = fieldTrans[field]
		if ok {
			// 替换掉英文字段为中文
			return errors.New(strings.Replace(tranStr, e.Field(), v, -1))
		}
		return errors.New(tranStr)
	}
	return nil
}
