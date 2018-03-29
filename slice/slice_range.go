/*
 * 说明：
 * 作者：zhe
 * 时间：2018-03-28 08:40
 * 更新：
 */

package main

import (
	"fmt"
	"runtime"
	"time"
)

type field struct {
	name string
}

// go 中的 method 本质上就是以 method 的 receiver 作为第一个参数的普通 function
// instance.method(x,y) <=> function(instance, x,y)

// print 是类型 *field 的一个方法
func (f *field) print() {
	fmt.Println(f.name)
}

// print1 是接受一个类型为 *field 作为参数的函数，与上面的方法是等价的
func print1(f *field) {
	fmt.Println(f.name)
}

func main() {
	runtime.GOMAXPROCS(1)

	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		fmt.Printf("data1：value: %v, addr: %p\n", v, v)
		go v.print()
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		fmt.Printf("data2：value: %v, addr: %p\n", v, &v)
		go v.print()
		go print1(&v)
	}

	// 上面启动的各个 child goroutine 在 main goroutine 执行到 Sleep 时才被调度执行
	time.Sleep(3 * time.Second)
}
