/*
 * 说明：
 * 作者：zhe
 * 时间：2018-06-18 22:39
 * 更新：
 */

package main

import (
	"log"
	"os"
	"sync"
)

func main() {
	ch := ParallelWrite([]byte("hello"))
	select {
	case err, ok := <-ch:
		if !ok {
			log.Println(err)
		}
	default:

	}
}

// race 两个go程数据竞争的demo
func race() {
	wait := make(chan struct{})
	n := 0

	go func() {
		n++ // func.1 goroutine: 一次访问要执行的操作：读，递增，写
		wait <- struct{}{}
	}()

	n++ // main goroutine: 访问冲突
	<-wait
	log.Println("n =", n)
}

// shareIsCaring 通过通信(channel)来共享内存
func shareIsCaring() {
	ch := make(chan int)

	go func() {
		n := 0 // 仅为一个 goroutine 可见的局部变量
		n++
		ch <- n // 数据从当前 goroutine 离开
	}()

	n := <-ch // 数据安全的到达了另一个 goroutine
	n++
	log.Println(n) // 输出：2
}

// raceOnLoopCounter
// https://tip.golang.org/doc/articles/race_detector.html#Race_on_loop_counter
func raceOnLoopCounter() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ { // i 在 for 循环迭代中会被重用
		go func() {
			log.Printf("[g]: p: %v, v: %v\n", &i, i) // 6个goroutine共享变量i
			wg.Done()
		}()
	}

	wg.Wait() // 阻塞等待所有goroutine执行结束，go程被调度执行
	println()
	// output:
	// 2018/06/21 15:40:05 [g]: p: 0xc042056058, v: 1
	// 2018/06/21 15:40:05 [g]: p: 0xc042056058, v: 5
	// 2018/06/21 15:40:05 [g]: p: 0xc042056058, v: 5
	// 2018/06/21 15:40:05 [g]: p: 0xc042056058, v: 5
	// 2018/06/21 15:40:05 [g]: p: 0xc042056058, v: 2
	// 输出结果不确定
}

// fixDataRace 改进 dataRace 函数
func fixRaceOnLoopCounter() {
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ { // i 在 for 循环迭代中会被重用
		go func(n int) { // i 作为参数传入闭包中
			log.Printf("[g]: p: %v, v: %v\n", &n, n)
			wg.Done()
		}(i)
	}

	wg.Wait() // 阻塞等待所有goroutine执行结束，go程被调度执行

	// Output:
	// 2018/06/21 15:40:05 [g]: p: 0xc0420100a0, v: 4
	// 2018/06/21 15:40:05 [g]: p: 0xc042092010, v: 1
	// 2018/06/21 15:40:05 [g]: p: 0xc042092020, v: 0
	// 2018/06/21 15:40:05 [g]: p: 0xc042092028, v: 2
	// 2018/06/21 15:40:05 [g]: p: 0xc042092038, v: 3
	// 一定会输出：01234 这几个数字，但其顺序不确定
}

// ParallelWrite 把 data 写入 file1 和 file2，返回 error
// https://tip.golang.org/doc/articles/race_detector.html#Accidentally_shared_variable
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)

	f1, err := os.Create("file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f1.Write(data) // 这里的 err 是和 main goroutine 共享的，所以下面的写操作会发生数据竞争
			if err != nil {
				res <- err
			}
			f1.Close()
		}()
	}

	f2, err := os.Create("file2") // 第二次写入 err 冲突
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data) // 这里的 err 是和 main goroutine 共享的，所以下面的写操作会发生数据竞争
			if err != nil {
				res <- err
			}
			f2.Close()
		}()
	}
	return res
}
