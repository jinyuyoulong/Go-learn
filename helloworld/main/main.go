package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "v5u.win/add/helloworld"
)

type HelloServer struct {
}

func (h HelloServer) SayHello(c context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	//创建gRPC 服务器，将我们实现的Greeter服务绑定到一个端口
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &HelloServer{})
	reflection.Register(s)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
