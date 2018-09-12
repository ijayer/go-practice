/*
 * 说明：方法的接收者为值类型和指针类型会有什么变化？
 * 作者：zhe
 * 时间：2018-04-08 10:30
 * 更新：
 */

package main

import "fmt"

type Win struct {
	Large int
	Width int
}

// 方法定义
func (w Win) Set() {
	w.SetWidth()
}

func (w *Win) SetWidth() {
	w.Width = 2
}

// 等价函数
func Set(w Win) {
	SetWidth(&w)
}

func SetWidth(w *Win) {
	w.Width = 2
}

func main() {
	w := &Win{}

	w.Set() // 方法调用
	fmt.Printf("w: %+v\n", w)

	Set(*w) // 函数调用
	fmt.Printf("w: %+v\n", w)
}
