/*
 * @Author: yhlyl
 * @Date: 2019-11-03 11:12:02
 * @LastEditTime: 2019-11-04 21:27:11
 * @LastEditors: yhlyl
 * @Description:
 * @FilePath: /gin_micro/socket/connection.go
 * @Github: https://github.com/android-coco/gin_micro
 */
package socket

import (
	"errors"
	"gin_micro/util"
	jasonlog "gin_micro/util/log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// 写超时时间
var writeWait = 10 * time.Second

var upGrader = websocket.Upgrader{
	HandshakeTimeout: 5 * time.Second,
	ReadBufferSize:   4096,
	WriteBufferSize:  4096,
	CheckOrigin: func(r *http.Request) bool { //允许跨域
		return true
	},
	EnableCompression: true, //压缩
}

type Connection struct {
	wsConn    *websocket.Conn
	inChan    chan []byte
	outChan   chan *Msg
	closeChan chan struct{}

	mutex    sync.Mutex
	isClosed bool
}

//const (
//	// TextMessage denotes a text data message. The text message payload is
//	// interpreted as UTF-8 encoded text data.
//	TextMessage = 1
//
//	// BinaryMessage denotes a binary data message.
//	BinaryMessage = 2
//
//	// CloseMessage denotes a close control message. The optional message
//	// payload contains a numeric code and text. Use the FormatCloseMessage
//	// function to format a close message payload.
//	CloseMessage = 8
//
//	// PingMessage denotes a ping control message. The optional message payload
//	// is UTF-8 encoded text.
//	PingMessage = 9
//
//	// PongMessage denotes a pong control message. The optional message payload
//	// is UTF-8 encoded text.
//	PongMessage = 10
//)

type Msg struct {
	MessageType int    //消息类型
	Data        []byte //数据
}

func InitConn(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*Connection, error) {
	var conn *Connection
	wsConn, err := upGrader.Upgrade(w, r, responseHeader)
	if err == nil {
		conn = &Connection{
			wsConn:    wsConn,
			inChan:    make(chan []byte, 0),
			outChan:   make(chan *Msg, 0),
			closeChan: make(chan struct{}, 1),
		}
		// 启动读协程
		go conn.readLoop()
		// 启动写协程
		go conn.writeLoop()
	}
	return conn, err
}

// TODO 修改创建msg的方法
func (conn *Connection) createMsg(cmd byte, data []byte, uid int32) (*Msg, error) {

	//TODO 根据uid查询相应的公私钥

	//1,组成命令包
	createCmd := util.CreateCmd(cmd, data, uid)
	// 2,注册msg
	msg := &Msg{
		MessageType: websocket.BinaryMessage,
		Data:        createCmd,
	}
	return msg, nil
}

func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) WriteMessage(cmd byte, data []byte, uid int32) (err error) {
	msg, err := conn.createMsg(cmd, data, uid)
	if err != nil {
		err = errors.New("msg create err" + err.Error())
	}
	select {
	case conn.outChan <- msg:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// 获取客户端ip
func (conn *Connection) GetIP() string {
	return conn.wsConn.RemoteAddr().String()
}

func (conn *Connection) Close() {

	time.Sleep(50 * time.Millisecond)
	//关闭chan 只能执行一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()

	// 线程安全,可以重复close
	conn.wsConn.Close()
}

// 读 协程
func (conn *Connection) readLoop() {
	for {
		if _, data, err := conn.wsConn.ReadMessage(); err != nil {
			goto ERROR
		} else {
			// 阻塞这里,等待inChan 空闲
			select {
			case conn.inChan <- data:
			case <-conn.closeChan:
				//closeChan关闭的时候
				goto ERROR
			}
		}
	}
ERROR:
	conn.Close()
}

//写  协程
func (conn *Connection) writeLoop() {
	for {
		select {
		case data := <-conn.outChan:
			//写超时时间
			conn.wsConn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.wsConn.WriteMessage(data.MessageType, data.Data); err != nil {
				jasonlog.Errorf("写失败：err:%v", err)
				goto ERROR
			}
		case <-conn.closeChan:
			goto ERROR
		}

	}
ERROR:
	conn.Close()
}
