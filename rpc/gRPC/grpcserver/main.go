// grpcserver project main.go
// grpc 服务端
package main

import (
	pb "grpcserver/greeter"
	"log"
	"net"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
)

//定义需监听的端口
const (
	port = ":50051"
)

//实现GreeterServer接口
type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{
		Message: "Hello" + in.GetName(),
	}, nil
}

func main() {
	//tcp方式监听端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", port)
	}
	s := grpc.NewServer()
	//注册GreeterServer接口的实现来响应客户端的请求
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
