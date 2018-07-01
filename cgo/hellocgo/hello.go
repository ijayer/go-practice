/*
 * 说明：
 * 作者：zhe
 * 时间：2018-06-26 4:17 PM
 * 更新：
 */

package main

/*
#include <stdio.h>

void SayHello(char* str);
*/
import "C"

func main() {
	C.SayHello(C.CString("hello cgo"))
}
