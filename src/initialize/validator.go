package initialize

import (
	"clockwerk/src/global"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"regexp"
)

/*
   说明：上传数据校验器初始化
*/
func Validate() {
	// 语言处理，中文
	translator := zh.New()
	uni := ut.New(translator, translator)
	trans, _ := uni.GetTranslator("zh")

	validate := validator.New()

	// 注册自定义验证
	validate, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// 注册失败错误信息模板
		msg := ""
		msgTmpl := "注册自定义参数校验方法错误："

		err := validate.RegisterValidation("is-username", validateUsername)
		if err != nil {
			msg = msgTmpl + "is-username"
		}

		err = validate.RegisterValidation("is-mobile", validateMobile)
		if err != nil {
			msg = msgTmpl + "is-mobile"
		}

		if msg != "" {
			log.Println(msg)
			panic(msg)
		}
	}

	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)
	global.Validate = validate
	global.Translator = trans
	log.Println("Validator.v10校验器初始化完成")
}

/*
   说明：自定义验证
*/
// 手机号验证
func validateMobile(fl validator.FieldLevel) bool {
	num, ok := fl.Field().Interface().(string)
	if ok {
		reg := `^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`
		rgx := regexp.MustCompile(reg)
		return rgx.MatchString(num)
	}
	return false
}

// 用户名验证：用户名必须为小写字母开头长度为4-30的字符串（只支持字母数字）
func validateUsername(fl validator.FieldLevel) bool {
	str, ok := fl.Field().Interface().(string)
	if ok {
		reg := `^[a-z][0-9a-z]{3,30}$`
		rgx := regexp.MustCompile(reg)
		return rgx.MatchString(str)
	}
	return false
}
