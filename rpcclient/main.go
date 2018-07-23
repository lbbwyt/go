// rpcclient project main.go
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

//我们把上面的服务端和客户端的代码分别编译，然后先把服务端开启，
//在客户端目录下运行 go run main.go 127.0.0.1
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]
	//	创建一个客户端，建立客户端和服务器端的连接:
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Args{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	//异步调用
	quot := new(Quotient)
	divCall := client.Go("Arith.Divide", args, quot, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	//quot = replyCall.Reply.(*Quotient)
	if replyCall.Error != nil {
		log.Fatal("arith error:", replyCall.Error)
	}
	fmt.Printf("异步调用 Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)

	// 同步调用
	var q Quotient
	err = client.Call("Arith.Divide", args, &q)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("同步调用 Arith: %d/%d=%d remainder  %d\n", args.A, args.B, q.Quo, q.Rem)
}
