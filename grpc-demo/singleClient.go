package main

import (
	"context"
	"distribute-task-scheduleule/proto/single"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	//grpc连接
	conn, err := grpc.Dial("localhost:8002", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//client
	client := single.NewGreetsClient(conn)
	_, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//调用
	reply, err := client.SayHello(context.Background(), &single.HelloRequest{Name: "李四", Message: "回来吃饭吗"})
	if err != nil {
		log.Fatalf("couldn not greet: %v", err)
	}
	log.Println(reply.Name, reply.Message)
}
