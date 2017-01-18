package main

import (
	"flag"
	"instance.golang.com/mymiddleware/handler"
	"instance.golang.com/mymiddleware/middleware"
)

// command line parameter
// set the command line parameters as follows
// usage:   --port=8000
// default: "8000"(string)
var Port = flag.String("port", "8000", "http port")

func main() {
	handler.MainHandler(Port)
	middleware.MainMiddleware()
}
