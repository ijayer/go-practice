// producter & consumer
package main

import "fmt"

func main() {
	var chnum chan int = make(chan int)
	var done chan bool = make(chan bool)

	go producter(0, 1000, chnum)
	go consumer(chnum, done)

	<-done
}

// integer producter
func producter(start, end int, out chan int) {
	for i := 0; i < end; i++ {
		out <- start
		start += 10
	}
	close(out)
}

// integer consumer
func consumer(in chan int, done chan bool) {
	for num := range in {
		fmt.Printf("#______consumer recv: %v\n", num)
	}
	done <- true
}
