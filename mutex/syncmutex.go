package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	sync.Mutex
	Num int
}

func (c *SafeCounter) Inc() {
	c.Lock()
	fmt.Println("#start inc...")
	c.Num++
	fmt.Println("#sleep 50us...")
	time.Sleep(50 * time.Microsecond)
	c.Unlock()
}

func (c *SafeCounter) Dec() {
	c.Lock()
	fmt.Println("#start dec...")
	c.Num--
	fmt.Println("#sleep 50us...")
	time.Sleep(50 * time.Microsecond)
	c.Unlock()
}

func (c SafeCounter) getValue() int {
	c.Lock()
	v := c.Num
	c.Unlock()
	return v
}

func main() {
	c := new(SafeCounter)
	for i := 0; i < 1000; i++ {
		go c.Inc()
		// go c.Dec()
	}
	time.Sleep(3 * time.Second)
	fmt.Printf("Result: value = %v\n", c.getValue())
}
