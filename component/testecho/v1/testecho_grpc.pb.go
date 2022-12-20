// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// TestEchoServiceClient is the client API for TestEchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestEchoServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
	EchoMultiple(ctx context.Context, in *EchoMultipleRequest, opts ...grpc.CallOption) (TestEchoService_EchoMultipleClient, error)
	EchoBiDi(ctx context.Context, opts ...grpc.CallOption) (TestEchoService_EchoBiDiClient, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
}

type testEchoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestEchoServiceClient(cc grpc.ClientConnInterface) TestEchoServiceClient {
	return &testEchoServiceClient{cc}
}

func (c *testEchoServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/viam.component.testecho.v1.TestEchoService/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testEchoServiceClient) EchoMultiple(ctx context.Context, in *EchoMultipleRequest, opts ...grpc.CallOption) (TestEchoService_EchoMultipleClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestEchoService_ServiceDesc.Streams[0], "/viam.component.testecho.v1.TestEchoService/EchoMultiple", opts...)
	if err != nil {
		return nil, err
	}
	x := &testEchoServiceEchoMultipleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestEchoService_EchoMultipleClient interface {
	Recv() (*EchoMultipleResponse, error)
	grpc.ClientStream
}

type testEchoServiceEchoMultipleClient struct {
	grpc.ClientStream
}

func (x *testEchoServiceEchoMultipleClient) Recv() (*EchoMultipleResponse, error) {
	m := new(EchoMultipleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testEchoServiceClient) EchoBiDi(ctx context.Context, opts ...grpc.CallOption) (TestEchoService_EchoBiDiClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestEchoService_ServiceDesc.Streams[1], "/viam.component.testecho.v1.TestEchoService/EchoBiDi", opts...)
	if err != nil {
		return nil, err
	}
	x := &testEchoServiceEchoBiDiClient{stream}
	return x, nil
}

type TestEchoService_EchoBiDiClient interface {
	Send(*EchoBiDiRequest) error
	Recv() (*EchoBiDiResponse, error)
	grpc.ClientStream
}

type testEchoServiceEchoBiDiClient struct {
	grpc.ClientStream
}

func (x *testEchoServiceEchoBiDiClient) Send(m *EchoBiDiRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testEchoServiceEchoBiDiClient) Recv() (*EchoBiDiResponse, error) {
	m := new(EchoBiDiResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testEchoServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, "/viam.component.testecho.v1.TestEchoService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestEchoServiceServer is the server API for TestEchoService service.
// All implementations must embed UnimplementedTestEchoServiceServer
// for forward compatibility
type TestEchoServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
	EchoMultiple(*EchoMultipleRequest, TestEchoService_EchoMultipleServer) error
	EchoBiDi(TestEchoService_EchoBiDiServer) error
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	mustEmbedUnimplementedTestEchoServiceServer()
}

// UnimplementedTestEchoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestEchoServiceServer struct {
}

func (UnimplementedTestEchoServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedTestEchoServiceServer) EchoMultiple(*EchoMultipleRequest, TestEchoService_EchoMultipleServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoMultiple not implemented")
}
func (UnimplementedTestEchoServiceServer) EchoBiDi(TestEchoService_EchoBiDiServer) error {
	return status.Errorf(codes.Unimplemented, "method EchoBiDi not implemented")
}
func (UnimplementedTestEchoServiceServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedTestEchoServiceServer) mustEmbedUnimplementedTestEchoServiceServer() {}

// UnsafeTestEchoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestEchoServiceServer will
// result in compilation errors.
type UnsafeTestEchoServiceServer interface {
	mustEmbedUnimplementedTestEchoServiceServer()
}

func RegisterTestEchoServiceServer(s grpc.ServiceRegistrar, srv TestEchoServiceServer) {
	s.RegisterService(&TestEchoService_ServiceDesc, srv)
}

func _TestEchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestEchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.testecho.v1.TestEchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestEchoServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestEchoService_EchoMultiple_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EchoMultipleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestEchoServiceServer).EchoMultiple(m, &testEchoServiceEchoMultipleServer{stream})
}

type TestEchoService_EchoMultipleServer interface {
	Send(*EchoMultipleResponse) error
	grpc.ServerStream
}

type testEchoServiceEchoMultipleServer struct {
	grpc.ServerStream
}

func (x *testEchoServiceEchoMultipleServer) Send(m *EchoMultipleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _TestEchoService_EchoBiDi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestEchoServiceServer).EchoBiDi(&testEchoServiceEchoBiDiServer{stream})
}

type TestEchoService_EchoBiDiServer interface {
	Send(*EchoBiDiResponse) error
	Recv() (*EchoBiDiRequest, error)
	grpc.ServerStream
}

type testEchoServiceEchoBiDiServer struct {
	grpc.ServerStream
}

func (x *testEchoServiceEchoBiDiServer) Send(m *EchoBiDiResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testEchoServiceEchoBiDiServer) Recv() (*EchoBiDiRequest, error) {
	m := new(EchoBiDiRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestEchoService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestEchoServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/viam.component.testecho.v1.TestEchoService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestEchoServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestEchoService_ServiceDesc is the grpc.ServiceDesc for TestEchoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestEchoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "viam.component.testecho.v1.TestEchoService",
	HandlerType: (*TestEchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _TestEchoService_Echo_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _TestEchoService_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EchoMultiple",
			Handler:       _TestEchoService_EchoMultiple_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EchoBiDi",
			Handler:       _TestEchoService_EchoBiDi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "component/testecho/v1/testecho.proto",
}