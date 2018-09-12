#!/bin/sh

# go install
echo "go install instance.golang.com/go-std-command/_go_install"
go install instance.golang.com/go-std-command/_go_install

echo ""
echo "ls ${GOPATH}/bin -a"
ls ${GOPATH}/bin -a
echo "_go_install.exe"
_go_build.exe
sleep 2s

echo ""
echo "go clean -x"
go clean -x