package main

import (
	"fmt"
	"instance.golang.com/utils"
	"time"
)

func main() {
	in1 := make(chan int)
	in2 := make(chan string)

	go produce1(in1)
	go produce2(in2)
	go consumers(in1, in2)

	time.Sleep(1000 * time.Millisecond)
}

func produce1(in1 chan int) {
	for i := 0; ; i++ {
		in1 <- i
	}
}

func produce2(in2 chan string) {
	for i := 0; ; i++ {
		in2 <- utils.Now()
	}
}

func consumers(out1 chan int, out2 chan string) {
	for {
		select {
		case v := <-out1:
			fmt.Printf("##_________received from channel [1]: %v\n", v)
		case vv := <-out2:
			fmt.Printf("##_________received from channel [2]: %v\n", vv)
		default:
			fmt.Println("##_________channel don't output")
		}
	}
}
