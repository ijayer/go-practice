package handler

import (
	"net/http"
	"log"
	"strings"
	"fmt"
)

// Empty struct type
type foo struct {}

// foo implement ServeHTTP method of the http.Handler interface
func (f foo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("#______________________Into ServeHTTP of foo")
	w.Write([]byte("serve http foo."))
}

type page struct {
	body string
}

// page implement ServeHTTP method of the http.Handler interface
func (p page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("#______________________Into ServeHTTP of page")
	w.Write([]byte(p.body))
}

// string is the url path and http.Handler is any type that has a ServeHTTP method
type multiplexer map[string]http.Handler

// multiplexer implement ServeHTTP method of the http.Handler interface
func (m multiplexer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("#______________________into ServeHTTP of multiplexer")

	if handler, ok := m[r.RequestURI]; ok {
		handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

var mux = multiplexer{
	"/": 		foo{},
	"/about/":	page{"about"},
	"/contact/":	page{"contact"},
}

func appendingTrailingSlash(handler http.Handler) http.Handler {
	log.Println("#______________________into middleware")
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("#______________________into return handler")
			fmt.Println(r.URL.Path)
			if !strings.HasSuffix(r.URL.Path, "/") {
				http.Redirect(w, r, r.URL.Path + "/", http.StatusNotFound)
			} else {
				handler.ServeHTTP(w, r)
			}
		},
	)
}

// MainHandler called in main function
func MainHandler(port *string) {
	// wrap mux with middleware
	http.ListenAndServe(":" + *port, appendingTrailingSlash(mux))
}

