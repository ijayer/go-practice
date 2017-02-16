// RPC TCP Server

package main

import (
	"errors"
	"net/rpc"
	"net"
	"os"
	"log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

//rpc:      function
//param:    args *Reveive   args *Reply
//return:   error
func (t *Arith) MultiplyTcp(receive *Args, reply *int) error{
	*reply = receive.A * receive.B
	return nil
}

//
func (t *Arith) DivideTcp(receive *Args, reply *Quotient) error{
	if receive.B == 0 {
		return errors.New("divide by zero")
	}
	reply.Quo = receive.A / receive.B
	reply.Rem = receive.A % receive.B
	return nil
}

func main(){
	// 分配内存空间
	arith := new(Arith)

	// 给airth类型注册rpc服务
	rpc.Register(arith)

	// 注册rpc的tcp服务
	// 1. 建立tcp服务
	// 2. 监听端口
	// 3. 阻塞接受

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil{
		log.Println("Fatal Error: ", err.Error())
		os.Exit(1)
	}

	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil{
		log.Println("Fatal Error: ", err.Error())
		os.Exit(1)
	}

	// 该tcp连接属于阻塞单用户模式, 设计多并发可使用goroutine
	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}