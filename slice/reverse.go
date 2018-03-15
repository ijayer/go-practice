/*
 * 说明：
 * 作者：zhe
 * 时间：2018-03-15 14:29
 * 更新：
 */

package main

import "fmt"

func main() {
	a := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := a[0:5]
	fmt.Println("before call reverse()")
	fmt.Printf("a = %v\n", a)
	fmt.Printf("s = %v\n", s)

	s2 := reverse(s)
	fmt.Println("after call reverse()")
	fmt.Printf("s2 = %v\n", s2)

	a2 := [9]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("before call reverseArray()")
	fmt.Printf("a2 = %v\n", a2)

	a22 := reverseArray(a2)
	fmt.Println("after call reverseArray()")
	fmt.Printf("a22 = %v\n", a22)
}

// reverse 将切片的相邻元素逆转
func reverse(s []int) []int {
	for i := 0; i < len(s)-1; i++ {
		s[i], s[i+1] = s[i+1], s[i]
	}
	return s
}

// reverseArray 将数组的相邻元素逆转
func reverseArray(a [9]int) [9]int {
	for i := 0; i < len(a)-1; i++ {
		a[i], a[i+1] = a[i+1], a[i]
	}
	return a
}
