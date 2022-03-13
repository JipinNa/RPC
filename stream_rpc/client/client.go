package main

import (
	"RPC/stream_rpc/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(`localhost:50052`, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer func() {
		conn.Close()
	}()
	c := proto.NewGreeterClient(conn)
	client, err := c.GetStream(context.Background(), &proto.StreamReqData{
		Data: "come on",
	})
	if err != nil {
		panic(err)
	}
	for {
		if rsp, err := client.Recv(); err == nil {
			fmt.Println(rsp)
		} else {
			break
		}
	}
}
