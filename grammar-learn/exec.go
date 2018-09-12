package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd.exe", "/c", "tasklist")
	var buf []uint8
	var err error

	err = nil
	if nil == err {
		buf, err = cmd.Output()
		if err != nil {
			fmt.Printf("%t", err)
			return
		}
	} else {
		fmt.Printf("fatal error:%t", err)
		return
	}

	var c uint8
	var line []uint8
	jump := false
	for i := 0; i < len(buf); i += 1 {
		c = buf[i]
		if jump == true {
			if c == uint8('\n') {
				jump = false
			}
			continue
		}
		if c > 0x7f {
			jump = true
			line = []uint8{}
			fmt.Println("One line skipped.")
			continue
		}
		if c == uint8('\n') {
			fmt.Printf("%s\n", line)
			line = []uint8{}
			continue
		}
		line = append(line, c)
	}

}
