package main

import (
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"grpc_stream_server.study/pb"
	"grpc_stream_server.study/services"
	"log"
	"net"
)

func runGrpcServer() {
	severCert, err := tls.LoadX509KeyPair("ssl/server-cert.pem", "ssl/server-key.pem")
	if err != nil {
		log.Fatal("can't create server cert")
		return
	}
	configTls := &tls.Config{
		ClientAuth:   tls.NoClientCert,
		Certificates: []tls.Certificate{severCert},
	}
	credential := credentials.NewTLS(configTls)
	grpcServer := grpc.NewServer(grpc.Creds(credential))
	server := services.Server{}
	pb.RegisterTodoServiceServer(grpcServer, &server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", ":1998")
	if err != nil {
		log.Fatal("cannot create listener")
	}
	fmt.Println("start server at :1998")
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
	fmt.Println("server start.....")
}

func main() {
	runGrpcServer()
}
