// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package echo_op

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// EchoServerClient is the client API for EchoServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServerClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoServerClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServerClient(cc grpc.ClientConnInterface) EchoServerClient {
	return &echoServerClient{cc}
}

func (c *echoServerClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/EchoServer/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServerServer is the server API for EchoServer service.
// All implementations must embed UnimplementedEchoServerServer
// for forward compatibility
type EchoServerServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedEchoServerServer()
}

// UnimplementedEchoServerServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServerServer struct {
}

func (UnimplementedEchoServerServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedEchoServerServer) mustEmbedUnimplementedEchoServerServer() {}

// UnsafeEchoServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServerServer will
// result in compilation errors.
type UnsafeEchoServerServer interface {
	mustEmbedUnimplementedEchoServerServer()
}

func RegisterEchoServerServer(s grpc.ServiceRegistrar, srv EchoServerServer) {
	s.RegisterService(&_EchoServer_serviceDesc, srv)
}

func _EchoServer_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServerServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EchoServer/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServerServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "EchoServer",
	HandlerType: (*EchoServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoServer_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}
