package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")

	time.Sleep(6 * time.Second)
	fmt.Println("End of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * time.Second)
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * time.Second)
	fmt.Println("End of shortWait()")
}