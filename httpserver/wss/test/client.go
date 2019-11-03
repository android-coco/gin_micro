package main

import (
	"encoding/json"
	"gin_micro/httpserver/wss/proto"
	"gin_micro/util"
	protoutil "github.com/gogo/protobuf/proto"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"time"
)

var url = "ws://127.0.0.1:13000/v1/wss"

func main() {
	createClient()
}

func createClient() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial(url, map[string][]string{
		"token": {"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQzMDcyOTEsImlzcyI6IjEwMCIsIm5iZiI6MTU3MTcxNTI5MX0.lIpO9H9fBZaawuAYjKJgt8ImdALAchwmBCI3e2sxyQVlIauuXddMEhovmJqlxZEJ7djsvZGNEyUYitY8j6HXM47gDhnamSc2Sn9C6f4LhMwlP-V1MBeiDisKho61_Fbui0fR-hqKQNbaOrvI2aLaVA695oKG7hJ7EuWJMHYVM0SOCwPBqFvFU8A086kCgzXcvSSmKRvctaJrbMfyGgLswHkBaaMnMP0XTpyoUo_0UO1PxpZqFlryapXIGd4MYsdYpFqH-K_NBXUJxznXGGVwDh8AWz73UnlCIfCzhyC4sOfW_QWx0xOhHf-krdKXhJauop22EtnXud0HFiZJHVCO5g"},
		"uid":   {"100"},
	})

	if err != nil {
		log.Fatal("dial:", err)
	} else {
		log.Printf("connecting to %s successful,waiting for heartbeat packet\n", url)
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			//校验
			if !util.IsCheckCmd(message, 100) {
				log.Println("客户端非法包", message)
				continue
			}
			var msg proto.Msg
			err = protoutil.Unmarshal(message[8:len(message)-2], &msg)
			//log.Print("recv:\n", message)

			bytes, _ := json.Marshal(msg)
			log.Print("recvMsg:\n", msg)
			log.Print("recvMessage:\n", string(message[8:len(message)-2]))
			log.Print("recvString:\n", string(msg.Data), string(bytes))
			//log.Print("recv:\n", err)
		}
	}()

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	var i = 0
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			i++
			if i == 2 {
				// 发送订阅信息
				data := []byte("哈哈")
				//心跳
				msg := &proto.Msg{
					ErrorNo:  util.SuccessCode,
					ErrorMsg: "测试",
					Data:     data,
				}
				dataMsg, _ := protoutil.Marshal(msg)
				err := c.WriteMessage(websocket.BinaryMessage, util.CreateCmd(0x01, dataMsg, 100))
				log.Println("写数据:", util.CreateCmd(0x01, dataMsg, 100), err)
				if err != nil {
					log.Println("write:", err)
					return
				}
			}

		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "bay bay"))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
