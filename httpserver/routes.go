package httpserver

import (
	"gin_micro/httpserver/handle"
	"gin_micro/httpserver/servermiddleware"
	"gin_micro/httpserver/user"
	"gin_micro/httpserver/wss"
	"gin_micro/module/apigw"
	"github.com/gin-gonic/gin"
)

func initRoutes(router *gin.Engine) {
	router.GET("/version", handle.Version)
	router.GET("/ping", handle.Ping)

	//websocket,
	router.GET("/v1/wss", wss.Wss)
	router.GET("/v1/client", wss.Client)
	v1 := router.Group("/v1", servermiddleware.Base())
	{
		//活动
		v1.POST("/registered", user.Registered)
	}

	//获取服务器列表
	router.POST("v1/public/host", apigw.GetHostHandler)
}
