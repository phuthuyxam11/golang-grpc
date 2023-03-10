// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: service_todo.proto

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

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodoStream(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (TodoService_CreateTodoStreamClient, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodoStream(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (TodoService_CreateTodoStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &TodoService_ServiceDesc.Streams[0], "/pb.TodoService/CreateTodoStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &todoServiceCreateTodoStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TodoService_CreateTodoStreamClient interface {
	Recv() (*CreateTodoResponse, error)
	grpc.ClientStream
}

type todoServiceCreateTodoStreamClient struct {
	grpc.ClientStream
}

func (x *todoServiceCreateTodoStreamClient) Recv() (*CreateTodoResponse, error) {
	m := new(CreateTodoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	CreateTodoStream(*CreateTodoRequest, TodoService_CreateTodoStreamServer) error
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) CreateTodoStream(*CreateTodoRequest, TodoService_CreateTodoStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateTodoStream not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_CreateTodoStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateTodoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TodoServiceServer).CreateTodoStream(m, &todoServiceCreateTodoStreamServer{stream})
}

type TodoService_CreateTodoStreamServer interface {
	Send(*CreateTodoResponse) error
	grpc.ServerStream
}

type todoServiceCreateTodoStreamServer struct {
	grpc.ServerStream
}

func (x *todoServiceCreateTodoStreamServer) Send(m *CreateTodoResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateTodoStream",
			Handler:       _TodoService_CreateTodoStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service_todo.proto",
}
