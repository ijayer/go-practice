package mygorilla

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/gorilla/handlers"
	"time"
	log "github.com/Sirupsen/logrus"
)

func MainGorilla(port *string) {
	r := mux.NewRouter().StrictSlash(false)

	// RESTFul (Representational State Transfer)
	// Resources资源, Representation表现层, State transfer状态转化
	// a api path: representational one network resources
	// a method:   representational a way of handling to the resources
	sourceStorage := NewSourceStorage()
	r.HandleFunc("/api/sources/songs", sourceStorage.GetMethodDemo).Methods("Get")
	r.HandleFunc("/api/sources/songs", sourceStorage.PostMethodDemo).Methods("Post")
	r.HandleFunc("/api/sources/songs", sourceStorage.PutMethodDemo).Methods("Put")
	r.HandleFunc("/api/sources/songs", sourceStorage.DeleteMethodDemo).Methods("Delete")

	// static file server
	http.Handle("/api/sources", http.FileServer(http.Dir("./router/")))

	server := &http.Server{
		Handler:handlers.ContentTypeHandler(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET","POST","PUT","DELETE"}),
			handlers.AllowedHeaders([]string{"Content-Type"}))(r), "application/json"),
		Addr:   ":"+*port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
