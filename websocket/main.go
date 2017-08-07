// websocket test: http://www.blue-zero.com/WebSocket/
package main

import (
	"flag"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
	"text/template"
)

var addr = flag.String("addr", ":8082", "http service addr")
var homeTemplate = template.Must(template.ParseFiles("D:/code/go_path/src/instance.golang.com/websocket/client/home.html"))

func main() {
	flag.Parse()

	router := httprouter.New()

	// listen channel semaphore
	go hub.run()

	userResource := NewUserResource()
	router.GET("/home", userResource.serveHome)
	router.GET("/ws/*name", userResource.serveWs)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})
	fmt.Printf("##_______________[Listen and serice on%s]\n", *addr)
	if err := http.ListenAndServe(*addr, c.Handler(router)); err != nil {
		logrus.Fatal("ListenAndServe: ", err)
	}
}
