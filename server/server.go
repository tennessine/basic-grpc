package main

import (
	"context"
	"github.com/tennessine/tennessine/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SimpleService struct {
}

func (s *SimpleService) Route(ctx context.Context, req *proto.SimpleRequest) (*proto.SimpleResponse, error) {
	res := proto.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}

const (
	Address string = ":8080"
	Network string = "tcp"
)

func main() {
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listening...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务\
	proto.RegisterSimpleServer(grpcServer, &SimpleService{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
