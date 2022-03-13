package server

import (
	"RPC/rpc/protobuf"
	"context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, request *protobuf.HelloRequest) (*protobuf.HelloReply, error) {
	return &protobuf.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}
