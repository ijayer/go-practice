/*
 * 说明：
 * 作者：zhe
 * 时间：2018-06-27 3:29 PM
 * 更新：
 */

package main

import "C"
import "fmt"

// export SayHello
// SayHello 将Go语言实现的函数SayHello导出为C语言函数
func SayHello(s *C.char) {
	fmt.Println(C.GoString(s)) // s 转化为 go 的 string 类型输出
}
