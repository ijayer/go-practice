.PYTHON:clean

gobuild:
	@echo Linux: c-shared: .so
	go build -buildmode=c-shared -o hello.so main.go

	@echo

	@echo Win: c-shared: .dll
	go build -buildmode=c-shared -o hello.dll main.go

build:
	gcc hello.def -o hello hello.c hello.so
	./hello

	@echo

	gcc hello.def -o hello.exe hello.c hello.dll
	./hello.exe

clean:
	go clean -x
	rm -rf *.h *.exe *.so *.dll
