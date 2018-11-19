/*
 * 说明：『Go核心36讲』| 07 | 数组和切片 Demo
 * 作者：zhe
 * 时间：2018-11-18 2:47 PM
 * 更新：
 */

package main

import "fmt"

// 问题：切片的底层数组什么时候会被替换？

func main() {
	// testAppend1()
	// testAppend2()
	// testAppend3()
}

func testAppend1() {
	// 初始切片
	s1 := make([]int, 3, 5)
	fmt.Printf("s1 length: \t%d\n", len(s1))
	fmt.Printf("s1 cap: \t%d\n", cap(s1))
	fmt.Printf("s1 value: \t%d\n", s1)
	fmt.Printf("s1 adds: \t%p\n", s1)

	// 添加 1 个元素，容量不超过初始化时设定值，即 5.
	s1 = append(s1, 1)
	println("添加新元素...")
	fmt.Printf("s1 length: \t%d\n", len(s1))
	fmt.Printf("s1 cap: \t%d\n", cap(s1))
	fmt.Printf("s1 value: \t%d\n", s1)
	fmt.Printf("s1 adds: \t%p\n", s1)

	// output:
	// s1 length:      3
	// s1 cap:         5
	// s1 value:       [0 0 0]
	// s1 adds:        0xc00007a030
	// 扩容...
	// s1 length:      4
	// s1 cap:         5
	// s1 value:       [0 0 0 1]
	// s1 adds:        0xc00007a030
	// Note: 无需扩容时，append() 函数返回的是指向底层数组的新切片

	// 扩容：添加 3 个元素，容量超过初始化时的设定值，即 5, 但小于其
	// 容量的 2 倍
	s1 = append(s1, 2, 3, 4)
	println("扩容...")
	fmt.Printf("s1 length: \t%d\n", len(s1))
	fmt.Printf("s1 cap: \t%d\n", cap(s1))
	fmt.Printf("s1 value: \t%d\n", s1)
	fmt.Printf("s1 adds: \t%p\n", s1)
	// 扩容...
	// s1 length:      7
	// s1 cap:         10
	// s1 value:       [0 0 0 1 2 3 4]
	// s1 adds:        0xc0000820a0
	// Note: 在扩容后，append() 返回的是指向新底层数组的新切片
}

func testAppend2() {
	// 初始切片
	s1 := make([]int, 3, 5)
	fmt.Printf("s1 length: \t%d\n", len(s1))
	fmt.Printf("s1 cap: \t%d\n", cap(s1))
	fmt.Printf("s1 value: \t%d\n", s1)
	fmt.Printf("s1 adds: \t%p\n", s1)

	// 扩容1: 添加元素后，切片的总长度大于原来容量的 2 倍
	println("扩容...")
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("s1 length: \t%d\n", len(s1))
	fmt.Printf("s1 cap: \t%d\n", cap(s1))
	fmt.Printf("s1 value: \t%d\n", s1)
	fmt.Printf("s1 adds: \t%p\n", s1)

	// output:
	// s1 length:      3
	// s1 cap:         5
	// s1 value:       [0 0 0]
	// s1 adds:        0xc00007a030
	// 扩容...
	// s1 length:      12
	// s1 cap:         12
	// s1 value:       [0 0 0 1 2 3 4 5 6 7 8 9]
	// s1 adds:        0xc000046060
	// Note: 另外，如果我们一次追加的元素过多，以至于使新长度比原容量
	// 的 2 倍还要大，那么新容量就会以新长度为基准。
}

func testAppend3() {
	// 初始切片
	s1 := make([]int, 3, 5)
	fmt.Printf("s1 length: \t%d\n", len(s1))
	fmt.Printf("s1 cap: \t%d\n", cap(s1))
	fmt.Printf("s1 value: \t%d\n", s1)
	fmt.Printf("s1 adds: \t%p\n", s1)

	// 扩容1: 添加元素后，切片的总长度等于原来容量的 2 倍
	println("扩容...")
	s1 = append(s1, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("s1 length: \t%d\n", len(s1))
	fmt.Printf("s1 cap: \t%d\n", cap(s1))
	fmt.Printf("s1 value: \t%d\n", s1)
	fmt.Printf("s1 adds: \t%p\n", s1)
}
