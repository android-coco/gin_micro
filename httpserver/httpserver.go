/*
 * @Author: yhlyl
 * @Date: 2019-11-03 11:03:09
 * @LastEditTime: 2019-11-06 15:33:53
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/httpserver/httpserver.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package httpserver

import (
	"gin_micro/config"
	"gin_micro/httpserver/servermiddleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(config.GetService().Mode)
	gin.ForceConsoleColor()
	router := gin.Default()
	//各种中间件
	router.Use(gin.Recovery())
	router.Use(gin.ErrorLogger())
	router.Use(servermiddleware.BaseLogger())
	router.Use(servermiddleware.EnableCors([]string{"*"}))
	initRoutes(router)
	return router
}
