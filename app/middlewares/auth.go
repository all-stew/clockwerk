package middleware

import (
	"clockwerk/app/global"
	"clockwerk/app/middlewares/handler"
	jwt "github.com/appleboy/gin-jwt/v2"
	"time"
)

/*
LoginHandler（中间件）：
    说明：用于登录操作。

    Authenticator（必需）：
        说明：使用 LoginHandler 就会触发该函数，从 ctx 中取出用户信息进行校验，然后返回一个包含用户信息的 struct 或 map。
        后续：
            验证成功，将返回作为参数传递给 PayloadFunc 函数。
            验证失败，触发 Unauthorized 函数。

    PayloadFunc（可选）：
        说明：接收从 Authenticator 传递的 struct 或 map，提取用户信息再封装成 MapClaims（类型：map[string]interface{}）。
        注意：
            1. MapClaims 必须包含属性 IdentityKey。
            2. MapClaims 将被嵌入 token 中作为 token claims。
            3. 之后的请求，用户传入 token，可以使用 ExtractClaims 来取出封装的数据。

    LoginResponse（可选）：
        说明：PayloadFunc 创建 token 后，如果 SendCookie 开启，则调用该函数对用户请求进行响应。

MiddlewareFunc（中间件）：
    	说明：处理登录后请求，通过 IdentityHandler 和 Authorizator 对 Token 进行检查，都通过则继续，失败则调用 Unauthorized。

    IdentityHandler（可选）：
        说明：从 token 中获取用户信息，将数据封装后传递给 Authorizator，返回的数据需要 IdentityKey 属性。

    Authorizator（可选）：
        说明：检查用户是否登录，通过 ExtractClaims 检查，而不是查数据库。登录返回 true，否则返回 false，并调用 Unauthorized。

LogoutHandler（中间件）：
    说明：如果配置了 SendCookie，将清除所有 cookie 并调用 LogoutResponse。

    LogoutResponse（可选）：
        说明：返回用户注销响应。

RefreshHandler（中间件）：
    说明：返回 token 的 json 给用户，和 LoginResponse 类似，使用默认的即可。
*/

// AuthInit jwt 认证初始化
func AuthInit() (*jwt.GinJWTMiddleware, error) {
	timeout := time.Hour
	maxRefresh := time.Hour

	if global.Conf.System.Mode == "dev" {
		timeout = time.Duration(876010) * time.Hour
		maxRefresh = time.Duration(876010) * time.Hour
	} else {
		timeout = time.Duration(global.Conf.Jwt.Timeout) * time.Second
		maxRefresh = time.Duration(global.Conf.Jwt.MaxRefresh) * time.Second
	}

	return jwt.New(&jwt.GinJWTMiddleware{

		Realm:           global.Conf.Jwt.Realm,       // JWT 标识
		Key:             []byte(global.Conf.Jwt.Key), // 服务端密钥
		Timeout:         timeout,                     // 有效时间
		MaxRefresh:      maxRefresh,                  // 最大刷新时间
		Authenticator:   handler.Authenticator,       // 用户登陆认证
		PayloadFunc:     handler.PayloadFunc,         // 用户登录信息封装
		LoginResponse:   handler.LoginResponse,       // 用户登录成功响应
		IdentityHandler: handler.IdentityHandler,     // 用户后续访问 Token 解封装，再封装
		Authorizator:    handler.Authorizator,        // 用户后续访问 Token 校验
		LogoutResponse:  handler.LogoutResponse,      // 用户登出处理
		Unauthorized:    handler.Unauthorized,        // 用户登录/验证失败响应
		RefreshResponse: handler.RefreshResponse,     // 刷新token成功响应
		TokenLookup:     "header: Authorization",     // token 检索模式，用于提取 token
		TokenHeadName:   "Bearer",                    // token 在请求头时的名称，默认值为 Bearer
		TimeFunc:        time.Now,
	})
}
