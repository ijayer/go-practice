package main

import (
	"errors"
	"reflect"
)

type Person struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Time   string `json:"time"`
	School School `json:"school"`
}

type School struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func InfoTag(v interface{}) error {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Struct {
		return errors.New("Invalid type")
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("json")
		println(tag)
	}
	return nil
}

func main() {
	p := Person{}
	InfoTag(p)
}

//import "fmt"
//
//type Users struct {
//	Name string
//}
//
//func main() {
//	u := Users{Name: "zhe"}
//	u.Write1()
//	fmt.Println(u)
//
//	uu := &Users{Name: "lie"}
//	uu.Write2()
//	fmt.Println(*uu)
//}
//
//func (u Users) Write1() {
//	u.Name = "zheer"
//}
//
//func (u *Users) Write2() {
//	u.Name = "lieer"
//}

//import "fmt"
//
//type Users struct {
//	Name string
//}
//
//func main() {
//	u := Users{Name: "zhe"}
//	u.Modify()
//	u.String()
//
//	uu := new(Users)
//	uu.Name = "lie"
//	uu.Modify1()
//	uu.String1()
//}
//
//func (u Users) String() {
//	fmt.Println("The users's name is: " + u.Name)
//}
//
//func (u Users) Modify() {
//	u.Name = "zheer"
//}
//
//func (u *Users) String1() {
//	fmt.Println("The users's name is: " + u.Name)
//}
//
//func (u *Users) Modify1() {
//	u.Name = "lieer"
//}
