#!/bin/sh

echo "go build instance.golang.com/go-std-command/_go_build"
go build instance.golang.com/go-std-command/_go_build

echo ""
echo "ls"
ls -a

echo ""
echo ./_go_build.exe
./_go_build.exe
sleep 2s

echo ""
echo "go clean -x"
go clean -x
go clean