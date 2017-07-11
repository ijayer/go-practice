package main

import "fmt"

type Simpler interface {
	Get() int
	Set(i int)
}

type Simple struct {
	Value int
}

func (s *Simple) Get() int {
	return s.Value
}

func (s *Simple) Set(i int) {
	s.Value = i
}

func main() {
	s := new(Simple)
	s.Value = 1

	fmt.Println(s.Get())
	s.Set(2)
	fmt.Println(s.Get())
}
