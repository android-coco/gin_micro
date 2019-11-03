package wss

import (
	"encoding/json"
	"gin_micro/config"
	"gin_micro/httpserver/wss/proto"
	"gin_micro/model"
	"gin_micro/socket"
	"gin_micro/util"
	"gin_micro/util/jwt"
	"github.com/gin-gonic/gin"
	protoutil "github.com/gogo/protobuf/proto"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Wss(c *gin.Context) {
	conn, err := socket.InitConn(c.Writer, c.Request, nil)
	var token, uid string
	token = c.Request.FormValue("token")
	uid = c.Request.FormValue("uid")
	if token == "" {
		token = c.GetHeader("token")
	}
	if uid == "" {
		uid = c.GetHeader("uid")
	}
	util.Logger.Infof("[QP WEB SERVER] ws parameter. token is: %s", token)
	if err != nil {
		util.Logger.Error("[QP WEB SERVER] upgrade:", err)
		return
	}
	if token == "" || uid == "" {
		msg := &proto.Msg{
			ErrorNo:  http.StatusForbidden,
			ErrorMsg: http.StatusText(http.StatusForbidden),
		}
		data, _ := protoutil.Marshal(msg)
		conn.WriteMessage(0x00, data, 100)
		conn.Close()
		return
	}
	// 验证token和uid
	b, s, e := jwt.EasyToken{}.ValidateToken(token)
	if !b || s != uid || e != nil {
		msg := &proto.Msg{
			ErrorNo:  http.StatusForbidden,
			ErrorMsg: http.StatusText(http.StatusForbidden),
		}
		data, _ := protoutil.Marshal(msg)
		conn.WriteMessage(0x00, data, 100)
		conn.Close()
		return
	}

	// TODO 测试暂时屏蔽 一个token 1 个连接
	//var m sync.Mutex
	//m.Lock()
	var ip = strings.Split(c.ClientIP(), ":")[0]
	//var isRepeatConnection = false
	//model.ClientList.Range(func(key, value interface{}) bool {
	//	tempClient := value.(*model.Client)
	//	if tempClient.Token == token && tempClient.IsOnLine {
	//		conn.WriteMessage(module.ApiResp{
	//			ErrorNo:  http.StatusNotAcceptable,
	//			ErrorMsg: "No repeat connection."})
	//		conn.Close()
	//		isRepeatConnection = true
	//		return false
	//	}
	//	return true
	//})
	//m.Unlock()
	//if isRepeatConnection {
	//	return
	//}
	uidInt, _ := strconv.Atoi(uid)

	// 当前客户端
	var currentClient *model.Client
	currentClient = &model.Client{
		Token:       token,
		Uid:         int32(uidInt),
		Ip:          ip,
		Conn:        conn,
		MessageChan: make(chan interface{}, 0),
		IsOnLine:    true,
		OnLineTime:  time.Now().Format("2006-01-02T15:04:05.000"),
		OfflineTime: "",
	}
	//加入map
	model.ClientList.Store(token, currentClient)
	go readMsg(currentClient)
	go heartBeat(currentClient)
}

// 读客户端消息
func readMsg(currentClient *model.Client) {
	defer currentClient.ReleaseClient()
	for {
		var (
			data []byte
			err  error
		)
		data, err = currentClient.Conn.ReadMessage()
		// TODO  客服端数据 要处理
		//fmt.Println(len(data),string(data), err)
		if err != nil {
			// 连接错误
			util.Logger.Errorf("读取错误：err:%v", err)
			break
		}
		cmd := util.IsCheckCmd(data, currentClient.Uid)
		if !cmd {
			util.Logger.Errorf("数据包错误：err:%v", err)
			//TODO  返回错误到客户端
			return
		}

		var msg proto.Msg
		err = protoutil.Unmarshal(data[8:len(data)-2], &msg)
		dataCallBack, _ := protoutil.Marshal(&msg)
		log.Print("服务端recvMsg:\n", msg, string(msg.Data))
		err = currentClient.Conn.WriteMessage(0x01, dataCallBack, currentClient.Uid)
		if err != nil {
			util.Logger.Errorf("回复错误：%v", err)
			// 回复错误
			break
		}
	}
}

// 心跳包
func heartBeat(currentClient *model.Client) {
	ticker := time.NewTicker(time.Duration(config.GetWss().HeartbeatTime) * time.Second)
	defer ticker.Stop()
	defer currentClient.ReleaseClient()
	for {
		select {
		case <-ticker.C:

			heart := model.Heartbeat{
				TimeStr: time.Now().Format("2006-01-02 15:04:05"),
			}
			dataByte, _ := json.Marshal(heart)
			//心跳
			msg := &proto.Msg{
				ErrorNo:  util.SuccessCode,
				ErrorMsg: "",
				Data:     dataByte,
			}
			data, _ := protoutil.Marshal(msg)
			//heartBeat := module.ApiResp{
			//	ErrorNo:  util.SuccessCode,
			//	ErrorMsg: "",
			//	Data: model.Heartbeat{
			//		TimeStr: time.Now().Format("2006-01-02 15:04:05"),
			//	}}
			//bytes, _ := json.Marshal(heartBeat)
			err := currentClient.Conn.WriteMessage(0x00, data, currentClient.Uid)
			if err != nil {
				// 某个客户端异常退出
				return
			}
		}
	}
}
