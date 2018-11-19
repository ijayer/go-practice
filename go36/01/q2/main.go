/*
 * 说明：『Go核心36讲』| 02 & 03 | *源码文件 Demo
 * 作者：zhe
 * 时间：2018-09-13 9:50 PM
 * 更新：
 */

package main

import (
	"go36/01/q2/lib"
	"go36/01/q2/lib/internal" // Note: Use of internal package is not allowed.
)

func main() {
	println(internal.Internal)
	println(lib.CallInternalPkg)
}
