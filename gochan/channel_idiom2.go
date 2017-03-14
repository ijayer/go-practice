package main

import (
	"fmt"
	"instance.golang.com/utils"
	"time"
)

func main() {
	suck1(pump1())
	time.Sleep(10 * time.Second)
}

// send channel
func pump1() chan string {
	ch := make(chan string)
	go func() {
		for i:=0; ; i++ {
			ch <- utils.Now()
			time.Sleep(1000 * time.Millisecond)
		}
	}()
	return ch
}


// recv channel
func suck1(ch chan string) {
	go func() {
		for v := range ch {
			fmt.Printf("#_________ch = %v\n", v)
		}
	}()
}
