// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: loadBalancer/loadBalancer.proto

package loadBalancer

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_loadBalancer_loadBalancer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_loadBalancer_loadBalancer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_loadBalancer_loadBalancer_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

var File_loadBalancer_loadBalancer_proto protoreflect.FileDescriptor

var file_loadBalancer_loadBalancer_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x32, 0x85, 0x01, 0x0a, 0x13, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0f, 0x2e, 0x77, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x0f, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x10, 0x5a,
	0x0e, 0x2e, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_loadBalancer_loadBalancer_proto_rawDescOnce sync.Once
	file_loadBalancer_loadBalancer_proto_rawDescData = file_loadBalancer_loadBalancer_proto_rawDesc
)

func file_loadBalancer_loadBalancer_proto_rawDescGZIP() []byte {
	file_loadBalancer_loadBalancer_proto_rawDescOnce.Do(func() {
		file_loadBalancer_loadBalancer_proto_rawDescData = protoimpl.X.CompressGZIP(file_loadBalancer_loadBalancer_proto_rawDescData)
	})
	return file_loadBalancer_loadBalancer_proto_rawDescData
}

var file_loadBalancer_loadBalancer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_loadBalancer_loadBalancer_proto_goTypes = []interface{}{
	(*Request)(nil),       // 0: worker.Request
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_loadBalancer_loadBalancer_proto_depIdxs = []int32{
	0, // 0: worker.loadBalancerService.Register:input_type -> worker.Request
	0, // 1: worker.loadBalancerService.DeRegister:input_type -> worker.Request
	1, // 2: worker.loadBalancerService.Register:output_type -> google.protobuf.Empty
	1, // 3: worker.loadBalancerService.DeRegister:output_type -> google.protobuf.Empty
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_loadBalancer_loadBalancer_proto_init() }
func file_loadBalancer_loadBalancer_proto_init() {
	if File_loadBalancer_loadBalancer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_loadBalancer_loadBalancer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_loadBalancer_loadBalancer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_loadBalancer_loadBalancer_proto_goTypes,
		DependencyIndexes: file_loadBalancer_loadBalancer_proto_depIdxs,
		MessageInfos:      file_loadBalancer_loadBalancer_proto_msgTypes,
	}.Build()
	File_loadBalancer_loadBalancer_proto = out.File
	file_loadBalancer_loadBalancer_proto_rawDesc = nil
	file_loadBalancer_loadBalancer_proto_goTypes = nil
	file_loadBalancer_loadBalancer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoadBalancerServiceClient is the client API for LoadBalancerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadBalancerServiceClient interface {
	Register(ctx context.Context, in *Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeRegister(ctx context.Context, in *Request, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type loadBalancerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoadBalancerServiceClient(cc grpc.ClientConnInterface) LoadBalancerServiceClient {
	return &loadBalancerServiceClient{cc}
}

func (c *loadBalancerServiceClient) Register(ctx context.Context, in *Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/worker.loadBalancerService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancerServiceClient) DeRegister(ctx context.Context, in *Request, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/worker.loadBalancerService/DeRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoadBalancerServiceServer is the server API for LoadBalancerService service.
type LoadBalancerServiceServer interface {
	Register(context.Context, *Request) (*emptypb.Empty, error)
	DeRegister(context.Context, *Request) (*emptypb.Empty, error)
}

// UnimplementedLoadBalancerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadBalancerServiceServer struct {
}

func (*UnimplementedLoadBalancerServiceServer) Register(context.Context, *Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedLoadBalancerServiceServer) DeRegister(context.Context, *Request) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeRegister not implemented")
}

func RegisterLoadBalancerServiceServer(s *grpc.Server, srv LoadBalancerServiceServer) {
	s.RegisterService(&_LoadBalancerService_serviceDesc, srv)
}

func _LoadBalancerService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.loadBalancerService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).Register(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancerService_DeRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancerServiceServer).DeRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/worker.loadBalancerService/DeRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancerServiceServer).DeRegister(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadBalancerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "worker.loadBalancerService",
	HandlerType: (*LoadBalancerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _LoadBalancerService_Register_Handler,
		},
		{
			MethodName: "DeRegister",
			Handler:    _LoadBalancerService_DeRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loadBalancer/loadBalancer.proto",
}
