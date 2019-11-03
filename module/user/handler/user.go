package handler

import (
	"context"
	"encoding/json"
	"gin_micro/module/cache/redis"
	dbProto "gin_micro/module/db/proto"
	"gin_micro/module/user/proto"
	"gin_micro/util"
	"github.com/micro/go-micro"
	"time"
)

type User struct{}

var (
	dbCli dbProto.DbService
)

func init() {
	service := micro.NewService()

	// 初始化， 解析命令行参数等
	service.Init()
	cli := service.Client()
	// 初始化一个account服务的客户端
	dbCli = dbProto.NewDbService(util.GinMicroBb, cli)
}

// GetHost : 获取服务器列表
func (u *User) Host(ctx context.Context, req *proto.ReqClientHost, res *proto.RespClientHost) error {

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
	var data []*proto.ClientHost
	for _, host := range respHost.Host {
		data = append(data, &proto.ClientHost{Id: host.Id,
			HostName: host.HostName,
			Ip:       host.Ip,
			Port:     host.Port})
	}
	bytes, _ := json.Marshal(data)
	redis.GetRedisDb().Set("host",string(bytes),10 * time.Minute)
	res.Code = 0
	res.Message = "获取成功"
	res.Host = data
	return nil
}
