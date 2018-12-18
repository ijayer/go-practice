/*
 * 说明：『Go核心36讲』| 08 | container包中的那些容器 Demo
 * 作者：zhe
 * 时间：2018-11-30 9:15 AM
 * 更新：
 */

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 声明一个链表
	var l = list.New()
	fmt.Printf("list: %+v\n", *l)

	// 在链表头部或尾部插入元素
	e0 := l.PushBack(42)
	e1 := l.PushFront(13)
	e2 := l.PushBack(7)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("|%4d| ", e.Value.(int))
	}
	println()

	// 在某个元素前或后插入元素
	l.InsertBefore(3, e0)
	l.InsertAfter(196, e1)
	l.InsertAfter(1729, e2)

	// 循环打印
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("|%4d| ", e.Value.(int))
	}
	println()
}
