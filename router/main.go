package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"instance.golang.com/router/myiris"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.InfoLevel)
}

var Port = flag.String("port", "8000", "http port")

func main() {
	fmt.Printf("##_______________[Listen and serice on：%s]\n", *Port)

	//myhttprouter.MainHttpRouter(Port)

	//mygorilla.MainGorilla(Port)

	myiris.MainIris(Port)
}
