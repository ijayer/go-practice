package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		url := "https://hangzhou.qixing-group.com:90/v1.0/users/login"

		payload := strings.NewReader("{\n  \"account\":\"admin\",\n  \"password\":\"admin\",\n  \"platform\":\"android\"\n}")

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("content-type", "application/json")
		req.Header.Add("platform", "android")
		req.Header.Add("cache-control", "no-cache")
		req.Header.Add("postman-token", "237794aa-c559-a298-050f-ca323a46e051")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Printf("length: %v, %x\n", len(body), body)

		time.Sleep(1 * time.Second)
	}
}
