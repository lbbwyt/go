// jsonrpcserver project main.go
//json-rpc是基于TCP协议实现的，目前它还不支持HTTP方式。
package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

//在mian.go目录下执行go main.go :1234
func main() {
	airth := new(Arith)
	rpc.Register(airth)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		//使用并发的方式处理客户端的服务请求调用
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	defer conn.Close()
	jsonrpc.ServeConn(conn)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
