/*
 * 说明：
 * 作者：zhe
 * 时间：2018-05-30 21:52
 * 更新：
 */

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v interface{}) error {
	rd, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(rd)
	return err
}

func ReadJSON(r *http.Request, obj interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, obj)
}
