package handler

import (
	"clockwerk/src/global"
	. "clockwerk/src/models"
	"clockwerk/src/views"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(SysUser)
		return jwt.MapClaims{
			jwt.IdentityKey: u.Id,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		"IdentityKey": claims["identity"],
	}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var loginView views.LoginView
	if err := c.ShouldBind(&loginView); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	user, err := global.UserStore.FindByUsernameAndPassword(c, loginView.Username, loginView.Password)
	// todo 验证码登陆
	// todo 登陆日志

	if err == nil {
		return map[string]interface{}{"user": user}, nil
	} else {
		return nil, jwt.ErrFailedAuthentication
	}
}

func LogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "退出成功",
	})

}

func Authorizator(data interface{}, c *gin.Context) bool {

	if v, ok := data.(map[string]interface{}); ok {
		u, _ := v["user"].(SysUser)
		c.Set("userId", u.Id)
		return true
	}
	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
