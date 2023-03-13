package main

import (
	jwt_handler "com.study.grpc.auth.jwt/auth/jwt"
	"com.study.grpc.auth.jwt/gapi"
	"com.study.grpc.auth.jwt/pb"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
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

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("unary interceptor---> is call", info.FullMethod)
	// extract metadata
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	authValues := md["authorization"]
	if len(authValues) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := authValues[0]
	jwtProvider := jwt_handler.NewTokenJWTProvider("secret")
	claims, err := jwtProvider.Validate(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	// set permission and role
	if claims.Role != "admin" {
		return nil, status.Error(codes.PermissionDenied, "no permission to access this RPC")
	}

	return handler(ctx, req)
}

func runGrpcServer() {
	//jwtProvider := jwt_handler.NewTokenJWTProvider("secret")

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor))
	/* with stream server

	grpc.StreamServerInterceptor()

	*/
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
