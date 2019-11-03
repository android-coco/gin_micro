package main

import (
	"gin_micro/module/public/handler"
	"gin_micro/module/public/proto"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/transport/rabbitmq"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name("qp_web_server.service.public"),
		micro.Version("latest"),
	)

	// 初始化service, 解析命令行参数等
	service.Init()
	err := proto.RegisterPublicHandler(service.Server(), new(handler.Public))
	if err != nil {
		log.Fatalf("微服务注册失败:%v", err)
	}
	if err := service.Run(); err != nil {
		log.Fatalf("微服务启动失败:%v", err)
	}
}
