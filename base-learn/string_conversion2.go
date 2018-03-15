package main

import (
	"fmt"
	"strconv"
)

func main() {
	var origStr = "sss"
	var newStr string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, err := strconv.Atoi(origStr)
	if err != nil {
		fmt.Printf("the origStr %s is not an integer - exiting with error\n", origStr)
		return
	}
	fmt.Printf("the integer is %d\n", an)
	an += 5

	newStr = strconv.Itoa(an)
	fmt.Printf("The new string is : %s\n", newStr)
}
