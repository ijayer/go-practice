/*
 * 说明：
 * 作者：zhe
 * 时间：2018-04-18 16:10
 * 更新：
 */

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8083", "server port")
	flag.Parse()
}

func main() {
	router := httprouter.New()
	router.GET("/ws", ServeWs)

	fmt.Printf("service listen and serve on: [:%v]\n", port)
	http.ListenAndServe(":"+port, router)
}
