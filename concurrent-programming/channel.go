/*
 * 说明：
 * 作者：zhe
 * 时间：2018-06-18 20:26
 * 更新：
 */

package main

import (
	"fmt"
	"log"
	"time"
)

/*
func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go func() {
		ch1 <- 10
		ch2 <- "hello"
		close(ch1)
		close(ch2)
	}()

	log.Println("recv from ch1:", <-ch1) // 输出：10
	log.Println("recv from ch2:", <-ch2) // 输出："hello"

	log.Println("recv finish, ch1 & ch2 has been closed")

	log.Println("recv form ch1:", <-ch1) // 输出零值：0, 不会阻塞
	log.Println("recv form ch2:", <-ch2) // 输出零值：""(空字符串), 不会阻塞
	log.Println("recv form ch1:", <-ch1) // 再次输出零值：0, 不会阻塞

	v, ok := <-ch2
	if !ok {
		// at this point: k is false
		log.Println("read on a closed chan")
	}
	log.Println("recv form ch2:", v) // 再次输出零值：""(空字符串), 不会阻塞
}
*/

func main() {
	wait := Publish2("A goroutine starts a new thread of execution.", 2*time.Second)
	log.Println("Let's hope the news will published before I leave.")

	<-wait // block, 等待发布新闻完成
	// wait has been closed, now received zero-value
	log.Println("Ten seconds later: I’m leaving now.")
}

// Publish prints text to stdout after the given time has expired.
// It closes the wait channel when the text has been published.
//
// wait 是一个只允许接收的通道类型
func Publish2(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{}) // ch 是一个双向通道类型
	go func() {
		time.Sleep(delay)
		log.Println("BREAKING NEWS:", text)
		close(ch)
	}() // 注意这里的括号，必须调用匿名函数

	return ch // Note: 双向通道类型可以赋值给单通道类型的
}

// demo 测试单通道和双向通道间相互赋值是否ok
func demo() {
	ch := make(chan int)
	sendCh := make(chan<- int)
	recvCh := make(<-chan int)

	// 双向通道赋值给单向通道，成功
	sendCh = ch
	recvCh = ch

	// 单向通道传递给双向通道，失败
	// ch = sendCh
	// ch = recvCh

	// 单向通道间相互传递，失败
	// sendCh = recvCh
	// recvCh = sendCh

	fmt.Sprintln(sendCh, recvCh)
}
