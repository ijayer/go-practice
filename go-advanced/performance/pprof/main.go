/*
 * 说明：
 * 作者：zhe
 * 时间：2018-06-06 10:14 AM
 * 更新：
 */

package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/julienschmidt/httprouter"
)

// bigBytes 函数每次分配 100M 内存
func bigBytes() *[]byte {
	s := make([]byte, 100000000)
	return &s
}

// 使用自定义的 endpoint 提供 pprof 分析, 采用 httprouter 提供路由
func main() {
	router := httprouter.New()
	router.HandlerFunc("GET", "/debug/pprof/", pprof.Index)
	router.HandlerFunc("GET", "/debug/pprof/profile", pprof.Profile)
	router.HandlerFunc("GET", "/debug/pprof/trace", pprof.Trace)
	router.HandlerFunc("GET", "/debug/pprof/symbol", pprof.Symbol)
	router.HandlerFunc("GET", "/debug/pprof/cmdline", pprof.Cmdline)

	router.HandlerFunc("GET", "/debug/pprof/heap", pprof.Handler("heap").ServeHTTP)
	router.HandlerFunc("GET", "/debug/pprof/goroutine", pprof.Handler("goroutine").ServeHTTP)
	router.HandlerFunc("GET", "/debug/pprof/block", pprof.Handler("block").ServeHTTP)
	router.HandlerFunc("GET", "/debug/pprof/treadcreate", pprof.Handler("threadcreate").ServeHTTP)
	router.HandlerFunc("GET", "/debug/pprof/mutex", pprof.Handler("mutex").ServeHTTP)

	log.Fatal(http.ListenAndServe(":8082", router))
}

// Case3: 通过 Web 服务远程分析
// func main() {
// 	var wg sync.WaitGroup
//
// 	go func() {
// 		log.Println(http.ListenAndServe(":8082", nil))
// 	}()
//
// 	for i := 0; i < 10; i++ {
// 		s := bigBytes()
// 		if s == nil {
// 			log.Println("oh noes")
// 		}
// 	}
//
// 	wg.Add(1)
// 	wg.Wait()
// }

// Case2: 内存分析
// func main() {
// 	for i := 0; i < 10; i++ {
// 		s := bigBytes()
// 		if s == nil {
// 			log.Println("oh noes")
// 		}
// 	}
// 	pprof.WriteHeapProfile(os.Stdout)
// }

// Case1: CPU分析
// func main() {
// 	err := pprof.StartCPUProfile(os.Stdout)
// 	if err != nil {
// 		log.Printf("[ERROR]StartCPUProfile: %v\n", err)
// 	}
// 	defer pprof.StopCPUProfile()
//
// 	for i := 0; i < 10; i++ {
// 		s := bigBytes()
// 		if s == nil {
// 			log.Println("oh noes")
// 		}
// 	}
// }
