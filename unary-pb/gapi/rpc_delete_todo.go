package gapi

import (
	"context"
	"unary_pb.study/pb"
)

func (server *Server) DeleteTodo(ctx context.Context, request *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	return &pb.DeleteTodoResponse{Message: "ok"}, nil
}
