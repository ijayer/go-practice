//bufio Stdin NewReader
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	inputString, err := inputReader.ReadString('\n')
	CheckError(err)

	fmt.Println("Read from Stdin: ", inputString)

	fmt.Printf("%s", "switch_1: ")
	switch inputString {
	case "www\n":
		fmt.Println("I am www")
	case "google\n":
		fmt.Println("I am google")
	default:
		fmt.Println("Not found we need string!")
	}

	fmt.Printf("%s", "switch_2: ")
	switch inputString {
	case "www\n":
		fallthrough
	case "google\n":
		fallthrough
	case "com\n":
		fmt.Println("I am com")
	default:
		fmt.Println("Not found we need string!")
	}
}

func CheckError(err error) {
	if err != nil {
		log.Println("Faltl error", err.Error())
		return
	}
}
