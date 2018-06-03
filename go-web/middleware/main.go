package main

import (
	"flag"
)

// command line parameter
// set the command line parameters as follows
// usage:   --port=8000
// default: "8000"(string)
var Port = flag.String("port", "8000", "http port")

func main() {
	//handler.MainHandler(Port)
	//middleware.MainMiddleware(Port)
	flag.Parse()
	MainLogMiddleware(Port)
}
