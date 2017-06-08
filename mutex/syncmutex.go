package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	sync.Mutex
	Num int
}

func (c *SafeCounter) Inc() {
	c.Lock()
	c.Num++
	c.Unlock()
}

func (c *SafeCounter) Dec() {
	c.Lock()
	c.Num--
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
		go c.Dec()
	}
	fmt.Printf("Result: value = %v\n", c.getValue())
}
