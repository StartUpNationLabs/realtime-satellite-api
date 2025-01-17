// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: proto/satellite/v1/satellite.proto

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
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SatelliteService_GetSatellitePositions_FullMethodName = "/satellite.v1.SatelliteService/GetSatellitePositions"
	SatelliteService_GetSatelliteDetail_FullMethodName    = "/satellite.v1.SatelliteService/GetSatelliteDetail"
	SatelliteService_GetSatelliteGroups_FullMethodName    = "/satellite.v1.SatelliteService/GetSatelliteGroups"
	SatelliteService_GetMinimalSatellites_FullMethodName  = "/satellite.v1.SatelliteService/GetMinimalSatellites"
	SatelliteService_GetSatellitePath_FullMethodName      = "/satellite.v1.SatelliteService/GetSatellitePath"
)

// SatelliteServiceClient is the client API for SatelliteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SatelliteServiceClient interface {
	GetSatellitePositions(ctx context.Context, in *GetSatellitePositionsRequest, opts ...grpc.CallOption) (*GetSatellitePositionsResponse, error)
	GetSatelliteDetail(ctx context.Context, in *SatelliteDetailRequest, opts ...grpc.CallOption) (*SatelliteDetail, error)
	GetSatelliteGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSatelliteGroupsResponse, error)
	GetMinimalSatellites(ctx context.Context, in *GetSatellitePositionsRequest, opts ...grpc.CallOption) (*GetMinimalSatellitesResponse, error)
	GetSatellitePath(ctx context.Context, in *SatellitePathRequest, opts ...grpc.CallOption) (*GetSatellitePathResponse, error)
}

type satelliteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSatelliteServiceClient(cc grpc.ClientConnInterface) SatelliteServiceClient {
	return &satelliteServiceClient{cc}
}

func (c *satelliteServiceClient) GetSatellitePositions(ctx context.Context, in *GetSatellitePositionsRequest, opts ...grpc.CallOption) (*GetSatellitePositionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSatellitePositionsResponse)
	err := c.cc.Invoke(ctx, SatelliteService_GetSatellitePositions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *satelliteServiceClient) GetSatelliteDetail(ctx context.Context, in *SatelliteDetailRequest, opts ...grpc.CallOption) (*SatelliteDetail, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SatelliteDetail)
	err := c.cc.Invoke(ctx, SatelliteService_GetSatelliteDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *satelliteServiceClient) GetSatelliteGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetSatelliteGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSatelliteGroupsResponse)
	err := c.cc.Invoke(ctx, SatelliteService_GetSatelliteGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *satelliteServiceClient) GetMinimalSatellites(ctx context.Context, in *GetSatellitePositionsRequest, opts ...grpc.CallOption) (*GetMinimalSatellitesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMinimalSatellitesResponse)
	err := c.cc.Invoke(ctx, SatelliteService_GetMinimalSatellites_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *satelliteServiceClient) GetSatellitePath(ctx context.Context, in *SatellitePathRequest, opts ...grpc.CallOption) (*GetSatellitePathResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSatellitePathResponse)
	err := c.cc.Invoke(ctx, SatelliteService_GetSatellitePath_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SatelliteServiceServer is the server API for SatelliteService service.
// All implementations must embed UnimplementedSatelliteServiceServer
// for forward compatibility
type SatelliteServiceServer interface {
	GetSatellitePositions(context.Context, *GetSatellitePositionsRequest) (*GetSatellitePositionsResponse, error)
	GetSatelliteDetail(context.Context, *SatelliteDetailRequest) (*SatelliteDetail, error)
	GetSatelliteGroups(context.Context, *emptypb.Empty) (*GetSatelliteGroupsResponse, error)
	GetMinimalSatellites(context.Context, *GetSatellitePositionsRequest) (*GetMinimalSatellitesResponse, error)
	GetSatellitePath(context.Context, *SatellitePathRequest) (*GetSatellitePathResponse, error)
	mustEmbedUnimplementedSatelliteServiceServer()
}

// UnimplementedSatelliteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSatelliteServiceServer struct {
}

func (UnimplementedSatelliteServiceServer) GetSatellitePositions(context.Context, *GetSatellitePositionsRequest) (*GetSatellitePositionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSatellitePositions not implemented")
}
func (UnimplementedSatelliteServiceServer) GetSatelliteDetail(context.Context, *SatelliteDetailRequest) (*SatelliteDetail, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSatelliteDetail not implemented")
}
func (UnimplementedSatelliteServiceServer) GetSatelliteGroups(context.Context, *emptypb.Empty) (*GetSatelliteGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSatelliteGroups not implemented")
}
func (UnimplementedSatelliteServiceServer) GetMinimalSatellites(context.Context, *GetSatellitePositionsRequest) (*GetMinimalSatellitesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMinimalSatellites not implemented")
}
func (UnimplementedSatelliteServiceServer) GetSatellitePath(context.Context, *SatellitePathRequest) (*GetSatellitePathResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSatellitePath not implemented")
}
func (UnimplementedSatelliteServiceServer) mustEmbedUnimplementedSatelliteServiceServer() {}

// UnsafeSatelliteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SatelliteServiceServer will
// result in compilation errors.
type UnsafeSatelliteServiceServer interface {
	mustEmbedUnimplementedSatelliteServiceServer()
}

func RegisterSatelliteServiceServer(s grpc.ServiceRegistrar, srv SatelliteServiceServer) {
	s.RegisterService(&SatelliteService_ServiceDesc, srv)
}

func _SatelliteService_GetSatellitePositions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSatellitePositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SatelliteServiceServer).GetSatellitePositions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SatelliteService_GetSatellitePositions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SatelliteServiceServer).GetSatellitePositions(ctx, req.(*GetSatellitePositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SatelliteService_GetSatelliteDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SatelliteDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SatelliteServiceServer).GetSatelliteDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SatelliteService_GetSatelliteDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SatelliteServiceServer).GetSatelliteDetail(ctx, req.(*SatelliteDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SatelliteService_GetSatelliteGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SatelliteServiceServer).GetSatelliteGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SatelliteService_GetSatelliteGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SatelliteServiceServer).GetSatelliteGroups(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SatelliteService_GetMinimalSatellites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSatellitePositionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SatelliteServiceServer).GetMinimalSatellites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SatelliteService_GetMinimalSatellites_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SatelliteServiceServer).GetMinimalSatellites(ctx, req.(*GetSatellitePositionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SatelliteService_GetSatellitePath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SatellitePathRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SatelliteServiceServer).GetSatellitePath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SatelliteService_GetSatellitePath_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SatelliteServiceServer).GetSatellitePath(ctx, req.(*SatellitePathRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SatelliteService_ServiceDesc is the grpc.ServiceDesc for SatelliteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SatelliteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "satellite.v1.SatelliteService",
	HandlerType: (*SatelliteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSatellitePositions",
			Handler:    _SatelliteService_GetSatellitePositions_Handler,
		},
		{
			MethodName: "GetSatelliteDetail",
			Handler:    _SatelliteService_GetSatelliteDetail_Handler,
		},
		{
			MethodName: "GetSatelliteGroups",
			Handler:    _SatelliteService_GetSatelliteGroups_Handler,
		},
		{
			MethodName: "GetMinimalSatellites",
			Handler:    _SatelliteService_GetMinimalSatellites_Handler,
		},
		{
			MethodName: "GetSatellitePath",
			Handler:    _SatelliteService_GetSatellitePath_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/satellite/v1/satellite.proto",
}
