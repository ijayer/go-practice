package main

import "fmt"

func main() {
	result := Add1()
	fmt.Printf("result = %v\n", result(3))

	result = Add2(2)
	fmt.Printf("result = %v\n", result(3))
}

func Add1() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Add2(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
