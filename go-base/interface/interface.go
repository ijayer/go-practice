/*
 * Overview:    interface demo
 * Author:      @zhe
 * History:     2017-10-23 11:29:20
 * Note:
 * Links:       http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go
 */

package main

type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "WoWoWo"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "MioMio"
}

type Lion struct {
}

func (l Lion) Speak() string {
	return "HoHoHo"
}

type Gopher struct {
}

func (g Gopher) Speak() string {
	return "Gopher"
}

func main() {
	ans := []Animal{Dog{}, Cat{}, Lion{}, Gopher{}} // Dog, Cat, Lion, Gopher实现了Animal接口
	for _, v := range ans {
		println(v.Speak())
	}
}
