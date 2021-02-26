// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package backend

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

// RouteClient is the client API for Route service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RouteClient interface {
	AddTrustedUser(ctx context.Context, opts ...grpc.CallOption) (Route_AddTrustedUserClient, error)
	UpdateTrustedUser(ctx context.Context, opts ...grpc.CallOption) (Route_UpdateTrustedUserClient, error)
	RemoveTrustedUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error)
}

type routeClient struct {
	cc grpc.ClientConnInterface
}

func NewRouteClient(cc grpc.ClientConnInterface) RouteClient {
	return &routeClient{cc}
}

func (c *routeClient) AddTrustedUser(ctx context.Context, opts ...grpc.CallOption) (Route_AddTrustedUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &Route_ServiceDesc.Streams[0], "/route.Route/AddTrustedUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeAddTrustedUserClient{stream}
	return x, nil
}

type Route_AddTrustedUserClient interface {
	Send(*User) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type routeAddTrustedUserClient struct {
	grpc.ClientStream
}

func (x *routeAddTrustedUserClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeAddTrustedUserClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeClient) UpdateTrustedUser(ctx context.Context, opts ...grpc.CallOption) (Route_UpdateTrustedUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &Route_ServiceDesc.Streams[1], "/route.Route/UpdateTrustedUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeUpdateTrustedUserClient{stream}
	return x, nil
}

type Route_UpdateTrustedUserClient interface {
	Send(*User) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type routeUpdateTrustedUserClient struct {
	grpc.ClientStream
}

func (x *routeUpdateTrustedUserClient) Send(m *User) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeUpdateTrustedUserClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeClient) RemoveTrustedUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/route.Route/RemoveTrustedUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouteServer is the server API for Route service.
// All implementations must embed UnimplementedRouteServer
// for forward compatibility
type RouteServer interface {
	AddTrustedUser(Route_AddTrustedUserServer) error
	UpdateTrustedUser(Route_UpdateTrustedUserServer) error
	RemoveTrustedUser(context.Context, *User) (*Empty, error)
	mustEmbedUnimplementedRouteServer()
}

// UnimplementedRouteServer must be embedded to have forward compatible implementations.
type UnimplementedRouteServer struct {
}

func (UnimplementedRouteServer) AddTrustedUser(Route_AddTrustedUserServer) error {
	return status.Errorf(codes.Unimplemented, "method AddTrustedUser not implemented")
}
func (UnimplementedRouteServer) UpdateTrustedUser(Route_UpdateTrustedUserServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateTrustedUser not implemented")
}
func (UnimplementedRouteServer) RemoveTrustedUser(context.Context, *User) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTrustedUser not implemented")
}
func (UnimplementedRouteServer) mustEmbedUnimplementedRouteServer() {}

// UnsafeRouteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RouteServer will
// result in compilation errors.
type UnsafeRouteServer interface {
	mustEmbedUnimplementedRouteServer()
}

func RegisterRouteServer(s grpc.ServiceRegistrar, srv RouteServer) {
	s.RegisterService(&Route_ServiceDesc, srv)
}

func _Route_AddTrustedUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteServer).AddTrustedUser(&routeAddTrustedUserServer{stream})
}

type Route_AddTrustedUserServer interface {
	SendAndClose(*Empty) error
	Recv() (*User, error)
	grpc.ServerStream
}

type routeAddTrustedUserServer struct {
	grpc.ServerStream
}

func (x *routeAddTrustedUserServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeAddTrustedUserServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Route_UpdateTrustedUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteServer).UpdateTrustedUser(&routeUpdateTrustedUserServer{stream})
}

type Route_UpdateTrustedUserServer interface {
	SendAndClose(*Empty) error
	Recv() (*User, error)
	grpc.ServerStream
}

type routeUpdateTrustedUserServer struct {
	grpc.ServerStream
}

func (x *routeUpdateTrustedUserServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeUpdateTrustedUserServer) Recv() (*User, error) {
	m := new(User)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Route_RemoveTrustedUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteServer).RemoveTrustedUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/route.Route/RemoveTrustedUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteServer).RemoveTrustedUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// Route_ServiceDesc is the grpc.ServiceDesc for Route service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Route_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "route.Route",
	HandlerType: (*RouteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoveTrustedUser",
			Handler:    _Route_RemoveTrustedUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddTrustedUser",
			Handler:       _Route_AddTrustedUser_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UpdateTrustedUser",
			Handler:       _Route_UpdateTrustedUser_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/route.proto",
}
