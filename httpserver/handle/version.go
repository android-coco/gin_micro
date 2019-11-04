/*
 * @Author: yhlyl
 * @Date: 2019-11-03 10:59:57
 * @LastEditTime: 2019-11-04 21:17:53
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/httpserver/handle/version.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package handle

import (
	"gin_micro/module"
	"gin_micro/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const version = "gin_micro_v1.0.0"

func Version(context *gin.Context) {
	context.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.SuccessCode, ErrorMsg: "", Data: version})
}
