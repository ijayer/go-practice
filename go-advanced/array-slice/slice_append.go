/*
 * 说明：
 * 作者：zhe
 * 时间：2018-11-29 3:40 PM
 * 更新：
 */

package main

import "fmt"

func main() {
	s := make([]int, 6, 10)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[3] = 40
	s[4] = 50

	fmt.Println(s)

	s1 := s[1:4]
	fmt.Println(s1)

	s2 := s[2:6]
	fmt.Println(s2)

	s3 := append(s, 1, 2, 3) // 扩容后，底层数组没变
	fmt.Println(s3)
	println()

	s3[2] = 99      // s3[2] 元素的修改将会影响基于 s 创建的所有切片的与其对应的元素的值，这将会导致 bug
	fmt.Println(s3) // 共享原数组导致的 bug
	fmt.Println(s2) // 共享原数组导致的 bug
	fmt.Println(s1) // 共享原数组导致的 bug

	// Output:
	// [10 20 30 40 50 0]
	// [20 30 40]
	// [30 40 50 0]
	// [10 20 30 40 50 0 1 2 3]
	// [10 20 99 40 50 0 1 2 3]
	// [99 40 50 0]
	// [20 99 40]
}
