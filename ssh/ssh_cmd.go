package main

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"time"
)

const (
	User = "xxx"
	Host = "xxx"
	Port = "22"
)

type SSH struct {
	Username string
	Password string
	Host     string
}

func NewSSH(pwd string) *SSH {
	return &SSH{
		Username: User,
		Password: pwd,
		Host:     Host,
	}
}

func main() {
	ssh := NewSSH("xxx")

	session, err := ssh.connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Run("ls /")

	time.Sleep(1 * time.Second)
	command(session)
}

func (s *SSH) connect() (*ssh.Session, error) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(s.Password))
	clientConfig = &ssh.ClientConfig{User: s.Username, Auth: auth, Timeout: 30 * time.Second}

	// connect to ssh
	addr = fmt.Sprintf("%s:%s", s.Host, Port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}
	return session, nil
}

func command(session *ssh.Session) {
	//fd := int(os.Stdin.Fd())
	//oldState, err := terminal.MakeRaw(fd)
	//if err != nil {
	//	panic(err)
	//}
	//defer terminal.Restore(fd, oldState)

	// execute command
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	//termWidth, termHeight, err := terminal.GetSize(fd)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}

	// Set up terminal modes
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // enable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	// Request pseudo terminal
	if err := session.RequestPty("xterm-256color", 25, 80, modes); err != nil {
		fmt.Println(err.Error())
		return
	}

	session.Run("vim /home/shell/api")

	// Hosted shell
	//session.Shell()
	//session.Wait()
}
