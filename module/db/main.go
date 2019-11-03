package main

import (
	"fmt"
	configWeb "gin_micro/config"
	"gin_micro/module/db/db"
	"gin_micro/module/db/handler"
	"gin_micro/module/db/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/source"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	"github.com/micro/go-plugins/config/source/consul"
	_ "github.com/micro/go-plugins/transport/rabbitmq"
	"log"
)

func main() {

	consulSource := consul.NewSource(
		// optionally specify etcd address; default to localhost:8500
		consul.WithAddress("localhost:8500"),
		// optionally specify prefix; defaults to /micro/config
		consul.WithPrefix("/my/db"),
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

	var configDb configWeb.Db

	err = conf.Get("user").Scan(&configDb)
	// consul kv get my/prefix/host
	fmt.Println(configDb, err)
	if err != nil {
		log.Fatalf("微服务启动失败:%v  %s", err, "读取配置文件")
	}

	fmt.Println(configDb)

	_, err = db.InitDB(configDb)
	if err != nil {
		log.Fatalf("微服务启动失败:%v  %s", err, "数据库链接")
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
			err = v.Scan(&configDb)
			if err != nil {
				// do something
				log.Fatalf("微服务配置watch失败:%v", err)
			}
			fmt.Println(configDb)
		}
	}()
	service := micro.NewService(
		micro.Name("qp_web_server.service.db"),
		micro.Version("latest"),
	)

	// 初始化service, 解析命令行参数等
	service.Init()
	err = proto.RegisterDbHandler(service.Server(), new(handler.DbService))
	if err != nil {
		log.Fatalf("微服务注册失败:%v", err)
	}
	if err := service.Run(); err != nil {
		log.Fatalf("微服务启动失败:%v", err)
	}
}
