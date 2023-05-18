package main

import (
	"context"
	"distribute-task-scheduleule/proto/single"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *single.HelloRequest) (*single.HelloReply, error) {
	return &single.HelloReply{Name: "张三", Message: "不回来了"}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":8002")
	if err != nil {
		fmt.Println(err)
		return
	}
	//server
	server := grpc.NewServer()
	//注册服务
	single.RegisterGreetsServer(server, &Server{})
	//启动服务
	err = server.Serve(lis)
	if err != nil {
		fmt.Println(err)
		return
	}
}
