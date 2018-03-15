package main

import (
	"fmt"
	"os"
)

func main() {
	var gopath = os.Getenv("GOPATH")
	fmt.Printf("The Go path is: %s\n", gopath)

	path := os.Getenv("PATH")
	fmt.Printf("The OS path is: %s\n", path)
}
