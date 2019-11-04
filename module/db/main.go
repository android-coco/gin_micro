/*
 * @Author: yhlyl
 * @Date: 2019-11-04 00:34:10
 * @LastEditTime: 2019-11-04 21:26:26
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/module/db/main.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package main

import (
	"gin_micro/module/config"
	"gin_micro/module/db/db"
	"gin_micro/module/db/handler"
	"gin_micro/module/db/proto"
	"gin_micro/util"
	"log"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	regEtcd "github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/transport/rabbitmq"
)

func main() {
	initDB, err := db.InitDB(config.GetDb())
	if err != nil {
		log.Fatalf("微服务启动失败:%v  %s", err, "数据库链接")
	}
	defer func() {
		err := initDB.Close()
		if err != nil {
			log.Fatalf("db close err %v", err)
		}
	}()

	newBroker := rabbitmq.NewBroker(
		broker.Addrs(config.GetPublic().RabbitmqUrl),
	)
	if err := newBroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := newBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	eTCDRegistry := regEtcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{config.ConfUrl}
	})
	service := micro.NewService(
		micro.Name(util.GinMicroBb),
		micro.Registry(eTCDRegistry),
		micro.Broker(newBroker),
		micro.Version("latest"),
	)
	// 初始化service, 解析命令行参数等
	service.Init()
	err = proto.RegisterDbHandler(service.Server(), new(handler.DbService))
	if err != nil {
		log.Fatalf("微服务注册失败:%v", err)
	}
	//订阅消息
	go util.Sub(newBroker, util.UserQueue)
	if err := service.Run(); err != nil {
		log.Fatalf("微服务启动失败:%v", err)
	}
}
