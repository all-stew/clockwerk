package initialize

import (
	"clockwerk/src/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

/*
   说明：初始化 Redis 连接
*/
func Redis() {
	// 打印连接信息
	dsn := fmt.Sprintf("%s:%d", global.Conf.Redis.Host, global.Conf.Redis.Port)
	global.Log.Debug(fmt.Sprintf("打开连接（Redis）：tcp://%s/%d", dsn, global.Conf.Redis.DB))

	client := redis.NewClient(&redis.Options{
		// 连接信息配置
		Network:  "tcp",                      // 连接协议
		Addr:     dsn,                        // 连接地址
		Password: global.Conf.Redis.Password, // 密码
		DB:       global.Conf.Redis.DB,       // 数据库

		// 连接池配置
		PoolSize:     15, // 连接池最大连接数，默认4倍CPU数
		MinIdleConns: 10, // 在启动阶段创建 Idle 连接数量，长期维持 idle 状态的连接数不少于该数量

		// 超时配置
		DialTimeout:  5 * time.Second, // 连接建立超时时间
		ReadTimeout:  3 * time.Second, // 读超时，默认3秒
		WriteTimeout: 3 * time.Second, // 写超时，默认等于读超时
		PoolTimeout:  4 * time.Second, // 连接占满，客户端等待可用连接的最大等待时长，默认为读超时+1秒

		// 闲置连接检查配置
		IdleCheckFrequency: 60 * time.Second, // 闲置连接检查的周期，默认为1分钟
		IdleTimeout:        5 * time.Minute,  // 闲置超时，默认5分钟
		MaxConnAge:         0 * time.Second,  // 连接存活时长，默认为 0，即不关闭存活时长较长的连接

		// 执行失败时的重试策略
		MaxRetries:      0,                      // 执行失败最多重试多少次，默认为 0 不重试
		MinRetryBackoff: 8 * time.Millisecond,   // 每次计算重试间隔时间的下限，默认8毫秒
		MaxRetryBackoff: 512 * time.Millisecond, // 每次计算重试间隔时间的上限，默认512毫秒
	})

	// 测试连接 Redis
	ctx := context.Background()
	err := client.Ping(ctx).Err()
	if err != nil {
		panic("Redis初始化连接失败")
	} else {
		log.Println("Redis初始化完成")
		global.Redis = client
	}
}
