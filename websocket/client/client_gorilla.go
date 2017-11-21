package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var WsURL = "ws://192.168.1.51:444/"

var dialer *websocket.Dialer

func main() {
	conn, _, err := dialer.Dial(WsURL, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conn.RemoteAddr())

	var v interface{}
	for {
		message := []byte("hello, world!")
		err = conn.WriteJSON(message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Send: %s\n", string(message))

		err := conn.ReadJSON(v)
		if err != nil {
			fmt.Println("read:", err)
		}
		fmt.Printf("received: %s\n", v)

		time.Sleep(time.Second * 5)
	}
}

func timeWriter(conn *websocket.Conn) {
	for {
		time.Sleep(time.Second * 2)
		conn.WriteMessage(websocket.TextMessage, []byte(time.Now().Format("2006-01-02 15:04:05")))
	}
}
