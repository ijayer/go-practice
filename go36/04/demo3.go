/*
 * 说明：『Go核心36讲』| 07 | 数组和切片 Demo
 * 作者：zhe
 * 时间：2018-11-19 11:25 AM
 * 更新：
 */

package main

import "fmt"

func main() {
	a := [6]int{1, 2, 3, 4, 5, 6}
	s2 := a[1:4]      // s2 的底层数组：a
	s3 := a[2:cap(a)] // s3 的底层数组：a
	fmt.Printf("aa: value: %v, \tlen: %v, cap: %v\n", a, len(a), cap(a))
	fmt.Printf("s2: value: %v, \t\tlen: %v, cap: %v\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: value: %v, \t\tlen: %v, cap: %v\n", s3, len(s3), cap(s3))
	println()

	s2[2] = 22 // 影响 s3, a 的元素
	fmt.Printf("aa: value: %v, \tlen: %v, cap: %v\n", a, len(a), cap(a))
	fmt.Printf("s2: value: %v, \t\tlen: %v, cap: %v\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: value: %v, \t\tlen: %v, cap: %v\n", s3, len(s3), cap(s3))
	println()

	s2 = append(s2, 11, 22, 33, 44, 55, 66) // s2 的底层数组有改变，不再是 a
	fmt.Printf("aa: value: %v, \tlen: %v, cap: %v\n", a, len(a), cap(a))
	fmt.Printf("s2: value: %v, \t\tlen: %v, cap: %v\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: value: %v, \t\tlen: %v, cap: %v\n", s3, len(s3), cap(s3))
	println()

	s2[2] = 99 // 不影响 s3, a 的元素
	fmt.Printf("aa: value: %v, \tlen: %v, cap: %v\n", a, len(a), cap(a))
	fmt.Printf("s2: value: %v, \t\tlen: %v, cap: %v\n", s2, len(s2), cap(s2))
	fmt.Printf("s3: value: %v, \t\tlen: %v, cap: %v\n", s3, len(s3), cap(s3))
	println()
}
