package main

import (
	"fmt"
)

func main() {
	funcB()
}

func trace(s string) {
	fmt.Println("entering: ", s)
}

func unTrace(s string) {
	fmt.Println("leaving: ", s)
}

func funcA() {
	trace("a")
	defer unTrace("a")
	fmt.Println("in a")
}

func funcB() {
	trace("b")
	defer unTrace("b")
	fmt.Println("in b")
	funcA()
}
