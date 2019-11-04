/*
 * @Author: yhlyl
 * @Date: 2019-11-04 00:30:50
 * @LastEditTime: 2019-11-04 21:22:13
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/module/cache/redis/redis.go
 * @Github: https://github.com/android-coco/gin_micro
 */
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
