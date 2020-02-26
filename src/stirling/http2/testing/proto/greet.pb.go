// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/stirling/http2/testing/proto/greet.proto

package greetpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type HelloRequest struct {
	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Count int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *HelloRequest) Reset()      { *m = HelloRequest{} }
func (*HelloRequest) ProtoMessage() {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d33060a754189a, []int{0}
}
func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return m.Size()
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

func (m *HelloRequest) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (m *HelloReply) Reset()      { *m = HelloReply{} }
func (*HelloReply) ProtoMessage() {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d6d33060a754189a, []int{1}
}
func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return m.Size()
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

func init() {
	proto.RegisterType((*HelloRequest)(nil), "pl.stirling.http2.testing.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "pl.stirling.http2.testing.HelloReply")
}

func init() {
	proto.RegisterFile("src/stirling/http2/testing/proto/greet.proto", fileDescriptor_d6d33060a754189a)
}

var fileDescriptor_d6d33060a754189a = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x53, 0x3f, 0x4b, 0x7b, 0x31,
	0x14, 0x7d, 0xf7, 0xc7, 0xaf, 0xb6, 0x5e, 0x14, 0x24, 0x88, 0x56, 0x87, 0x50, 0x0a, 0x6a, 0x07,
	0x79, 0xaf, 0x3c, 0x17, 0xc1, 0x41, 0xec, 0xa2, 0x73, 0x3b, 0x29, 0x22, 0xa4, 0x35, 0x3c, 0x03,
	0xef, 0x9f, 0x49, 0x2a, 0x76, 0xf3, 0x23, 0xf8, 0x31, 0xfc, 0x22, 0x82, 0x38, 0x75, 0xec, 0x68,
	0xd3, 0xc5, 0xb1, 0x7e, 0x03, 0x69, 0x6c, 0x8a, 0x8b, 0x74, 0x69, 0xb7, 0x7b, 0xe0, 0x9c, 0x7b,
	0xce, 0x49, 0xb8, 0x78, 0xa8, 0x64, 0x27, 0x50, 0x5a, 0xc8, 0x58, 0xa4, 0x51, 0x70, 0xa7, 0x75,
	0x1e, 0x06, 0x9a, 0x2b, 0x3d, 0x41, 0xb9, 0xcc, 0x74, 0x16, 0x44, 0x92, 0x73, 0xed, 0xdb, 0x99,
	0xec, 0xe4, 0xb1, 0xef, 0xc8, 0xbe, 0x25, 0xfb, 0x53, 0x72, 0xf5, 0x18, 0xd7, 0x2e, 0x78, 0x1c,
	0x67, 0x4d, 0x7e, 0xdf, 0xe5, 0x4a, 0x13, 0x82, 0xff, 0x53, 0x96, 0xf0, 0x32, 0x54, 0xa0, 0xb6,
	0xda, 0xb4, 0x33, 0xd9, 0xc4, 0x42, 0x27, 0xeb, 0xa6, 0xba, 0xfc, 0xaf, 0x02, 0xb5, 0x42, 0xf3,
	0x07, 0x54, 0xf7, 0x11, 0xa7, 0xca, 0x3c, 0xee, 0x91, 0x32, 0x16, 0x13, 0xae, 0x14, 0x8b, 0x9c,
	0xd4, 0xc1, 0xf0, 0x1d, 0xb0, 0x78, 0x3e, 0x09, 0xc3, 0x25, 0xb9, 0xc6, 0x52, 0x8b, 0xf5, 0xac,
	0x8c, 0x1c, 0xf8, 0x7f, 0xa6, 0xf2, 0x7f, 0x47, 0xda, 0xdd, 0x9b, 0x4f, 0xcc, 0xe3, 0x5e, 0xd5,
	0x23, 0x0c, 0xd7, 0xdd, 0xf6, 0xb3, 0x88, 0x89, 0x74, 0xf1, 0x16, 0xe1, 0x2b, 0x60, 0x69, 0x5a,
	0x26, 0x24, 0x97, 0x58, 0x98, 0xf8, 0x89, 0x25, 0x54, 0xb9, 0x41, 0xb4, 0xab, 0x97, 0xd5, 0xe3,
	0x0b, 0x70, 0xa3, 0xa5, 0x25, 0x67, 0x89, 0x48, 0x23, 0xf7, 0x3b, 0x09, 0x6e, 0xbb, 0xf7, 0x6b,
	0x71, 0xf9, 0xc0, 0xe5, 0x8c, 0xb1, 0xf8, 0x04, 0x75, 0x20, 0x29, 0x6e, 0x39, 0xbb, 0x86, 0xb8,
	0x15, 0xcb, 0x74, 0xab, 0x41, 0x1d, 0x1a, 0xdd, 0xfe, 0x90, 0x7a, 0x83, 0x21, 0xf5, 0xc6, 0x43,
	0x0a, 0x4f, 0x86, 0xc2, 0x8b, 0xa1, 0xf0, 0x66, 0x28, 0xf4, 0x0d, 0x85, 0x0f, 0x43, 0xe1, 0xd3,
	0x50, 0x6f, 0x6c, 0x28, 0x3c, 0x8f, 0xa8, 0xd7, 0x1f, 0x51, 0x6f, 0x30, 0xa2, 0xde, 0xd5, 0x69,
	0x2e, 0x1e, 0x05, 0x8f, 0x59, 0x5b, 0xf9, 0x4c, 0x04, 0x33, 0x10, 0xcc, 0x3b, 0xc0, 0x13, 0x7b,
	0x80, 0x79, 0xbb, 0xbd, 0x62, 0xe1, 0xd1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xaf, 0x9a,
	0xc5, 0xb3, 0x03, 0x00, 0x00,
}

func (this *HelloRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelloRequest)
	if !ok {
		that2, ok := that.(HelloRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Count != that1.Count {
		return false
	}
	return true
}
func (this *HelloReply) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HelloReply)
	if !ok {
		that2, ok := that.(HelloReply)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	return true
}
func (this *HelloRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&greetpb.HelloRequest{")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Count: "+fmt.Sprintf("%#v", this.Count)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HelloReply) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&greetpb.HelloReply{")
	s = append(s, "Message: "+fmt.Sprintf("%#v", this.Message)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringGreet(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
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
	SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/pl.stirling.http2.testing.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/pl.stirling.http2.testing.Greeter/SayHelloAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	SayHelloAgain(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedGreeterServer) SayHelloAgain(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloAgain not implemented")
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
		FullMethod: "/pl.stirling.http2.testing.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pl.stirling.http2.testing.Greeter/SayHelloAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pl.stirling.http2.testing.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "SayHelloAgain",
			Handler:    _Greeter_SayHelloAgain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/stirling/http2/testing/proto/greet.proto",
}

// Greeter2Client is the client API for Greeter2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Greeter2Client interface {
	SayHi(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	SayHiAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeter2Client struct {
	cc *grpc.ClientConn
}

func NewGreeter2Client(cc *grpc.ClientConn) Greeter2Client {
	return &greeter2Client{cc}
}

func (c *greeter2Client) SayHi(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/pl.stirling.http2.testing.Greeter2/SayHi", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeter2Client) SayHiAgain(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/pl.stirling.http2.testing.Greeter2/SayHiAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Greeter2Server is the server API for Greeter2 service.
type Greeter2Server interface {
	SayHi(context.Context, *HelloRequest) (*HelloReply, error)
	SayHiAgain(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeter2Server can be embedded to have forward compatible implementations.
type UnimplementedGreeter2Server struct {
}

func (*UnimplementedGreeter2Server) SayHi(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHi not implemented")
}
func (*UnimplementedGreeter2Server) SayHiAgain(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHiAgain not implemented")
}

func RegisterGreeter2Server(s *grpc.Server, srv Greeter2Server) {
	s.RegisterService(&_Greeter2_serviceDesc, srv)
}

func _Greeter2_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Greeter2Server).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pl.stirling.http2.testing.Greeter2/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Greeter2Server).SayHi(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter2_SayHiAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Greeter2Server).SayHiAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pl.stirling.http2.testing.Greeter2/SayHiAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Greeter2Server).SayHiAgain(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pl.stirling.http2.testing.Greeter2",
	HandlerType: (*Greeter2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _Greeter2_SayHi_Handler,
		},
		{
			MethodName: "SayHiAgain",
			Handler:    _Greeter2_SayHiAgain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/stirling/http2/testing/proto/greet.proto",
}

// StreamingGreeterClient is the client API for StreamingGreeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamingGreeterClient interface {
	SayHelloServerStreaming(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (StreamingGreeter_SayHelloServerStreamingClient, error)
	SayHelloBidirStreaming(ctx context.Context, opts ...grpc.CallOption) (StreamingGreeter_SayHelloBidirStreamingClient, error)
}

type streamingGreeterClient struct {
	cc *grpc.ClientConn
}

func NewStreamingGreeterClient(cc *grpc.ClientConn) StreamingGreeterClient {
	return &streamingGreeterClient{cc}
}

func (c *streamingGreeterClient) SayHelloServerStreaming(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (StreamingGreeter_SayHelloServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamingGreeter_serviceDesc.Streams[0], "/pl.stirling.http2.testing.StreamingGreeter/SayHelloServerStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingGreeterSayHelloServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamingGreeter_SayHelloServerStreamingClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type streamingGreeterSayHelloServerStreamingClient struct {
	grpc.ClientStream
}

func (x *streamingGreeterSayHelloServerStreamingClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingGreeterClient) SayHelloBidirStreaming(ctx context.Context, opts ...grpc.CallOption) (StreamingGreeter_SayHelloBidirStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamingGreeter_serviceDesc.Streams[1], "/pl.stirling.http2.testing.StreamingGreeter/SayHelloBidirStreaming", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingGreeterSayHelloBidirStreamingClient{stream}
	return x, nil
}

type StreamingGreeter_SayHelloBidirStreamingClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type streamingGreeterSayHelloBidirStreamingClient struct {
	grpc.ClientStream
}

func (x *streamingGreeterSayHelloBidirStreamingClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingGreeterSayHelloBidirStreamingClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingGreeterServer is the server API for StreamingGreeter service.
type StreamingGreeterServer interface {
	SayHelloServerStreaming(*HelloRequest, StreamingGreeter_SayHelloServerStreamingServer) error
	SayHelloBidirStreaming(StreamingGreeter_SayHelloBidirStreamingServer) error
}

// UnimplementedStreamingGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedStreamingGreeterServer struct {
}

func (*UnimplementedStreamingGreeterServer) SayHelloServerStreaming(req *HelloRequest, srv StreamingGreeter_SayHelloServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStreaming not implemented")
}
func (*UnimplementedStreamingGreeterServer) SayHelloBidirStreaming(srv StreamingGreeter_SayHelloBidirStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBidirStreaming not implemented")
}

func RegisterStreamingGreeterServer(s *grpc.Server, srv StreamingGreeterServer) {
	s.RegisterService(&_StreamingGreeter_serviceDesc, srv)
}

func _StreamingGreeter_SayHelloServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingGreeterServer).SayHelloServerStreaming(m, &streamingGreeterSayHelloServerStreamingServer{stream})
}

type StreamingGreeter_SayHelloServerStreamingServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type streamingGreeterSayHelloServerStreamingServer struct {
	grpc.ServerStream
}

func (x *streamingGreeterSayHelloServerStreamingServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamingGreeter_SayHelloBidirStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingGreeterServer).SayHelloBidirStreaming(&streamingGreeterSayHelloBidirStreamingServer{stream})
}

type StreamingGreeter_SayHelloBidirStreamingServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type streamingGreeterSayHelloBidirStreamingServer struct {
	grpc.ServerStream
}

func (x *streamingGreeterSayHelloBidirStreamingServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingGreeterSayHelloBidirStreamingServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StreamingGreeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pl.stirling.http2.testing.StreamingGreeter",
	HandlerType: (*StreamingGreeterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloServerStreaming",
			Handler:       _StreamingGreeter_SayHelloServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloBidirStreaming",
			Handler:       _StreamingGreeter_SayHelloBidirStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "src/stirling/http2/testing/proto/greet.proto",
}

func (m *HelloRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintGreet(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintGreet(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *HelloReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HelloReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *HelloReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintGreet(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGreet(dAtA []byte, offset int, v uint64) int {
	offset -= sovGreet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *HelloRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovGreet(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovGreet(uint64(m.Count))
	}
	return n
}

func (m *HelloReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovGreet(uint64(l))
	}
	return n
}

func sovGreet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGreet(x uint64) (n int) {
	return sovGreet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HelloRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HelloRequest{`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Count:` + fmt.Sprintf("%v", this.Count) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HelloReply) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HelloReply{`,
		`Message:` + fmt.Sprintf("%v", this.Message) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGreet(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HelloRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGreet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HelloRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGreet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGreet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGreet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGreet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGreet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGreet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGreet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HelloReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGreet
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HelloReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HelloReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGreet
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGreet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGreet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGreet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGreet
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGreet
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGreet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGreet
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGreet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGreet
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGreet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGreet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGreet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGreet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGreet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGreet = fmt.Errorf("proto: unexpected end of group")
)
