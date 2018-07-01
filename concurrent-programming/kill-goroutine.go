/*
 * 说明：how to kill a goroutine
 * 作者：zhe
 * 时间：2018-07-01 15:02
 * 更新：update demo
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// test goWithQuitChan
	wg.Add(1)
	quit := make(chan struct{})
	goWithQuitChan(quit)
	time.AfterFunc(time.Second*2, func() {
		quit <- struct{}{}
	})
	wg.Wait()

	// test Generator
	num := Generator()
	fmt.Println(<-num)
	fmt.Println(<-num)

	// close(num) 操作可能会引起 panic: send on a closed channel
	// Generator 中的 go 程在 select 选择上可能会出现: case ch <- n: 发送了，紧接着 case <- ch: 收到信号 return
	// 而 ch <- n 之前发送的值没有被接收就关闭了 channel, 从而就会是在一个关闭了的通道上执行发送操作 panic.
	//
	// Note: 有数据竞争
	close(num)

	fmt.Println("main end...")
}

// goWithQuitChan 带有 stop signal 的 Goroutine
func goWithQuitChan(quit chan struct{}) {
	defer func() {
		fmt.Println("[goroutine]goWithQuitChan end...")
		close(quit)
		wg.Done()
	}()

	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				// ...
			}
		}
	}()
}

// Generator returns a channel that produces the numbers 1, 2 ...
// To stop underlying goroutine, close the channel
func Generator() chan int {
	ch := make(chan int)

	go func() {
		n := 1
		for {
			select {
			case ch <- n: // produces 1, 2, 3 ...
				n++
			case <-ch: // receive close signal
				return
			}
		}
	}()

	return ch
}
