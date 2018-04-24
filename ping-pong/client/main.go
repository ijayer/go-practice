/*
 * 说明：
 * 作者：zhe
 * 时间：2018-04-18 16:34
 * 更新：
 */

package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", "10.0.0.120:8083", "server address [ip:port]")
	flag.Parse()
	initWsUrl()
}

var wsUrl string

func initWsUrl() {
	wsUrl = fmt.Sprintf("ws://%v/ws", addr)
}

func main() {
	fmt.Printf("client is connecting to: [%v]\n", wsUrl)

	// 连接服务器
	dialer := websocket.Dialer{}
	dialer.HandshakeTimeout = 60 * time.Second
	wsConn, _, err := dialer.Dial(wsUrl, nil)
	if err != nil {
		fmt.Println("Error: websocket dial failed.", err)
		return
	}
	fmt.Println("connected")

	// 初始化 Connection, 开启读写 G 程
	conn := NewConnection(wsConn)
	// go conn.Ping()
	go conn.Writer()
	go conn.Reader()

	for {
		time.Sleep(2 * time.Second)
		conn.Send <- []byte("biu°biu°")
	}
	select {}
}
