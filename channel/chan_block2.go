package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan time.Time)
	go receive(ch)

	ch <- time.Now().UTC()
	time.Sleep(15 * time.Second)

	ch <- time.Now().UTC()
	time.Sleep(1 * time.Second)

	ch <- time.Now().UTC()
	time.Sleep(5 * time.Second)
}

func receive(ch chan time.Time) {
	for {
		fmt.Printf("#_________output=%v, time=%v\n", <-ch, time.Now().UTC())
	}
}
