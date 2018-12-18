/*
 * 说明：https://blog.csdn.net/yunlilang/article/details/78458812
 * 作者：zhe
 * 时间：2018-12-04 9:56 PM
 * 更新：测试 QQ 群里的一个问题
 */

package main

import "fmt"

func main() {
	var s []int

	s = append(s, 0)
	printlnSlice(s)

	s = append(s, 1)
	printlnSlice(s)

	s = append(s, 2, 3, 4, 5, 7, 8, 9)
	printlnSlice(s)
	// Note: 解释这里 cap=10 的原因
	// 另外，如果我们一次追加的元素过多，以至于使新长度比原容量的 2 倍还要大，
	// 那么新容量就会以新长度为基准(Note: 有时会出现与预期计算的不相等的情况，因为
	// 底层算法还要考虑到按字节对齐)。

	// Output:
	// len:  1, cap:  1, ptr: 0xc000048400, val: [0]
	// len:  2, cap:  2, ptr: 0xc000048460, val: [0 1]
	// len:  9, cap: 10, ptr: 0xc0000484a0, val: [0 1 2 3 4 5 7 8 9]
}

func printlnSlice(s []int) {
	fmt.Printf("len: %2d, cap: %2d, ptr: %p, val: %v\n", len(s), cap(s), &s, s)
}
