package mygorilla

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v interface{}) error {
	rd, _ := json.Marshal(v)
	//gorilla cross origin para
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(rd)
	return err
}

func ReadJSON(r *http.Request, jsonObject interface{}) error {
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &jsonObject)
	return err
}
