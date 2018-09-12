package main

import (
	"fmt"
)

type stockPosition struct {
	tracker    string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

// showValue.1
func showValue(v valuable) {
	fmt.Printf("the value is %v\n", v.getValue())
}

// showValue.2
// refactor showValue
func showValue2(i interface{}) {
	switch v := i.(type) {
	case stockPosition:
		fmt.Printf("the %v's %v value is %v$\n", v.tracker, "stock", v.getValue())
	case car:
		fmt.Printf("the %v's %v value is %v$\n", v.make, v.model, v.getValue())
	default:
		fmt.Printf("error: unknow type '%v'\n", v)
		return
	}
}

func main() {
	sp := stockPosition{"Google", 99.99, 0.8}
	showValue2(sp)

	c := car{"OOOO", "A8L", 9999}
	showValue2(c)

	showValue2("balabala")
}
