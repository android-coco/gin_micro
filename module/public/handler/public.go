package handler

import (
	"context"
	dbProto "gin_micro/module/db/proto"
	"gin_micro/module/public/proto"
	"gin_micro/util"
	"github.com/micro/go-micro"
)

type Public struct{}

var (
	dbCli dbProto.DbService
)

func init() {
	service := micro.NewService()

	// 初始化， 解析命令行参数等
	service.Init()
	cli := service.Client()
	// 初始化一个account服务的客户端
	dbCli = dbProto.NewDbService("qp_web_server.service.db", cli)
}

// GetHost : 获取服务器列表
func (u *Public) GetHost(ctx context.Context, req *proto.ReqHost, res *proto.RespHost) error {

	appVersion := req.AppVersion
	appName := req.AppName
	// 参数简单校验
	if appName == "" || appVersion == "" {
		res.Code = 1
		res.Message = "参数无效"
		return nil
	}
	respHost, err := dbCli.GetHost(ctx, &dbProto.ReqHost{HostName: appName})
	if err != nil {
		return err
	}
	if respHost.Code != 0 {
		res.Code = util.ErrorSqlCode
		res.Message = respHost.Message
		return nil
	}
	var data []*proto.Host
	for _, host := range respHost.Host {
		data = append(data, &proto.Host{Id: host.Id,
			HostName: host.HostName,
			Ip:       host.Ip,
			Port:     host.Port})
	}
	//data = append(data, &proto.Host{Id: 1, HostName: "webGw", Ip: "192.168.1.5", Port: "13000"})
	//data = append(data, &proto.Host{Id: 2, HostName: "webGw", Ip: "192.168.1.5", Port: "13000"})
	res.Code = 0
	res.Message = "获取成功"
	res.Host = data
	return nil
}
