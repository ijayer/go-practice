package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

type UserResource struct {
	UserStorage *UserStorage
	Connection  *Connection
}

func NewUserResource() *UserResource {
	return &UserResource{}
}

// load home.html
func (s UserResource) serveHome(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
func (s *UserResource) serveWs(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Println("##_______________into serveWs handle")
	names := strings.Split(params.ByName("name"), "/")
	if len(names) != 3 {
		http.Error(w, fmt.Sprintf("%s", "Unprocesable fields or parameters"), http.StatusUnprocessableEntity)
		return
	}
	id := names[1]
	platform := names[2]

	// upgrade web socket protocol
	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// init instance of Conn
	conn := &Connection{
		id:       id,
		send:     make(chan []byte, 256),
		ws:       ws,
		platform: platform,
	}
	// register to map[connections]
	hub.register <- conn

	// goroutine: send msg from hub to websocket(client)
	go s.writer(conn)

	// loop for read msg from websocket(client) to hub
	s.Reader(conn)
}
