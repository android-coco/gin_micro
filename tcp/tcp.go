package tcp

import (
	"encoding/json"
	"fmt"
	"gin_micro/config"
	"gin_micro/httpserver/wss/proto"
	"gin_micro/model"
	"gin_micro/socket/tcp"
	"gin_micro/util"
	"gin_micro/util/uuid"
	protoutil "github.com/gogo/protobuf/proto"
	"log"
	"time"
)

func Run(address string) {
	err := socket.InitConn(address, func(conn *socket.Connection) {
		defer conn.Close()
		client(conn)
	})
	if err != nil {
		log.Fatalf("tcp err:%v", err)
	}
}

func client(conn *socket.Connection) {
	//TODO 服务器判断
	currentClient := &model.TCPClient{
		Id:          uuid.NewUUID().Hex32(),
		Uid:         100,
		Ip:          conn.GetIP(),
		Conn:        conn,
		MessageChan: make(chan interface{}, 0),
		IsOnLine:    true,
		OnLineTime:  time.Now().Format("2006-01-02T15:04:05.000"),
		OfflineTime: "",
	}
	fmt.Println(currentClient.Id)
	//加入map
	model.ClientList.Store(currentClient.Id, currentClient)
	go heartBeat(currentClient)
	readMsg(currentClient)
}

// 读客户端消息
func readMsg(currentClient *model.TCPClient) {
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
			return
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
func heartBeat(currentClient *model.TCPClient) {
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
			fmt.Println("写数据：", data)
			err := currentClient.Conn.WriteMessage(0x00, data, currentClient.Uid)
			if err != nil {
				fmt.Println(err)
				// 某个客户端异常退出
				return
			}
		}
	}
}
