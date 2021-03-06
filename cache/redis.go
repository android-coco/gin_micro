/*
 * @Author: yhlyl
 * @Date: 2019-11-03 11:56:26
 * @LastEditTime: 2019-11-06 15:28:31
 * @LastEditors: yhlyl
 * @Description: In User Settings Edit
 * @FilePath: /gin_micro/cache/redis.go
 */
package redis

import (
	"gin_micro/config"

	"github.com/go-redis/redis/v7"
)

var db *redis.Client

// InitRedis 初始化 redis
func InitRedis(redisConfig config.Redis) error {
	db = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host,
		Password: redisConfig.PassWd,
		DB:       redisConfig.Db,
	})
	return db.Ping().Err()
}

// GetRedisDb 获取Redis连接实例
func GetRedisDb() *redis.Client {
	return db
}
