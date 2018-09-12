package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	School string `json:"school"`
}

func (u User) Hello() {
	fmt.Println("Hello World.")
}

func Info(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println(t.Name())

	v := reflect.ValueOf(i)
	// fmt.Println(v.Interface())
	fmt.Println("Fields:", "num:", t.NumField())
	for i := 0; i < v.NumField(); i++ {
		t := t.Field(i)
		fmt.Printf("%10s: %v = %v %v\n", t.Name, t.Type, v.Field(i).Interface(), t.Tag)
	}

	fmt.Println("Methods:", "num:", t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		t := t.Method(i)
		fmt.Printf("%10s: %v\n", t.Name, t.Type)
	}
}

func main() {
	user := User{"123", "nnnn", 12, "hz"}
	Info(user)
}
