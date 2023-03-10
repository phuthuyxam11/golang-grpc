// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: rpc_delete_todo.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeleteType int32

const (
	DeleteType_SOFT_DELETE DeleteType = 0
	DeleteType_HARD_DELETE DeleteType = 1
)

// Enum value maps for DeleteType.
var (
	DeleteType_name = map[int32]string{
		0: "SOFT_DELETE",
		1: "HARD_DELETE",
	}
	DeleteType_value = map[string]int32{
		"SOFT_DELETE": 0,
		"HARD_DELETE": 1,
	}
)

func (x DeleteType) Enum() *DeleteType {
	p := new(DeleteType)
	*p = x
	return p
}

func (x DeleteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeleteType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpc_delete_todo_proto_enumTypes[0].Descriptor()
}

func (DeleteType) Type() protoreflect.EnumType {
	return &file_rpc_delete_todo_proto_enumTypes[0]
}

func (x DeleteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeleteType.Descriptor instead.
func (DeleteType) EnumDescriptor() ([]byte, []int) {
	return file_rpc_delete_todo_proto_rawDescGZIP(), []int{0}
}

type DeleteTodoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DeleteType DeleteType `protobuf:"varint,2,opt,name=delete_type,json=deleteType,proto3,enum=pb.DeleteType" json:"delete_type,omitempty"`
}

func (x *DeleteTodoRequest) Reset() {
	*x = DeleteTodoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_delete_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoRequest) ProtoMessage() {}

func (x *DeleteTodoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_delete_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoRequest.ProtoReflect.Descriptor instead.
func (*DeleteTodoRequest) Descriptor() ([]byte, []int) {
	return file_rpc_delete_todo_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteTodoRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteTodoRequest) GetDeleteType() DeleteType {
	if x != nil {
		return x.DeleteType
	}
	return DeleteType_SOFT_DELETE
}

type DeleteTodoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteTodoResponse) Reset() {
	*x = DeleteTodoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_delete_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTodoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTodoResponse) ProtoMessage() {}

func (x *DeleteTodoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_delete_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTodoResponse.ProtoReflect.Descriptor instead.
func (*DeleteTodoResponse) Descriptor() ([]byte, []int) {
	return file_rpc_delete_todo_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteTodoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_rpc_delete_todo_proto protoreflect.FileDescriptor

var file_rpc_delete_todo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x72, 0x70, 0x63, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x64,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x54, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2f, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x22, 0x2e, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x64, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2a, 0x2e, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x4f, 0x46, 0x54, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x48, 0x41, 0x52, 0x44, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10,
	0x01, 0x42, 0x16, 0x5a, 0x14, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x73, 0x74, 0x75, 0x64, 0x79, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_delete_todo_proto_rawDescOnce sync.Once
	file_rpc_delete_todo_proto_rawDescData = file_rpc_delete_todo_proto_rawDesc
)

func file_rpc_delete_todo_proto_rawDescGZIP() []byte {
	file_rpc_delete_todo_proto_rawDescOnce.Do(func() {
		file_rpc_delete_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_delete_todo_proto_rawDescData)
	})
	return file_rpc_delete_todo_proto_rawDescData
}

var file_rpc_delete_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpc_delete_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_delete_todo_proto_goTypes = []interface{}{
	(DeleteType)(0),            // 0: pb.DeleteType
	(*DeleteTodoRequest)(nil),  // 1: pb.DeleteTodoRequest
	(*DeleteTodoResponse)(nil), // 2: pb.DeleteTodoResponse
}
var file_rpc_delete_todo_proto_depIdxs = []int32{
	0, // 0: pb.DeleteTodoRequest.delete_type:type_name -> pb.DeleteType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_delete_todo_proto_init() }
func file_rpc_delete_todo_proto_init() {
	if File_rpc_delete_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_delete_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_delete_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTodoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_delete_todo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_delete_todo_proto_goTypes,
		DependencyIndexes: file_rpc_delete_todo_proto_depIdxs,
		EnumInfos:         file_rpc_delete_todo_proto_enumTypes,
		MessageInfos:      file_rpc_delete_todo_proto_msgTypes,
	}.Build()
	File_rpc_delete_todo_proto = out.File
	file_rpc_delete_todo_proto_rawDesc = nil
	file_rpc_delete_todo_proto_goTypes = nil
	file_rpc_delete_todo_proto_depIdxs = nil
}
