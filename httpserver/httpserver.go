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
	router.Use(gin.Logger())
	router.Use(gin.ErrorLogger())
	router.Use(servermiddleware.BaseLogger())
	router.Use(servermiddleware.EnableCors([]string{"*"}))
	initRoutes(router)
	return router
}
