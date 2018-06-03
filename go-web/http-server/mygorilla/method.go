/*
 * 说明：
 * 作者：zhe
 * 时间：2018-05-25 2:42 PM
 * 更新：
 */

package mygorilla

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

func WriteJson(w http.ResponseWriter, v interface{}) error {
	rd, _ := json.Marshal(v)
	// gorilla cross origin para
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(rd)
	return err
}

func ReadJSON(r *http.Request, jsonObject interface{}) error {
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &jsonObject)
	return err
}
