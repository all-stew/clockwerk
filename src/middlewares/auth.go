package middleware

import (
	"clockwerk/src/global"
	"clockwerk/src/middlewares/handler"
	jwt "github.com/appleboy/gin-jwt/v2"
	"time"
)

// AuthInit jwt 认证初始化
func AuthInit() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour

	if global.Conf.System.Mode == "dev" {
		timeout = time.Duration(876010) * time.Hour
	} else {
		timeout = time.Duration(global.Conf.Jwt.Timeout) * time.Second
	}

	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:           global.Conf.Jwt.Realm,
		Key:             []byte(global.Conf.Jwt.Key),
		Timeout:         timeout,
		MaxRefresh:      time.Hour,
		PayloadFunc:     handler.PayloadFunc,
		IdentityHandler: handler.IdentityHandler,
		Authenticator:   handler.Authenticator,
		Authorizator:    handler.Authorizator,
		Unauthorized:    handler.Unauthorized,
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})
}
