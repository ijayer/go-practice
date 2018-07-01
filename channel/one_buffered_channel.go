/*
 * 说明：缓冲 chan 和 无缓冲 chan 测试
 * 作者：zhe
 * 时间：2018-06-21 9:51 PM
 * 更新：
 */

package main

import "fmt"

func main() {
	// unBufferedChan()
	oneBufferedChan()
}

func unBufferedChan() {
	ch := make(chan int) // or <=> ch := make(chan int, 0)
	// go func(i chan int) {
	// 	println(<-i)
	// }(ch)
	ch <- 1             // 阻塞，且会死锁；无缓冲的 chan 必须在另一个 goroutine 中读取该 chan 的数据，否则会造成死锁
	fmt.Println("send") // 不会输出
}

func oneBufferedChan() {
	ch := make(chan int, 1) // 可容纳一个元素
	ch <- 1                 // 不会阻塞, 且不会有死锁
	fmt.Println("send")     // 会输出：send
	fmt.Println(<-ch)       // 输出：1
	ch <- 2                 // 不会阻塞, 且不会有死锁
	fmt.Println("send")     // 会输出：send
	ch <- 3                 // 阻塞，且死锁
}
