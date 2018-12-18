/*
 * 说明：https://zhezh09.github.io//post/tech/code/golang/20181212-go语言机制-01-栈和指针/
 * 作者：zhe
 * 时间：2018-12-12 1:46 PM
 * 更新：
 */

package main

func main() {
	// Declare variable of type int with a value of 10
	count := 10

	// Display the "value of" and "address of" count.
	println("count:\tValue of[", count, "]\tAddr of [", &count, "]")

	// Pass the "address of" the count.
	increment2(&count)
	println("count:\tValue of[", count, "]\tAddr of [", &count, "]")

	// Output:
	// count:  Value of[ 10 ]  Addr of [ 0xc000035f80 ]
	// inc:    Value of[ 11 ]  Addr of [ 0xc000035f80 ]        Value Points To[ 11 ]
	// count:  Value of[ 11 ]  Addr of [ 0xc000035f80 ]
}

//go:noinline
func increment2(inc *int) {
	*inc++ // Increment the "value of" inc.
	println("inc:\tValue of[", *inc, "]\tAddr of [", inc, "]\tValue Points To[", *inc, "]")
}
