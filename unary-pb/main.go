package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"unary_pb.study/gapi"
	"unary_pb.study/pb"
)

func runGatewayServer() {
	fmt.Println("start server ...")
	server := gapi.Server{}
	grpcMux := runtime.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	err := pb.RegisterTodoServiceHandlerServer(ctx, grpcMux, &server)
	if err != nil {
		log.Fatalln("error register server: ", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listen, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalln("fail to listen port")
	}

	if err := http.Serve(listen, mux); err != nil {
		log.Fatalf("failed to start http serve: %v", err)
	}
}

func runGrpcServer() {
	grpcServer := grpc.NewServer()
	server := gapi.Server{}

	pb.RegisterTodoServiceServer(grpcServer, &server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":1997")
	if err != nil {
		log.Fatal("cannot create listener")
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func main() {
	go runGatewayServer()
	runGrpcServer()
}
