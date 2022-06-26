package v1

import (
	"clockwerk/app/global"
	"clockwerk/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// Store base64Captcha  缓存对象
var Store = base64Captcha.DefaultMemStore

type captchaAPIHandler struct {
}

func NewCaptchaAPIHandler() *captchaAPIHandler {
	return &captchaAPIHandler{}
}

// GetCaptcha 获取验证码
func (u *captchaAPIHandler) GetCaptcha(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, Store)
	// b64s是图片的base64编码
	id, b64s, err := cp.Generate()
	if err != nil {
		global.Log.Errorf("生成验证码错误,:%s ", err.Error())
		response.NewResult(ctx).Fail(500, "生成验证码错误")
		return
	}
	captchaObj := make(map[string]interface{})
	captchaObj["captchaId"] = id
	captchaObj["captcha"] = b64s

	response.NewResult(ctx).Success(captchaObj)
}
