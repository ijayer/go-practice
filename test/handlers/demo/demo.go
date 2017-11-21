// official demo of goland

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Method, r.URL, r.Host)
		io.WriteString(w, "hello world")
	}

	req := httptest.NewRequest("GET", "http://localhost:8090", nil)
	resp := httptest.NewRecorder()
	handler(resp, req)

	result := resp.Result()
	body, _ := ioutil.ReadAll(result.Body)
	fmt.Println(string(body))
}
