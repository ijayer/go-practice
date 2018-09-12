package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer func() {
			wg.Done()
			println("#Goroutine 1 end.")
		}()
		time.Sleep(1 * time.Second)
		println("#Goroutine 1 starting...")
	}()

	go func() {
		defer func() {
			wg.Done()
			println("#Goroutine 2 end.")
		}()
		time.Sleep(1 * time.Second)
		println("#Goroutine 2 starting...")
	}()

	wg.Add(-1)
	wg.Wait()
}
