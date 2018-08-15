// grpcclient project main.go
package main

import (
	"flag"
	"fmt"
	pb "grpcclient/greeter"
	"log"
	"sync"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

//是否异步调用
var isAsync bool

func init() {
	flag.BoolVar(&isAsync, "a", true, "是否异步调用，默认是")
}

//GreeterClient is the client API for Greeter service
//GreeterClient中定义Greeter服务中可以远程调用的方法
func invoke(c pb.GreeterClient, name string) {
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet :%v", err)
	}
	_ = r
}

//同步调用
func syncTest(c pb.GreeterClient, name string) {
	i := 10000
	t := time.Now().UnixNano()
	for ; i > 0; i-- {
		invoke(c, name)
	}
	//计算服务调用耗时，本地测试大概230ms
	fmt.Println("took", (time.Now().UnixNano()-t)/1000000, "ms")
}

func asyncTest(c [20]pb.GreeterClient, name string) {
	var wg sync.WaitGroup
	wg.Add(10000)
	t := time.Now().UnixNano()
	i := 10000
	for ; i > 0; i-- {
		go func() {
			invoke(c[i%20], name)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("took", (time.Now().UnixNano()-t)/1000000, "ms")

}

func main() {
	flag.Parse()
	//创建客户端连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	var c [20]pb.GreeterClient
	name := defaultName

	i := 0
	for ; i < 20; i++ {
		c[i] = pb.NewGreeterClient(conn)
		invoke(c[i], name)
	}

	if isAsync {
		asyncTest(c, name)
	} else {
		syncTest(c[0], name)
	}

}
