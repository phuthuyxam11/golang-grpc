package gapi

import (
	"context"
	"unary_pb.study/pb"
)

func (server *Server) CreateTodo(ctx context.Context, request *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	return &pb.CreateTodoResponse{
		Status: "ok",
	}, nil
}
