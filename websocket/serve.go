package main

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

// load home.html
func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("##_________into serveHome handle")
	if r.URL.Path != "/home" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTemplate.Execute(w, r.Host)
}


// serveWs handles webSocket requests from the peer
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("##_________into serveWs handle")
	// upgrade web socket protocol
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// init instance of Conn
	conn := &Conn{send: make(chan []byte, 256), ws: ws}

	// register to map[connections]
	hub.register <- conn

	// goroutine: send msg from hub to websocket(client)
	go conn.writePump()

	// loop for read msg from websocket(client) to hub
	conn.readPump()
}