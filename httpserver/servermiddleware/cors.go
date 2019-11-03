package servermiddleware

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func EnableCors(allowedDomain []string) gin.HandlerFunc {
	return cors.New(initCorsConf(allowedDomain))
}

func initCorsConf(allowedDomain []string) cors.Config {
	config := cors.DefaultConfig()
	if len(allowedDomain) == 0 {
		return config
	}

	allowedOrigins := make([]string, 0, len(allowedDomain))
	originPrefixList := make([]string, 0, len(allowedDomain)*2)
	if allowedDomain[0] == "*" {
		allowedOrigins = []string{"*"}
	} else {
		for _, domain := range allowedDomain {
			allowedOrigins = append(allowedOrigins, "https://"+domain)
			originPrefixList = append(originPrefixList, "."+domain, "//"+domain)
		}
	}
	config.AllowOrigins = allowedOrigins
	config.AllowOriginFunc = func(origin string) bool {
		for _, originPrefix := range originPrefixList {
			if strings.HasSuffix(origin, originPrefix) {
				return true
			}
		}
		return false
	}
	return config
}
