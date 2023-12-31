// protoc -I third_party -I fuxi/api/bff/v1 --go_out=. --go-grpc_out=. --go-http_out=. fuxi/api/bff/v1/fuxi_bff.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.18.0
// source: fuxi_bff.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	FuxiBFF_SetUrl_FullMethodName        = "/v1.FuxiBFF/SetUrl"
	FuxiBFF_GetUrl_FullMethodName        = "/v1.FuxiBFF/GetUrl"
	FuxiBFF_AllocationUrl_FullMethodName = "/v1.FuxiBFF/AllocationUrl"
)

// FuxiBFFClient is the client API for FuxiBFF service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FuxiBFFClient interface {
	SetUrl(ctx context.Context, in *SetUrlRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUrl(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error)
	AllocationUrl(ctx context.Context, in *AllocationRequest, opts ...grpc.CallOption) (*AllocationResponse, error)
}

type fuxiBFFClient struct {
	cc grpc.ClientConnInterface
}

func NewFuxiBFFClient(cc grpc.ClientConnInterface) FuxiBFFClient {
	return &fuxiBFFClient{cc}
}

func (c *fuxiBFFClient) SetUrl(ctx context.Context, in *SetUrlRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, FuxiBFF_SetUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fuxiBFFClient) GetUrl(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error) {
	out := new(GetUrlResponse)
	err := c.cc.Invoke(ctx, FuxiBFF_GetUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fuxiBFFClient) AllocationUrl(ctx context.Context, in *AllocationRequest, opts ...grpc.CallOption) (*AllocationResponse, error) {
	out := new(AllocationResponse)
	err := c.cc.Invoke(ctx, FuxiBFF_AllocationUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FuxiBFFServer is the server API for FuxiBFF service.
// All implementations must embed UnimplementedFuxiBFFServer
// for forward compatibility
type FuxiBFFServer interface {
	SetUrl(context.Context, *SetUrlRequest) (*emptypb.Empty, error)
	GetUrl(context.Context, *GetUrlRequest) (*GetUrlResponse, error)
	AllocationUrl(context.Context, *AllocationRequest) (*AllocationResponse, error)
	mustEmbedUnimplementedFuxiBFFServer()
}

// UnimplementedFuxiBFFServer must be embedded to have forward compatible implementations.
type UnimplementedFuxiBFFServer struct {
}

func (UnimplementedFuxiBFFServer) SetUrl(context.Context, *SetUrlRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUrl not implemented")
}
func (UnimplementedFuxiBFFServer) GetUrl(context.Context, *GetUrlRequest) (*GetUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrl not implemented")
}
func (UnimplementedFuxiBFFServer) AllocationUrl(context.Context, *AllocationRequest) (*AllocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocationUrl not implemented")
}
func (UnimplementedFuxiBFFServer) mustEmbedUnimplementedFuxiBFFServer() {}

// UnsafeFuxiBFFServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FuxiBFFServer will
// result in compilation errors.
type UnsafeFuxiBFFServer interface {
	mustEmbedUnimplementedFuxiBFFServer()
}

func RegisterFuxiBFFServer(s grpc.ServiceRegistrar, srv FuxiBFFServer) {
	s.RegisterService(&FuxiBFF_ServiceDesc, srv)
}

func _FuxiBFF_SetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuxiBFFServer).SetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuxiBFF_SetUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuxiBFFServer).SetUrl(ctx, req.(*SetUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuxiBFF_GetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuxiBFFServer).GetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuxiBFF_GetUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuxiBFFServer).GetUrl(ctx, req.(*GetUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FuxiBFF_AllocationUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FuxiBFFServer).AllocationUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FuxiBFF_AllocationUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FuxiBFFServer).AllocationUrl(ctx, req.(*AllocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FuxiBFF_ServiceDesc is the grpc.ServiceDesc for FuxiBFF service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FuxiBFF_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.FuxiBFF",
	HandlerType: (*FuxiBFFServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetUrl",
			Handler:    _FuxiBFF_SetUrl_Handler,
		},
		{
			MethodName: "GetUrl",
			Handler:    _FuxiBFF_GetUrl_Handler,
		},
		{
			MethodName: "AllocationUrl",
			Handler:    _FuxiBFF_AllocationUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fuxi_bff.proto",
}
