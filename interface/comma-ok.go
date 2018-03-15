/*
 * 类型断言测试
 */
package main

import (
	"fmt"
	"strconv"
)

type Element interface{}

type List []Element

type Person struct {
	age  int
	name string
}

func (p Person) String() string {
	return "(name:" + p.name + "age:" + strconv.Itoa(p.age) + "years)"
}

func main() {
	list := make(List, 3)
	list[0] = 123
	list[1] = "hello world"
	list[2] = Person{age: 12, name: "bob"}

	//if-else test
	fmt.Println("@_@____________if-else test")
	for i, v := range list {
		if value, ok := v.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", i, value)
		} else if value, ok := v.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", i, value)
		} else if value, ok := v.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %v\n", i, value)
		} else {
			fmt.Printf("list[%d] is of a different type\n", i)
		}
	}

	//switch-case test
	fmt.Println("@_@____________switch-case test")
	for i, v := range list {
		switch value := v.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", i, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", i, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %v\n", i, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", i)
		}
	}
}
