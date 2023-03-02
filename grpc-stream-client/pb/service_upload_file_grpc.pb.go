// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service_upload_file.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadServiceClient interface {
	UploadFileStream(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadFileStreamClient, error)
}

type uploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServiceClient(cc grpc.ClientConnInterface) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) UploadFileStream(ctx context.Context, opts ...grpc.CallOption) (UploadService_UploadFileStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &UploadService_ServiceDesc.Streams[0], "/pb.UploadService/UploadFileStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploadServiceUploadFileStreamClient{stream}
	return x, nil
}

type UploadService_UploadFileStreamClient interface {
	Send(*UploadFileRequest) error
	CloseAndRecv() (*UploadFileResponse, error)
	grpc.ClientStream
}

type uploadServiceUploadFileStreamClient struct {
	grpc.ClientStream
}

func (x *uploadServiceUploadFileStreamClient) Send(m *UploadFileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploadServiceUploadFileStreamClient) CloseAndRecv() (*UploadFileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadServiceServer is the server API for UploadService service.
// All implementations must embed UnimplementedUploadServiceServer
// for forward compatibility
type UploadServiceServer interface {
	UploadFileStream(UploadService_UploadFileStreamServer) error
	mustEmbedUnimplementedUploadServiceServer()
}

// UnimplementedUploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUploadServiceServer struct {
}

func (UnimplementedUploadServiceServer) UploadFileStream(UploadService_UploadFileStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFileStream not implemented")
}
func (UnimplementedUploadServiceServer) mustEmbedUnimplementedUploadServiceServer() {}

// UnsafeUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServiceServer will
// result in compilation errors.
type UnsafeUploadServiceServer interface {
	mustEmbedUnimplementedUploadServiceServer()
}

func RegisterUploadServiceServer(s grpc.ServiceRegistrar, srv UploadServiceServer) {
	s.RegisterService(&UploadService_ServiceDesc, srv)
}

func _UploadService_UploadFileStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploadServiceServer).UploadFileStream(&uploadServiceUploadFileStreamServer{stream})
}

type UploadService_UploadFileStreamServer interface {
	SendAndClose(*UploadFileResponse) error
	Recv() (*UploadFileRequest, error)
	grpc.ServerStream
}

type uploadServiceUploadFileStreamServer struct {
	grpc.ServerStream
}

func (x *uploadServiceUploadFileStreamServer) SendAndClose(m *UploadFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploadServiceUploadFileStreamServer) Recv() (*UploadFileRequest, error) {
	m := new(UploadFileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploadService_ServiceDesc is the grpc.ServiceDesc for UploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadFileStream",
			Handler:       _UploadService_UploadFileStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service_upload_file.proto",
}