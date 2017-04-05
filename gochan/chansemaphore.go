// 协程中使用通道来输出结果
package main

import (
	"fmt"
)

func main() {
	fmt.Println("#_________main() is running")

	ch := make(chan uint64)
	go sum(ch)
	fmt.Println("#_________waiting for the result")

	result := <-ch
	fmt.Printf("#_________result = %v\n", result)
}

func sum(ch chan uint64) {
	var sum uint64 = 0
	var i uint64 = 0

	for i = 0; i < 10000000000; i++ {
		sum += i
	}
	ch <- sum
}
