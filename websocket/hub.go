package main

import "fmt"

type Hub struct {
	// registered connections
	connections 	map[*Conn]bool
	// inbound messages from the connections
	broadcast   	chan []byte
	// register requests from the connections
	register 	chan *Conn
	// unregister requests from the connections
	unregister	chan *Conn
}

var hub = Hub{
	broadcast:	make(chan []byte),
	register:	make(chan *Conn),
	unregister: 	make(chan *Conn),
	connections: 	make(map[*Conn]bool),
}

func NewHub() *Hub {
	return &Hub{
		connections: 	make(map[*Conn]bool),
		broadcast: 	make(chan []byte),
		register: 	make(chan *Conn),
		unregister: 	make(chan *Conn),
	}
}

func (h *Hub) run() {
	for {
		select {
		case conn := <- h.register:
			fmt.Println("#Registered")
			h.connections[conn] = true
		case conn := <- h.unregister:
			fmt.Println("#Unregistered")
			if _, ok := h.connections[conn]; ok {
				delete(h.connections, conn)
				close(conn.send)
			}
		case message := <- h.broadcast:
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

