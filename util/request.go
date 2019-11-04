/*
 * @Author: yhlyl
 * @Date: 2019-11-03 11:18:04
 * @LastEditTime: 2019-11-04 21:28:19
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/util/request.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package util

import (
	"gin_micro/module"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

var requestAgent = gorequest.New()

func GetGoRequestAgent() *gorequest.SuperAgent {
	return requestAgent
}

func RequestErrRsp(resp gorequest.Response, body string, errors []error, c *gin.Context) {
	if resp != nil && resp.StatusCode != http.StatusOK {
		c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{
			ErrorNo:  int64(resp.StatusCode),
			ErrorMsg: http.StatusText(resp.StatusCode),
		})
		return
	}
	if errors != nil {
		c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{
			ErrorNo:  http.StatusInternalServerError,
			ErrorMsg: http.StatusText(http.StatusInternalServerError),
		})
		return
	}
}

//gin get or post param
func GetOrPost(c *gin.Context, key string) string {
	get := c.Query(key)
	if get != "" {
		return get
	} else {
		post := c.PostForm(key)
		if post != "" {
			return post
		}
	}
	return ""
}
