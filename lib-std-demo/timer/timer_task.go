package main

import (
	"fmt"
	"time"

	"github.com/zhezh09/go-practice/utils"
)

var code = make(map[int]string)

func main() {
	cacheCode(1, utils.RandNumMath())
	time.Sleep(1 * time.Second)

	cacheCode(2, utils.RandNumMath())
	time.Sleep(1 * time.Second)

	cacheCode(3, utils.RandNumMath())
	time.Sleep(20 * time.Second)
	fmt.Println("##______________main return")
}

func cacheCode(key int, value string) {
	code[key] = value
	fmt.Printf("##______________cached code=%v\n", code)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("##______________into timer task func")
		delete(code, key)
		fmt.Printf("##______________deleted code=%v\n", code)
	})
	fmt.Println("##______________cacheCode return")
}
