package gapi

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
	"unary_pb.study/pb"
)

func (server *Server) DeleteTodoWithDeadline(ctx context.Context, request *pb.DeleteTodoRequest) (*pb.DeleteTodoResponse, error) {
	id := request.GetId()
	if id < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID is not < 0")
	}

	for i := 0; i < 3; i++ {
		// check client cancel
		if ctx.Err() == context.Canceled {
			log.Println("context.canceled..")
			return nil, status.Errorf(codes.Canceled, "client canceled req")
		}
		time.Sleep(1 * time.Second)
	}

	return &pb.DeleteTodoResponse{Message: "ok"}, nil
}
