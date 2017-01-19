package main

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"flag"
	"instance.golang.com/myrouter/myhttprouter"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.InfoLevel)
}

var Port = flag.String("port", "8000", "http port")

func main() {
	myhttprouter.MainHttpRouter(Port)
}