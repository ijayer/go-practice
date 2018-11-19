/*
 * 说明：
 * 作者：zhe
 * 时间：2018-09-13 9:51 PM
 * 更新：
 */

package lib

import "go-puzzlers/demo_1/q2/lib/internal" // Use of internal package

func Hi(s string) {
	println(s)
}

func CallInternalPkg() {
	println(internal.Internal)
	println(internal.II) // internal 包, 也遵循大写字母为：公开，小写字母为：包级私有
}
