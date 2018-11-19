/*
 * 说明：
 * 作者：zhe
 * 时间：2018-10-23 9:43 AM
 * 更新：
 */

package main

import (
	"flag"

	"go-web/api"
)

func init() {
	flag.StringVar(&api.DBName, "db_name", "mgo", "mongo name")
	flag.StringVar(&api.DBAddr, "db_addr", "localhost:27017", "mongo addr")
	flag.StringVar(&api.Addr, "port", "8090", "http server port")
	flag.Parse()
}

func main() {
	app := api.NewApp()
	app.InitMgo(api.DBName, api.DBAddr)
	app.Run(api.Addr)
}
