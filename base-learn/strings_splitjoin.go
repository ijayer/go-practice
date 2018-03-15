package main

import (
	"log"
	"strings"
)

func main() {
	//str := "The quick brown fox jumps over the lazy dog"
	//s1 := strings.Fields(str)
	//fmt.Printf("Splitted in slice: %v\n", s1)
	//for _, val := range s1 {
	//	fmt.Printf("%s - ", val)
	//}
	//fmt.Println()
	//
	//str2 := "GO1|The ABC of Go|25"
	//sl2 := strings.Split(str2, "|")
	//fmt.Printf("Splitted in slice: %v\n", sl2)
	//for _,	val := range sl2  {
	//	fmt.Printf("%s - ", val)
	//}
	//fmt.Println()
	//str3 := strings.Join(sl2, ";")
	//fmt.Printf("sl2 joined by ;: %s\n", str3)

	//---------------------------------------------------------------------
	//startTime := "2016-11-25"

	curTime := "2016-11-24"
	endTime := "2016-12-25"

	//st := strings.Split(startTime, "-")
	et := strings.Split(endTime, "-")
	ct := strings.Split(curTime, "-")

	//my-log-test.Printf("##Info________endTime = %v", et)
	//my-log-test.Printf("##Info________curTime = %v", ct)

	var a, b string
	if len(ct) == len(et) {
		for i := 0; i < 3; i++ {
			//my-log-test.Println("##Info______________compare Object: ", et[i], dt[i])
			a += ct[i]
			b += et[i]
		}
		//my-log-test.Println("##Info______________________________________________compare result: ", status)

		log.Printf("##Info________curTime = %v", a)
		log.Printf("##Info________endTime = %v", b)

		status := strings.Compare(a, b)
		if status >= 0 {
			log.Println(false)
		} else {
			log.Println(true)
		}
	}
}

func DateCompare(currentTime, endTime string) bool {
	ct := strings.Split(currentTime, "-")
	et := strings.Split(endTime, "-")

	var a, b string
	if len(ct) == len(et) {
		for i := 0; i < 3; i++ {
			a += ct[i]
			b += et[i]
		}
	} else {
		return false
	}

	if status := strings.Compare(a, b); status >= 0 {
		return false
	}
	return true
}
