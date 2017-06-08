package main

import (
	"fmt"
	"reflect"
)

type Company struct {
	Employee
	Title string `json:"title"`
}

type Employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	c := Company{Employee: Employee{"zhe", 2}, Title: "hzqx"}

	t := reflect.TypeOf(c)
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0}))

	Set(&c)
	fmt.Printf("%s: %#v\n", "Modiyied", c)
}

func Set(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println(t.Name())

	v := reflect.ValueOf(i)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Warning")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Title")
	if !f.IsValid() {
		fmt.Println("bad")
	}

	if f.Kind() == reflect.String {
		fmt.Println("setting...")
		f.SetString("byebye")
	}
}
