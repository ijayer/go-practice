package myhttprouter

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
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

	log.Fatal(http.ListenAndServe(":"+*port, router))
}
