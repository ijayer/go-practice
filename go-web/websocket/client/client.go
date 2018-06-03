// websocket client
package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var host = flag.String("host", "192.168.1.51:444", "http service host address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	done := make(chan struct{})

	conn := connection()
	reader(conn, done)

	type Data struct {
		ID       string `json:"id"`
		Platform string `json:"platform"`
	}
	dt := Data{"58bf9acbf608351f5c42395f", "windows"}
	conn.WriteJSON(&dt)

	time.Sleep(1 * time.Second)
	writer(conn, done, interrupt)
}

func connection() *websocket.Conn {
	u := url.URL{Scheme: "ws", Host: *host, Path: ""}
	log.Printf("##_________url = %v\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("#Dial error: %v\n", err)
	}

	return conn
}

func reader(conn *websocket.Conn, done chan struct{}) {
	go func() {
		defer conn.Close()
		defer close(done)

		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("#Read error: %v\n", err)
				return
			}
			log.Printf("##_________Recv msg = %v\n", message)
		}
	}()
}

func writer(conn *websocket.Conn, done chan struct{}, interrupt chan os.Signal) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case tick := <-ticker.C:
			err := conn.WriteMessage(websocket.TextMessage, []byte(tick.String()))
			if err != nil {
				log.Printf("#Send error: %v\n", err)
				return
			}
		case <-interrupt:
			log.Println("##_________Interrupt")
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Printf("#Write close error: %v\n", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
				conn.Close()
				return
			}
		}
	}
}
