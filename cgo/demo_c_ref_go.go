package main

import "C"
import "fmt"

//export go_print
func go_print(value string) {
	fmt.Printf("Go print: %v\n", value)
}

// main函数是必须的，有main函数才能让Cgo编译器去把包编译成C的库
func main() {
}
