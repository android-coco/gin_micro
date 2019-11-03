package handle

import (
	"gin_micro/module"
	"gin_micro/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

const version = "gin_micro_v1.0.0"

func Version(context *gin.Context) {
	context.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.SuccessCode, ErrorMsg: "", Data: version})
}
