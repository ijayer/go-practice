/**
 * //3. 采用缓冲和flag解析命令行参数来读取文件
 * //以下代码实现了类似linux下的 cat 命令
 * //添加了行号输出
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var lineNum = flag.Bool("n", false, "print lineNum")

func main() {
	fmt.Println("\n@_@_______________bufio && flag readwrite file demo")
	flag.PrintDefaults()
	flag.Parse()
	//没有参数：则获取命令行输入然后输出
	if flag.NArg() == 0 {
		fmt.Println("command-line arguments is nil")
		Cat(bufio.NewReader(os.Stdin))
	} else if flag.NArg() > 0 {
		for i := 0; i < flag.NArg(); i++ {
			f, err := os.Open(flag.Arg(i))
			if err != nil {
				fmt.Println("[error] ", err.Error())
				continue
			}
			Cat(bufio.NewReader(f))
		}
	}
	return
}

func Cat(r *bufio.Reader) {
	for {
		b, err := r.ReadBytes('\n')

		if *lineNum { //输出行号
			fmt.Printf("%d ", i)
			i = i + 1
		}
		fmt.Fprintf(os.Stdout, "%s\n", b)
		if err == io.EOF {
			fmt.Println("[info]: end of file")
			break
		}
	}
	return
}
