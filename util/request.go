package util

import (
	"gin_micro/module"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	"net/http"
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
