package main

import (
	"RPC/rpc/protobuf"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func() {
		conn.Close()
	}()
	c := protobuf.NewGreeterClient(conn)
	rsp, err := c.SayHello(context.Background(), &protobuf.HelloRequest{
		Name: "oj8k",
	})
	fmt.Println(rsp)
}
