// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/services/user_interest_service.proto

package services // import "google.golang.org/genproto/googleapis/ads/googleads/v1/services"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import resources "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Request message for [UserInterestService.GetUserInterest][google.ads.googleads.v1.services.UserInterestService.GetUserInterest].
type GetUserInterestRequest struct {
	// Resource name of the UserInterest to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserInterestRequest) Reset()         { *m = GetUserInterestRequest{} }
func (m *GetUserInterestRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserInterestRequest) ProtoMessage()    {}
func (*GetUserInterestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_interest_service_f0f53b0efc9abbb8, []int{0}
}
func (m *GetUserInterestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserInterestRequest.Unmarshal(m, b)
}
func (m *GetUserInterestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserInterestRequest.Marshal(b, m, deterministic)
}
func (dst *GetUserInterestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserInterestRequest.Merge(dst, src)
}
func (m *GetUserInterestRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserInterestRequest.Size(m)
}
func (m *GetUserInterestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserInterestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserInterestRequest proto.InternalMessageInfo

func (m *GetUserInterestRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserInterestRequest)(nil), "google.ads.googleads.v1.services.GetUserInterestRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserInterestServiceClient is the client API for UserInterestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserInterestServiceClient interface {
	// Returns the requested user interest in full detail
	GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error)
}

type userInterestServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserInterestServiceClient(cc *grpc.ClientConn) UserInterestServiceClient {
	return &userInterestServiceClient{cc}
}

func (c *userInterestServiceClient) GetUserInterest(ctx context.Context, in *GetUserInterestRequest, opts ...grpc.CallOption) (*resources.UserInterest, error) {
	out := new(resources.UserInterest)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v1.services.UserInterestService/GetUserInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserInterestServiceServer is the server API for UserInterestService service.
type UserInterestServiceServer interface {
	// Returns the requested user interest in full detail
	GetUserInterest(context.Context, *GetUserInterestRequest) (*resources.UserInterest, error)
}

func RegisterUserInterestServiceServer(s *grpc.Server, srv UserInterestServiceServer) {
	s.RegisterService(&_UserInterestService_serviceDesc, srv)
}

func _UserInterestService_GetUserInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v1.services.UserInterestService/GetUserInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInterestServiceServer).GetUserInterest(ctx, req.(*GetUserInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInterestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v1.services.UserInterestService",
	HandlerType: (*UserInterestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserInterest",
			Handler:    _UserInterestService_GetUserInterest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v1/services/user_interest_service.proto",
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/services/user_interest_service.proto", fileDescriptor_user_interest_service_f0f53b0efc9abbb8)
}

var fileDescriptor_user_interest_service_f0f53b0efc9abbb8 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x49, 0x2e, 0x5c, 0xb8, 0xe1, 0x5e, 0x2e, 0x44, 0x90, 0x52, 0x5c, 0x94, 0xda, 0x85,
	0x74, 0x31, 0x43, 0x14, 0x51, 0x46, 0xbb, 0x48, 0x37, 0xd5, 0x8d, 0x94, 0x8a, 0x5d, 0x48, 0xa0,
	0x8c, 0xcd, 0x21, 0x04, 0x9a, 0x99, 0x3a, 0x67, 0xd2, 0x8d, 0xb8, 0xf1, 0x15, 0x7c, 0x03, 0x97,
	0xee, 0x7d, 0x09, 0xc1, 0x95, 0xaf, 0xe0, 0xca, 0x97, 0x50, 0xd2, 0xc9, 0x84, 0xaa, 0x2d, 0xdd,
	0x7d, 0x9c, 0x7c, 0xbf, 0xef, 0xfc, 0xc9, 0x78, 0xc7, 0x89, 0x94, 0xc9, 0x04, 0x28, 0x8f, 0x91,
	0x1a, 0x59, 0xa8, 0x59, 0x40, 0x11, 0xd4, 0x2c, 0x1d, 0x03, 0xd2, 0x1c, 0x41, 0x8d, 0x52, 0xa1,
	0x41, 0x01, 0xea, 0x51, 0x59, 0x26, 0x53, 0x25, 0xb5, 0xf4, 0x1b, 0x06, 0x21, 0x3c, 0x46, 0x52,
	0xd1, 0x64, 0x16, 0x10, 0x4b, 0xd7, 0xf7, 0x57, 0xe5, 0x2b, 0x40, 0x99, 0xab, 0x1f, 0x0d, 0x4c,
	0x70, 0x7d, 0xcb, 0x62, 0xd3, 0x94, 0x72, 0x21, 0xa4, 0xe6, 0x3a, 0x95, 0x02, 0xcd, 0xd7, 0x66,
	0xc7, 0xdb, 0xec, 0x81, 0xbe, 0x40, 0x50, 0xa7, 0x25, 0x36, 0x80, 0xeb, 0x1c, 0x50, 0xfb, 0xdb,
	0xde, 0x3f, 0x1b, 0x3c, 0x12, 0x3c, 0x83, 0x9a, 0xd3, 0x70, 0x76, 0xfe, 0x0c, 0xfe, 0xda, 0xe2,
	0x19, 0xcf, 0x60, 0xf7, 0xc5, 0xf1, 0x36, 0x16, 0xe1, 0x73, 0x33, 0xac, 0xff, 0xe4, 0x78, 0xff,
	0xbf, 0xe5, 0xfa, 0x87, 0x64, 0xdd, 0x8a, 0x64, 0xf9, 0x28, 0x75, 0xba, 0x92, 0xac, 0x56, 0x27,
	0x8b, 0x5c, 0xf3, 0xe0, 0xee, 0xf5, 0xed, 0xde, 0x0d, 0x7c, 0x5a, 0x9c, 0xe7, 0xe6, 0xcb, 0x1a,
	0x9d, 0x71, 0x8e, 0x5a, 0x66, 0xa0, 0x90, 0xb6, 0xe7, 0xf7, 0xb2, 0x10, 0xd2, 0xf6, 0x6d, 0xf7,
	0xc3, 0xf1, 0x5a, 0x63, 0x99, 0xad, 0x9d, 0xb4, 0x5b, 0x5b, 0xb2, 0x75, 0xbf, 0xb8, 0x68, 0xdf,
	0xb9, 0x3c, 0x29, 0xe9, 0x44, 0x4e, 0xb8, 0x48, 0x88, 0x54, 0x09, 0x4d, 0x40, 0xcc, 0xef, 0x6d,
	0x7f, 0xdc, 0x34, 0xc5, 0xd5, 0xef, 0xe4, 0xc8, 0x8a, 0x07, 0xf7, 0x57, 0x2f, 0x0c, 0x1f, 0xdd,
	0x46, 0xcf, 0x04, 0x86, 0x31, 0x12, 0x23, 0x0b, 0x35, 0x0c, 0x48, 0xd9, 0x18, 0x9f, 0xad, 0x25,
	0x0a, 0x63, 0x8c, 0x2a, 0x4b, 0x34, 0x0c, 0x22, 0x6b, 0x79, 0x77, 0x5b, 0xa6, 0xce, 0x58, 0x18,
	0x23, 0x63, 0x95, 0x89, 0xb1, 0x61, 0xc0, 0x98, 0xb5, 0x5d, 0xfd, 0x9e, 0xcf, 0xb9, 0xf7, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x92, 0xe6, 0xf7, 0x12, 0xce, 0x02, 0x00, 0x00,
}
