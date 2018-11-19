/*
 * 说明：『Go核心36讲』| 04 | 程序实体的那些事儿（上） Demo
 * 作者：zhe
 * 时间：2018-11-08 9:39 AM
 * 更新：
 */

package main

import . "fmt"

// import . "flag"
//
// var ErrHelp string // Error: ErrHelp redeclared

var a = 0 // a 属于 main 包代码块

func main() {
	a := 5 // int: a 属于 main 函数代码块
	Printf("a: %p, %v\n", &a, a)

	if true {
		i, a := return1and2() // int8, int8
		Printf("a: %p, %v\n", &a, a)
		Printf("i: %p, %v\n", &i, i)
		// output: i 声明新变量；a 声明新变量(a 属于 if 语句的代码块
		// 中, 且 a 与 main 函数代码块中的变量重名，此时并没有覆盖外
		// 层代码块中 a 的值，而是属于新变量)
		// a: 0xc0000500a8, 2
		// i: 0xc0000500a0, 1

		for true {
			i := 4.4
			Printf("i: %p, %v\n", &i, i)
			// output: i 声明新变量——可重名变量(此时 i 属于 for 语句
			// 的代码块中，与上级代码块中的 i 重名，但属于不同的变量，
			// 类型也不相同)
			// i: 0xc0000500c0, 4
			break
		}

		// var a = 3 // Error: 'a' redeclared

		Printf("i: %p, %v\n", &i, i)
	}
	Printf("a: %p, %v\n", &a, a)

	j, a := return6and8() // 对同一代码块中的变量 a 进行复用(变量重声明，
	// 类型必须为 int,否则报错)
	Printf("a: %p, %v\n", &a, a)
	Printf("j: %p, %v\n", &j, j)
	// output: a 变量复用；j 声明新变量
	// a: 0xc000050058, 5
	// a: 0xc000050058, 8
	// j: 0xc0000500a0, 6
}

func return1and2() (int8, int8) {
	return 1, 2
}

func return6and8() (int, int) {
	return 6, 8
}
