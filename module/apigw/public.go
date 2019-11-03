package apigw

import (
	"context"
	"gin_micro/module"
	publicProto "gin_micro/module/public/proto"
	"gin_micro/module/selector"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/config/cmd"
	"log"
	"net/http"
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
	appVersion := c.Request.FormValue("app_version")
	appName := c.Request.FormValue("app_name")
	// Create new request to service go.micro.srv.example, method Example.Call
	req := client.NewRequest("qp_web_server.service.public", "Public.GetHost", &publicProto.ReqHost{
		AppName:    appName,
		AppVersion: appVersion,
	})

	resp := &publicProto.RespHost{}

	// Call service
	err := client.Call(context.TODO(), req, resp)

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
