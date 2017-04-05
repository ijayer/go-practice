package myiris

import (
	"github.com/kataras/iris"
	"fmt"
)

type UserAPI struct{
	*iris.Context
}

func (u UserAPI) Get() {
	u.SetHeader("Accept-Language", "zh-CN,zh;q=0.8")
	u.Write("Get from /users")

	fmt.Println("header:", string(u.Request.Header.Header()))
	fmt.Println("header:", string(u.Request.Header.String()))
	fmt.Println("header:", string(u.RequestHeader("Accept-Language")))
}

func (u UserAPI) Post() {
	u.Write("Get from /users")

	fmt.Println("header:", string(u.Request.Header.String()))
	fmt.Println("header:", string(u.Request.Header.Header()))
}

func (u UserAPI) GetBy(id string) {
	u.Write("Get from /user/%s", id)
}

func (u UserAPI) Put() {
	name := u.FormValue("name")
	println(string(name))
	println("Put from /users")
}

func (u UserAPI) PostBy(id string) {
	name := u.FormValue("name")
	println(string(name))
	println("Put from /users/" + id)
}

func (u UserAPI) DeleteBy(id string) {
	println("Delete from /" + id)
}
