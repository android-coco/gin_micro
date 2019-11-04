/*
 * @Author: yhlyl
 * @Date: 2019-11-03 11:05:07
 * @LastEditTime: 2019-11-04 21:20:54
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/httpserver/wss/home.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package wss

import (
	"gin_micro/model"
	"gin_micro/module"
	"gin_micro/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type clientListData struct {
	Total      int64        `json:"total"`
	ClientList []clientData `json:"client_list"`
}

type clientData struct {
	Type        string `json:"type,omitempty"`
	Token       string `json:"token,omitempty"`
	Id          string `json:"id,omitempty"`
	Uid         int32  `json:"uid,omitempty"`
	Ip          string `json:"ip,omitempty"`
	IsOnline    bool   `json:"is_online,omitempty"`
	OnlineTime  string `json:"online_time,omitempty"`
	OfflineTime string `json:"offline_time,omitempty"`
}

//统计用户数量
func Client(c *gin.Context) {
	var clientList []clientData
	model.ClientList.Range(func(key, value interface{}) bool {
		switch value.(type) {
		case *model.Client:
			var currentClient = value.(*model.Client)
			var client clientData
			//if !currentClient.IsOnLine {
			//	return false
			//}
			client = clientData{
				Type:        "WebSocket",
				Token:       currentClient.Token,
				Uid:         currentClient.Uid,
				Ip:          currentClient.Ip,
				IsOnline:    currentClient.IsOnLine,
				OnlineTime:  currentClient.OnLineTime,
				OfflineTime: currentClient.OfflineTime,
			}
			clientList = append(clientList, client)
		case *model.TCPClient:
			var currentClient = value.(*model.TCPClient)
			var client clientData
			//if !currentClient.IsOnLine {
			//	return false
			//}
			client = clientData{
				Type:        "TCP",
				Id:          currentClient.Id,
				Uid:         currentClient.Uid,
				Ip:          currentClient.Ip,
				IsOnline:    currentClient.IsOnLine,
				OnlineTime:  currentClient.OnLineTime,
				OfflineTime: currentClient.OfflineTime,
			}
			clientList = append(clientList, client)
		}

		return true
	})
	c.JSON(http.StatusOK, module.ApiResp{
		ErrorNo:  util.SuccessCode,
		ErrorMsg: "",
		Data:     clientListData{Total: int64(len(clientList)), ClientList: clientList},
	})
}
