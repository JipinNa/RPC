package main

import (
	"RPC/rpc/protobuf"
	"RPC/rpc/server"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"runtime"
)

func main() {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error:
			fmt.Println(err)
		}
	}()
	g := grpc.NewServer()
	protobuf.RegisterGreeterServer(g, &server.Server{})
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = g.Serve(listen)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
