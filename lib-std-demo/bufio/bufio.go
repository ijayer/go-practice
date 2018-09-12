// GO BufIO 库学习

package main

import (
	"bufio"
	"bytes"
	"log"
	"strings"
)

/**
 * reader和writer基本结构
 * Reader implments buffering for an io.Reader object
 *
 * type Reader struct{
 *	buf 		[]byte
 *　	rd 		io.Reader
 *	r, w		int
 *	err 		error
 *	lastByte 	int
 *	lastRuneSize	int
 * }
 *
 * type Writer struct {
 *	err 		error
 *	buf 		[]byte
 *	n 		int
 *	wr 		io.Writer
 * }
 *
 * ReadWriter 集成了bufio.Reader和bufio.Writer, 实现了io.ReadWriter接口
 * type ReadWriter struct{
 *	*Reader
 *	*Writer
 * }
 */

func main() {
	//1. construct a Reader use bufio.NewReader: 读缓冲区
	ReadBuf := strings.NewReader("1234567890")
	reader := bufio.NewReader(ReadBuf)

	//2. construct a Writer use bufio.NewWriter: 写缓冲区
	WriteBuf := bytes.NewBuffer(make([]byte, 0))
	writer := bufio.NewWriter(WriteBuf)

	//3. peek(): return forward n bytes in buffer
	log.Println("Peek() demo...")
	buff, err := reader.Peek(5)
	CheckError(err)

	buff[0] = 'a'
	buff, _ = reader.Peek(5)

	writer.Write(buff)
	writer.Flush() //remember call flush to finshed write operation

	log.Println("buf---changed: ", buff)
	log.Println("buf-unchanged: ", ReadBuf)

	//4. read(): read n byte on ReadBuf once, n determined by the buff len
	//we can use for to read data cycle, n<=0, read over.
	log.Println("Read() demo...")
	buff1 := make([]byte, 2)
	for {
		n, _ := reader.Read(buff1)
		//CheckError(err)

		if n <= 0 {
			log.Println("read end... ")
			break
		}
		log.Printf("read %d bytes is %s", n, string(buff1))
	}

	//5. ReadByte() && UnReadByte()
	//ReadByte() read 1 byte data and return, or return error
	//UnReadByte set the last char is unread which by ReadByte() get
	log.Println("ReadByte() demo...")
	readBuf2 := strings.NewReader("1234567890")
	reader2 := bufio.NewReader(readBuf2)
	//read byte
	bt, err := reader2.ReadByte()
	CheckError(err)
	log.Println(string(bt))

	bt3, err3 := reader2.ReadByte()
	CheckError(err3)
	log.Println(string(bt3))
	//unread byte
	reader2.UnreadByte()
	bt2, err := reader2.ReadByte()
	log.Println(string(bt2))

	//6. ReadRune() && UnReadRune()

	//7. Buffered(): 读取缓冲区中数据字节数(只有执行读才会使用到缓冲区，否则buffered不可用)

	//8. ReadBytes()
	log.Println("ReadBytes() demo...")
	readBuf3 := strings.NewReader("中文;123;456;789")
	reader3 := bufio.NewReader(readBuf3)
	for {
		line, err := reader3.ReadBytes(';')
		if err != nil { //not found delim, return err!=nil
			log.Println("not found delim, read end")
			break
		}
		log.Println(line)
	}

	//9. Flush: submit data, update data at once
	log.Println("write demo...")
	bt4 := bytes.NewBuffer(make([]byte, 0))
	writer4 := bufio.NewWriter(bt4) //create write-buffer and default: 4096

	writer4.WriteString("1234567890") //write 10 bytes in write-buffer, last 4086
	writer4.WriteRune(rune('号'))
	//Buffered returns the number of bytes that have been written into the current buffer.
	//Available returns how many bytes are unused in the buffer.
	log.Println("Buffered:", writer4.Buffered(), ";  Available:", writer4.Available(), ";  bt4:", bt4)

	//执行Flush后，当前的缓冲区被清空，数据将写入到bt4数组中
	writer4.Flush()
	log.Println("Buffered:", writer4.Buffered(), " ;  Available:", writer4.Available(), ";  bt4:", bt4.Len())
	log.Println(bt4)

	//10. writestring(), write(), writebyte(), writerune()

	//11. readfrom
	Readbuf4 := strings.NewReader("www.google.com")
	bt6 := bytes.NewBuffer(make([]byte, 0))

	writer5 := bufio.NewWriter(bt6)
	writer5.ReadFrom(Readbuf4)

	log.Println(bt6)
}

func CheckError(err error) {
	if err != nil {
		log.Println("Faltal error ---> ", err.Error())
		return
	}
}
