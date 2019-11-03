package config

import (
	"fmt"
	"gin_micro/util"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	"github.com/micro/go-micro/config/source/etcd"
	"log"
)

const (
	//配置中心地址
	ConfUrl = "127.0.0.1:2379"

	//公用配置路径
	PublicPath = "public"
	//mysql
	MySQLPath = "mysql"
	//Redis
	RedisPath = "redis"
	//配置文件名称
	ConfName = "config"

	//配置文件版本
	Version = "version"
)

func init() {
	eTCDSource := etcd.NewSource(
		// optionally specify etcd address; default to localhost:8500
		etcd.WithAddress(ConfUrl),
		// optionally specify prefix; defaults to /micro/config
		etcd.WithPrefix(util.ConfigRoot),
		// optionally strip the provided prefix from the keys, defaults to false
		etcd.StripPrefix(true),
		source.WithEncoder(json.NewEncoder()),
	)

	// Create new config
	conf := config.NewConfig()

	// Load file source
	err := conf.Load(eTCDSource)
	if err != nil {
		log.Fatalf("微服务配置加载失败:%v", err)
	}
	InitConfig(conf)
	go func() {
		for {
			//观测目录的变化。当文件有改动时，新值便可生效。监听配置文件版本变化
			w, err := conf.Watch(Version)
			fmt.Println("======init config ======",w, err)
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
			fmt.Println(string(v.Bytes()))
			//重新配置
			InitConfig(conf)
		}
	}()

}

func InitConfig(conf config.Config) {
	//public
	if err := conf.Get(PublicPath, ConfName).Scan(&configSrv.Public); err != nil {
		log.Fatalf("微服务配置加载失败:%v", err)
	}
	//mysql
	if err := conf.Get(MySQLPath, ConfName).Scan(&configSrv.DB); err != nil {
		log.Fatalf("微服务配置加载失败:%v", err)
	}

	//redis
	if err := conf.Get(RedisPath, ConfName).Scan(&configSrv.Redis); err != nil {
		log.Fatalf("微服务配置加载失败:%v", err)
	}
}

var configSrv Config

// Config 配置
type Config struct {
	DB     Db        `json:"db"`
	Redis  Redis     `json:"redis"`
	Public Public    `json:"public"`
}

// 公共配置
type Public struct {
	RabbitmqUrl string `json:"rabbitmq_url"`
}

// DB 数据库配置
type Db struct {
	EnableLog          bool   `json:"enable_log"`
	Dialect            string `json:"dialect"`
	Host               string `json:"host"`
	User               string `json:"user"`
	PassWd             string `json:"pass" json:"pass"`
	Db                 string `json:"db"`
	MaxOpenConnections int    `json:"max_open_connections"`
	MaxIdleConnections int    `json:"max_idle_connections"`
}

// redis
type Redis struct {
	Host   string `json:"host"`
	PassWd string `json:"pass_wd"`
	Db     int    `json:"db"`
}


func GetDb() Db {
	return configSrv.DB
}

func GetRedis() Redis {
	return configSrv.Redis
}

func GetPublic() Public {
	return configSrv.Public
}

