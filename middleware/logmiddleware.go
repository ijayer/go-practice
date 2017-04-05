package main

import (
	"log"
	"net/http"
)

func MiddlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("#______________________into middleware one")
			next.ServeHTTP(w, r)
			log.Println("#______________________into middleware one again")
		},
	)
}

func MiddlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("#______________________into middleware two")
			if r.URL.Path != "/" {
				w.Write([]byte("r.URL.Path is not `/`"))
				return
			}
			next.ServeHTTP(w, r)
			log.Println("#______________________into middleware two again")
		},
	)
}

func FinalMiddleware(w http.ResponseWriter, r *http.Request) {
	log.Println("#______________________into final middleware")
	w.Write([]byte("ok."))
}

func MainLogMiddleware(port *string) {
	finalHandler := http.HandlerFunc(FinalMiddleware)
	http.ListenAndServe(":"+*port, MiddlewareOne(MiddlewareTwo(finalHandler)))
}
