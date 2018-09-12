package main

import "fmt"

type TZ int //使用TZ来操作int类型的数据

func main() {
	var a, b TZ = 3, 4
	c := a + b
	fmt.Printf("c = %d \n", c)
}
