package main

import (
	"fmt"
	"math"
)

type RectAngel struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// 申明一个方法，用结构体类型作为函数的接收者
func (r RectAngel) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := RectAngel{12, 2}
	c1 := Circle{3}

	fmt.Println("Area of r1 is: ", r1.area())
	fmt.Println("Area of c1 is: ", c1.area())
}
