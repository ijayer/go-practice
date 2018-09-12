package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++ //ret作为f()的返回值
	}()
	return 1
}

func main() {
	fmt.Println(f())
}
