// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package registration

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

// ServiceRAClient is the client API for ServiceRA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceRAClient interface {
	HandleUsersCreate(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseCreate, error)
	HandleSessionsCreate(ctx context.Context, in *RequestCreateSessions, opts ...grpc.CallOption) (*ResponseCreateSessions, error)
}

type serviceRAClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceRAClient(cc grpc.ClientConnInterface) ServiceRAClient {
	return &serviceRAClient{cc}
}

func (c *serviceRAClient) HandleUsersCreate(ctx context.Context, in *RequestCreate, opts ...grpc.CallOption) (*ResponseCreate, error) {
	out := new(ResponseCreate)
	err := c.cc.Invoke(ctx, "/grpc_service.ServiceRA/handleUsersCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceRAClient) HandleSessionsCreate(ctx context.Context, in *RequestCreateSessions, opts ...grpc.CallOption) (*ResponseCreateSessions, error) {
	out := new(ResponseCreateSessions)
	err := c.cc.Invoke(ctx, "/grpc_service.ServiceRA/handleSessionsCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceRAServer is the server API for ServiceRA service.
// All implementations must embed UnimplementedServiceRAServer
// for forward compatibility
type ServiceRAServer interface {
	HandleUsersCreate(context.Context, *RequestCreate) (*ResponseCreate, error)
	HandleSessionsCreate(context.Context, *RequestCreateSessions) (*ResponseCreateSessions, error)
	mustEmbedUnimplementedServiceRAServer()
}

// UnimplementedServiceRAServer must be embedded to have forward compatible implementations.
type UnimplementedServiceRAServer struct {
}

func (UnimplementedServiceRAServer) HandleUsersCreate(context.Context, *RequestCreate) (*ResponseCreate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleUsersCreate not implemented")
}
func (UnimplementedServiceRAServer) HandleSessionsCreate(context.Context, *RequestCreateSessions) (*ResponseCreateSessions, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleSessionsCreate not implemented")
}
func (UnimplementedServiceRAServer) mustEmbedUnimplementedServiceRAServer() {}

// UnsafeServiceRAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceRAServer will
// result in compilation errors.
type UnsafeServiceRAServer interface {
	mustEmbedUnimplementedServiceRAServer()
}

func RegisterServiceRAServer(s grpc.ServiceRegistrar, srv ServiceRAServer) {
	s.RegisterService(&ServiceRA_ServiceDesc, srv)
}

func _ServiceRA_HandleUsersCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRAServer).HandleUsersCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_service.ServiceRA/handleUsersCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRAServer).HandleUsersCreate(ctx, req.(*RequestCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _ServiceRA_HandleSessionsCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCreateSessions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceRAServer).HandleSessionsCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_service.ServiceRA/handleSessionsCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceRAServer).HandleSessionsCreate(ctx, req.(*RequestCreateSessions))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceRA_ServiceDesc is the grpc.ServiceDesc for ServiceRA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceRA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_service.ServiceRA",
	HandlerType: (*ServiceRAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "handleUsersCreate",
			Handler:    _ServiceRA_HandleUsersCreate_Handler,
		},
		{
			MethodName: "handleSessionsCreate",
			Handler:    _ServiceRA_HandleSessionsCreate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc_server.proto",
}