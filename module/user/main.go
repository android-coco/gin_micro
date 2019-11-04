package main

import (
	"gin_micro/module/config"
	userHandler "gin_micro/module/user/handler"
	"gin_micro/module/user/proto"
	"gin_micro/util"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/transport/rabbitmq"
	"log"
)

func main() {
	newBroker := rabbitmq.NewBroker(
		broker.Addrs(config.GetPublic().RabbitmqUrl),
	)
	if err := newBroker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}
	if err := newBroker.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
	eTCDRegistry := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{config.ConfUrl}
	})
	service := micro.NewService(
		micro.Name(util.GinMicroUser),
		micro.Registry(eTCDRegistry),
		micro.Metadata(map[string]string{
			"type": "helloWorld",
		}),
		micro.Broker(newBroker),
		micro.Version("latest"),
	)
	// 初始化service, 解析命令行参数等
	service.Init()
	err := proto.RegisterUserHandler(service.Server(), new(userHandler.User))
	//发送消息
	go util.Pub(newBroker, util.UserQueue)

	if err != nil {
		log.Fatalf("微服务注册失败:%v", err)
	}
	if err := service.Run(); err != nil {
		log.Fatalf("微服务启动失败:%v", err)
	}
}
