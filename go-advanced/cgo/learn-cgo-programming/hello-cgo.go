/*
 * 说明：
 * 作者：zhe
 * 时间：2018-06-26 3:27 PM
 * 更新：
 */

package main

/*
#include <stdio.h>
*/
import "C"
import "fmt"

func main() {
	CallCStdFunc()
}

// CallCStdFunc 调用 C 的标准库函数输出字符串
func CallCStdFunc() {
	C.puts(C.CString("hello cgo"))
	// output:
	// $ go run hello-cgo.go
	// hello cgo
}

// CVarTypes 测试变量的类型
func CVarTypes() {
	gStr := "hello cgo"
	cStr := C.CString(gStr)
	fmt.Printf("GType: %T, CType: %T\n", gStr, cStr)

	// output:
	// GType: string, CType: *main._Ctype_char
}
