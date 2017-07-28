package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	host := flag.String("h", "", "Remote host address")
	user := flag.String("u", "", "Remote host account")
	pwd := flag.String("p", "", "Remote host password")
	flag.Parse()

	if *host == " " && *user == "" && *pwd == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var (
		err        error
		sftpClient *sftp.Client
	)

	sftpClient, err = connect(*user, *pwd, *host)
	if err != nil {
		fmt.Println("#Error(connect): ", err.Error())
		return
	}
	defer sftpClient.Close()

	var localFilePath = `D:\share\app`
	var remoteDir = "/home/sftp"

	srcFile, err := os.Open(localFilePath)
	if err != nil {
		fmt.Println("#Error(Open): ", err.Error())
		return
	}
	defer srcFile.Close()

	var remoteFileName = filepath.Base(localFilePath)
	dstFile, err := sftpClient.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		fmt.Println("#Error(Create): ", err.Error())
		return
	}
	defer dstFile.Close()

	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, srcFile)
	if err != nil {
		fmt.Println("#Error(Copy): ", err.Error())
		return
	}
	_, err = dstFile.Write(buf.Bytes())
	if err != nil {
		fmt.Println("#Error(Write): ", err.Error())
		return
	}
}

func connect(user, pwd, host string) (*sftp.Client, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		sshClient    *ssh.Client
		sftpClient   *sftp.Client
		err          error
	)

	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(pwd))
	clientConfig = &ssh.ClientConfig{
		User:    user,
		Auth:    auth,
		Timeout: 60 * time.Second,
	}

	addr = fmt.Sprintf("%s:%d", host, 22)
	fmt.Printf("Addr: %s\n", addr)
	if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		fmt.Println("#Error(Dial): ", err.Error())
		return nil, err
	}

	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		fmt.Println("#Error(NewClient): ", err.Error())
		return nil, err
	}
	return sftpClient, nil
}
