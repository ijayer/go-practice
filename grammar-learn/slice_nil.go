/*
 * 说明：
 * 作者：zhe
 * 时间：2018-03-16 10:03
 * 更新：
 */

package main

import "fmt"

func main() {
	var s []string
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		println("nil slice!")
	}
}
