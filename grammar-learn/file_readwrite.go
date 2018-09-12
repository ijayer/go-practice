/**
 * 文件读取函数的测试范例
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

var inputFile = "z_demo.txt"
var outputFile = "z_demo_tmp.txt"
var objFile = "zssleay32.dll"

func main() {
	//1. os包和bufio包读取文件
	fmt.Println("--------------------------------------os bufio demo")
	file, err := os.Open("readme.md")
	CheckError(err)
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, readError := reader.ReadString('\n')
		if readError == io.EOF {
			fmt.Println("end of file")
			break
		}
		fmt.Println("read content: ", str)
	}

	//2. 将整个文件的内容读取到一个字符串中
	//ioutil包提供ioutil.ReadFile() && ioutil.WriteFile()
	fmt.Println("--------------------------------------ioutil demo")

	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("@_@_______end of file", err.Error())
		return
	}
	fmt.Println("read z_demo.txt content: ", string(b))

	err = ioutil.WriteFile(outputFile, b, 0x666)
	CheckError(err)

	//3. 带缓冲的读取
	//如果文件不是按行划分，则不能使用以上方法处理，使用bufio.Reader 的Read()函数将文件读取到一个缓冲区
	fmt.Println("--------------------------------------read obj file demo")
	buffer := make([]byte, 1024)

	file1, err1 := os.Open(GetCurrentPath() + "/file/" + objFile)
	CheckError(err1)
	defer file1.Close()

	reader1 := bufio.NewReader(file1)
	for {
		n, _ := reader1.Read(buffer)
		if n == 0 {
			fmt.Println("@_@_______file read finished")
			break
		}
		fmt.Println("buffer content: ", string(buffer))
	}

	//4. 按列读取文件中的数据
	//fmt包提供Fscan开头的一系列函数

	//5. path/filepath包提供了跨平台的文件名和路径处理函数Base()
	fmt.Println("--------------------------------------filepath package demo")
	str1 := filepath.Base("F:/SvnFiles/SDK/src/mydemo.com")
	fmt.Println("filename: ", str1)

	//6. 写文件
	//文件句柄、bufio的读取器和bufio的写入器
	fmt.Println("--------------------------------------write file demo")
	filename := "zoutput.bat"
	str2 := "www.google.com\n"

	file2, err2 := os.OpenFile(GetCurrentPath()+"/file/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	CheckError(err2)
	defer file2.Close()

	writer := bufio.NewWriter(file2)
	for i := 0; i < 10; i++ {
		_, err := writer.WriteString(str2)
		CheckError(err)
	}
	writer.Flush()

	//7. fmt包里面可以使用以F开头的print函数直接写入任何io.writer, 包括文件
	//fmt.Fprintf
	fmt.Println("--------------------------------------fmt Fprintf demo")
	fmt.Fprint(file2, "some test data.\n")
	//fmt.Fprintf(file2, "some test data.f\n")
	os.Stdout.WriteString("hello! www.google.com!")
}

func CheckErrors(err error) {
	if err != nil {
		fmt.Println("@_@_______faltal error", err.Error())
		return
	}
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	CheckErrors(err)
	return dir
}
