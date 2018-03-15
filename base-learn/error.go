package main

import (
	//"errors"
	"fmt"
)

type errorString struct {
	s string
}

func main() {
	f, err := Sqrt(-1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
}

func (e *errorString) Error() string {
	return e.s
}

func New(text string) error {
	return &errorString{text}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, New("math: square root of negative number")
	}
	return f * f, nil
}
