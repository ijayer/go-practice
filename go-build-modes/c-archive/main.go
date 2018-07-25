/*
 * 说明：
 * 作者：zhe
 * 时间：2018-07-24 9:38 AM
 * 更新：
 */

package main

import "C"
import (
	"fmt"

	"instance.golang.com/go-build-modes/c-archive/utils"
)

func main() {
	utils.Version()
	fmt.Println()
	Hello()
}

// export Hello
func Hello() {
	fmt.Println("hello, world!")
}
