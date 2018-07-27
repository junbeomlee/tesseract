// Code generated by protoc-gen-go. DO NOT EDIT.
// source: default_service.proto

package legacy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_default_service_09c5a8ed21e91654, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Request struct {
	Tx                   []byte   `protobuf:"bytes,1,opt,name=Tx,proto3" json:"Tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_default_service_09c5a8ed21e91654, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type Response struct {
	Result               string   `protobuf:"bytes,1,opt,name=Result" json:"Result,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=Method" json:"Method,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=Error" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_default_service_09c5a8ed21e91654, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Response) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *Response) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Response) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "pb.Empty")
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Response)(nil), "pb.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DefaultService service

type DefaultServiceClient interface {
	RunICode(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type defaultServiceClient struct {
	cc *grpc.ClientConn
}

func NewDefaultServiceClient(cc *grpc.ClientConn) DefaultServiceClient {
	return &defaultServiceClient{cc}
}

func (c *defaultServiceClient) RunICode(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/pb.DefaultService/RunICode", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *defaultServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/pb.DefaultService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DefaultService service

type DefaultServiceServer interface {
	RunICode(context.Context, *Request) (*Response, error)
	Ping(context.Context, *Empty) (*Empty, error)
}

func RegisterDefaultServiceServer(s *grpc.Server, srv DefaultServiceServer) {
	s.RegisterService(&_DefaultService_serviceDesc, srv)
}

func _DefaultService_RunICode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).RunICode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DefaultService/RunICode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).RunICode(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DefaultService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DefaultServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.DefaultService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DefaultServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _DefaultService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.DefaultService",
	HandlerType: (*DefaultServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunICode",
			Handler:    _DefaultService_RunICode_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _DefaultService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "default_service.proto",
}

func init() {
	proto.RegisterFile("default_service.proto", fileDescriptor_default_service_09c5a8ed21e91654)
}

var fileDescriptor_default_service_09c5a8ed21e91654 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xbb, 0x4e, 0xc4, 0x30,
	0x10, 0x45, 0x77, 0x4d, 0xf6, 0x35, 0xac, 0xb6, 0x18, 0x01, 0x32, 0x29, 0x10, 0x72, 0x03, 0x55,
	0x0a, 0xf8, 0x04, 0x92, 0x82, 0x02, 0x09, 0x99, 0x34, 0x54, 0x28, 0xc1, 0x03, 0x44, 0x0a, 0xb1,
	0xf1, 0x03, 0x85, 0xbf, 0x47, 0xb1, 0x23, 0x6d, 0x77, 0xcf, 0xb5, 0x35, 0x3a, 0x17, 0xce, 0x15,
	0x7d, 0x34, 0xa1, 0xf7, 0x6f, 0x8e, 0xec, 0x6f, 0xf7, 0x4e, 0x85, 0xb1, 0xda, 0x6b, 0x64, 0xa6,
	0x15, 0x1b, 0x58, 0x55, 0xdf, 0xc6, 0xff, 0x89, 0x4b, 0xd8, 0x48, 0xfa, 0x09, 0xe4, 0x3c, 0x1e,
	0x80, 0xd5, 0x23, 0x5f, 0x5e, 0x2f, 0x6f, 0xf7, 0x92, 0xd5, 0xa3, 0x50, 0xb0, 0x95, 0xe4, 0x8c,
	0x1e, 0x1c, 0xe1, 0x05, 0xac, 0x25, 0xb9, 0xd0, 0xfb, 0xf8, 0xbe, 0x93, 0x33, 0x4d, 0xfd, 0x13,
	0xf9, 0x2f, 0xad, 0x38, 0x4b, 0x7d, 0x22, 0x44, 0xc8, 0xca, 0xc6, 0x37, 0xfc, 0x24, 0x5e, 0x8b,
	0x19, 0xcf, 0x60, 0x55, 0x59, 0xab, 0x2d, 0xcf, 0xe2, 0xd7, 0x04, 0x77, 0xaf, 0x70, 0x28, 0x93,
	0xe6, 0x4b, 0xb2, 0xc4, 0x1b, 0xd8, 0xca, 0x30, 0x3c, 0x3e, 0x68, 0x45, 0x78, 0x5a, 0x98, 0xb6,
	0x98, 0x05, 0xf3, 0x7d, 0x82, 0xa4, 0x24, 0x16, 0x78, 0x05, 0xd9, 0x73, 0x37, 0x7c, 0xe2, 0x6e,
	0xea, 0xe3, 0x9c, 0xfc, 0x18, 0xc5, 0xa2, 0x5d, 0xc7, 0xbd, 0xf7, 0xff, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x5f, 0x06, 0x2b, 0x72, 0x08, 0x01, 0x00, 0x00,
}