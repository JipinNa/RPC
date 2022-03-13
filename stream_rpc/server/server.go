package main

import (
	"RPC/stream_rpc/proto"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

const PORT = ":50052"

type Server struct {
}

func (s *Server) GetStream(req *proto.StreamReqData, server proto.Greeter_GetStreamServer) error {
	var count int
	for {
		err := server.Send(&proto.StreamRspData{
			Data: fmt.Sprintf("%v;count:%d;data:%s", time.Now().Unix(), count, req.Data),
		})
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second)
		if count > 10 {
			return nil
		}
		count++
	}
	return nil
}

func (s *Server) PutStream(server proto.Greeter_PutStreamServer) error {
	return nil
}

func (s *Server) AllStream(server proto.Greeter_AllStreamServer) error {
	return nil
}

func main() {
	lis, err := net.Listen(`tcp`, PORT)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
