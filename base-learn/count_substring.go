package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "Hello, how is it going , Hugo!"
	var manyG = "gggggggggggggg"

	fmt.Printf("The Times H's is in the %s is : ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("The times of \"gg\" is in the %s is : ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

}
