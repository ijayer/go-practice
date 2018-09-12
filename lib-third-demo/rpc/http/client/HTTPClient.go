// RPC http client

package client

import (
	"fmt"
	"net/rpc"
	"os"

	"github.com/labstack/gommon/log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	// 获取命令行参数(os/Args包)
	// go run HTTPClient.go localhost
	if len(os.Args) != 1 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	//serverAddress := os.Args[1]
	serverAddress := "localhost"

	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing", err)
	}

	// Synchronous call 同步调用
	args := Args{17, 8}

	// Call函数参数：
	// 1. 要调用的函数名字
	// 2. 传递的参数
	// 3. 返回的参数，指针类型
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error: ", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error: ", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
