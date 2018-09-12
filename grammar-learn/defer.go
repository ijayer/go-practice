// defer 逆序执行

package main

import "fmt"

func main() {
	defer function1() // 第一个注册，最后执行
	defer function2() // 第二个注册，倒数第二个执行
	fmt.Println("This is third statement!")
}

func function1() {
	fmt.Println("This is first statement!")
}

func function2() {
	fmt.Println("This is second statement!")
}
