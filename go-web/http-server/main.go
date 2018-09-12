package main

import (
	"flag"
	"os"

	"github.com/zhezh09/go-practice/utils"

	"github.com/sirupsen/logrus"
	"github.com/zhezh09/go-practice/go-web/http-server/myhttprouter"
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
