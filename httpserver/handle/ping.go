/*
 * @Author: yhlyl
 * @Date: 2019-11-03 11:00:07
 * @LastEditTime: 2019-11-04 21:17:11
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/httpserver/handle/ping.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package handle

import (
	"gin_micro/module"
	"gin_micro/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.SuccessCode, ErrorMsg: "pong"})
}
