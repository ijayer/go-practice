/*
 * 说明：Session & Cookie 实际应用
 * 作者：zhe
 * 时间：2018-05-25 3:39 PM
 * 更新：Session 管理设计
 *		  - 全局 Session 管理器
 *		  - 保证 SessionID 的全局唯一性
 *		  - 为每个客户关联一个 Session
 *		  - Session 的存储(内存|文件|数据库)
 *		  - Session 的过期处理
 */

package main

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"

	"instance.golang.com/utils"
)

const maxLifeTime = 36 // unit: second

// 初始化logrus, 全局有效
func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: utils.TimeLayout,
		ForceColors:     true,
	})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stderr)
}

// 定义基于内存的Session存储器
var globalMemProvider = new(MemoryProvider)

func init() {
	globalMemProvider.sessions = make(map[string]interface{}, 0)
	Register("memory", globalMemProvider)
}

// 定义并初始化全局Session管理器
var globalSessionManager = new(Manager)

func init() {
	var err error
	// NewManager 根据 providerName 在全局变量 providers 中提取 Session 存储器对象完成初始化
	// Note: 在调用NewManager时，必须保证参数 providerName 所表示的 Session 存储器已经完成注册
	globalSessionManager, err = NewManager("memory", "gosessionid", maxLifeTime)
	if err != nil {
		panic(err)
	}
	go globalSessionManager.GC()
}

func main() {
	router := httprouter.New()

	user := NewUser()
	router.Handle(http.MethodPost, "/signin", user.Signin)
	router.Handle(http.MethodPost, "/signout", user.Signout)
	router.Handle(http.MethodGet, "/welcome", user.Welcome)
	router.Handle(http.MethodGet, "/refresh", Refresh)

	logrus.Infof("service listen and serve on: [:%v]", 8081)
	if err := http.ListenAndServe(":8081", router); err != nil {
		logrus.Panicf("http server: %v", err)
	}
}
