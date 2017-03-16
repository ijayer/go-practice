package main

import (
	"flag"
	"net/http"
	"text/template"
	"github.com/Sirupsen/logrus"
)

var addr = flag.String("addr", ":8082", "http service addr")
var homeTemplate = template.Must(template.ParseFiles("D:/code/go_path/src/instance.golang.com/websocket/home.html"))

func main() {
	flag.Parse()

	// listen channel semaphore
	go hub.run()

	http.HandleFunc("/home", serveHome)
	http.HandleFunc("/ws", serveWs)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		logrus.Fatal("ListenAndServe: ", err)
	}
}
