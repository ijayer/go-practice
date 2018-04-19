/*
 * 说明：
 * 作者：zhe
 * 时间：2018-04-19 17:09
 * 更新：
 */

package myhttprouter

import (
	"net/http"

	"context"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// MyHandle
type MyHandle func(http.ResponseWriter, *http.Request, httprouter.Params) (Response, error)

// Logger 中间件: 记录接口请求日志(请求入口)
//
// Logger 接收一个自定义的Handle处理器，处理服务的业务关系，但正真被注册到路由器中的Handle函数是下面return的这个匿名函数：
//      func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//          // handle logic
//          // ...
//      }
// 实际上，该匿名函数只是表面上的匿名，在注册路由的过程中，Go会在底层为该匿名函数添加一个 新的名称，例如
//      (qx-api/src/middleware/Logger.func1), 即 Handle:qx-api/src/middleware/Logger.func1
// 最终会在 ServeHTTP(w http.ResponseWriter, req *http.Request) 中调用执行; 而参数列表中的自定义 Handle 会在该匿名函
// 数中被调用
func Logger(next MyHandle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// request entry
		var ctx = r.Context()
		start := time.Now()
		ctx = context.WithValue(ctx, "start", start)

		// return from auth middleware
		result, err := next(w, r.WithContext(ctx), ps)
		if err != nil {
			logrus.Error(err)
		}

		// handle response
		ctx = result.Ctx
		if ctx != nil {
			if v := ctx.Value("key"); v != nil {
				logrus.Info("key", v.(string))
			}
		}
		logrus.Info(result.Data)

		// request end
		end := time.Now()
		logrus.Infof("time cost: %v", end.Sub(start))
	}
}

// Auth middleware 接口授权验证
func Auth(next MyHandle) MyHandle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (Response, error) {
		// auth token ...
		var ctx = r.Context()
		ctx = context.WithValue(ctx, "some params", "some params")

		// return from our custom handle
		return next(w, r.WithContext(ctx), ps)
	}
}
