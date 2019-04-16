// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/hello.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9418a89e4940c19, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Karma                float32  `protobuf:"fixed32,2,opt,name=karma,proto3" json:"karma,omitempty"`
	Dlt                  int64    `protobuf:"varint,3,opt,name=dlt,proto3" json:"dlt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9418a89e4940c19, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetKarma() float32 {
	if m != nil {
		return m.Karma
	}
	return 0
}

func (m *HelloReply) GetDlt() int64 {
	if m != nil {
		return m.Dlt
	}
	return 0
}

type TimRequest struct {
	Year                 int64    `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimRequest) Reset()         { *m = TimRequest{} }
func (m *TimRequest) String() string { return proto.CompactTextString(m) }
func (*TimRequest) ProtoMessage()    {}
func (*TimRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9418a89e4940c19, []int{2}
}

func (m *TimRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimRequest.Unmarshal(m, b)
}
func (m *TimRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimRequest.Marshal(b, m, deterministic)
}
func (m *TimRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimRequest.Merge(m, src)
}
func (m *TimRequest) XXX_Size() int {
	return xxx_messageInfo_TimRequest.Size(m)
}
func (m *TimRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TimRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TimRequest proto.InternalMessageInfo

func (m *TimRequest) GetYear() int64 {
	if m != nil {
		return m.Year
	}
	return 0
}

type TimReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimReply) Reset()         { *m = TimReply{} }
func (m *TimReply) String() string { return proto.CompactTextString(m) }
func (*TimReply) ProtoMessage()    {}
func (*TimReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9418a89e4940c19, []int{3}
}

func (m *TimReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimReply.Unmarshal(m, b)
}
func (m *TimReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimReply.Marshal(b, m, deterministic)
}
func (m *TimReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimReply.Merge(m, src)
}
func (m *TimReply) XXX_Size() int {
	return xxx_messageInfo_TimReply.Size(m)
}
func (m *TimReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TimReply.DiscardUnknown(m)
}

var xxx_messageInfo_TimReply proto.InternalMessageInfo

func (m *TimReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
	proto.RegisterType((*TimRequest)(nil), "helloworld.TimRequest")
	proto.RegisterType((*TimReply)(nil), "helloworld.TimReply")
}

func init() { proto.RegisterFile("pb/hello.proto", fileDescriptor_b9418a89e4940c19) }

var fileDescriptor_b9418a89e4940c19 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x49, 0xa2, 0x6d, 0x7d, 0x88, 0xda, 0x67, 0xd1, 0x18, 0x5c, 0x84, 0xc1, 0x45, 0x56,
	0x09, 0xe8, 0x21, 0x74, 0xa5, 0x90, 0x06, 0x5c, 0x4f, 0xe9, 0x23, 0x06, 0x67, 0x32, 0x71, 0x32,
	0x22, 0xa1, 0x74, 0xe3, 0x15, 0x3c, 0x8a, 0x47, 0xf1, 0x0a, 0x1e, 0x44, 0xf2, 0x6a, 0x69, 0x45,
	0x71, 0xf7, 0xff, 0x33, 0xdf, 0xfc, 0x6f, 0xfe, 0x07, 0x07, 0xcd, 0x2c, 0x7b, 0x20, 0xa5, 0x4c,
	0xda, 0x58, 0xe3, 0x0c, 0x02, 0x9b, 0x17, 0x63, 0xd5, 0x3c, 0x3a, 0x2f, 0x8d, 0x29, 0x15, 0x65,
	0xb2, 0xa9, 0x32, 0x59, 0xd7, 0xc6, 0x49, 0x57, 0x99, 0xba, 0x5d, 0x91, 0x42, 0xc0, 0xfe, 0x4d,
	0xcf, 0xe6, 0xf4, 0xf4, 0x4c, 0xad, 0x43, 0x84, 0x9d, 0x5a, 0x6a, 0x0a, 0xbd, 0xd8, 0x4b, 0xf6,
	0x72, 0xd6, 0xe2, 0x16, 0xe0, 0x9b, 0x69, 0x54, 0x87, 0x21, 0x0c, 0x35, 0xb5, 0xad, 0x2c, 0xd7,
	0xd0, 0xda, 0xe2, 0x04, 0x76, 0x1f, 0xa5, 0xd5, 0x32, 0xf4, 0x63, 0x2f, 0xf1, 0xf3, 0x95, 0xc1,
	0x23, 0x08, 0xe6, 0xca, 0x85, 0x41, 0xec, 0x25, 0x41, 0xde, 0x4b, 0x11, 0x03, 0x14, 0x95, 0xde,
	0x9a, 0xd8, 0x91, 0xb4, 0x1c, 0x16, 0xe4, 0xac, 0xc5, 0x05, 0x8c, 0x98, 0xf8, 0x77, 0xde, 0xe5,
	0xbb, 0x07, 0xc3, 0x6b, 0x4b, 0xe4, 0xc8, 0xe2, 0x3d, 0x8c, 0xa6, 0xb2, 0xe3, 0x6f, 0x62, 0x98,
	0x6e, 0xea, 0xa7, 0xdb, 0xed, 0xa2, 0x93, 0x3f, 0x6e, 0x1a, 0xd5, 0x89, 0xb3, 0xd7, 0x8f, 0xcf,
	0x37, 0xff, 0x18, 0xc7, 0xbc, 0x25, 0x66, 0xb2, 0x45, 0xdf, 0x7d, 0x89, 0x77, 0x30, 0x98, 0xca,
	0xae, 0xa8, 0x34, 0xfe, 0x78, 0xbc, 0x29, 0x10, 0x4d, 0x7e, 0x9d, 0xf7, 0x91, 0xa7, 0x1c, 0x39,
	0xc6, 0x43, 0x8e, 0x2c, 0x2a, 0x9d, 0x2d, 0xfa, 0x6a, 0xcb, 0xd9, 0x80, 0x17, 0x7f, 0xf5, 0x15,
	0x00, 0x00, 0xff, 0xff, 0xd5, 0xe7, 0x9a, 0x24, 0xb4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayTim(ctx context.Context, in *TimRequest, opts ...grpc.CallOption) (*TimReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayTim(ctx context.Context, in *TimRequest, opts ...grpc.CallOption) (*TimReply, error) {
	out := new(TimReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayTim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayTim(context.Context, *TimRequest) (*TimReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedGreeterServer) SayTim(ctx context.Context, req *TimRequest) (*TimReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayTim not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayTim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayTim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayTim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayTim(ctx, req.(*TimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayTim",
			Handler:    _Greeter_SayTim_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/hello.proto",
}
