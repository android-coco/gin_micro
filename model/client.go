package model

import (
	"fmt"
	"gin_micro/socket"
	tcp "gin_micro/socket/tcp"
	"gin_micro/util"
	"sync"
	"time"
)

// 记录客户Map    [apiKey,*model.Client]
var ClientList sync.Map

//客户端信息
type Client struct {
	Token       string //token
	Uid         int32  // uid
	Ip          string //客户端ip地址
	Conn        *socket.Connection
	MessageChan chan interface{}
	IsOnLine    bool   //是否在线
	OnLineTime  string //上线时间
	OfflineTime string //下线时间
}

func (c *Client) String() string {
	return fmt.Sprintf("token %s ip %s  IsOneLine  %t OnLineTime %s",
		c.Token, c.Ip, c.IsOnLine, c.OnLineTime)
}

func (c *Client) ReleaseClient() {
	c.IsOnLine = false
	c.Conn.Close()
	c.OfflineTime = time.Now().Format("2006-01-02T15:04:05.000")
	ClientList.Store(c.Token, c)
	util.Logger.Errorf("客户端token:%s,客户端ip:%s Bay Bay ", c.Token, c.Ip)
}

//客户端信息
type TCPClient struct {
	Id          string // 连接id
	Uid         int32  // uid
	Ip          string //客户端ip地址
	Conn        *tcp.Connection
	MessageChan chan interface{}
	IsOnLine    bool   //是否在线
	OnLineTime  string //上线时间
	OfflineTime string //下线时间
}

func (c *TCPClient) String() string {
	return fmt.Sprintf("Id %s ip %s  IsOneLine  %t OnLineTime %s",
		c.Id, c.Ip, c.IsOnLine, c.OnLineTime)
}

func (c *TCPClient) ReleaseClient() {
	c.IsOnLine = false
	c.Conn.Close()
	c.OfflineTime = time.Now().Format("2006-01-02T15:04:05.000")
	ClientList.Store(c.Id, c)
	util.Logger.Errorf("客户端Id:%s,客户端ip:%s Bay Bay ", c.Id, c.Ip)
}
