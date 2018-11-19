/*
 * 说明：『Go核心36讲』| 07 | 数组和切片 Demo
 * 作者：zhe
 * 时间：2018-11-17 4:19 PM
 * 更新：
 */

package main

import "fmt"

// 问题：怎样正确估算切片的长度和容量？

func main() {
	// Demo 1
	// 初始化切片 s1, 指定其大小 5，但没有指定其容量，此时
	// make 在初始化过程中会指定容量等于其长度，即 5；
	s1 := make([]int, 5)
	fmt.Printf("s1 length: \t%d\n", len(s1))
	fmt.Printf("s1 cap: \t%d\n", cap(s1))
	fmt.Printf("s1 value: \t%d\n", s1)
	// output
	// s1 length:      5
	// s1 cap:         5
	// s1 value:       [0 0 0 0 0]

	println()

	// Demo2
	// 初始化切片s2, 指定其大小 5，容量 8, 即底层数组容量也
	// 将是 8
	//
	// 由于 切片 s2 的长度为 5，所以你可以操作的对应的底层数
	// 组元素也将是第 1 个元素到第 5 个元素，即索引 [0,4]；
	// 且 s2中的索引从 0 到 4 所指向的元素恰恰就是其底层数组
	// 中索引从 0 到 4 代表的那 5 个元素
	s2 := make([]int, 5, 8)
	fmt.Printf("s2 length: \t%d\n", len(s2))
	fmt.Printf("s2 cap: \t%d\n", cap(s2))
	fmt.Printf("s2 value: \t%d\n", s2)
	// output
	// s2 length:      5
	// s2 cap:         8
	// s2 value:       [0 0 0 0 0]

	// s2 := make([]int, 5, 3)      // len larger than cap in make([]int)
	println()

	// Demo 3
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("s3 length: \t%d\n", len(s3))
	fmt.Printf("s3 cap: \t%d\n", cap(s3))
	fmt.Printf("s3 value: \t%d\n", s3)
	// output
	// s3 length:      8
	// s3 cap:         8
	// s3 value:       [1 2 3 4 5 6 7 8]

	println()

	// Demo 4
	// s4 是在 s3 上施加切片操作得来的，所以 s3 的底层数组就
	// 是 s4 的底层数组。
	s4 := s3[3:6] // [3, 6)
	fmt.Printf("s4 length: \t%d\n", len(s4))
	fmt.Printf("s4 cap: \t%d\n", cap(s4))
	fmt.Printf("s4 value: \t%d\n", s4)
	println()
	// output:
	// s4 length:      3
	// s4 cap:         5
	// s4 value:       [4 5 6]
	//
	// 在底层数组不变的情况下，切片代表的窗口可以向右扩展，直
	// 至其底层数组的末尾，所以 s4 的容量就是底层数组的长度 8
	// 减去切片 s4 的起始索引 3，即 5.
	//
	s4 = append(s4, 12)
	fmt.Printf("s4 length: \t%d\n", len(s4))
	fmt.Printf("s4 cap: \t%d\n", cap(s4))
	fmt.Printf("s4 value: \t%d\n", s4)
	// Note: 切片代表的窗口是无法向左扩展的, 也就是说，我们永远
	// 无法透过 s4 看到 s3 中最左边的那 3 个元素
}
