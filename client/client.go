package main

import (
	"context"
	"github.com/tennessine/tennessine/proto"
	"google.golang.org/grpc"
	"log"
)

const (
	Address string = ":8080"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient := proto.NewSimpleClient(conn)
	req := proto.SimpleRequest{Data: "grpc"}

	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call route err: %v", err)
	}
	log.Println(res)
}
