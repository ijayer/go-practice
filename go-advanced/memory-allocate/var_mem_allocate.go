/*
 * 说明：Go 逃逸分析
 * 作者：zhe
 * 时间：2018-12-06 9:37 AM
 * 更新：
 */

package main

import "fmt"

func main() {
	i := 1
	fmt.Printf("i  addr: %p\n", &i)

	r1 := plus(i)
	fmt.Printf("r1 addr: %p\n", &r1)
}

func plus(i0 int) *int {
	var i1 = 11
	var i2 = 22
	var i3 = 33
	var i4 = 44

	fmt.Printf("i0 addr: %p\n", &i0)
	fmt.Printf("i1 addr: %p\n", &i1)
	fmt.Printf("i2 addr: %p\n", &i2)
	fmt.Printf("i3 addr: %p\n", &i3)
	fmt.Printf("i4 addr: %p\n", &i4)

	return &i2
}
