/*
 * 说明：
 * 作者：zhe
 * 时间：2018-08-20 1:46 PM
 * 更新：
 */

package main

import (
	"os"

	"service.robot.com/util"

	"github.com/rifflock/lfshook"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
)

const TimeLayout = "2006-01-02 15:04:05"

var logger = logrus.NewEntry(nil)

func main() {
	configLogger("prod")
	logger.Infof("xxx")
}

// configLogger
func configLogger(env string) {
	var l = logrus.New()

	// 设置默认输出：stderr
	l.Out = ansicolor.NewAnsiColorWriter(os.Stderr)

	// 设置默认格式：TextFormatter
	l.Formatter = &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: TimeLayout,
	}

	switch env {
	case "test":
		l.Out = os.Stdout
		l.SetLevel(logrus.DebugLevel)

	case "prod":
		pathMap := lfshook.PathMap{
			logrus.InfoLevel:  "log/" + util.Date() + "/x-change.info.log",
			logrus.ErrorLevel: "log/" + util.Date() + "/x-change.erro.log",
		}
		l.AddHook(lfshook.NewHook(
			pathMap,
			&logrus.TextFormatter{
				FullTimestamp:   true,
				TimestampFormat: TimeLayout,
			}),
		)
		l.SetLevel(logrus.InfoLevel)

		// logger.Formatter = &logrus.TextFormatter{
		// 	FullTimestamp:   true,
		// 	TimestampFormat: util.TimeLayout,
		// }
		// file, err := os.OpenFile("x-change.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		// if err == nil {
		// 	logger.Out = file
		// } else {
		// 	logger.Info("Failed to log to file, using default stderr")
		// }
		// logger.SetLevel(logrus.InfoLevel)
	}

	filenameHook := util.NewHook()
	filenameHook.Field = "caller"
	l.AddHook(filenameHook)

	logger.Logger = l
	logger = logger.WithField("Module", "[X-CHANGE]")
}
