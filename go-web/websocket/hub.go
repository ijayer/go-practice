package main

import "fmt"

type Hub struct {
	// registered connections
	connections map[*Connection]bool
	// inbound messages from the connections
	broadcast chan []byte
	// register requests from the connections
	register chan *Connection
	// unregister requests from the connections
	unregister chan *Connection
}

var hub = Hub{
	broadcast:   make(chan []byte),
	register:    make(chan *Connection),
	unregister:  make(chan *Connection),
	connections: make(map[*Connection]bool),
}

func NewHub() *Hub {
	return &Hub{
		connections: make(map[*Connection]bool),
		broadcast:   make(chan []byte),
		register:    make(chan *Connection),
		unregister:  make(chan *Connection),
	}
}

func (h *Hub) run() {
	for {
		select {
		case conn := <-h.register:
			fmt.Println("#Registered")
			h.connections[conn] = true
		case conn := <-h.unregister:
			fmt.Println("#Unregistered")
			if _, ok := h.connections[conn]; ok {
				delete(h.connections, conn)
				close(conn.send)
			}
		case message := <-h.broadcast:
			fmt.Println("#Broadcasted")
			for conn := range h.connections {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(hub.connections, conn)
				}
			}
		}
	}
}
