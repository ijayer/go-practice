package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig = "666"
	var an int
	var news string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	//strconv.Atoi()提供两个返回值，一个是转换后的结果，另一个是
	//可能出现的错误
	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5

	news = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", news)
}
