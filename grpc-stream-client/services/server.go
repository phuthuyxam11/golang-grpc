package services

import "grpc.stream.client/pb"

type Server struct {
	pb.UnimplementedUploadServiceServer
}
