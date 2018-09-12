/*
 * 说明：
 * 作者：zhe
 * 时间：2018-03-16 09:20
 * 更新：
 */

package main

import "fmt"

func main() {
	arr2 := [5]int{1, 2}            // 1, 2, 0, 0, 0
	arr5 := arr2                    // 1, 2, 0, 0, 0
	fmt.Printf("arr2: %p\n", &arr2) // 0xc042072030
	fmt.Printf("arr5: %p\n", &arr5) // 0xc042072030

	arr5[0] = 5 // 5, 2, 0, 0, 0
	arr2[4] = 2 // 1, 2, 0, 0, 2
	fmt.Printf("arr5= %d \narr2= %d \narr5[0]==arr2[0]= %v \n", arr5, arr2, arr5[0] == arr2[0])

	slice3 := []string{"a", "b", "c"}
	fmt.Println("call before: ", slice3) // a, b, c
	sliceAsParam(slice3)
	fmt.Println("call  after: ", slice3) // d, b, c

	fmt.Println("call before: ", arr2) // [1 2 0 0 2]
	arrayAsParam(arr2)
	fmt.Println("call  after: ", arr2) // [1 2 0 0 2]
}

func sliceAsParam(s []string) {
	s[0] = "d"
}

func arrayAsParam(a [5]int) {
	a[0] = 5
}
