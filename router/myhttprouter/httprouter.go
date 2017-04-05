package myhttprouter

import (
	"github.com/julienschmidt/httprouter"
	log "github.com/Sirupsen/logrus"
	"net/http"

)

func MainHttpRouter(port *string) {
	router := httprouter.New()

	// api
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/parameters/*name", AllParaHello)
	router.GET("/protected/", BasicAuth(Protected, "username", "secret"))

	// static file server
	router.ServeFiles("/static/*filepath", http.Dir("./router/static_sources"))

	log.Fatal(http.ListenAndServe(":" + *port, router))
}
