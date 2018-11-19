/*
 * 说明：切片缩容
 * 作者：zhe
 * 时间：2018-11-19 1:22 PM
 * 更新：
 */

package main

import "fmt"

func main() {
	s := make([]int, 100)

	s1 := s[:]
	fmt.Printf("s1: len: %v, cap: %v\n", len(s1), cap(s1))

	a := [50]int{}

	s1 = a[:]
	fmt.Printf("s1: len: %v, cap: %v\n", len(s1), cap(s1))
}
