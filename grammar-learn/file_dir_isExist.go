package main

import (
	"log"
	"os"
)

func main() {
	// 查看当前的工作目录路径，得到测试文件的绝对路径
	currentDir, _ := os.Getwd()
	log.Println(currentDir)

	filename := currentDir + "/file/readme.md"
	log.Println(filename)

	if IsExists(filename) {
		log.Println(filename + "存在")
		// 返回，提示
		return

	} else {
		log.Println(filename + "不存在")
		// 创建该文件
	}
}

/*
  检查文件或目录是否存在
  参数： filename 指定的文件或目录
  返回： filename 存在 true 不存在 false
*/
func IsExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
