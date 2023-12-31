// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: jomtx/jomtx/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac2036fd62cc6f0, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac2036fd62cc6f0, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetTxnRequest struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryGetTxnRequest) Reset()         { *m = QueryGetTxnRequest{} }
func (m *QueryGetTxnRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetTxnRequest) ProtoMessage()    {}
func (*QueryGetTxnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac2036fd62cc6f0, []int{2}
}
func (m *QueryGetTxnRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTxnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTxnRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTxnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTxnRequest.Merge(m, src)
}
func (m *QueryGetTxnRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTxnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTxnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTxnRequest proto.InternalMessageInfo

func (m *QueryGetTxnRequest) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type QueryGetTxnResponse struct {
	Txn Txn `protobuf:"bytes,1,opt,name=Txn,proto3" json:"Txn"`
}

func (m *QueryGetTxnResponse) Reset()         { *m = QueryGetTxnResponse{} }
func (m *QueryGetTxnResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetTxnResponse) ProtoMessage()    {}
func (*QueryGetTxnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac2036fd62cc6f0, []int{3}
}
func (m *QueryGetTxnResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTxnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTxnResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTxnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTxnResponse.Merge(m, src)
}
func (m *QueryGetTxnResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTxnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTxnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTxnResponse proto.InternalMessageInfo

func (m *QueryGetTxnResponse) GetTxn() Txn {
	if m != nil {
		return m.Txn
	}
	return Txn{}
}

//	message QueryAllTxnResponse {
//	  repeated Txn                                    Txn        = 1 [(gogoproto.nullable) = false];
//	           cosmos.base.query.v1beta1.PageResponse pagination = 2;
//	}
type QueryGetSystemInfoRequest struct {
}

func (m *QueryGetSystemInfoRequest) Reset()         { *m = QueryGetSystemInfoRequest{} }
func (m *QueryGetSystemInfoRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetSystemInfoRequest) ProtoMessage()    {}
func (*QueryGetSystemInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac2036fd62cc6f0, []int{4}
}
func (m *QueryGetSystemInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSystemInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSystemInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSystemInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSystemInfoRequest.Merge(m, src)
}
func (m *QueryGetSystemInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSystemInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSystemInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSystemInfoRequest proto.InternalMessageInfo

type QueryGetSystemInfoResponse struct {
	SystemInfo SystemInfo `protobuf:"bytes,1,opt,name=SystemInfo,proto3" json:"SystemInfo"`
}

func (m *QueryGetSystemInfoResponse) Reset()         { *m = QueryGetSystemInfoResponse{} }
func (m *QueryGetSystemInfoResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetSystemInfoResponse) ProtoMessage()    {}
func (*QueryGetSystemInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ac2036fd62cc6f0, []int{5}
}
func (m *QueryGetSystemInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSystemInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSystemInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSystemInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSystemInfoResponse.Merge(m, src)
}
func (m *QueryGetSystemInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSystemInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSystemInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSystemInfoResponse proto.InternalMessageInfo

func (m *QueryGetSystemInfoResponse) GetSystemInfo() SystemInfo {
	if m != nil {
		return m.SystemInfo
	}
	return SystemInfo{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "jomtx.jomtx.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "jomtx.jomtx.QueryParamsResponse")
	proto.RegisterType((*QueryGetTxnRequest)(nil), "jomtx.jomtx.QueryGetTxnRequest")
	proto.RegisterType((*QueryGetTxnResponse)(nil), "jomtx.jomtx.QueryGetTxnResponse")
	proto.RegisterType((*QueryGetSystemInfoRequest)(nil), "jomtx.jomtx.QueryGetSystemInfoRequest")
	proto.RegisterType((*QueryGetSystemInfoResponse)(nil), "jomtx.jomtx.QueryGetSystemInfoResponse")
}

func init() { proto.RegisterFile("jomtx/jomtx/query.proto", fileDescriptor_7ac2036fd62cc6f0) }

var fileDescriptor_7ac2036fd62cc6f0 = []byte{
	// 457 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x41, 0xef, 0xd2, 0x30,
	0x18, 0xc6, 0xb7, 0xf9, 0x97, 0x43, 0xff, 0x89, 0x31, 0x05, 0x03, 0x0e, 0x19, 0xa4, 0x2a, 0x12,
	0x0f, 0x6b, 0x80, 0x18, 0x4f, 0x46, 0xc3, 0x45, 0xbd, 0x29, 0x72, 0xd2, 0x83, 0xe9, 0xa0, 0xcc,
	0x1a, 0xd7, 0x0e, 0x5a, 0x74, 0xc4, 0x78, 0xf1, 0xe2, 0xd5, 0xc4, 0x2f, 0xc5, 0x91, 0xc4, 0x8b,
	0x27, 0x63, 0xc0, 0x8f, 0xe0, 0x07, 0x30, 0x74, 0x9d, 0x6c, 0x32, 0xe3, 0xa5, 0x21, 0xef, 0xfb,
	0xf4, 0xf9, 0xbd, 0xef, 0x53, 0x06, 0xea, 0xaf, 0x45, 0xa4, 0x12, 0x9c, 0x9e, 0x8b, 0x15, 0x5d,
	0xae, 0xfd, 0x78, 0x29, 0x94, 0x80, 0xe7, 0xba, 0xe4, 0xeb, 0xd3, 0xad, 0x85, 0x22, 0x14, 0xba,
	0x8e, 0x0f, 0xbf, 0x52, 0x89, 0x7b, 0x2d, 0x14, 0x22, 0x7c, 0x43, 0x31, 0x89, 0x19, 0x26, 0x9c,
	0x0b, 0x45, 0x14, 0x13, 0x5c, 0x9a, 0xee, 0xed, 0xa9, 0x90, 0x91, 0x90, 0x38, 0x20, 0x92, 0xa6,
	0xce, 0xf8, 0x6d, 0x3f, 0xa0, 0x8a, 0xf4, 0x71, 0x4c, 0x42, 0xc6, 0xb5, 0xd8, 0x68, 0x1b, 0xf9,
	0x29, 0x62, 0xb2, 0x24, 0x51, 0xe6, 0x72, 0x25, 0xdf, 0x51, 0x49, 0x76, 0xa1, 0x95, 0x2f, 0xcb,
	0xb5, 0x54, 0x34, 0x7a, 0xc9, 0xf8, 0xdc, 0x4c, 0x86, 0x6a, 0x00, 0x3e, 0x3d, 0x10, 0x9f, 0x68,
	0xab, 0x31, 0x5d, 0xac, 0xa8, 0x54, 0xe8, 0x11, 0xa8, 0x16, 0xaa, 0x32, 0x16, 0x5c, 0x52, 0xd8,
	0x07, 0x95, 0x14, 0xd9, 0xb0, 0x3b, 0x76, 0xef, 0x7c, 0x50, 0xf5, 0x73, 0xab, 0xfb, 0xa9, 0x78,
	0x74, 0xb6, 0xf9, 0xde, 0xb6, 0xc6, 0x46, 0x88, 0x6e, 0x18, 0xff, 0x87, 0x54, 0x4d, 0x12, 0x6e,
	0xfc, 0xe1, 0x25, 0xe0, 0xb0, 0x99, 0x36, 0x39, 0x1b, 0x3b, 0x6c, 0x86, 0xee, 0x1b, 0x5e, 0xa6,
	0x32, 0xbc, 0x1e, 0xb8, 0x30, 0x49, 0xb8, 0x81, 0x5d, 0x2e, 0xc0, 0x26, 0x09, 0x37, 0xa4, 0x83,
	0x04, 0x35, 0xc1, 0xd5, 0xcc, 0xe0, 0x99, 0xde, 0xf1, 0x31, 0x9f, 0x8b, 0x6c, 0x9b, 0x17, 0xc0,
	0x2d, 0x6b, 0x1a, 0xc8, 0x3d, 0x00, 0x8e, 0x55, 0xc3, 0xaa, 0x17, 0x58, 0xc7, 0xb6, 0x41, 0xe6,
	0x2e, 0x0c, 0x7e, 0x39, 0xe0, 0xa2, 0x76, 0x87, 0x31, 0xa8, 0xa4, 0x11, 0xc0, 0x76, 0xe1, 0xfa,
	0x69, 0xbe, 0x6e, 0xe7, 0xdf, 0x82, 0x74, 0x2a, 0x74, 0xfd, 0xe3, 0xd7, 0x9f, 0x5f, 0x9c, 0x16,
	0x6c, 0xe2, 0x77, 0x91, 0x1c, 0xdc, 0x19, 0xde, 0xc5, 0xa7, 0x0f, 0x0f, 0x85, 0xce, 0xa7, 0x0c,
	0x57, 0x88, 0xbb, 0x0c, 0x57, 0x4c, 0x1a, 0xdd, 0xd4, 0xb8, 0x36, 0x6c, 0x95, 0xe2, 0x54, 0xc2,
	0xf1, 0x7b, 0x36, 0xfb, 0x00, 0x3f, 0xd9, 0xf9, 0xb0, 0x60, 0xb7, 0xd4, 0xf7, 0xe4, 0x01, 0xdc,
	0x5b, 0xff, 0xd5, 0x99, 0x31, 0x7a, 0x7a, 0x0c, 0x04, 0x3b, 0xa5, 0x63, 0xe4, 0xfe, 0xbd, 0xa3,
	0x07, 0x9b, 0x9d, 0x67, 0x6f, 0x77, 0x9e, 0xfd, 0x63, 0xe7, 0xd9, 0x9f, 0xf7, 0x9e, 0xb5, 0xdd,
	0x7b, 0xd6, 0xb7, 0xbd, 0x67, 0x3d, 0xef, 0x86, 0x4c, 0xbd, 0x5a, 0x05, 0xfe, 0x54, 0x44, 0x7f,
	0xb9, 0xfc, 0x59, 0x67, 0x1d, 0x53, 0x19, 0x54, 0xf4, 0x07, 0x30, 0xfc, 0x1d, 0x00, 0x00, 0xff,
	0xff, 0x1a, 0xe6, 0xd4, 0x07, 0xd8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Txn items.
	Txn(ctx context.Context, in *QueryGetTxnRequest, opts ...grpc.CallOption) (*QueryGetTxnResponse, error)
	// Queries a SystemInfo by index.
	SystemInfo(ctx context.Context, in *QueryGetSystemInfoRequest, opts ...grpc.CallOption) (*QueryGetSystemInfoResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/jomtx.jomtx.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Txn(ctx context.Context, in *QueryGetTxnRequest, opts ...grpc.CallOption) (*QueryGetTxnResponse, error) {
	out := new(QueryGetTxnResponse)
	err := c.cc.Invoke(ctx, "/jomtx.jomtx.Query/Txn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SystemInfo(ctx context.Context, in *QueryGetSystemInfoRequest, opts ...grpc.CallOption) (*QueryGetSystemInfoResponse, error) {
	out := new(QueryGetSystemInfoResponse)
	err := c.cc.Invoke(ctx, "/jomtx.jomtx.Query/SystemInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Txn items.
	Txn(context.Context, *QueryGetTxnRequest) (*QueryGetTxnResponse, error)
	// Queries a SystemInfo by index.
	SystemInfo(context.Context, *QueryGetSystemInfoRequest) (*QueryGetSystemInfoResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Txn(ctx context.Context, req *QueryGetTxnRequest) (*QueryGetTxnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Txn not implemented")
}
func (*UnimplementedQueryServer) SystemInfo(ctx context.Context, req *QueryGetSystemInfoRequest) (*QueryGetSystemInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemInfo not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jomtx.jomtx.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Txn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Txn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jomtx.jomtx.Query/Txn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Txn(ctx, req.(*QueryGetTxnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSystemInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jomtx.jomtx.Query/SystemInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SystemInfo(ctx, req.(*QueryGetSystemInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "jomtx.jomtx.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Txn",
			Handler:    _Query_Txn_Handler,
		},
		{
			MethodName: "SystemInfo",
			Handler:    _Query_SystemInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jomtx/jomtx/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetTxnRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTxnRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTxnRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetTxnResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTxnResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTxnResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Txn.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetSystemInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSystemInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSystemInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetSystemInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSystemInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSystemInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SystemInfo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetTxnRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovQuery(uint64(m.Id))
	}
	return n
}

func (m *QueryGetTxnResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Txn.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetSystemInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetSystemInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SystemInfo.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetTxnRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetTxnRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTxnRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetTxnResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetTxnResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTxnResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Txn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetSystemInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetSystemInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSystemInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetSystemInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetSystemInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSystemInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SystemInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SystemInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
