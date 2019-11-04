/*
 * @Author: yhlyl
 * @Date: 2019-11-03 11:13:38
 * @LastEditTime: 2019-11-04 21:27:22
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/tcp/test/client.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package main

import (
	"encoding/json"
	"fmt"
	"gin_micro/httpserver/wss/proto"
	"gin_micro/util"
	"io"
	"log"
	"net"
	"time"

	protoutil "github.com/gogo/protobuf/proto"
)

const (
	addr = "127.0.0.1:13001"
)

func main() {
	tcpClient()
}
func tcpClient() {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("连接服务端失败:", err.Error())
		return
	}
	fmt.Println("已连接服务器")
	defer conn.Close()
	go sender(conn)
	read(conn)

}

func sender(conn net.Conn) {
	for {
		time.Sleep(3 * time.Second)
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"

		// 发送订阅信息
		data := []byte(words)
		//心跳
		msg := &proto.Msg{
			ErrorNo:  util.SuccessCode,
			ErrorMsg: "测试",
			Data:     data,
		}
		dataMsg, _ := protoutil.Marshal(msg)
		n, err := conn.Write(util.CreateCmd(0x01, dataMsg, 100))
		if err != nil {
			log.Println("write:", err)
			return
		}
		log.Println("写数据:", util.CreateCmd(0x01, dataMsg, 100), n)
	}
}

func read(conn net.Conn) {
	for {
		var message = make([]byte, 1024)
		n, err := conn.Read(message)
		if err != nil && err != io.EOF || len(message) == 0 {
			log.Println("read:", err)
			return
		}
		message = message[:n]
		//校验
		if !util.IsCheckCmd(message, 100) {
			log.Println("客户端非法包", message)
			continue
		}
		var msg proto.Msg
		err = protoutil.Unmarshal(message[8:len(message)-2], &msg)
		bytes, _ := json.Marshal(msg)
		log.Print("recvString:\n", string(msg.Data), string(bytes))
	}
}
