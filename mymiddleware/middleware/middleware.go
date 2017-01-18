package middleware

import (
	"net/http"
)

type Middleware struct{}

func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(message))
		},
	)
}

func exampleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// our middleware logic goes here...
			next.ServeHTTP(w, r)
		},
	)
}

func MainMiddleware(port *string) {
	http.ListenAndServe(":" + *port, exampleMiddleware(messageHandler("Hello World")))
}