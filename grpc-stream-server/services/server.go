package services

import "grpc_stream_server.study/pb"

type Server struct {
	pb.UnimplementedTodoServiceServer
}
