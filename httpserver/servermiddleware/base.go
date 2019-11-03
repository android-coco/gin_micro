package servermiddleware

import (
	"bytes"
	"encoding/base64"
	"gin_micro/module"
	"gin_micro/util"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

type BaseReq struct {
	AppVersion        string `form:"app_version" json:"app_version" binding:"required"`
	AppName           string `form:"app_name" json:"app_name" binding:"required"`
	NumRegisterOrigin int    `form:"num_register_origin" json:"num_register_origin" binding:"required"`
	MachineID         string `form:"machine_id" json:"machine_id" binding:"required"`
}

//加密验证
func Base() gin.HandlerFunc {
	return func(c *gin.Context) {
		sign := c.GetHeader("sign")

		if sign == "" {
			//权限异常
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{
				ErrorNo:  http.StatusForbidden,
				ErrorMsg: http.StatusText(http.StatusForbidden),
			})
			return
		}

		req, err := c.GetRawData()
		//c.Set("req", string(req))
		//传递参数到下个中间件
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(req)) // 关键点
		if len(req) == 0 || err != nil {
			//token
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{
				ErrorNo:  http.StatusForbidden,
				ErrorMsg: "the pam is err!",
			})
			return
		}
		// 验证签名
		if sign != util.VerifySign(base64.StdEncoding.EncodeToString(req), util.DesKey) {
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{
				ErrorNo:  http.StatusForbidden,
				ErrorMsg: "the sign is err!",
			})
			return
		}
	}

}
