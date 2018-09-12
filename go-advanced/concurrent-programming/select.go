/*
 * 说明：
 * 作者：zhe
 * 时间：2018-06-23 15:25
 * 更新：
 */

package main

import (
	"log"
	"time"
)

func main() {
	// ch1, ch2 := make(chan int), make(chan int)
	// go Select(ch1, ch2)
	// time.Sleep(2 * time.Millisecond) // 阻塞 2ms, 等待 go 程被调度

	// random := RandomBits()
	// for i := 0; i < 10; i++ {
	// 	print(<-random)
	// }
	// println()

	ch := make(chan string)
	go selectWithTimeout(ch)
	// ch <- "news: boom boom boom"
	time.Sleep(2 * time.Second) // 阻塞2秒，等待 go 程调度

	log.Println("main end...")
}

// Select
func Select(ch1 chan int, ch2 chan int) {
	for {
		select {
		case <-ch1:
			log.Println("ch1")
		case <-ch2:
			log.Println("ch2")
		default:
			log.Println("nothing available")
		}
	}
}

// RandomBits 生成无限的随机二进制序列
func RandomBits() <-chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 1:
			case ch <- 2:
			}
		}
	}()
	return ch
}

// selectWithTimeout 一个带有超时机制的阻塞通道
func selectWithTimeout(ch chan string) {
	select {
	case s := <-ch:
		log.Println("[recv]", s)
	case <-time.After(time.Second): // 一秒超时
		log.Println("timeout")
	}
}
