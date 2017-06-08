package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A string
	C int
}

func (t T) Echoer() {
	fmt.Printf("Echo: A=%v, C=%v\n", t.A, t.C)
}

func (t T) Seter(str string, int int) {
	fmt.Printf("Echo: oldstr=%v, newstr=%v\n", t.A, str)
	fmt.Printf("Echo: oldint=%v, newint=%v\n", t.C, int)
}

func main() {
	t := T{A: "A", C: 3}

	v := reflect.ValueOf(t)

	fmt.Println("Method Num: ", v.NumMethod())
	args := []reflect.Value{reflect.ValueOf("aaaaa"), reflect.ValueOf(-333)}

	md0 := v.Method(0)
	md0.Call(nil)
	md1 := v.MethodByName("Seter")
	md1.Call(args)
}
