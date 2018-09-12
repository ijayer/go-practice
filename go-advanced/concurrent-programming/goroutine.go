/*
 * 说明：Go并发编程基础-执行线程(Goroutine)
 * 作者：zhe
 * 时间：2018-06-18 18:00
 * 更新：
 */

package main

import (
	"log"
	"time"
)

// demo_1.go
/*
// The following program will print “Hello from main goroutine”. It might also print
// “Hello from another goroutine”, depending on which of the two goroutines finish first.

func main() {
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	// 至此，程序执行结束，且所有活跃的 go程都将被 killed
}
*/

// demo_2.go
/*
// The next program will, most likely, print both “Hello from main goroutine” and “Hello
// from another goroutine”. They may be printed in any order. Yet another possibility is
// that the second goroutine is extremely slow and doesn’t print its message before the program ends.

func main() {
	go fmt.Println("Hello from another goroutine")
	fmt.Println("Hello from main goroutine")

	// main goroutine sleep 1s...
	// wait for other goroutine to finish
	time.Sleep(time.Second)
}
*/

// demo_3.go

func main() {
	Publish("A goroutine starts a new thread of execution.", 5*time.Second)
	log.Println("Let's hope the news will published before I leave.")

	// 等待发布新闻
	time.Sleep(10 * time.Second)
	log.Println("Ten seconds later: I’m leaving now.")
}

// Publish 函数在延迟delay到期后，打印text内容到标准输出
// 且该函数不会阻塞而是执行完成后立即返回
func Publish(text string, delay time.Duration) {
	go func() {
		time.Sleep(delay)
		log.Println("BREAKING NEWS:", text)
	}() // 注意这里的括号，必须调用匿名函数
}
