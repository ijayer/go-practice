package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jlaffaye/ftp"
)

const (
	User = "xxxxx"
	Pwd  = "xxxxx"
	Addr = "xxxxx"
)

func main() {
	url := "ftp://192.168.1.8/测试版/v1.0.4/terminal-setup.exe"
	path := strings.Replace(url, "ftp://192.168.1.8", "", -1)
	// connect
	connect(Addr, path)
}

// Connect connect to ftp server
func connect(addr, path string) {
	conn, err := ftp.Connect(addr)
	if err != nil {
		fmt.Println("Error: connect ftp server.", err.Error())
		return
	}

	if err := conn.Login(User, Pwd); err != nil {
		fmt.Println("Error: can't login ftp server.", err.Error())
		return
	}

	i, err := conn.FileSize("/测试版/v1.0.4/terminal-setup.exe")
	if err != nil {
		fmt.Println(err.Error())
	}
	size := float64(i) / (1024 * 1024)
	mb := fmt.Sprintf("%.1f", size)
	fmt.Printf("size = %sMB\n", mb)

	download(conn, path)
}

// Download download a file from remote ftp server
func download(conn *ftp.ServerConn, path string) {
	resp, err := conn.Retr("/测试版/v1.0.4/terminal-setup.exe")
	if err != nil {
		fmt.Println("Error: download file from ftp server.", err.Error())
		return
	}
	defer resp.Close()

	b, err := ioutil.ReadAll(resp)
	if err != nil {
		fmt.Println("Error: read file.", err.Error())
		return
	}
	mds := fmt.Sprintf("%x", sha256.Sum256(b))
	fmt.Println(mds)
}
