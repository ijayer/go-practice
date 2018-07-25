.PYTHON:clean

install:

	CGO_ENABLED=0 go install -buildmode=shared std

build:

	go build -linkshared hello.go

clean:

	go clean -x
	rm -rf *.exe