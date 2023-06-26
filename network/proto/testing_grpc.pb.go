// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	SayHello(ctx context.Context, in *GenericMessage, opts ...grpc.CallOption) (*GenericMessage, error)
	GetChattyServer(ctx context.Context, in *ChattyRequest, opts ...grpc.CallOption) (TestService_GetChattyServerClient, error)
	GetChattyClient(ctx context.Context, opts ...grpc.CallOption) (TestService_GetChattyClientClient, error)
	GetChattyBidi(ctx context.Context, opts ...grpc.CallOption) (TestService_GetChattyBidiClient, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) SayHello(ctx context.Context, in *GenericMessage, opts ...grpc.CallOption) (*GenericMessage, error) {
	out := new(GenericMessage)
	err := c.cc.Invoke(ctx, "/v1.TestService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) GetChattyServer(ctx context.Context, in *ChattyRequest, opts ...grpc.CallOption) (TestService_GetChattyServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[0], "/v1.TestService/GetChattyServer", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceGetChattyServerClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestService_GetChattyServerClient interface {
	Recv() (*GenericMessage, error)
	grpc.ClientStream
}

type testServiceGetChattyServerClient struct {
	grpc.ClientStream
}

func (x *testServiceGetChattyServerClient) Recv() (*GenericMessage, error) {
	m := new(GenericMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) GetChattyClient(ctx context.Context, opts ...grpc.CallOption) (TestService_GetChattyClientClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[1], "/v1.TestService/GetChattyClient", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceGetChattyClientClient{stream}
	return x, nil
}

type TestService_GetChattyClientClient interface {
	Send(*GenericMessage) error
	CloseAndRecv() (*GenericMessage, error)
	grpc.ClientStream
}

type testServiceGetChattyClientClient struct {
	grpc.ClientStream
}

func (x *testServiceGetChattyClientClient) Send(m *GenericMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceGetChattyClientClient) CloseAndRecv() (*GenericMessage, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(GenericMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testServiceClient) GetChattyBidi(ctx context.Context, opts ...grpc.CallOption) (TestService_GetChattyBidiClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[2], "/v1.TestService/GetChattyBidi", opts...)
	if err != nil {
		return nil, err
	}
	x := &testServiceGetChattyBidiClient{stream}
	return x, nil
}

type TestService_GetChattyBidiClient interface {
	Send(*GenericMessage) error
	Recv() (*GenericMessage, error)
	grpc.ClientStream
}

type testServiceGetChattyBidiClient struct {
	grpc.ClientStream
}

func (x *testServiceGetChattyBidiClient) Send(m *GenericMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testServiceGetChattyBidiClient) Recv() (*GenericMessage, error) {
	m := new(GenericMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	SayHello(context.Context, *GenericMessage) (*GenericMessage, error)
	GetChattyServer(*ChattyRequest, TestService_GetChattyServerServer) error
	GetChattyClient(TestService_GetChattyClientServer) error
	GetChattyBidi(TestService_GetChattyBidiServer) error
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) SayHello(context.Context, *GenericMessage) (*GenericMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedTestServiceServer) GetChattyServer(*ChattyRequest, TestService_GetChattyServerServer) error {
	return status.Errorf(codes.Unimplemented, "method GetChattyServer not implemented")
}
func (UnimplementedTestServiceServer) GetChattyClient(TestService_GetChattyClientServer) error {
	return status.Errorf(codes.Unimplemented, "method GetChattyClient not implemented")
}
func (UnimplementedTestServiceServer) GetChattyBidi(TestService_GetChattyBidiServer) error {
	return status.Errorf(codes.Unimplemented, "method GetChattyBidi not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenericMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.TestService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).SayHello(ctx, req.(*GenericMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_GetChattyServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ChattyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).GetChattyServer(m, &testServiceGetChattyServerServer{stream})
}

type TestService_GetChattyServerServer interface {
	Send(*GenericMessage) error
	grpc.ServerStream
}

type testServiceGetChattyServerServer struct {
	grpc.ServerStream
}

func (x *testServiceGetChattyServerServer) Send(m *GenericMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _TestService_GetChattyClient_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).GetChattyClient(&testServiceGetChattyClientServer{stream})
}

type TestService_GetChattyClientServer interface {
	SendAndClose(*GenericMessage) error
	Recv() (*GenericMessage, error)
	grpc.ServerStream
}

type testServiceGetChattyClientServer struct {
	grpc.ServerStream
}

func (x *testServiceGetChattyClientServer) SendAndClose(m *GenericMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceGetChattyClientServer) Recv() (*GenericMessage, error) {
	m := new(GenericMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TestService_GetChattyBidi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).GetChattyBidi(&testServiceGetChattyBidiServer{stream})
}

type TestService_GetChattyBidiServer interface {
	Send(*GenericMessage) error
	Recv() (*GenericMessage, error)
	grpc.ServerStream
}

type testServiceGetChattyBidiServer struct {
	grpc.ServerStream
}

func (x *testServiceGetChattyBidiServer) Send(m *GenericMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testServiceGetChattyBidiServer) Recv() (*GenericMessage, error) {
	m := new(GenericMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _TestService_SayHello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetChattyServer",
			Handler:       _TestService_GetChattyServer_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetChattyClient",
			Handler:       _TestService_GetChattyClient_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetChattyBidi",
			Handler:       _TestService_GetChattyBidi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "testing.proto",
}
