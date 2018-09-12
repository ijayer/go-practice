package main

import (
	"fmt"
	//	"strings"
)

func main() {
	var str = "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))

	for pos, char := range str {
		fmt.Printf("Character on position %d is %c\n", pos, char)
	}
	fmt.Println()

	str1 := "杭州琦星"
	fmt.Printf("The len of str is: %d\n", len(str1))
	for pos, char := range str {
		fmt.Printf("character %c starts at position %d\n", char, pos)
	}
	fmt.Println()

	fmt.Println("index int(rune) rune char bytes")
	for index, rune := range str1 {
		fmt.Printf("%-2d	%d	 %U '%c' %X\n", index, rune, rune, rune, []byte(string(rune)))
	}
}
