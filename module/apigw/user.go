/*
 * @Author: yhlyl
 * @Date: 2019-11-03 23:18:20
 * @LastEditTime: 2019-11-04 18:15:09
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/module/apigw/user.go
 * @https://github.com/android-coco/gin_micro
 */
package apigw

import (
	"context"
	"gin_micro/httpserver/servermiddleware"
	"gin_micro/module"
	"gin_micro/module/selector"
	userProto "gin_micro/module/user/proto"
	"gin_micro/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
)

func init() {
	cmd.Init()
	client.DefaultClient = client.NewClient(
		//自定义选择器
		client.Selector(selector.FirstNodeSelector()),
	)
}

// GetHostHandler : 获取服务器列表
func GetHostHandler(c *gin.Context) {
	var baseReq servermiddleware.BaseReq
	err := c.ShouldBindJSON(&baseReq)
	if err != nil {
		util.Logger.Errorf("GetHostHandler 接口  参数绑定 出错 err: %s ", err.Error())
		c.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.ErrorLackCode, ErrorMsg: err.Error()})
		return
	}
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest(util.GinMicroUser, "User.Host", &userProto.ReqClientHost{
		AppName:    baseReq.AppName,
		AppVersion: baseReq.AppVersion,
	})

	resp := &userProto.RespClientHost{}

	// Call service
	err = client.Call(context.TODO(), req, resp)

	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, module.ApiResp{
		ErrorNo:  int64(resp.Code),
		ErrorMsg: resp.Message,
		Data:     resp.Host,
	})
}
