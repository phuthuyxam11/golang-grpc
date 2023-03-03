package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_stream_server.study/pb"
	"grpc_stream_server.study/services"
	"log"
	"net"
)

func runGrpcServer() {
	grpcServer := grpc.NewServer()
	server := services.Server{}
	pb.RegisterTodoServiceServer(grpcServer, &server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", ":1998")
	if err != nil {
		log.Fatal("cannot create listener")
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
	fmt.Println("server start.....")
}

func main() {
	runGrpcServer()
}
