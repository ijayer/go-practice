// RPC TCP Client

package main

import (
	"os"
	"fmt"
	"net/rpc"
	"log"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}


func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: ", os.Args[0], "server:port")
		os.Exit(1)
	}
	//service := os.Args[1]
	service := "localhost"

	//创建tcp连接
	client, err := rpc.Dial("tcp", service + ":1234")

	args := Args{8, 2}

	var reply int

	err = client.Call("Arith.MultiplyTcp", args, &reply)
	if err != nil{
		log.Println("Fatal Error: ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient

	err = client.Call("Arith.DivideTcp", args, &quot)
	if err != nil{
		log.Println("Fatal Error: ", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}