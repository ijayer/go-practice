/*
 * 说明：『Go核心36讲』| 04 | 程序实体的那些事儿（下） Demo
 * 作者：zhe
 * 时间：2018-11-16 10:22 AM
 * 更新：
 */

package main

import (
	"fmt"
	"reflect"
)

// MyString 定义为 string 类型的别名
type MyString = string

// MyString2 属于类型再定义，其潜在类型为 string 类型
type MyString2 string

func main() {
	printInvalidUnicode()
	println()

	testTypeConvert()
	println()

	printChineseAndCharacterLen()
	println()

	testTypeAlias()
}

func printInvalidUnicode() {
	i := -1
	println(string(i))
	// output:
	// �
}

func printChineseAndCharacterLen() {
	var s = "你"
	var c = "n"

	println(len(s))
	println(len(c))

	fmt.Printf("%v\n", []byte(s))
	fmt.Printf("%v\n", []byte(c))
}

func testTypeConvert() {
	i := uint8(0)
	println(i)

	j := uint8(255)
	println(j)

	k := int8(-128)
	println(k)

	o := int8(127)
	println(o)

	var srcInt = int16(-255) // 原[-255]: 10000000 11111111 | 反[-255]：11111111 00000000 | 补[-255]：11111111 00000001
	println(srcInt)          // -255
	dstInt := int8(srcInt)   // srcInt 被强转为 int8 后，Go 会把在较高位置(或者说最左边位置)上的 8 位二进制数直接截掉
	println(dstInt)          // 即得到：0000 0001, 符号为 0 表示正数，由于正数的原、反、补码相同，故得到 1
}

func testTypeAlias() {
	var s MyString
	var ss MyString2

	println("s:", reflect.TypeOf(s).String())
	println("ss", reflect.TypeOf(ss).String())

	// output:
	// string
	// main.MyString2

	s = "s"
	ss = "ss"

	// s == ss // Invalid operation: s == ss (mismatched types MyString and MyString2)
	//
	// s = ss // Cannot use 'ss' (type MyString2) as type MyString in assignment
	// ss = s // Cannot use 's' (type MyString) as type MyString2 in assignment

	s = MyString(ss)
	ss = MyString2(s)

	println("s:", reflect.TypeOf(s).String())
	println("ss", reflect.TypeOf(ss).String())
}
