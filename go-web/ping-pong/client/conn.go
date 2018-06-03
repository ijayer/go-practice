/*
 * 说明：
 * 作者：zhe
 * 时间：2018-04-18 16:05
 * 更新：
 */

package main

import (
	"fmt"
	"time"

	"sync"

	"github.com/gorilla/websocket"
)

// Websocket读写相关配置
const (
	WriteWait      = 6 * time.Second // 写数据超时时间
	PingWait       = 5 * time.Second // 读取Ping消息超时时间
	PongWait       = 4 * time.Second // 读取Pong消息超时时间
	PingPeriod     = 3 * time.Second // 写Ping消息的时间频率(PingPeriod必须小于PongWait)
	MaxMessageSize = 512             // 消息体最大容量
)

// Connection is a middleman between the websocket connection and the hub.
type Connection struct {
	Send chan []byte     // Buffered channel of outbound messages.
	Ws   *websocket.Conn // The websocket connection.
	Mu   sync.RWMutex
}

// NewConnection 初始化 Connection
func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{
		Send: make(chan []byte, 256),
		Ws:   conn,
		Mu:   sync.RWMutex{},
	}
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

	for {
		mt, b, err := c.Ws.ReadMessage()
		if err != nil {
			fmt.Printf("Error: read msg failed. MsgType: %v. %v\n", mt, err)
			return
		}
		fmt.Printf("read msg: %v\n", string(b))
	}

	fmt.Println("goroutine reader closed")
}

// Ping send heartbeat messages to client
func (c *Connection) Ping() {
	ticker := time.NewTicker(PingPeriod)
	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()

	for {
		select {
		case <-ticker.C:
			c.Ws.SetWriteDeadline(time.Now().Add(WriteWait))
			// fmt.Printf("ping %v msg: %v\n", c.Ws.RemoteAddr(), websocket.PingMessage)
			err := c.Ws.WriteMessage(websocket.PingMessage, []byte{})
			if err != nil {
				fmt.Printf("Error: ping %v failed.%v\n", c.Ws.RemoteAddr(), err)
			}
		}
	}
}

// reconnect 当Websocket连接关闭后，尝试重连
// Issue: panic: repeated read on failed websocket connection
//
// func (c *Connection) reconnect() {
// 	fmt.Println("reconnecting...")
//
// 	// 连接服务器
// 	dialer := websocket.Dialer{}
// 	wsConn, _, err := dialer.Dial(wsUrl, nil)
// 	if err != nil {
// 		fmt.Println("Error: websocket dial failed.", err)
// 		fmt.Println("reconnectiong after 3s ...")
// 		time.Sleep(3 * time.Second)
//
// 		c.reconnect()
// 	}
// 	fmt.Println("reconnected")
//
// 	// 初始化 Connection, 开启读写 G 程
// 	c.Mu.Lock()
// 	conn := NewConnection(wsConn)
// 	c.Mu.Unlock()
//
// 	go conn.Ping()
// 	go conn.Reader()
// 	go conn.Writer()
// }
