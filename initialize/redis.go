package initialize

import (
	"clockwerk/global"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-redis/redis"
)

func InitRedis() {
	addr := fmt.Sprintf("%s:%d", global.ServerSetting.RedisConfig.Host, global.ServerSetting.RedisConfig.Port)
	// 生成redis客户端
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// 链接redis
	_, err := global.Redis.Ping().Result()
	if err != nil {
		color.Red("[InitRedis] 链接redis异常:")
		color.Yellow(err.Error())
	}
}
