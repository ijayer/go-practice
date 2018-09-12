package main

import (
	"fmt"
	"reflect"
)

func main() {
	var n = 3.14

	t := reflect.TypeOf(n)
	fmt.Println(t)

	fmt.Println("#Copy")
	v := reflect.ValueOf(n)
	fmt.Println("Is CanSet? ", v.CanSet())
	// v.SetFloat(3.14159) // panic

	fmt.Println("#Quote")
	v = reflect.ValueOf(&n)
	fmt.Println("Is CanSet? ", v.CanSet())

	v = v.Elem() // &p -> *p
	fmt.Println("Is CanSet? ", v.CanSet())
	v.SetFloat(3.1415)
}
