/*
 * 说明：
 * 作者：zhe
 * 时间：2018-03-16 11:32
 * 更新：
 */

package main

// 不能为内建类型声明方法
//func (m int) add() { // error
//
//}

// 内建类型声明方法可以为内建类型定义别名
type myInt int

func (m myInt) add() {

}
