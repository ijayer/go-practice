package main

import (
	"errors"
	"fmt"
	"log"
)

// 声明USB接口类型
type USB interface {
	Connecter
	Name() string
}

type Connecter interface {
	Connect()
}

// 声明一个结构体 PhoneConnecter 实现USB接口
type PhoneConnecter struct {
	name string
}

// 实现Name方法
func (p PhoneConnecter) Name() string {
	return p.name
}

// 实现Connect方法
func (p PhoneConnecter) Connect() {
	fmt.Printf("Connected: %v\n", p.name)
}

// 定义Disconnect函数，接收一个USB类型的变量
func Disconnect(usb interface{}) {
	pc, ok := usb.(PhoneConnecter)
	if !ok {
		log.Printf("%v\n", errors.New("unknow device"))
		return
	}
	fmt.Printf("Disconnected: %v\n", pc.Name())
}

func main() {
	var u USB
	pc := PhoneConnecter{name: "iPhone"}
	u = pc
	u.Connect()
	Disconnect(u)
}
