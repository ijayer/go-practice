package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var newLine = flag.Bool("n", false, "print newLine")

const (
	Space   = " "
	NewLine = "\n"
)

func main() {
	//1. os包获取命令行参数
	//命令行执行：./a.out.exe args1 args2 args3
	//命令行输出：args1, args2, args3
	fmt.Println("\n@_@_______________os args demo")
	args := ""

	if len(os.Args) > 1 {
		args += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("os args is: ", args)

	//2. flag包获取和解析命令行参数
	//以下代码模拟一个echo打印程序
	fmt.Println("\n@_@_______________flag package demo")
	flag.PrintDefaults()
	flag.Parse() //scans the args list and sets up flags

	var s = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *newLine {
				s += NewLine
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
