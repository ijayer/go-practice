/*
 * 说明：『Go核心36讲』| 04 | 程序实体的那些事儿（中） Demo
 * 作者：zhe
 * 时间：2018-11-08 2:31 PM
 * 更新：
 */

package main

import (
	"fmt"
	"reflect"
)

var container = []string{"zero", "one", "two"} // main 包

func main() {
	println("> main func")
	container := map[int]string{0: "zero", 1: "one", 2: "two"} // main 函数

	fmt.Printf("The element is %q. \n", container[1])
	fmt.Printf("The container's type is %q. \n", reflect.TypeOf(container).String())

	{
		println("> inner")
		container := map[string]string{"1": "one"}
		fmt.Printf("The element is %q. \n", container["1"])
		fmt.Printf("The container's type is %q. \n", reflect.TypeOf(container).String())
	}
}
