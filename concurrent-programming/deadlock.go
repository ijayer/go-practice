/*
 * 说明：
 * 作者：zhe
 * 时间：2018-06-23 15:45
 * 更新：
 */

package main

import (
	"log"
	"time"
)

func main() {
	wait := PublishOk("PublishOk", time.Second)
	<-wait

	waitNoCloseReturn := PublishWithDeadlock("PublishWithDeadlock", time.Second)
	<-waitNoCloseReturn

	log.Println("main end...")
}

// PublishOk
func PublishOk(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	time.AfterFunc(delay, func() {
		log.Println("[NEWS]", text)
		// 发布成功后，通知 wait 结束等待
		close(ch)
	})
	return ch
}

// PublishWithDeadlock
func PublishWithDeadlock(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	time.AfterFunc(delay, func() {
		log.Println("[NEWS]", text)
		// 注释 close(ch), 即 wait 通道会永久阻塞等待，造成死锁
		// close(ch)
	})
	return ch
}
