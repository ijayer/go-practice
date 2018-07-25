.PYTHON:clean

TARGET=c-archive

go-build:
	go build -x -v -buildmode=c-archive .

build:
	gcc -o hello.exe c-archive.c c-archive.a -lWinMM -lntdll -lWS2_32

clean:
	go clean -x
	rm -rf *.h *.c *.exe
