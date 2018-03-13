/*
 * 说明：logrus wrapper
 * 作者：zhe
 * 时间：2018-03-07 21:15
 * 更新：
 */

package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

/* truncated */

// Decorate appends line, file and function context to the logger and returns a function to call before
// each log
func Decorate(logger *logrus.Entry) func() *logrus.Entry {
	return func() *logrus.Entry {
		if pc, f, line, ok := runtime.Caller(1); ok {
			fnName := runtime.FuncForPC(pc).Name()
			_, file := path.Split(f) // 取文件名
			caller := fmt.Sprintf("%s:%v %s", file, line, fnName)

			return logrus.WithField("caller", caller)
		}
		return logger
	}
}

// Log appends line, file and function context to the logger
func Log() *logrus.Entry {
	if pc, f, line, ok := runtime.Caller(1); ok {
		fnName := runtime.FuncForPC(pc).Name()
		_, file := path.Split(f) // 取文件名
		caller := fmt.Sprintf("%s:%v %s", file, line, fnName)

		return logrus.WithField("caller", caller)
	}
	return &logrus.Entry{}
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)

	Log().WithFields(logrus.Fields{
		"field":  "something",
		"field2": "something else",
	}).WithError(errors.New("error")).Error("error upgrading websocket")

	logCtx := Log()
	entry := logCtx.WithFields(logrus.Fields{
		"name":       "zhe",
		"request_id": "xx123",
	})
	entry.Info("xxx")
	entry.Errorf("%s", "xxx")

	Log().WithFields(logrus.Fields{
		"name":       "zhe",
		"request_id": "xx123",
	}).Info("info")

	Log().Warn("xxx")
}
