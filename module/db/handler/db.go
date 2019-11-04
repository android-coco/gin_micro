/*
 * @Author: yhlyl
 * @Date: 2019-11-03 23:21:33
 * @LastEditTime: 2019-11-04 21:24:00
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/module/db/handler/db.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package handler

import (
	"context"
	"gin_micro/model"
	"gin_micro/module/db/db"
	"gin_micro/module/db/proto"
	"gin_micro/util"
)

// Db : 用于实现DbServiceHandler接口的对象
type DbService struct{}

func (dbs *DbService) GetHost(ctx context.Context, req *proto.ReqHost, res *proto.RespHost) error {

	hosts, err := model.GetHost(db.GetDB())
	if err != nil {
		res.Code = 1
		res.Message = err.Error()
		return nil
	}

	var respHosts []*proto.Host

	for _, host := range hosts {
		respHosts = append(respHosts, &proto.Host{Id: host.Id, HostName: host.Name, Ip: host.Ip, Port: host.Port})
	}

	res.Code = util.SuccessCode
	res.Message = "获取成功"
	res.Host = respHosts

	return nil
}
