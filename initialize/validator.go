package initialize

import (
	"clockwerk/global"
	"clockwerk/utils"
	"fmt"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"

	"reflect"
	"strings"
)

// InitTrans validator信息翻译
func InitTrans(locale string) (err error) {
	color.Blue("init trans")
	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, enT)
		global.Trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			_ = en_translations.RegisterDefaultTranslations(v, global.Trans)
		case "zh":
			_ = zh_translations.RegisterDefaultTranslations(v, global.Trans)
		default:
			_ = en_translations.RegisterDefaultTranslations(v, global.Trans)
		}

		utils.RegisterValidatorFunc(v, "mobile", "手机号非法", utils.ValidateMobile)
		return
	}
	return
}
