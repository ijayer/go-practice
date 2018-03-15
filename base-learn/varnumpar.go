package main

import "fmt"

//func main() {
//	x := Min(1, 3, 2, 0)
//	fmt.Printf("The mininum is： %d\n", x)
//	arr := []int{7, 9, 3, 5, 1}
//	x = Min(arr ...)
//	fmt.Printf("The mininum in the arr is: %d", x)
//}
//
//func Min(a ...int) int {
//	if len(a) == 0 {
//		return 0
//	}
//	min := a[0]
//	for _, v := range a {
//		if v < min {
//			min = v
//		}
//	}
//	return min
//}

func main() {
	num := Min(4, 5, 1, 9)
	fmt.Printf("The mininum is: %d\n", num)

	arr := []int{7, 4, 6, 1, 0, 9}
	num = Min(arr...)
	fmt.Printf("The mininum in arr is: %d", num)

}

func Min(arr ...int) int {
	if len(arr) == 0 {
		return 0
	}
	min := arr[0]

	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
