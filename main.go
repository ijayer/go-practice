package main

import (
	"instance.golang.com/mymiddleware"
	"instance.golang.com/mycontext"
)

func main() {
	mycontext.MainContext()
	mymiddleware.MainMiddleware()
}
