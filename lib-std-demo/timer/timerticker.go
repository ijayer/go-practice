package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	// Timer(定时器)
	logrus.Println("#___________timer starting timing...")

	Timer1 := time.NewTimer(2 * time.Second)
	// 主函数阻塞 2s
	<-Timer1.C
	logrus.Println("#___________time 1 is over...")

	timer2 := time.NewTimer(2 * time.Second)
	go func() {
		<-timer2.C
		logrus.Println("#___________into time 2 func")
	}()
	stop := timer2.Stop()
	if stop {
		logrus.Println("#___________time 2 is over...")
	}

	time.Sleep(3 * time.Second)
	logrus.Println("#___________the timer is over...")

	// Ticker(打点器)
	logrus.Println("#___________ticker starting timing...")
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Printf("Ticker at: %v\n", t)
		}
	}()

	time.Sleep(10 * time.Second)
	ticker.Stop()
	logrus.Println("#___________ticker is over...")
}
