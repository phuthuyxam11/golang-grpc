package gapi

import (
	"com.study.grpc.auth.jwt/pb"
	"context"
)

func (server *Server) CreateTodo(ctx context.Context, request *pb.CreateTodoRequest) (*pb.CreateTodoResponse, error) {
	return &pb.CreateTodoResponse{
		Status: "ok",
	}, nil
}
