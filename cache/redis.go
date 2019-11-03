package redis

import (
	"gin_micro/config"
	"github.com/go-redis/redis"
)

var db *redis.Client

func InitRedis(redisConfig config.Redis) {
	db = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.PassWd,
		DB:       redisConfig.Db,
	})
}

// GetRedisDb 获取Redis连接实例
func GetRedisDb() *redis.Client {
	return db
}

