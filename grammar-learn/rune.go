package main

import "fmt"

func main() {
	s := "Hello 世界！"

	r := []rune(s) // 转换为 []rune 自动复制数据

	r[6] = '中' // 修改 []rune

	r[7] = '国' // 修改 []rune

	fmt.Println(s) // s不能被修改，内容保持不变

	fmt.Println(string(r)) // 转换为字符串，有一次复制数据

}
