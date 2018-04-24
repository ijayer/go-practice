/*
 * 说明：
 * 作者：zhe
 * 时间：2018-04-18 16:05
 * 更新：
 */

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

// Websocket读写相关配置
const (
	WriteWait      = 5 * time.Second // 写数据超时时间
	PingWait       = 5 * time.Second // 读取Ping消息超时时间
	MaxMessageSize = 512             // 消息体最大容量
)

// Upgrade upgrading an HTTP connection to a WebSocket connection.
var Upgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Connection is a middleman between the websocket connection and the hub.
type Connection struct {
	Ws   *websocket.Conn // The websocket connection.
	Send chan []byte     // Buffered channel of outbound messages.
}

// NewConnection 初始化 Connection
func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{
		Send: make(chan []byte, 256),
		Ws:   conn,
	}
}

// ServeWs handles Websocket request from the peer
func ServeWs(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	wsConn, err := Upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error: upgrade websocket.", err)
		return
	}
	conn := NewConnection(wsConn)

	go conn.Writer()
	go conn.Reader()
}

// Writer write msg to the websocket
func (c *Connection) Writer() {
	defer func() {
		c.Ws.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.Ws.SetWriteDeadline(time.Now().Add(WriteWait))
				c.Ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.Ws.SetWriteDeadline(time.Now().Add(WriteWait))
			w, err := c.Ws.NextWriter(websocket.TextMessage)
			if err != nil {
				fmt.Println("Error: c.Ws.NextWriter.", err)
				return
			}
			if _, err := w.Write(msg); err != nil {
				fmt.Println("Error: write msg failed.", err)
				return
			}
			if err := w.Close(); err != nil {
				fmt.Println("Error: close ws failed.", err)
				return
			}

			fmt.Printf("write msg: %v\n", string(msg))
		}
	}
}

// Reader read msg from the websocket
func (c *Connection) Reader() {
	defer func() {
		c.Ws.Close()
	}()

	// c.Ws.SetPingHandler(func(string) error {
	// 	// 设置PingHandler，对客户端发来的Ping消息回应Pong
	// 	c.Ws.SetReadDeadline(time.Now().Add(PingWait))
	// 	if err := c.Ws.WriteControl(websocket.PongMessage, []byte{}, time.Now().Add(WriteWait)); err != nil {
	// 		return err
	// 	}
	// 	fmt.Printf("pong %v msg: %v\n", c.Ws.RemoteAddr(), websocket.PongMessage)
	//
	// 	return nil
	// })

	c.Ws.SetReadLimit(MaxMessageSize)
	for {
		mt, b, err := c.Ws.ReadMessage()
		if err != nil {
			fmt.Printf("Error: read msg from the websocket failed. MsgType: %v. %v\n", mt, err)
			return
		}
		fmt.Printf("read  msg: %v\n", string(b))
	}
}

// Ping send heartbeat messages to client
// func (c *Connection) Ping() {
// 	ticker := time.NewTicker(PingPeriod)
// 	defer func() {
// 		ticker.Stop()
// 		c.Ws.Close()
// 	}()
//
// 	for {
// 		select {
// 		case <-ticker.C:
// 			c.Ws.SetWriteDeadline(time.Now().Add(WriteWait))
// 			fmt.Printf("ping %v msg: %v\n", c.Ws.RemoteAddr(), websocket.PingMessage)
// 			err := c.Ws.WriteMessage(websocket.PingMessage, []byte{})
// 			if err != nil {
// 				fmt.Printf("Error: ping %v failed. %v", c.Ws.RemoteAddr(), err)
// 				return
// 			}
// 		}
// 	}
// }
