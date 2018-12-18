/*
 * 说明：
 * 作者：zhe
 * 时间：2018-12-05 9:31 AM
 * 更新：
 */

package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 创建一个环, 包含 3 个元素
	r := ring.New(3)
	fmt.Printf("var  ring: %+v\n\n", *r)

	// 初始化
	for i := 1; i <= 3; i++ {
		r.Value = i
		r = r.Next()
	}
	fmt.Printf("init ring: %+v\n\n", *r)

	// println element
	fmt.Println("printlnRingElement")
	printlnRingElement(r)
	println()

	// sum
	s := sumRingElement(r)
	fmt.Printf("sum  ring: %d\n\n", s)

	// range ring
	fmt.Println("rangeRing")
	rangeRing(r)
	println()
}

// printlnRingElement
func printlnRingElement(r *ring.Ring) {
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})
}

// sumRingElement
func sumRingElement(r *ring.Ring) int {
	s := 0
	r.Do(func(i interface{}) {
		s += i.(int)
	})
	return s
}

// rangeRing
func rangeRing(r *ring.Ring) {
	p := r.Next()
	fmt.Println(p.Value)

	for p != r {
		p = p.Next()
		fmt.Println(p.Value)
	}
}
