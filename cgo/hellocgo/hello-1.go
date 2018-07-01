/*
 * 说明：
 * 作者：zhe
 * 时间：2018-06-26 4:17 PM
 * 更新：
 */

package main

// #include "hello.h"
import "C"

func main() {
	C.SayHello(C.CString("hello cgo")) // C.SayHello 调用的是 hello.h 中定义的 SayHello 函数
}
