/*
 * 说明：
 * 作者：zhe
 * 时间：2018-06-26 3:27 PM
 * 更新：
 */

package main

/*
#include <stdio.h>

void SayHello(char* str) {
	puts(str);
}
*/
import "C"

func main() {
	CallCFuncWithCustomized()
}

// CallCFuncWithCustomized 使用自己的 C 函数
// 此时，CGO 域需要实现 SayHello 函数
func CallCFuncWithCustomized() {
	C.SayHello(C.CString("hello cgo"))
	// output:
	// $ go run hello-cgo.go
	// hello cgo
}
