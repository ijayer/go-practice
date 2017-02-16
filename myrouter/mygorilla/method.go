package mygorilla

import "net/http"

type SourceStorage struct{}

func NewSourceStorage() *SourceStorage {
	return &SourceStorage{}
}

func (s SourceStorage) GetMethodDemo(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query().Get("hello")
	WriteJson(w, v)
}

func (s SourceStorage) PostMethodDemo(w http.ResponseWriter, r *http.Request) {
	var v interface{}
	ReadJSON(r, &v)
	// do some logical processing
	WriteJson(w, v)
}

func (s SourceStorage) PutMethodDemo(w http.ResponseWriter, r *http.Request) {
	var v interface{}
	ReadJSON(r, &v)
	// do some logical processing
	WriteJson(w, v)
}

func (s SourceStorage) DeleteMethodDemo(w http.ResponseWriter, r *http.Request) {
	var v interface{}
	ReadJSON(r, &v)
	// do some logical processing
	WriteJson(w, v)
}