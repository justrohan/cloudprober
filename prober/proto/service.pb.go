// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/prober/proto/service.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import proto1 "github.com/google/cloudprober/probes/proto"

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

type AddProbeRequest struct {
	Probe                *proto1.ProbeDef `protobuf:"bytes,1,opt,name=probe" json:"probe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AddProbeRequest) Reset()         { *m = AddProbeRequest{} }
func (m *AddProbeRequest) String() string { return proto.CompactTextString(m) }
func (*AddProbeRequest) ProtoMessage()    {}
func (*AddProbeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ba7f7c5047aeec33, []int{0}
}
func (m *AddProbeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProbeRequest.Unmarshal(m, b)
}
func (m *AddProbeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProbeRequest.Marshal(b, m, deterministic)
}
func (dst *AddProbeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProbeRequest.Merge(dst, src)
}
func (m *AddProbeRequest) XXX_Size() int {
	return xxx_messageInfo_AddProbeRequest.Size(m)
}
func (m *AddProbeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProbeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddProbeRequest proto.InternalMessageInfo

func (m *AddProbeRequest) GetProbe() *proto1.ProbeDef {
	if m != nil {
		return m.Probe
	}
	return nil
}

type AddProbeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddProbeResponse) Reset()         { *m = AddProbeResponse{} }
func (m *AddProbeResponse) String() string { return proto.CompactTextString(m) }
func (*AddProbeResponse) ProtoMessage()    {}
func (*AddProbeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ba7f7c5047aeec33, []int{1}
}
func (m *AddProbeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddProbeResponse.Unmarshal(m, b)
}
func (m *AddProbeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddProbeResponse.Marshal(b, m, deterministic)
}
func (dst *AddProbeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddProbeResponse.Merge(dst, src)
}
func (m *AddProbeResponse) XXX_Size() int {
	return xxx_messageInfo_AddProbeResponse.Size(m)
}
func (m *AddProbeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddProbeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddProbeResponse proto.InternalMessageInfo

type RemoveProbeRequest struct {
	ProbeName            *string  `protobuf:"bytes,1,opt,name=probe_name,json=probeName" json:"probe_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveProbeRequest) Reset()         { *m = RemoveProbeRequest{} }
func (m *RemoveProbeRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveProbeRequest) ProtoMessage()    {}
func (*RemoveProbeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ba7f7c5047aeec33, []int{2}
}
func (m *RemoveProbeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveProbeRequest.Unmarshal(m, b)
}
func (m *RemoveProbeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveProbeRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveProbeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveProbeRequest.Merge(dst, src)
}
func (m *RemoveProbeRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveProbeRequest.Size(m)
}
func (m *RemoveProbeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveProbeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveProbeRequest proto.InternalMessageInfo

func (m *RemoveProbeRequest) GetProbeName() string {
	if m != nil && m.ProbeName != nil {
		return *m.ProbeName
	}
	return ""
}

type RemoveProbeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveProbeResponse) Reset()         { *m = RemoveProbeResponse{} }
func (m *RemoveProbeResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveProbeResponse) ProtoMessage()    {}
func (*RemoveProbeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ba7f7c5047aeec33, []int{3}
}
func (m *RemoveProbeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveProbeResponse.Unmarshal(m, b)
}
func (m *RemoveProbeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveProbeResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveProbeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveProbeResponse.Merge(dst, src)
}
func (m *RemoveProbeResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveProbeResponse.Size(m)
}
func (m *RemoveProbeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveProbeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveProbeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AddProbeRequest)(nil), "cloudprober.AddProbeRequest")
	proto.RegisterType((*AddProbeResponse)(nil), "cloudprober.AddProbeResponse")
	proto.RegisterType((*RemoveProbeRequest)(nil), "cloudprober.RemoveProbeRequest")
	proto.RegisterType((*RemoveProbeResponse)(nil), "cloudprober.RemoveProbeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CloudproberClient is the client API for Cloudprober service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CloudproberClient interface {
	// AddProbe adds a probe to cloudprober. Error is returned if probe is already
	// defined or there is an error during initialization of the probe.
	AddProbe(ctx context.Context, in *AddProbeRequest, opts ...grpc.CallOption) (*AddProbeResponse, error)
	// RemoveProbe stops the probe and removes it from the in-memory database.
	RemoveProbe(ctx context.Context, in *RemoveProbeRequest, opts ...grpc.CallOption) (*RemoveProbeResponse, error)
}

type cloudproberClient struct {
	cc *grpc.ClientConn
}

func NewCloudproberClient(cc *grpc.ClientConn) CloudproberClient {
	return &cloudproberClient{cc}
}

func (c *cloudproberClient) AddProbe(ctx context.Context, in *AddProbeRequest, opts ...grpc.CallOption) (*AddProbeResponse, error) {
	out := new(AddProbeResponse)
	err := c.cc.Invoke(ctx, "/cloudprober.Cloudprober/AddProbe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudproberClient) RemoveProbe(ctx context.Context, in *RemoveProbeRequest, opts ...grpc.CallOption) (*RemoveProbeResponse, error) {
	out := new(RemoveProbeResponse)
	err := c.cc.Invoke(ctx, "/cloudprober.Cloudprober/RemoveProbe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudproberServer is the server API for Cloudprober service.
type CloudproberServer interface {
	// AddProbe adds a probe to cloudprober. Error is returned if probe is already
	// defined or there is an error during initialization of the probe.
	AddProbe(context.Context, *AddProbeRequest) (*AddProbeResponse, error)
	// RemoveProbe stops the probe and removes it from the in-memory database.
	RemoveProbe(context.Context, *RemoveProbeRequest) (*RemoveProbeResponse, error)
}

func RegisterCloudproberServer(s *grpc.Server, srv CloudproberServer) {
	s.RegisterService(&_Cloudprober_serviceDesc, srv)
}

func _Cloudprober_AddProbe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProbeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudproberServer).AddProbe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprober.Cloudprober/AddProbe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudproberServer).AddProbe(ctx, req.(*AddProbeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cloudprober_RemoveProbe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProbeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudproberServer).RemoveProbe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cloudprober.Cloudprober/RemoveProbe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudproberServer).RemoveProbe(ctx, req.(*RemoveProbeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cloudprober_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cloudprober.Cloudprober",
	HandlerType: (*CloudproberServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProbe",
			Handler:    _Cloudprober_AddProbe_Handler,
		},
		{
			MethodName: "RemoveProbe",
			Handler:    _Cloudprober_RemoveProbe_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/google/cloudprober/prober/proto/service.proto",
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/prober/proto/service.proto", fileDescriptor_service_ba7f7c5047aeec33)
}

var fileDescriptor_service_ba7f7c5047aeec33 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0xce,
	0xc9, 0x2f, 0x4d, 0x29, 0x28, 0xca, 0x4f, 0x4a, 0x2d, 0xd2, 0x47, 0x50, 0x25, 0xf9, 0xfa, 0xc5,
	0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x7a, 0x60, 0x9e, 0x10, 0x37, 0x92, 0x3a, 0x29, 0x73, 0x22,
	0x8c, 0x29, 0x86, 0x1a, 0x93, 0x9c, 0x9f, 0x97, 0x96, 0x99, 0x0e, 0x31, 0x45, 0xc9, 0x95, 0x8b,
	0xdf, 0x31, 0x25, 0x25, 0x00, 0x24, 0x1f, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x64, 0xc4,
	0xc5, 0x0a, 0x56, 0x2f, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa3, 0x87, 0x64, 0x92, 0x1e,
	0xc4, 0x24, 0x3d, 0xb0, 0x06, 0x97, 0xd4, 0xb4, 0x20, 0x88, 0x52, 0x25, 0x21, 0x2e, 0x01, 0x84,
	0x31, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x4a, 0xc6, 0x5c, 0x42, 0x41, 0xa9, 0xb9, 0xf9, 0x65,
	0xa9, 0x28, 0xa6, 0xcb, 0x72, 0x71, 0x81, 0xb5, 0xc4, 0xe7, 0x25, 0xe6, 0x42, 0xac, 0xe0, 0x0c,
	0xe2, 0x04, 0x8b, 0xf8, 0x25, 0xe6, 0xa6, 0x2a, 0x89, 0x72, 0x09, 0xa3, 0x68, 0x82, 0x98, 0x65,
	0xb4, 0x86, 0x91, 0x8b, 0xdb, 0x19, 0xe1, 0x0c, 0x21, 0x4f, 0x2e, 0x0e, 0x98, 0x7d, 0x42, 0xa8,
	0x0e, 0x44, 0xf3, 0x8d, 0x94, 0x2c, 0x0e, 0x59, 0xa8, 0x23, 0x19, 0x84, 0x82, 0xb8, 0xb8, 0x91,
	0x6c, 0x14, 0x92, 0x47, 0x51, 0x8f, 0xe9, 0x01, 0x29, 0x05, 0xdc, 0x0a, 0x60, 0x66, 0x02, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x38, 0x24, 0xef, 0x67, 0xd6, 0x01, 0x00, 0x00,
}
