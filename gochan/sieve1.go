package main

import (
	"fmt"
)

func main() {

	ch := make(chan int) // 创建一个新通道

	go generate(ch) // 开启一个协程，作为数据生成器

	for {
		prime := <-ch
		fmt.Printf("p=%v\n", prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		//time.Sleep(500 * time.Millisecond)
	}
}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // send `i` to channel `ch`
	}
}

func filter(in, out chan int, prime int) {
	for {
		k := <-in
		fmt.Printf("k=%v\n", k)
		if k%prime != 0 {
			out <- k
		}
	}
}
