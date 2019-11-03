package module

import (
	"github.com/gorilla/websocket"
	"net/http"
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
