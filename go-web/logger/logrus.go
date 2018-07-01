/*
 * 说明：
 * 作者：zhe
 * 时间：2018-03-07 15:52
 * 更新：
 */

package main

import (
	"os"

	"instance.golang.com/utils"

	"github.com/omidnikta/logrus"
	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
)

// 设置默认的 Logger 参数
func init() {
	// 设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: utils.TimeLayout,
	})
	// 设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	// 设置最低log level
	log.SetLevel(log.DebugLevel)
}

func main() {
	LogrusLevel()
	LogrusTextFormatter()
}

// Logrus 等级输出, 默认Info及以上
func LogrusLevel() {
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// log.Fatal("Bye.")         // log之后会调用os.Exit(1)
	// log.Panic("I'm bailing.") // log之后会panic()
}

// Logrus 的非结构化日志定制输出(TextFormatter{}格式)
//
// 默认输出格式：time="2018-04-05T06:08:53+08:00" level=info msg="logrus log to lumberjack in normal text formatter"
//
// 如果isColored为false，输出的就是带有time, msg, level的结构化日志；只有isColored为true才能输出我们想要的普通日志。
// isColored的值与三个属性有关：ForceColors 、isTerminal和DisableColors。我们按照让isColored为true的条件组合重新
// 设置一下这三个属性，因为输出到file，因此isTerminal自动为false (https://tonybai.com/2018/04/06/the-problems-i-
// encountered-when-writing-go-code-issue-3rd/?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io)
//
// 定制 TextFormatter：
//      1. 设置时间格式为：TimestampFormat = 2006-01-02 15:04:05
//      2. 启用 Colors:   ForceColors = true
func LogrusTextFormatter() {
	l := logrus.NewEntry(&logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			TimestampFormat: utils.TimeLayout,
			FullTimestamp:   true,
			ForceColors:     true,
		},
		Level: logrus.InfoLevel,
	})
	l.Info("info info info")
	l.Warn("warn warn warn")
	l.Error("error error error")
	l.Fatal("fatal fatal fatal")
	l.Panic("panic panic panic")
}

// Logrus 使用Fields
func LogrusFields() {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")
}

// Logrus 使用固定的 Fields
// 有时候我们需要固定的fields,不需要向每行都重复写,只需要生成一 logrus.Entry
//
// logrus.WithFields会自动返回一个 *Entry，Entry里面的有些变量会被自动加上
//
// time:entry被创建时的时间戳
// msg:在调用.Info()等方法时被添加
// level
func LogrusStableFields() {
	entry := log.WithFields(log.Fields{"request_id": "fdjldoKLJLKD332480", "user_ip": "192.168.1.169"})
	entry.Info("something happened on that request")
	entry.Warn("something not great happened")
}

// Logrus 创建新实例, 新实例的配置与不影响原logger
// 我们可以发现，独立的Logger,拥有自己的各个参数，比如直接使用logrus.Panic("GG")这是使用默认的Logger,
// 上面提到的init函数里面的各项设置，是设置默认Logger的，不会对自己生成的Logger有影响
func LogrusNewLogger() {
	l := log.New()
	l.SetLevel(log.DebugLevel)
	l.Out = os.Stderr

	l.Debug("Useful debugging information.")
	l.Info("Something noteworthy happened!")
	l.Warn("You should probably take a look at this.")
	l.Error("Something failed but I'm not quitting.")
	l.Fatal("Bye.")         // log之后会调用os.Exit(1)
	l.Panic("I'm bailing.") // log之后会panic()
}
