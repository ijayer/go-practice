/*
 * 说明：『Go核心36讲』| 02 & 03 | *源码文件 Demo
 * 作者：zhe
 * 时间：2018-09-13 11:08 PM
 * 更新：
 */

package main

import (
	"fmt"
	"io"
	"os"
)

func Say() {
	var err error
	fmt.Printf("err: %p\n", err)

	n, err := io.WriteString(os.Stderr, "hello")
	fmt.Printf("err: %p\n", err)
	fmt.Printf("n: %v\n", n)
}
