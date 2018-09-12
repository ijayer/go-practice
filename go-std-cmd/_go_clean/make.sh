#!/bin/sh

# go build
echo go build ./main.go
go build ./main.go

echo ""
echo ls -a
ls -a
sleep 1s

# go clean
echo ""
echo "go clean -x"
go clean -x

echo ""
echo ls -a
ls -a
sleep 1s

# go install
echo ""
echo go install -v
go install -v

echo ""
echo ls ${GOPATH}/bin -a
ls ${GOPATH}/bin -a
sleep 1s

# go clean
echo ""
echo "go clean -x -i"
go clean -x -i

echo ""
echo ls ${GOPATH}/bin -a
ls ${GOPATH}/bin -a
sleep 1s