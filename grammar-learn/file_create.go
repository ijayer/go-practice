package main

import (
	"log"
	"os"
)

func main() {
	//查看当前的工作目录路径，得到测试文件的绝对路径
	//在程序运行过程中的到该.exe文件所在的目录
	currentDir, _ := os.Getwd()
	log.Println(currentDir)

	path := currentDir + "/readme.md"

	f, err := os.Create(path)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(f.Name())
	defer f.Close()

	//dir := "D:/zhang/wang"
	//
	//err = os.MkdirAll(dir, 0755)
	//if err != nil {
	//	my-log-test.Println(err)
	//	return
	//}
}
