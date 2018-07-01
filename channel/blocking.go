package main

import (
	"fmt"
	"time"
)

func main() {
	output := make(chan int, 3)

	// 无缓冲通道必须要有一个 `接收者` 准备好接受
	// 通道的数据，然后 `发送者` 才可以将数据发送
	// 给接收者，否则就会产生死锁：
	// fatal error: all goroutines are asleep - deadlock!

	output <- 3 //
	output <- 3 //
	output <- 3 //
	// output <- 3	// 程序阻塞
	go f(output)

	time.Sleep(1 * time.Millisecond)
}

func f(in chan int) {
	for {
		fmt.Printf("#________input = %v\n", <-in)
	}
}
