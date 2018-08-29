/*
 * 说明：理解 Go 的 nil
 * 作者：zhe
 * 时间：2018-08-23 5:32 PM
 * 更新：
 */

package main

import (
	"fmt"
)

func main() {
	nilOfChannels()
}

func untypedZero() {
	a := false
	fmt.Printf("%T\n", a)

	b := ""
	fmt.Printf("%T\n", b)

	c := 0
	fmt.Printf("%T\n", c)

	d := 0.0
	fmt.Printf("%T\n", d)
}

// ----------------------------------------- pointers
func pointers() {
	var p *int
	fmt.Printf("%v\n", p)

	println(p == nil) // equal to nil
	println(*p)       // dereference
}

// ----------------------------------------- channels

func nilOfChannels() {
	var c chan int

	go func() {
		select {
		case v, ok := <-c:
			fmt.Printf("v: %v, ok:%v\n", v, ok)
		default:

		}
	}()

	c <- 2
	close(c)
}

// ----------------------------------------- interfaces
type doError struct {
	msg string
}

func (d *doError) Error() string {
	return d.msg
}

func do() error {
	var e *doError
	fmt.Printf("type: %T, value: %v\n", e, e)

	return e
}

func nilOfInterfaces() {
	err := do()

	fmt.Printf("type: %T, value: %v\n", err, err)
	fmt.Printf("type: %T, value: %v\n", nil, nil)

	fmt.Println(err == nil)

	// output:
	// type: *main.doError, value: <nil>
	// type: *main.doError, value: <nil>
	// type: <nil>, value: <nil>
	// false
}
