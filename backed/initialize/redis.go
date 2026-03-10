package initialize

import (
	"fmt"
	"myApp/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitRedis() {
	global.Redis = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			global.ServerConfig.RedisInfo.Host,
			global.ServerConfig.RedisInfo.Port),
		Password: global.ServerConfig.RedisInfo.Password,
		DB:       global.ServerConfig.RedisInfo.DB,
		PoolSize: global.ServerConfig.RedisInfo.PoolSize,
	})
	_, err := global.Redis.Ping().Result()
	zap.S().Infof("Redis connected: %v", err == nil)
}

func CloseRedis() {
	_ = global.Redis.Close()
}
