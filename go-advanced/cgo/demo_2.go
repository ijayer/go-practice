/*
 * 说明：cgo 目前还不支持调用 C 函数指针，不过你可以声明存放 C 函数指针的 Go 变量，
 *       并在 Go 和 C 之间传递。C 函数可以调用从 Go 中获取到的函数指针，例如：
 * 作者：zhe
 * 时间：2018-07-23 3:12 PM
 * 更新：
 */

package main

/*
typedef int(*intFunc) ();

int bridge_int_func(intFunc f)
{
    return f();
}

int fortyTwo()
{
    return 42;
}
*/
import "C"

import "fmt"

func main() {
	// 声明存放 C 函数指针的 Go 变量 fn
	fn := C.intFunc(C.fortyTwo)

	// 将表示 C 函数指针的 Go 变量传递给 C 函数
	fmt.Println(int(C.bridge_int_func(fn)))

	// 输出：42
}
