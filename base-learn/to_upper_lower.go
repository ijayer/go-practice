package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "Hi, How are you google, I'm Apple!"
	var lower string
	var upper string

	fmt.Printf("original: %s\n", str)

	lower = strings.ToLower(str)
	fmt.Printf("ToLower: %s\n", lower)

	upper = strings.ToUpper(str)
	fmt.Printf("ToUpper: %s\n", upper)
}
