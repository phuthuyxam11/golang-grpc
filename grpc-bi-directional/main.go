package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc.bi.directional/pb"
	"grpc.bi.directional/services"
	"log"
	"net"
)

func main() {
	grpcServer := grpc.NewServer()
	server := services.Server{}
	pb.RegisterSplitFileServiceServer(grpcServer, &server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", ":1998")
	if err != nil {
		log.Fatal("cannot create listener")
	}
	fmt.Println("server start.....")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}

}
