/*
 * 说明：在 Go 文件中声明一个带有特殊类型 _GoString_ 类型的 C 函数。
 * 作者：zhe
 * 时间：2018-07-23 2:48 PM
 * 更新：
 */

package main

/*
size_t _GoStringLen(_GoString_ s);
const char *_GoStringPtr(_GoString_ s);
*/
import "C"
import "fmt"

func main() {
	var str = "hello, world"
	fmt.Println("len from Go func: ", len(str))
	fmt.Println("ptr from Go func: ", &str)
	fmt.Println()

	fmt.Println("len from C  func: ", C._GoStringLen(str))
	fmt.Println("ptr from C  func: ", C._GoStringPtr(str))
}
