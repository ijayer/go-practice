package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Define
var DBName string
var DBAddr string
var Addr string

// define struct App, to hold our applications
type App struct {
	ver     string
	Router  *mux.Router
	Session *mgo.Session
}

// responder for the http response
type Responder struct {
	Total int         `json:"total"`
	Data  interface{} `json:"data"`
}

// NewApp return an instance of App
func NewApp() *App {
	return &App{ver: "v0"}
}

// Initialize init the application
func (a *App) InitMgo(db, addr string) {
	log.Printf("mgo listen and serve on [%v, %v]\n", db, addr)
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{addr},
		Timeout:  time.Second * 30,
		Database: db,
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}
	a.Session = session
	a.Router = mux.NewRouter()
	a.initRoutes()
}

// Run the application
func (a *App) Run(addr string) {
	log.Printf("api listen and serve on [:%v]\n", addr)
	log.Fatal(http.ListenAndServe(":"+addr, a.Router))
}

// Clean up the database
func (a *App) Clean(dbName string) {
	session := a.Session.Copy()
	defer session.Close()

	err := a.Session.DB(dbName).DropDatabase()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("the database %v has been cleaned\n", dbName)
}

// Initialize api routes
func (a *App) initRoutes() {
	a.Router.HandleFunc("/"+a.ver+"/products", a.Find).Methods(http.MethodGet)
	a.Router.HandleFunc("/"+a.ver+"/products", a.Create).Methods(http.MethodPost)
	a.Router.HandleFunc("/"+a.ver+"/products/{id:[0-9]+}", a.FindOne).Methods(http.MethodGet)
	a.Router.HandleFunc("/"+a.ver+"/products/{id:[0-9]+}", a.Update).Methods(http.MethodPut)
	a.Router.HandleFunc("/"+a.ver+"/products/{id:[0-9]+}", a.Delete).Methods(http.MethodDelete)
}

// Find list of products
func (a *App) Find(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	offset, _ := strconv.Atoi(r.FormValue("offset"))

	if offset < 0 {
		offset = 0
	}
	if limit > 10 || limit < 1 {
		limit = 10
	}

	results, err := find(bson.M{}, a.Session, offset, limit)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if len(results) == 0 {
		respondWithJson(w, http.StatusOK, Responder{Total: 0, Data: []struct{}{}})
		return
	}
	respondWithJson(w, http.StatusOK, Responder{Total: len(results), Data: results})
}

// Create a product
func (a *App) Create(w http.ResponseWriter, r *http.Request) {
	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.create(p, a.Session); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, Responder{Total: 1, Data: p})
}

// Find a product
func (a *App) FindOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid id")
		return
	}

	p := Product{PID: id}
	if p, err = p.findOne(id, a.Session); err != nil {
		switch err {
		case mgo.ErrNotFound:
			respondWithError(w, http.StatusNotFound, err.Error())
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJson(w, http.StatusOK, Responder{Total: 1, Data: p})
}

// Update a product
func (a *App) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid product id")
		return
	}
	var p Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid request payload")
		return
	}
	defer r.Body.Close()

	p.PID = id

	b, _ := json.Marshal(p)
	var update bson.M
	json.Unmarshal(b, &update)
	delete(update, "id")
	if err := p.update(bson.M{"pid": id}, update, a.Session); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, Responder{Total: 1, Data: p})
}

// Delete a product
func (a *App) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	p := Product{PID: id}
	if err := p.delete(bson.M{"pid": id}, a.Session); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// Respond with error to client
func respondWithError(w http.ResponseWriter, code int, err string) {
	respondWithJson(w, code, map[string]string{"error": err})
}

// Respond with json to client
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("json marshal error: %v\n", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
