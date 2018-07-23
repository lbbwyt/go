// 使用http实现的go rpc 服务端
package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

//需要定义传入参数和返回参数的数据结构：
type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

//定义一个服务对象
type Arith int

//实现这个类型的两个方法
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("分母不能为0")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	//注册了一个Arith的RPC服务
	rpc.Register(arith)
	//然后通过rpc.HandleHTTP函数把该服务注册到了HTTP协议上
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
