package gapi

import "com.study.grpc.auth.jwt/pb"

// server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedTodoServiceServer
}
