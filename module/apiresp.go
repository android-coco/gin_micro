/*
 * @Author: yhlyl
 * @Date: 2019-10-21 15:48:18
 * @LastEditTime: 2019-11-04 21:27:05
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/module/apiresp.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package module

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var UpGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { //允许跨域
		return true
	},
}

type ApiResp struct {
	ErrorNo  int64       `json:"errno"`
	ErrorMsg string      `json:"errmsg"`
	Data     interface{} `json:"data,omitempty"`
}
