package gapi

import "unary_pb.study/pb"

// server is used to implement helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedTodoServiceServer
}
