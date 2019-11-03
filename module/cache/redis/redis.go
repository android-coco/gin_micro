package redis

import (
	configWeb "gin_micro/config"
	"github.com/go-redis/redis"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/consul"
	"log"
)

var db *redis.ClusterClient

func init() {

	consulSource := consul.NewSource(
		// optionally specify etcd address; default to localhost:8500
		consul.WithAddress("localhost:8500"),
		// optionally specify prefix; defaults to /micro/config
		consul.WithPrefix("/my/redis"),
		// optionally strip the provided prefix from the keys, defaults to false
		consul.StripPrefix(true),
		source.WithEncoder(json.NewEncoder()),
	)

	// Create new config
	conf := config.NewConfig()

	// Load file source
	err := conf.Load(consulSource)

	if err != nil {
		log.Fatalf("微服务配置加载失败:%v", err)
	}

	var configRedis configWeb.Redis

	err = conf.Get("user").Scan(&configRedis)
	if err != nil {
		log.Fatalf("微服务启动失败:%v  %s", err, "读取配置文件")
	}
	go func() {
		for {
			//观测目录的变化。当文件有改动时，新值便可生效。
			w, err := conf.Watch("user")
			if err != nil {
				// do something
				log.Fatalf("微服务配置watch失败:%v", err)
			}
			// wait for next value
			v, err := w.Next()
			if err != nil {
				// do something
				log.Fatalf("微服务配置watch失败:%v", err)
			}
			err = v.Scan(&configRedis)
			if err != nil {
				// do something
				log.Fatalf("微服务配置watch失败:%v", err)
			}
		}
	}()
	db = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    configRedis.Hosts,
		Password: configRedis.PassWd,
	})
}

// GetRedisDb 获取Redis连接实例
func GetRedisDb() *redis.ClusterClient {
	return db
}
