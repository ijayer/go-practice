/*
 * 说明：Go Module Demo
 * 作者：zhe
 * 时间：2018-08-29 9:55 AM
 * 更新：
 */

package main

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Cache struct {
	sync.Map
}

func main() {
	cache := Cache{Map:sync.Map{}}

	cache.Store("i", "1")
	cache.Store("j", "2")
	cache.Store("k", "3")

	cache.Range(func(key, value interface{}) bool {
		logrus.Infoln(key, value)
		return true
	})
}
