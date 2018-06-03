package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

var origin = "http://192.168.1.51/"
var wsURL = "ws://192.168.1.51:444"

func main() {
	ws, err := websocket.Dial(wsURL, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	for {
		message := []byte("hello, world!")
		_, err = ws.Write(message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Send: %s\n", message)

		var msg = make([]byte, 512)
		_, err = ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Receive: %s\n", msg)

		time.Sleep(time.Second * 5)
	}
}
