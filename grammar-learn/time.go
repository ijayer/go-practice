package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	t.Format("2006-01-02 15:04:05")

	// 记录开始时间
	start := time.Now().Nanosecond()

	// 计算过程
	sum := 0
	for i := 0; i <= 100000000; i++ {
		sum += i
	}

	// 记录结束时间
	end := time.Now().Nanosecond()

	// 输出执行时间，单位为毫秒。
	fmt.Println((end - start) / 1000000)

	//输出执行结果
	fmt.Println(sum)

	//sendTimes := Date()
	//fmt.Println("[sendTime] ", sendTimes)
	//
	//for i:=1; i<10; i++{
	//	if i == 9 {
	//		break
	//	}
	//}
	//
	//recvTimes := Date()
	//fmt.Println("[recvTime] ", recvTimes)
	//
	//seconds := DelayTimes(sendTimes, recvTimes)
	//fmt.Println("[delay seconds] ", seconds)

	//fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	//
	//t = time.Now().UTC()
	//fmt.Println(t)
	//fmt.Println(time.Now())
	//
	//week = 60 * 60 * 24 * 7 * 1e9
	//week_from_now := t.Add(week)
	//fmt.Println(week_from_now)
	//
	//fmt.Println(t.Format(time.RFC822))
	//fmt.Println(t.Format(time.ANSIC))
	//fmt.Println(t.Format("02 Jan 2006 15:04"))
	//s := t.Format("20160714")
	//fmt.Println(t, "=>", s)
}

func Date() int {
	return time.Now().Second()
}

func DelayTimes(sendTimes, recvTimes int) int {
	//s, _ := strconv.Atoi(sendTimes)
	//fmt.Println("[s] ", sendTimes)
	//r, _ := strconv.Atoi(recvTimes)
	//fmt.Println("[r] ", recvTimes)

	return recvTimes - sendTimes
}
