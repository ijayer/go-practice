package main

import (
	"log"
	"os"
)

func main() {
	os.Mkdir("astaxie", 0777)
	os.MkdirAll("astaxie/test1/test2", 0777)

	err := os.Remove("astaxie")
	if err != nil {
		log.Println(err)
	}

	os.RemoveAll("astaxie")
}
