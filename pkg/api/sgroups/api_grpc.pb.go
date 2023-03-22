// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: sgroups/api.proto

package sgroups

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

// SecGroupServiceClient is the client API for SecGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecGroupServiceClient interface {
	Sync(ctx context.Context, in *SyncReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SyncStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SyncStatusResp, error)
	ListNetworks(ctx context.Context, in *ListNetworksReq, opts ...grpc.CallOption) (*ListNetworksResp, error)
	ListSecurityGroups(ctx context.Context, in *ListSecurityGroupsReq, opts ...grpc.CallOption) (*ListSecurityGroupsResp, error)
	GetSgSubnets(ctx context.Context, in *GetSgSubnetsReq, opts ...grpc.CallOption) (*GetSgSubnetsResp, error)
	GetRules(ctx context.Context, in *GetRulesReq, opts ...grpc.CallOption) (*RulesResp, error)
	FindRules(ctx context.Context, in *FindRulesReq, opts ...grpc.CallOption) (*RulesResp, error)
	GetSecGroupForAddress(ctx context.Context, in *GetSecGroupForAddressReq, opts ...grpc.CallOption) (*SecGroup, error)
}

type secGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecGroupServiceClient(cc grpc.ClientConnInterface) SecGroupServiceClient {
	return &secGroupServiceClient{cc}
}

func (c *secGroupServiceClient) Sync(ctx context.Context, in *SyncReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hbf.v1.sgroups.SecGroupService/Sync", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secGroupServiceClient) SyncStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SyncStatusResp, error) {
	out := new(SyncStatusResp)
	err := c.cc.Invoke(ctx, "/hbf.v1.sgroups.SecGroupService/SyncStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secGroupServiceClient) ListNetworks(ctx context.Context, in *ListNetworksReq, opts ...grpc.CallOption) (*ListNetworksResp, error) {
	out := new(ListNetworksResp)
	err := c.cc.Invoke(ctx, "/hbf.v1.sgroups.SecGroupService/ListNetworks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secGroupServiceClient) ListSecurityGroups(ctx context.Context, in *ListSecurityGroupsReq, opts ...grpc.CallOption) (*ListSecurityGroupsResp, error) {
	out := new(ListSecurityGroupsResp)
	err := c.cc.Invoke(ctx, "/hbf.v1.sgroups.SecGroupService/ListSecurityGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secGroupServiceClient) GetSgSubnets(ctx context.Context, in *GetSgSubnetsReq, opts ...grpc.CallOption) (*GetSgSubnetsResp, error) {
	out := new(GetSgSubnetsResp)
	err := c.cc.Invoke(ctx, "/hbf.v1.sgroups.SecGroupService/GetSgSubnets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secGroupServiceClient) GetRules(ctx context.Context, in *GetRulesReq, opts ...grpc.CallOption) (*RulesResp, error) {
	out := new(RulesResp)
	err := c.cc.Invoke(ctx, "/hbf.v1.sgroups.SecGroupService/GetRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secGroupServiceClient) FindRules(ctx context.Context, in *FindRulesReq, opts ...grpc.CallOption) (*RulesResp, error) {
	out := new(RulesResp)
	err := c.cc.Invoke(ctx, "/hbf.v1.sgroups.SecGroupService/FindRules", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secGroupServiceClient) GetSecGroupForAddress(ctx context.Context, in *GetSecGroupForAddressReq, opts ...grpc.CallOption) (*SecGroup, error) {
	out := new(SecGroup)
	err := c.cc.Invoke(ctx, "/hbf.v1.sgroups.SecGroupService/GetSecGroupForAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecGroupServiceServer is the server API for SecGroupService service.
// All implementations must embed UnimplementedSecGroupServiceServer
// for forward compatibility
type SecGroupServiceServer interface {
	Sync(context.Context, *SyncReq) (*emptypb.Empty, error)
	SyncStatus(context.Context, *emptypb.Empty) (*SyncStatusResp, error)
	ListNetworks(context.Context, *ListNetworksReq) (*ListNetworksResp, error)
	ListSecurityGroups(context.Context, *ListSecurityGroupsReq) (*ListSecurityGroupsResp, error)
	GetSgSubnets(context.Context, *GetSgSubnetsReq) (*GetSgSubnetsResp, error)
	GetRules(context.Context, *GetRulesReq) (*RulesResp, error)
	FindRules(context.Context, *FindRulesReq) (*RulesResp, error)
	GetSecGroupForAddress(context.Context, *GetSecGroupForAddressReq) (*SecGroup, error)
	mustEmbedUnimplementedSecGroupServiceServer()
}

// UnimplementedSecGroupServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSecGroupServiceServer struct {
}

func (UnimplementedSecGroupServiceServer) Sync(context.Context, *SyncReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sync not implemented")
}
func (UnimplementedSecGroupServiceServer) SyncStatus(context.Context, *emptypb.Empty) (*SyncStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncStatus not implemented")
}
func (UnimplementedSecGroupServiceServer) ListNetworks(context.Context, *ListNetworksReq) (*ListNetworksResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNetworks not implemented")
}
func (UnimplementedSecGroupServiceServer) ListSecurityGroups(context.Context, *ListSecurityGroupsReq) (*ListSecurityGroupsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecurityGroups not implemented")
}
func (UnimplementedSecGroupServiceServer) GetSgSubnets(context.Context, *GetSgSubnetsReq) (*GetSgSubnetsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSgSubnets not implemented")
}
func (UnimplementedSecGroupServiceServer) GetRules(context.Context, *GetRulesReq) (*RulesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRules not implemented")
}
func (UnimplementedSecGroupServiceServer) FindRules(context.Context, *FindRulesReq) (*RulesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindRules not implemented")
}
func (UnimplementedSecGroupServiceServer) GetSecGroupForAddress(context.Context, *GetSecGroupForAddressReq) (*SecGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecGroupForAddress not implemented")
}
func (UnimplementedSecGroupServiceServer) mustEmbedUnimplementedSecGroupServiceServer() {}

// UnsafeSecGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecGroupServiceServer will
// result in compilation errors.
type UnsafeSecGroupServiceServer interface {
	mustEmbedUnimplementedSecGroupServiceServer()
}

func RegisterSecGroupServiceServer(s grpc.ServiceRegistrar, srv SecGroupServiceServer) {
	s.RegisterService(&SecGroupService_ServiceDesc, srv)
}

func _SecGroupService_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecGroupServiceServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hbf.v1.sgroups.SecGroupService/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecGroupServiceServer).Sync(ctx, req.(*SyncReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecGroupService_SyncStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecGroupServiceServer).SyncStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hbf.v1.sgroups.SecGroupService/SyncStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecGroupServiceServer).SyncStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecGroupService_ListNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworksReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecGroupServiceServer).ListNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hbf.v1.sgroups.SecGroupService/ListNetworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecGroupServiceServer).ListNetworks(ctx, req.(*ListNetworksReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecGroupService_ListSecurityGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecurityGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecGroupServiceServer).ListSecurityGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hbf.v1.sgroups.SecGroupService/ListSecurityGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecGroupServiceServer).ListSecurityGroups(ctx, req.(*ListSecurityGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecGroupService_GetSgSubnets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSgSubnetsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecGroupServiceServer).GetSgSubnets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hbf.v1.sgroups.SecGroupService/GetSgSubnets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecGroupServiceServer).GetSgSubnets(ctx, req.(*GetSgSubnetsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecGroupService_GetRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRulesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecGroupServiceServer).GetRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hbf.v1.sgroups.SecGroupService/GetRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecGroupServiceServer).GetRules(ctx, req.(*GetRulesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecGroupService_FindRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRulesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecGroupServiceServer).FindRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hbf.v1.sgroups.SecGroupService/FindRules",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecGroupServiceServer).FindRules(ctx, req.(*FindRulesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecGroupService_GetSecGroupForAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecGroupForAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecGroupServiceServer).GetSecGroupForAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hbf.v1.sgroups.SecGroupService/GetSecGroupForAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecGroupServiceServer).GetSecGroupForAddress(ctx, req.(*GetSecGroupForAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SecGroupService_ServiceDesc is the grpc.ServiceDesc for SecGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hbf.v1.sgroups.SecGroupService",
	HandlerType: (*SecGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sync",
			Handler:    _SecGroupService_Sync_Handler,
		},
		{
			MethodName: "SyncStatus",
			Handler:    _SecGroupService_SyncStatus_Handler,
		},
		{
			MethodName: "ListNetworks",
			Handler:    _SecGroupService_ListNetworks_Handler,
		},
		{
			MethodName: "ListSecurityGroups",
			Handler:    _SecGroupService_ListSecurityGroups_Handler,
		},
		{
			MethodName: "GetSgSubnets",
			Handler:    _SecGroupService_GetSgSubnets_Handler,
		},
		{
			MethodName: "GetRules",
			Handler:    _SecGroupService_GetRules_Handler,
		},
		{
			MethodName: "FindRules",
			Handler:    _SecGroupService_FindRules_Handler,
		},
		{
			MethodName: "GetSecGroupForAddress",
			Handler:    _SecGroupService_GetSecGroupForAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sgroups/api.proto",
}
