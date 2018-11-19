/*
 * 说明：『Go核心36讲』| 02 & 03 | *源码文件 Demo
 * 作者：zhe
 * 时间：2018-09-16 9:48 PM
 * 更新：
 */

package main

import (
	"fmt"
)

func main() {
	i := ii(1)
	fmt.Printf("main: %v, %p\n", i, &i)

	if true {
		i := ii(2)
		fmt.Printf("InIf：%v, %p\n", i, &i)
	}

	fmt.Printf("main: %v, %p\n", i, &i)
}

func ii(i int) int {
	return i
}
