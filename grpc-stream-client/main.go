package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc.stream.client/pb"
	"grpc.stream.client/services"
	"log"
	"math"
	"net"
)

type ServerOption interface {
	// contains filtered or unexported methods
}

func main() {
	grpcServer := grpc.NewServer(grpc.MaxSendMsgSize(math.MaxInt), grpc.MaxRecvMsgSize(math.MaxInt))
	server := services.Server{}

	pb.RegisterUploadServiceServer(grpcServer, &server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":1999")
	if err != nil {
		log.Fatal("cannot create listener")
	}
	fmt.Println("server start...")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}
