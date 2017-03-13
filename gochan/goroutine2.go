package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("main() is running")

	// 声明一个string类型的通道
	ch := make(chan string)

	// 开启发送和接受数据的协程
	// 协程通信必须传入相同参数
	go sendData(ch)
	go getData(ch)

	// 主函数休眠, 等待协程通信完成
	// 否则主函数退出, 所有操作都结束执行
	time.Sleep(1 * time.Second)
	fmt.Println("main() is over")
}

func sendData(ch chan string) {
	ch <- "aaa"
	ch <- "bbb"
	ch <- "ccc"
	ch <- "ddd"
	ch <- "eee"
}

func getData(ch chan string) {
	var input string

	// time.Sleep(1 * time.Second)

	for {
		input = <- ch
		fmt.Printf("##_________input = %v\n", input)
	}
}