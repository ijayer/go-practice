package main

import (
	"flag"
	"os"

	"qx-api/src/utils"

	"github.com/sirupsen/logrus"
	"instance.golang.com/http-server/myhttprouter"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "http port")
	flag.Parse()

	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: utils.TimeLayout,
	})
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	logrus.Infof("service listen and serve on: [:%s]", port)
	myhttprouter.MainHttpRouter(port)
}
