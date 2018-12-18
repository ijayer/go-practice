/*
 * 说明：
 * 作者：zhe
 * 时间：2018-11-19 2:47 PM
 * 更新：
 */

package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
)

var chunkSize = 4

func main() {
	ctx, _ := ioutil.ReadFile("./a.txt")
	a := fmt.Sprintf("a.txt: %x\n", sha256.Sum256(ctx))

	var txt = []string{
		"0123",
		"4567",
		"890a",
		"bcde",
		"fg",
	}

	f, _ := os.OpenFile("./b.txt", os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer f.Close()

	l := len(ctx)
	index := l / chunkSize
	if l%chunkSize != 0 {
		index += 1
	}
	println("index:", index)

	var off int64
	for i := 0; i < index; i++ {
		println("i:", i)

		start := i * chunkSize
		// end := (i+1)*chunkSize - 1
		// if i == (int(index) - 1) {
		// 	end = l
		// }

		data := []byte(txt[i])

		off = int64(start)
		f.Seek(off, 0)
		println("off:", off)

		n, err := f.WriteAt([]byte(data), off)
		if err != nil {
			println(err)
		}
		println("written:", n)
	}

	ctx, _ = ioutil.ReadFile("./b.txt")
	b := fmt.Sprintf("a.txt: %x\n", sha256.Sum256(ctx))
	println(a == b)
}
