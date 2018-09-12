/*
 * 说明：Golang基础库学习
 * 作者：zhe
 * 时间：2018-06-13 21:08
 * 更新：package runtime
 */

package runtime

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

func main() {
	TestCallerFunc()
}

func runtimeInfo() {
	log.Println("cpu num:", runtime.NumCPU())

	if runtime.GOOS == "windows" {
		log.Println("os: windows")
	}
	if runtime.GOOS == "linux" {
		log.Println("os: linux")
	}
}

func runOnSingleCPU() {
	runtime.GOMAXPROCS(1)

	go func() {
		defer func() {
			log.Println("goroutine end...")
		}()
		go func() {
			log.Println("goroutine")
		}()
	}()

	for i := 0; i < 10; i++ {
		log.Println("i=", i)
		if i == 4 {
			runtime.Gosched() // main 进程让出 CPU(切换任务)
		}
	}
	log.Println("main end...")
}

func runOnMultipleCPU() {
	runtime.GOMAXPROCS(2)
	exit := make(chan int)

	go func() {
		defer close(exit)
		log.Println("goroutine num: ", runtime.NumGoroutine())

		go func() {
			log.Println("goroutine")
		}()
	}()

	for i := 0; i < 10; i++ {
		log.Println("i =", i)
		if i == 4 {
			runtime.Gosched() // main 进程让出 CPU(切换任务)
		}
	}
	log.Println("main end...")

	<-exit
}

// funcCallerInfo runtime.Caller 用法 demo: 输出 goroutine 栈上的函数调用信息
func funcCallerInfo() {
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			// 无法获得信息，ok=false
			break
		}
		log.Printf("skip: %v, pc: %v, file: %v, line: %v\n", skip, pc, file, line)
	}

	// output:
	// 2018/06/14 10:11:24 skip: 0, pc: 4812470, file: D:/code/Go_Path/src/instance.golang.com/base-lib/runtime.go, line: 77
	// 2018/06/14 10:11:24 skip: 1, pc: 4812054, file: D:/code/Go_Path/src/instance.golang.com/base-lib/runtime.go, line: 16
	// 2018/06/14 10:11:24 skip: 2, pc: 4366477, file: D:/Go/src/runtime/proc.go, line: 198
	// 2018/06/14 10:11:24 skip: 3, pc: 4516640, file: D:/Go/src/runtime/asm_amd64.s, line: 2361
}

// funcCallersDemo 输出每个栈帧的 pc(program counters) 信息
func funcCallersDemo() {
	pc := make([]uintptr, 1024)
	for skip := 0; ; skip++ {
		i := runtime.Callers(skip, pc)
		if i <= 0 {
			// 未记录任何 pc 信息
			break
		}
		log.Printf("skip: %v, pc: %v\n", skip, pc[:i])
	}

	// output:
	// 2018/06/14 14:33:30 skip: 0, pc: [4219048 4812589 4812199 4366622 4516785]
	// 2018/06/14 14:33:30 skip: 1, pc: [4812589 4812199 4366622 4516785]
	// 2018/06/14 14:33:30 skip: 2, pc: [4812199 4366622 4516785]
	// 2018/06/14 14:33:30 skip: 3, pc: [4366622 4516785]
	// 2018/06/14 14:33:30 skip: 4, pc: [4516785]
}

// funcForPCDemo
func funcForPCDemo() {
	for skip := 0; ; skip++ {
		pc, _, _, ok := runtime.Caller(skip)
		if !ok {
			// 无法获得信息
			break
		}
		f := runtime.FuncForPC(pc)
		file, line := f.FileLine(0)
		log.Printf("skip: %v, pc: %v\n", skip, pc)
		log.Printf("\t file: %v, line: %v\n", file, line)
		log.Printf("\t entry: %v\n", f.Entry())
		log.Printf("\t name: %v\n", f.Name())
	}
	println()

	// output:
	// 2018/06/14 15:42:25 skip: 0, pc: 4813275
	// 2018/06/14 15:42:25      file: D:/code/Go_Path/src/instance.golang.com/base-lib/runtime.go, line: 113
	// 2018/06/14 15:42:25      entry: 4812368
	// 2018/06/14 15:42:25      name: main.funcForPCDemo
	// 2018/06/14 15:42:25 skip: 1, pc: 4812342
	// 2018/06/14 15:42:25      file: D:/code/Go_Path/src/instance.golang.com/base-lib/runtime.go, line: 15
	// 2018/06/14 15:42:25      entry: 4812304
	// 2018/06/14 15:42:25      name: main.main
	// 2018/06/14 15:42:25 skip: 2, pc: 4366621
	// 2018/06/14 15:42:25      file: D:/Go/src/runtime/proc.go, line: 109
	// 2018/06/14 15:42:25      entry: 4366096
	// 2018/06/14 15:42:25      name: runtime.main
	// 2018/06/14 15:42:25 skip: 3, pc: 4516928
	// 2018/06/14 15:42:25      file: D:/Go/src/runtime/asm_amd64.s, line: 2361
	// 2018/06/14 15:42:25      entry: 4516928
	// 2018/06/14 15:42:25      name: runtime.goexit

	pc := make([]uintptr, 1024)
	for skip := 0; ; skip++ {
		i := runtime.Callers(skip, pc)
		if i <= 0 {
			break
		}
		log.Printf("skip: %v, pc: %v\n", skip, pc[:i])

		for j := 0; j < i; j++ {
			f := runtime.FuncForPC(pc[j])
			file, line := f.FileLine(0)
			log.Printf("skip: %v, pc: %v\n", skip, pc[j])
			log.Printf("\t file: %v, line: %v\n", file, line)
			log.Printf("\t entry: %v\n", f.Entry())
			log.Printf("\t name: %v\n", f.Name())
		}
		break
	}

	// output:
	// 2018/06/14 15:42:25 skip: 0, pc: [4219048 4813404 4812343 4366622 4516929]
	// 2018/06/14 15:42:25 skip: 0, pc: 4219048
	// 2018/06/14 15:42:25      file: D:/Go/src/runtime/extern.go, line: 205
	// 2018/06/14 15:42:25      entry: 4218944
	// 2018/06/14 15:42:25      name: runtime.Callers
	// 2018/06/14 15:42:25 skip: 0, pc: 4813404
	// 2018/06/14 15:42:25      file: D:/code/Go_Path/src/instance.golang.com/base-lib/runtime.go, line: 113
	// 2018/06/14 15:42:25      entry: 4812368
	// 2018/06/14 15:42:25      name: main.funcForPCDemo
	// 2018/06/14 15:42:25 skip: 0, pc: 4812343
	// 2018/06/14 15:42:25      file: D:/code/Go_Path/src/instance.golang.com/base-lib/runtime.go, line: 15
	// 2018/06/14 15:42:25      entry: 4812304
	// 2018/06/14 15:42:25      name: main.main
	// 2018/06/14 15:42:25 skip: 0, pc: 4366622
	// 2018/06/14 15:42:25      file: D:/Go/src/runtime/proc.go, line: 109
	// 2018/06/14 15:42:25      entry: 4366096
	// 2018/06/14 15:42:25      name: runtime.main
	// 2018/06/14 15:42:25 skip: 0, pc: 4516929
	// 2018/06/14 15:42:25      file: D:/Go/src/runtime/asm_amd64.s, line: 2361
	// 2018/06/14 15:42:25      entry: 4516928
	// 2018/06/14 15:42:25      name: runtime.goexit
}

// CallerInfo 返回调用者的：文件名、函数名、调用行号等用户友好的信息
//
// skip 表示要跳过的栈帧数; 代码中的 skip+1 表示跳过 CallerInfo 函数自身的调用
func CallerInfo(skip int) (name, file string, line int, ok bool) {
	var pc uintptr
	if pc, file, line, ok = runtime.Caller(skip + 1); !ok {
		return
	}
	name = runtime.FuncForPC(pc).Name()
	return
}

// CallerFunc 格式化返回 CallerInfo 调用的结果信息
var CallerFunc = func(skip int) string {
	name, file, line, ok := CallerInfo(skip + 1)
	if !ok {
		return fmt.Sprintf("unkonw")
	}
	file = filepath.Base(file)
	return fmt.Sprintf("[%s:%v %v]", file, line, name)
}

func TestCallerFunc() {
	log.Printf("%v\n", CallerFunc(0))
}
