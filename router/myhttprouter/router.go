package myhttprouter

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

// MainHttpRouter
//
// Logger & Auth middleware 调用方式：Logger(Auth(Hello))
// 注册过程：
//      func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error
//                                       ↓
//                                       ↓¹ hello 传给 next
//                                Auth(next MyHandle) MyHandle
//                                                     ⇡    ↓
//      底层绑定 ...Auth.func1, 并返回给 MyHandle     ² ⇡    ↓
//      instance.golang.com/router/myhttprouter/Auth.func1  ↓
//                                                          ↓
//                                                          ↓³ Auth.func1 传给 next
//                                                 Logger(next MyHandle) httprouter.Handle
//                                                                             ⇡    ↓
//                       底层绑定 ...Logger.func1, 并返回给 httprouter.Handle  ⁴⇡    ↓
//                            instance.golang.com/router/myhttprouter/Logger.func1  ↓
//                                                                                  ↓
//                                                                                  ↓⁵
//                                                                          Logger.func1 最终被注册到
//                                                                          了httprouter 的路由节点中
//
// 调用过程：与注册的流程刚好相反
func MainHttpRouter(port string) {
	router := httprouter.New()
	router.GET("/", Logger(Index))
	router.GET("/hello/:name", Logger(Auth(Hello)))
	router.GET("/parameters/*name", AllParaHello)
	router.GET("/protected/", BasicAuth(Protected, "username", "secret"))
	router.ServeFiles("/static/*filepath", http.Dir("./router/static_sources"))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
