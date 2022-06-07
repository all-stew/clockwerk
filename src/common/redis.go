package common

/*
   说明：Redis Key 相关配置项
   推荐 Redis 连接工具：https://github.com/qishibo/AnotherRedisDesktopManager/
*/
type RedisKeyConfig struct {
	Tag                      string `json:"tag"`                      // 分隔符，使用工具连接时候能根据分隔符归类
	TokenKeyPrefix           string `json:"tokenKeyPrefix"`           // Token 前缀
	TokenExpiresKeyPrefix    string `json:"tokenExpiresKeyPrefix"`    // Token 超时时间前缀
	CurrentUserInfoKeyPrefix string `json:"currentUserInfoKeyPrefix"` // 用户信息前缀
}

var RedisKeyPrefix = RedisKeyConfig{
	Tag:                      ":",
	TokenKeyPrefix:           "Token",
	TokenExpiresKeyPrefix:    "TokenExpires",
	CurrentUserInfoKeyPrefix: "CurrentUserInfo",
}

// RedisKeyExpiresTimeConfig Redis Key 超时时间配置
type RedisKeyExpiresTimeConfig struct {
	CurrentUserInfoKeyExpiresTime int `json:"userExpiresTime"` // 用户信息保存超时时间，单位秒
}

var RedisKeyExpiresTime = RedisKeyExpiresTimeConfig{
	CurrentUserInfoKeyExpiresTime: 30,
}
