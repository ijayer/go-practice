package main

import (
	"log"
	"strings"
)

func main() {

	s := "Hello, workd! Hello!"

	i := strings.LastIndex(s, "h")

	log.Println(i)

	i = strings.LastIndex(s, "H")

	log.Println(i)

	i = strings.LastIndex(s, "")

	log.Println(i)

}
