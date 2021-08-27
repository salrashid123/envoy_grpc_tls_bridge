// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.13.0
// source: echo/echo.proto

package echo

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_echo_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_echo_echo_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type EchoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoReply) Reset() {
	*x = EchoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoReply) ProtoMessage() {}

func (x *EchoReply) ProtoReflect() protoreflect.Message {
	mi := &file_echo_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoReply.ProtoReflect.Descriptor instead.
func (*EchoReply) Descriptor() ([]byte, []int) {
	return file_echo_echo_proto_rawDescGZIP(), []int{1}
}

func (x *EchoReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_echo_echo_proto protoreflect.FileDescriptor

var file_echo_echo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x22, 0x21, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x09, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x78, 0x0a, 0x0a, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12,
	0x30, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x11, 0x2e, 0x65, 0x63,
	0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x0e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x45, 0x63,
	0x68, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1e, 0x5a, 0x1c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6c, 0x72, 0x61, 0x73,
	0x68, 0x69, 0x64, 0x31, 0x32, 0x33, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_echo_echo_proto_rawDescOnce sync.Once
	file_echo_echo_proto_rawDescData = file_echo_echo_proto_rawDesc
)

func file_echo_echo_proto_rawDescGZIP() []byte {
	file_echo_echo_proto_rawDescOnce.Do(func() {
		file_echo_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_echo_proto_rawDescData)
	})
	return file_echo_echo_proto_rawDescData
}

var file_echo_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_echo_echo_proto_goTypes = []interface{}{
	(*EchoRequest)(nil), // 0: echo.EchoRequest
	(*EchoReply)(nil),   // 1: echo.EchoReply
}
var file_echo_echo_proto_depIdxs = []int32{
	0, // 0: echo.EchoServer.SayHello:input_type -> echo.EchoRequest
	0, // 1: echo.EchoServer.SayHelloStream:input_type -> echo.EchoRequest
	1, // 2: echo.EchoServer.SayHello:output_type -> echo.EchoReply
	1, // 3: echo.EchoServer.SayHelloStream:output_type -> echo.EchoReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_echo_echo_proto_init() }
func file_echo_echo_proto_init() {
	if File_echo_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_echo_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoReply); i {
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
			RawDescriptor: file_echo_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_echo_proto_goTypes,
		DependencyIndexes: file_echo_echo_proto_depIdxs,
		MessageInfos:      file_echo_echo_proto_msgTypes,
	}.Build()
	File_echo_echo_proto = out.File
	file_echo_echo_proto_rawDesc = nil
	file_echo_echo_proto_goTypes = nil
	file_echo_echo_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EchoServerClient is the client API for EchoServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoServerClient interface {
	SayHello(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
	SayHelloStream(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (EchoServer_SayHelloStreamClient, error)
}

type echoServerClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServerClient(cc grpc.ClientConnInterface) EchoServerClient {
	return &echoServerClient{cc}
}

func (c *echoServerClient) SayHello(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := c.cc.Invoke(ctx, "/echo.EchoServer/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServerClient) SayHelloStream(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (EchoServer_SayHelloStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EchoServer_serviceDesc.Streams[0], "/echo.EchoServer/SayHelloStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServerSayHelloStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EchoServer_SayHelloStreamClient interface {
	Recv() (*EchoReply, error)
	grpc.ClientStream
}

type echoServerSayHelloStreamClient struct {
	grpc.ClientStream
}

func (x *echoServerSayHelloStreamClient) Recv() (*EchoReply, error) {
	m := new(EchoReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EchoServerServer is the server API for EchoServer service.
type EchoServerServer interface {
	SayHello(context.Context, *EchoRequest) (*EchoReply, error)
	SayHelloStream(*EchoRequest, EchoServer_SayHelloStreamServer) error
}

// UnimplementedEchoServerServer can be embedded to have forward compatible implementations.
type UnimplementedEchoServerServer struct {
}

func (*UnimplementedEchoServerServer) SayHello(context.Context, *EchoRequest) (*EchoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedEchoServerServer) SayHelloStream(*EchoRequest, EchoServer_SayHelloStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStream not implemented")
}

func RegisterEchoServerServer(s *grpc.Server, srv EchoServerServer) {
	s.RegisterService(&_EchoServer_serviceDesc, srv)
}

func _EchoServer_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServerServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.EchoServer/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServerServer).SayHello(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoServer_SayHelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EchoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServerServer).SayHelloStream(m, &echoServerSayHelloStreamServer{stream})
}

type EchoServer_SayHelloStreamServer interface {
	Send(*EchoReply) error
	grpc.ServerStream
}

type echoServerSayHelloStreamServer struct {
	grpc.ServerStream
}

func (x *echoServerSayHelloStreamServer) Send(m *EchoReply) error {
	return x.ServerStream.SendMsg(m)
}

var _EchoServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.EchoServer",
	HandlerType: (*EchoServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _EchoServer_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloStream",
			Handler:       _EchoServer_SayHelloStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "echo/echo.proto",
}
