/*
 * 说明：Cache Service based on gRPC
 * 作者：zhe
 * 时间：2018-09-12 8:42 AM
 * 更新：
 */

package main

import (
	"cache-service/client"
	"cache-service/server"
)

func main() {
	go func() {
		server.Run()
	}()
	client.Run()
}
