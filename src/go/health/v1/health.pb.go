// Code generated by protoc-gen-go. DO NOT EDIT.
// source: health/v1/health.proto

package health

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_health_d23b8bcecb448c77, []int{1, 0}
}

type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_d23b8bcecb448c77, []int{0}
}
func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (dst *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(dst, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=sagittarius.health.v1.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	Version              string                            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_health_d23b8bcecb448c77, []int{1}
}
func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (dst *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(dst, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func (m *HealthCheckResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthCheckRequest)(nil), "sagittarius.health.v1.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "sagittarius.health.v1.HealthCheckResponse")
	proto.RegisterEnum("sagittarius.health.v1.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/sagittarius.health.v1.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sagittarius.health.v1.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sagittarius.health.v1.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health/v1/health.proto",
}

func init() { proto.RegisterFile("health/v1/health.proto", fileDescriptor_health_d23b8bcecb448c77) }

var fileDescriptor_health_d23b8bcecb448c77 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0x48, 0x4d, 0xcc,
	0x29, 0xc9, 0xd0, 0x2f, 0x33, 0xd4, 0x87, 0xb0, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44,
	0x8b, 0x13, 0xd3, 0x33, 0x4b, 0x4a, 0x12, 0x8b, 0x32, 0x4b, 0x8b, 0xf5, 0xa0, 0x32, 0x65, 0x86,
	0x4a, 0x7a, 0x5c, 0x42, 0x1e, 0x60, 0x8e, 0x73, 0x46, 0x6a, 0x72, 0x76, 0x50, 0x6a, 0x61, 0x69,
	0x6a, 0x71, 0x89, 0x90, 0x04, 0x17, 0x7b, 0x71, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0xab, 0xb4, 0x97, 0x91, 0x4b, 0x18, 0x45, 0x43, 0x71, 0x41,
	0x7e, 0x5e, 0x71, 0xaa, 0x50, 0x00, 0x17, 0x5b, 0x71, 0x49, 0x62, 0x49, 0x69, 0x31, 0x58, 0x03,
	0x9f, 0x91, 0x85, 0x1e, 0x56, 0xfb, 0xf4, 0xb0, 0xe8, 0xd5, 0x0b, 0x06, 0x99, 0x9d, 0x97, 0x1e,
	0x0c, 0xd6, 0x1f, 0x04, 0x35, 0x07, 0xe4, 0x86, 0xb2, 0xd4, 0xa2, 0xe2, 0xcc, 0xfc, 0x3c, 0x09,
	0x26, 0x88, 0x1b, 0xa0, 0x5c, 0x25, 0x2b, 0x2e, 0x5e, 0x14, 0x2d, 0x42, 0xdc, 0x5c, 0xec, 0xa1,
	0x7e, 0xde, 0x7e, 0xfe, 0xe1, 0x7e, 0x02, 0x0c, 0x20, 0x4e, 0xb0, 0x6b, 0x50, 0x98, 0xa7, 0x9f,
	0xbb, 0x00, 0xa3, 0x10, 0x3f, 0x17, 0xb7, 0x9f, 0x7f, 0x48, 0x3c, 0x4c, 0x80, 0xc9, 0x28, 0x83,
	0x8b, 0x0d, 0xe2, 0x04, 0xa1, 0x38, 0x2e, 0x56, 0xb0, 0x33, 0x84, 0x34, 0x89, 0x71, 0x2a, 0x38,
	0x5c, 0xa4, 0xb4, 0x88, 0xf7, 0x95, 0x93, 0x37, 0x97, 0x44, 0x62, 0x26, 0x76, 0xf5, 0x4e, 0xdc,
	0x10, 0x0d, 0x01, 0xa0, 0x98, 0x09, 0x60, 0x8c, 0x62, 0x83, 0xc8, 0xfc, 0x60, 0x64, 0x5c, 0xc5,
	0x24, 0x12, 0x0c, 0xd7, 0x01, 0x35, 0x56, 0x2f, 0xcc, 0x30, 0x89, 0x0d, 0x1c, 0x89, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xfa, 0x11, 0x05, 0xde, 0x01, 0x00, 0x00,
}
