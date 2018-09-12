package main

import (
	"os/exec"
	//"strings"
	//"fmt"
	"log"
	//"os"
)

func main() {

	var arg []string

	cmd := exec.Command("D:/setup/lingoes_2.9.2_cn.exe", arg...)

	err := cmd.Start()
	if err != nil {
		log.Fatal(err.Error())
	}
}
