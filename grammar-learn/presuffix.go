package main

import (
	"fmt"
	"strings"
)

/*
func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\"have prefix %s?	", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
*/

func main() {
	var str = "lan zhou jiao tong da xue!"
	fmt.Printf(" T/F ? does the string \"%s\" is end of suffixx %s ?", str, "da xue!")
	fmt.Printf("%t\n", strings.HasSuffix(str, "da xue!"))
}
