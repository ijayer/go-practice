/*
 * 说明：Go·Cron定时任务
 * 作者：zhe
 * 时间：2018-03-13 10:02
 * 更新：添加测试Demo
 */

package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {
	i := 0
	c := cron.New()
	spec := `*/5 * * * * ?`

	c.AddFunc(spec, func() {
		i++
		log.Printf("cron is running: %v\n", i)
	})
	c.Start()

	select {}
}
