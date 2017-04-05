package main

import (
	"fmt"
	"log"
	"flag"
	"net/http"
)

var Port = flag.String("port", "8000", "http port")

func main() {
	flag.Parse()
	fmt.Printf("##_______________[Listen and serice on：%s]\n", *Port)
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/upload", uploadServer)
	err := http.ListenAndServe(":" + *Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
