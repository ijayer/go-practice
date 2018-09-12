package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go pump(ch)

	// 生产者会在第一次发送后阻塞
	// 通道数据输出后再无接收者，导致通道非空
	fmt.Printf("#_________output = %v\n", <-ch)

	// 开启一个协程在无限循环中接受数据
	go suck(ch)

	time.Sleep(1 * time.Microsecond)
}

// 生产者，无线循环产生整形数
func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

// 消费者，无线循环接受整形数
func suck(ch chan int) {
	for {
		fmt.Printf("#_________output = %v\n", <-ch)
	}
}
