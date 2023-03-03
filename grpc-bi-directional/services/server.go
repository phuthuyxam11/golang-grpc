package services

import "grpc.bi.directional/pb"

type Server struct {
	pb.UnimplementedSplitFileServiceServer
}
