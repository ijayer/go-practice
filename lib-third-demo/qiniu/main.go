/*
 * 说明：
 * 作者：zhe
 * 时间：2018-03-01 14:37
 * 更新：
 */

package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/zhezh09/go-practice/lib-third-demo/qiniu/server"
)

func main() {
	router := httprouter.New()
	router.GET("/api/qiniu/upload/token", server.UploadToken)
	router.GET("/api/qiniu/download/url", server.DownloadURL)
	router.POST("/api/qiniu/callback", server.QiNiuCallback)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
		ExposedHeaders: []string{"X-Something-Special", "SomethingElse"},
	})
	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
