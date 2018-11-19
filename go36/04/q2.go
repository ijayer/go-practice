/*
 * 说明：
 * 作者：zhe
 * 时间：2018-11-19 10:55 AM
 * 更新：
 */

package main

import "fmt"

func main() {
	// 示例1。
	a1 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%p, a1: %v (len: %d, cap: %d)\n", &a1, a1, len(a1), cap(a1))

	s9 := a1[1:4]
	// s9[0] = 1
	fmt.Printf("s9: %v (len: %d, cap: %d)\n", s9, len(s9), cap(s9))

	for i := 1; i <= 5; i++ {
		s9 = append(s9, i)
		fmt.Printf("s9(%d): %v (len: %d, cap: %d)\n", i, s9, len(s9), cap(s9))
	}

	fmt.Printf("%p, a1: %v (len: %d, cap: %d)\n", &a1, a1, len(a1), cap(a1))
	fmt.Println()
}
