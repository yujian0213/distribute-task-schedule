package main

import (
	"distribute-task-scheduleule/proto/client_stream"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

// ClientStreamService 定义服务
type ClientStreamService struct {
}

func (s *ClientStreamService) RouteList(srv client_stream.ClientStream_RouteListServer) error {
	for {
		//从流中获取消息
		res, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&client_stream.SimpleResponse{Value: "ok"})
		}
		if err != nil {
			return err
		}
		log.Println(res.StreamData)
	}
}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	lis, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	client_stream.RegisterClientStreamServer(grpcServer, &ClientStreamService{})
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
