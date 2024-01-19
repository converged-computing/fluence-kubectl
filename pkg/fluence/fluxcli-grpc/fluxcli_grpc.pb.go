// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: pkg/fluence/fluxcli-grpc/fluxcli.proto

package fluxcli

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExternalPluginServiceClient is the client API for ExternalPluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalPluginServiceClient interface {
	ListGroups(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
	GetGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error)
}

type externalPluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalPluginServiceClient(cc grpc.ClientConnInterface) ExternalPluginServiceClient {
	return &externalPluginServiceClient{cc}
}

func (c *externalPluginServiceClient) ListGroups(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, "/fluxcli.ExternalPluginService/ListGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *externalPluginServiceClient) GetGroup(ctx context.Context, in *GroupRequest, opts ...grpc.CallOption) (*GroupResponse, error) {
	out := new(GroupResponse)
	err := c.cc.Invoke(ctx, "/fluxcli.ExternalPluginService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalPluginServiceServer is the server API for ExternalPluginService service.
// All implementations must embed UnimplementedExternalPluginServiceServer
// for forward compatibility
type ExternalPluginServiceServer interface {
	ListGroups(context.Context, *GroupRequest) (*GroupResponse, error)
	GetGroup(context.Context, *GroupRequest) (*GroupResponse, error)
	mustEmbedUnimplementedExternalPluginServiceServer()
}

// UnimplementedExternalPluginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExternalPluginServiceServer struct {
}

func (UnimplementedExternalPluginServiceServer) ListGroups(context.Context, *GroupRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGroups not implemented")
}
func (UnimplementedExternalPluginServiceServer) GetGroup(context.Context, *GroupRequest) (*GroupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedExternalPluginServiceServer) mustEmbedUnimplementedExternalPluginServiceServer() {}

// UnsafeExternalPluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalPluginServiceServer will
// result in compilation errors.
type UnsafeExternalPluginServiceServer interface {
	mustEmbedUnimplementedExternalPluginServiceServer()
}

func RegisterExternalPluginServiceServer(s grpc.ServiceRegistrar, srv ExternalPluginServiceServer) {
	s.RegisterService(&ExternalPluginService_ServiceDesc, srv)
}

func _ExternalPluginService_ListGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalPluginServiceServer).ListGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fluxcli.ExternalPluginService/ListGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalPluginServiceServer).ListGroups(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExternalPluginService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalPluginServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fluxcli.ExternalPluginService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalPluginServiceServer).GetGroup(ctx, req.(*GroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalPluginService_ServiceDesc is the grpc.ServiceDesc for ExternalPluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalPluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "fluxcli.ExternalPluginService",
	HandlerType: (*ExternalPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListGroups",
			Handler:    _ExternalPluginService_ListGroups_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _ExternalPluginService_GetGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/fluence/fluxcli-grpc/fluxcli.proto",
}
