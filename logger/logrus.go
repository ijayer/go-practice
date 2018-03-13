/*
 * 说明：
 * 作者：zhe
 * 时间：2018-03-07 15:52
 * 更新：
 */

package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	LogrusLevel()
}

// 设置默认的 Logger 参数
func init() {
	// 设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	log.SetFormatter(&log.TextFormatter{})
	// 设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
	log.SetOutput(os.Stdout)
	// 设置最低log level
	log.SetLevel(log.DebugLevel)
}

// Logrus 等级输出, 默认Info及以上
func LogrusLevel() {
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	log.Fatal("Bye.")         //log之后会调用os.Exit(1)
	log.Panic("I'm bailing.") //log之后会panic()
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
	l.Fatal("Bye.")         //log之后会调用os.Exit(1)
	l.Panic("I'm bailing.") //log之后会panic()
}
