package main

/*
#include <stdio.h>
#include <stdlib.h>

// define "hellofoo"
char *c_foo = "hellofoo";

// define print
void c_print(char *str) {
	printf("C  Test: print: %s\n", str);
}
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Printf("Go Test: %v\n", Random())
	GoStrToCStr()
	CStrToGoStr()
}

// Random call C's rand()
func Random() int {
	return int(C.rand())
}

// GoStrToCStr go type convert to C type
func GoStrToCStr() {
	str := "hello world"
	fmt.Printf("Go Test: %T, %v\n", str, str)

	cs := C.CString(str) // Go's string to C's String
	defer C.free(unsafe.Pointer(cs))

	fmt.Printf("Go Test: %T, %v\n", cs, *cs)
	C.c_print(cs) // Call C's print
}

// CStrToGoStr c type convert to Go type
func CStrToGoStr() {
	gs := C.GoString(C.c_foo) // C's char to Go's string
	fmt.Printf("Go Test: %T, %v\n", gs, gs)
}
