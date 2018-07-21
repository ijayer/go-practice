#!/usr/bin/env bash

# tags_test_darwin.go   编译约束：只能在 darwin  平台下被编译
# tags_test_linux.go    编译约束：只能在 linux   平台下被编译
# tags_test_win.go      编译约束：只能在 windows 平台下被编译
# hello_darwin_amd64.go 编译约束：只能在 darwin, amd64 平台下被编译
# hello_linux_amd64.go  编译约束：只能在 linux, amd64  平台下被编译
# hello_windows_386.go  编译约束：只能在 windows, 386  平台下被编译

#
echo "$ ls"
ls
echo -e "\n"

# 列出当前目录下，可以被编译的 go 文件列表
echo "$ go list -f {{.GoFiles}} ./"
go list -f {{.GoFiles}} ./
echo -e "\n"
# output: [tags_test_win.go]

# 编译当前包下的文件
echo "$ go build -o out && ./out"
echo -e "\n"
go build -o out && ./out

# output: hello, windows