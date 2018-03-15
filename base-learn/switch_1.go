package main

import (
	"fmt"
)

func main() {
	var num1 = 12
	switch num1 {
	case 1, 2, 3:
		fmt.Println("第一季度")
	case 4, 5, 6:
		fmt.Println("第二季度")
	case 7, 8, 9:
		fmt.Println("第三季度") //fallthrough只执行后面的一个case语句
	case 10, 11, 12:
		fmt.Println("第四季度")
	default:
		fmt.Println("default case")
	}
}
