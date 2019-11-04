/*
 * @Author: yhlyl
 * @Date: 2019-11-03 11:00:22
 * @LastEditTime: 2019-11-04 21:18:04
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/httpserver/servermiddleware/auth.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package servermiddleware

import (
	"bytes"
	"encoding/json"
	redis "gin_micro/cache"
	"gin_micro/module"
	"gin_micro/util"
	"gin_micro/util/jwt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseAuthReq struct {
	BaseReq
	Token string `form:"token" json:"token"  binding:"required"`
	Uid   string `form:"uid" json:"uid"  binding:"required"`
}

//token验证
func BaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		var authReq BaseAuthReq
		//err := c.ShouldBindJSON(&authReq)

		req, err := c.GetRawData()
		if err != nil {
			util.Logger.Errorf("BaseAuth  参数绑定 出错 err: %s ", err.Error())
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{ErrorNo: 8, ErrorMsg: err.Error()})
			return
		}
		//传递参数到下个中间件
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(req)) // 关键点

		err = json.Unmarshal(req, &authReq)

		if err != nil {
			util.Logger.Errorf("BaseAuth  参数绑定 出错 err: %s ", err.Error())
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{ErrorNo: 8, ErrorMsg: err.Error()})
			return
		}
		token := authReq.Token
		uid := authReq.Uid

		//token := c.GetHeader("token")
		//uid := c.GetHeader("uid")
		if token == "" || uid == "" {
			//权限异常
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{
				ErrorNo:  http.StatusForbidden,
				ErrorMsg: http.StatusText(http.StatusForbidden),
			})
			return
		}
		//TODO 验证token是否存在
		//1,查询redis
		tokenRedis := redis.GetRedisDb().Get(util.RedisKeyToken + uid).String()

		et := jwt.EasyToken{}
		valid, tokenUid, _ := et.ValidateToken(token)
		if !valid || tokenRedis != token || tokenUid != uid {
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{
				ErrorNo:  http.StatusForbidden,
				ErrorMsg: "登录失效,请重新登录！",
			})
			return
		}

		c.Set("uid", uid)
	}

}
