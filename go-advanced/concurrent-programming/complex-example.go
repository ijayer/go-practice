/*
 * 说明：Go 并发编程基础综合示例
 * 作者：zhe[Reference: http://blog.xiayf.cn/2015/05/20/fundamentals-of-concurrent-programming/](9. 综合所有示例)
 * 时间：2018-07-01 18:33
 * 更新：Updating demo
 */

package main

import (
	"fmt"
	"sync"
)

func main() {
	people := []string{"KeBe", "Wade", "Paul", "HaDen", "James"}

	match := make(chan string, 1) // 为一个未匹配的发送操作提供空间

	wg := new(sync.WaitGroup)
	wg.Add(len(people))
	for _, name := range people {
		go Seek(name, match, wg)
		// Step1: 第一个被调度的 goroutine 必定执行分支 `case match <- name:` 给通道写入数据，且不阻塞
		// Step2: 第二个被调度的 goroutine 必定执行分支 `case peer :<- match:` 读取通道的数据
		// Step3: 同 Step1
		// Step4: 同 Step2
		// Step5: 同 Step1
	}
	wg.Wait() // 等待所有 goroutine 结束

	select {
	case name := <-match:
		// 接收来自最后一个被调度的 goroutine 发送的数据
		fmt.Printf("No one received %s’s message.\n", name)
	default:
		// 没有待处理的发送操作
	}
}

// Seek 发送一个name到match管道或从match管道接收一个peer，结束时通知wait group
func Seek(name string, match chan string, wg *sync.WaitGroup) {
	fmt.Printf("[g:%6s] ", name)
	select {
	case peer := <-match:
		fmt.Printf("%s sent a message to %s.\n", peer, name)
	case match <- name:
		// 等待某个goroutine接收我的消息
		fmt.Println("Send")
	}
	wg.Done()
}
