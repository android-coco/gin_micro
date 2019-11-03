package handle

import (
	"gin_micro/module"
	"gin_micro/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.SuccessCode, ErrorMsg: "pong",})
}
