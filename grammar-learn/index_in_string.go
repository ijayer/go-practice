package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "Hi, I'm a string, Hi."

	fmt.Printf("The position of \"string\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "string"))

	fmt.Printf("The position of \"Hi\" first in the instance is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of \"Hi\" last in the instance is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"whisper\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "whisper"))
}
