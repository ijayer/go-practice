package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name float64 = 3.4

	t := reflect.TypeOf(name)
	fmt.Println(t)
	fmt.Println(t.String())

	v := reflect.ValueOf(name)
	fmt.Println(v)
	fmt.Println(v.String())

	v.Elem()
}
