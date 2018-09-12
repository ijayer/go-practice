package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/zhezh09/go-practice/utils"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Conn is an middleman between the webSocket connection and the hub
type Connection struct {
	// id
	id string
	// platform
	platform string
	// the web socket connection
	ws *websocket.Conn
	// buffered channel of outbound messages
	send chan []byte
}

// reader pumps messages from the webSocket connection to the hub
func (s *UserResource) Reader(c *Connection) {
	fmt.Println("##_________into reader")
	defer func() {
		hub.unregister <- c
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				logrus.Printf("error: %v\n", err)
			}
			fmt.Printf("##_________Read msg error: %v\n", err)
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))
		hub.broadcast <- message
		fmt.Printf("#Send msg=%v to hub.broadcast\n", message)
	}
}

// write writes a message with the given message type and payload
func (s *UserResource) write(c *Connection, mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePum pumps messages from the hub to the webSocket connection
func (s *UserResource) writer(c *Connection) {
	fmt.Println("##_________into writer")
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// the hub closed the channel
				s.write(c, websocket.CloseMessage, []byte{})
				return
			}

			c.ws.SetWriteDeadline(time.Now().Add(writeWait))
			w, err := c.ws.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			fmt.Printf("#Send msg=%v to client\n", message)

			// add queued chat messages to the current webSocket message
			n := len(c.send)
			for i := 0; i < n; i++ {
				fmt.Println("#Send queued messages to the current webSocket message")
				w.Write(newLine)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			fmt.Printf("##_________Send ping at: %v\n", utils.Now())
			if err := s.write(c, websocket.PingMessage, []byte{'p'}); err != nil {
				return
			}
		}
	}
}
