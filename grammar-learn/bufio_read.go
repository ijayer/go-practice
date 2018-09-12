//bufio包 带有缓冲的IO操作

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//construct a reader(读写器)
	inputReader := bufio.NewReader(os.Stdin)
	string, err := inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input from os.Stdin: %s", string)
	}
}
