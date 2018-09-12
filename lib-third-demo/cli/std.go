package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	Space = " "
	Delim = '\n'
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please input <username> <password> <host>(Separated by spaces): ")
	str, _ := reader.ReadString(Delim)

	strS := strings.Split(str, Space)
	if len(strS) != 3 {
		fmt.Println(errors.New("Invalid parameter"))
		os.Exit(2)
	}

	fmt.Printf("%v %v %v\n", strS[0], strS[1], strS[2])
}
