package redis

import (
	"gin_micro/module/config"
	"github.com/go-redis/redis"
)

var db *redis.Client

func init() {
	configRedis := config.GetRedis()
	db = redis.NewClient(&redis.Options{
		Addr:     configRedis.Host,
		Password: configRedis.PassWd,
		DB:       configRedis.Db,
	})
}

// GetRedisDb 获取Redis连接实例
func GetRedisDb() *redis.Client {
	return db
}
