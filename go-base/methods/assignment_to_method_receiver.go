/*
 * 说明：
 * 作者：zhe
 * 时间：2018-07-26 9:43 AM
 * 更新：
 */

package main

import "fmt"

type Obj struct {
	Name string
}

func (o *Obj) do() error {
	// Warning：Assignment to method receiver propagates(传递)
	// only to callees(被调用者) but not to caller
	o = &Obj{Name: "obj_"}

	// Note:
	//
	// 假设有一个新的 Obj 对象
	//
	//      tmp := &Obj{Name: "temp"}
	//
	// 给方法接收者赋值，只能通过被调用的方式，即：
	//
	//      o.Name = tmp.Name
	//
	// 而不是直接赋值给调用者, 即：
	//
	//      o = tmp

	return nil
}

func main() {
	o := &Obj{}
	if err := o.do(); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("obj: %+v", o)
}
