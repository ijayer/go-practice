package main

import "flag"

var DBName string
var Addr string

func init() {
	flag.StringVar(&DBName, "db", "restapi", "mgo db name")
	flag.StringVar(&Addr, "addr", "localhost:8090", "http port")
	flag.Parse()
}

func main() {
	app := NewApp()
	app.Init(DBName)
	app.Run(Addr)
}
