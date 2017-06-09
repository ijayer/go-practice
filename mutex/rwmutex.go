package main

import (
	"math/rand"
	"sync"
	"time"
)

type DataStorage struct {
	sync.RWMutex
	Data int
}

func main() {
	d := DataStorage{sync.RWMutex{}, 0}
	for i := 0; i < 1; i++ {
		go d.Reader()
		go d.Writer(5)
		time.Sleep(100 * time.Microsecond)
		go d.Reader()
	}
	time.Sleep(1 * time.Second)
}

func (s *DataStorage) Reader() {
	println("#start reading...")
	for i := 0; i < 1; i++ {
		println("rlock.")
		s.RWMutex.RLock()

		println("#read ", s.Data)
		time.Sleep(1 * time.Microsecond)

		println("runlock.")
		s.RWMutex.RUnlock()
	}
	println("#end   reading...")
}

func (s *DataStorage) Writer(n int) {
	println("#start writing...")
	for i := 1; i < 2; i++ {
		s.RWMutex.Lock()
		println("lock.")

		s.Data = n
		println("#write ", i*rand.Intn(n))
		time.Sleep(1000 * time.Microsecond)

		println("unlock.")
		s.RWMutex.Unlock()
	}
	println("#end   writing...")
}
