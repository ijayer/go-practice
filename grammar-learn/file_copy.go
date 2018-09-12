/**
 * 文件拷贝
 */
package main

import (
	"fmt"
	"io"
	"os"
)

var (
	srcFile = "/zoutput.bat"
	dstFile = "/z_demo_tmp.txt"
)

func main() {
	//copy file
	dstFile = GetCurrentPath1() + dstFile
	srcFile = GetCurrentPath1() + srcFile

	CopyFile(dstFile, srcFile)
	fmt.Println("copy file finished!")
}

func CopyFile(dst, src string) (written int64, err error) {
	fs, err := os.Open(src)
	CheckError1(err)
	defer fs.Close()

	fd, err1 := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0655)
	CheckError1(err1)
	defer fd.Close()

	written, err = io.Copy(fd, fs)

	return written, err
}

func CheckError1(err error) {
	if err != nil {
		fmt.Println("fatal error ---> ", err.Error())
		return
	}
}

func GetCurrentPath1() string {
	dir, err := os.Getwd()
	CheckError1(err)

	return dir
}
