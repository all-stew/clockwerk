package handler

import (
	"clockwerk/app/global"
	"clockwerk/app/views"
	"clockwerk/pkg/response"
	"clockwerk/pkg/validator"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// Authenticator 用户登陆验证
func Authenticator(c *gin.Context) (interface{}, error) {
	var loginView views.LoginView
	valid, _ := validator.BindAndValid(c, &loginView)
	if !valid {
		return nil, jwt.ErrMissingLoginValues
	}

	// 用户名密码验证
	user, err := global.UserStore.FindByUsernameAndPassword(c, loginView.Username, loginView.Password)
	// todo 验证码登陆
	// todo 登陆日志

	if err == nil {
		data := map[string]interface{}{
			"user": map[string]interface{}{
				"id":       user.Id,
				"username": user.Username,
				// TODO: 用户的角色信息
				"role": []string{"admin"},
			},
		}
		return data, nil
	} else {
		return nil, jwt.ErrFailedAuthentication
	}
}

// PayloadFunc 封装用户信息
func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		return jwt.MapClaims{
			jwt.IdentityKey: v["user"],
			"user":          v["user"],
		}
	}
	return jwt.MapClaims{}
}

func LoginResponse(ctx *gin.Context, code int, token string, expires time.Time) {
	response.NewResult(ctx).Success(map[string]interface{}{
		"token":  token,
		"expire": expires.Format("2006-01-02 15:04:05"),
	})
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
		"user":        claims["user"],
	}
}

func Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		c.Set("user", v["user"])
		return true
	}
	return false
}

func LogoutResponse(ctx *gin.Context, code int) {
	response.NewResult(ctx).Success(nil)
}

func Unauthorized(ctx *gin.Context, code int, message string) {
	response.NewResult(ctx).Fail(uint(code), message)
}

func RefreshResponse(ctx *gin.Context, code int, token string, expire time.Time) {
	response.NewResult(ctx).Success(map[string]interface{}{
		"token":  token,
		"expire": expire.Format("2006-01-02 15:04:05"),
	})
}
