/*
 * 说明：
 * 作者：zhe
 * 时间：2018-11-29 2:57 PM
 * 更新：
 */

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := varSlice()
	println(s1 == nil)

	s2 := makeSlice()
	println(s2 == nil)

	// Output:
	//   var s []int:
	//   addr=0x0
	//   len=0
	//   cap=0
	//   s == nil?true
	//   s := make([]int, 0):
	//   addr=0x5771c8
	//   len=0
	//   cap=0
	//   s == nil?false
}

// varSlice 声明一个 nil 切片
//      [ ptr,    len,    cap ]
//         |       |       |
//        0x0      0       0
func varSlice() []int {
	var s []int // nil 切片

	println("var s []int:")
	fmt.Printf("  addr=%p\n  len=%v\n  cap=%v\n  s == nil?", *(*[]byte)(unsafe.Pointer(&s)), len(s), cap(s))

	return s
}

// makeSlice 创建一个空切片，通过源码可知 makeslice 时会在底层申请内存，
// 即 slice.Ptr 会指向某一块内存空间
//
//      [ ptr,    len,    cap ]
//         |       |       |
//      0x5771c8   0       0
//
//      Source file src/runtime/slice.go
//
//      func makeslice(et *_type, len, cap int) slice {
//      	// NOTE: The len > maxElements check here is not strictly necessary,
//      	// but it produces a 'len out of range' error instead of a 'cap out of range' error
//      	// when someone does make([]T, bignumber). 'cap out of range' is true too,
//      	// but since the cap is only being supplied implicitly, saying len is clearer.
//      	// See issue 4085.
//      	maxElements := maxSliceCap(et.size)
//      	if len < 0 || uintptr(len) > maxElements {
//      		panicmakeslicelen()
//      	}
//
//      	if cap < len || uintptr(cap) > maxElements {
//      		panicmakeslicecap()
//      	}
//
//          // 申请一块内存
//      	p := mallocgc(et.size*uintptr(cap), et, true)
//      	return slice{p, len, cap}
//      }
func makeSlice() []int {
	s := make([]int, 0) // 空切片，

	println("s := make([]int, 0): ")
	fmt.Printf("  addr=%p\n  len=%v\n  cap=%v\n  s == nil?", *(*[]byte)(unsafe.Pointer(&s)), len(s), cap(s))

	return s
}
