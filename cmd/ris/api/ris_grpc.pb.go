// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// RoutingInformationServiceClient is the client API for RoutingInformationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoutingInformationServiceClient interface {
	LPM(ctx context.Context, in *LPMRequest, opts ...grpc.CallOption) (*LPMResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetRouters(ctx context.Context, in *GetRoutersRequest, opts ...grpc.CallOption) (*GetRoutersResponse, error)
	GetLonger(ctx context.Context, in *GetLongerRequest, opts ...grpc.CallOption) (*GetLongerResponse, error)
	ObserveRIB(ctx context.Context, in *ObserveRIBRequest, opts ...grpc.CallOption) (RoutingInformationService_ObserveRIBClient, error)
	DumpRIB(ctx context.Context, in *DumpRIBRequest, opts ...grpc.CallOption) (RoutingInformationService_DumpRIBClient, error)
	ObserveRouters(ctx context.Context, in *ObserveRoutersRequest, opts ...grpc.CallOption) (RoutingInformationService_ObserveRoutersClient, error)
}

type routingInformationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutingInformationServiceClient(cc grpc.ClientConnInterface) RoutingInformationServiceClient {
	return &routingInformationServiceClient{cc}
}

func (c *routingInformationServiceClient) LPM(ctx context.Context, in *LPMRequest, opts ...grpc.CallOption) (*LPMResponse, error) {
	out := new(LPMResponse)
	err := c.cc.Invoke(ctx, "/bio.ris.RoutingInformationService/LPM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingInformationServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/bio.ris.RoutingInformationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingInformationServiceClient) GetRouters(ctx context.Context, in *GetRoutersRequest, opts ...grpc.CallOption) (*GetRoutersResponse, error) {
	out := new(GetRoutersResponse)
	err := c.cc.Invoke(ctx, "/bio.ris.RoutingInformationService/GetRouters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingInformationServiceClient) GetLonger(ctx context.Context, in *GetLongerRequest, opts ...grpc.CallOption) (*GetLongerResponse, error) {
	out := new(GetLongerResponse)
	err := c.cc.Invoke(ctx, "/bio.ris.RoutingInformationService/GetLonger", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routingInformationServiceClient) ObserveRIB(ctx context.Context, in *ObserveRIBRequest, opts ...grpc.CallOption) (RoutingInformationService_ObserveRIBClient, error) {
	stream, err := c.cc.NewStream(ctx, &RoutingInformationService_ServiceDesc.Streams[0], "/bio.ris.RoutingInformationService/ObserveRIB", opts...)
	if err != nil {
		return nil, err
	}
	x := &routingInformationServiceObserveRIBClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoutingInformationService_ObserveRIBClient interface {
	Recv() (*RIBUpdate, error)
	grpc.ClientStream
}

type routingInformationServiceObserveRIBClient struct {
	grpc.ClientStream
}

func (x *routingInformationServiceObserveRIBClient) Recv() (*RIBUpdate, error) {
	m := new(RIBUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routingInformationServiceClient) DumpRIB(ctx context.Context, in *DumpRIBRequest, opts ...grpc.CallOption) (RoutingInformationService_DumpRIBClient, error) {
	stream, err := c.cc.NewStream(ctx, &RoutingInformationService_ServiceDesc.Streams[1], "/bio.ris.RoutingInformationService/DumpRIB", opts...)
	if err != nil {
		return nil, err
	}
	x := &routingInformationServiceDumpRIBClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoutingInformationService_DumpRIBClient interface {
	Recv() (*DumpRIBReply, error)
	grpc.ClientStream
}

type routingInformationServiceDumpRIBClient struct {
	grpc.ClientStream
}

func (x *routingInformationServiceDumpRIBClient) Recv() (*DumpRIBReply, error) {
	m := new(DumpRIBReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routingInformationServiceClient) ObserveRouters(ctx context.Context, in *ObserveRoutersRequest, opts ...grpc.CallOption) (RoutingInformationService_ObserveRoutersClient, error) {
	stream, err := c.cc.NewStream(ctx, &RoutingInformationService_ServiceDesc.Streams[2], "/bio.ris.RoutingInformationService/ObserveRouters", opts...)
	if err != nil {
		return nil, err
	}
	x := &routingInformationServiceObserveRoutersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RoutingInformationService_ObserveRoutersClient interface {
	Recv() (*RouterUpdate, error)
	grpc.ClientStream
}

type routingInformationServiceObserveRoutersClient struct {
	grpc.ClientStream
}

func (x *routingInformationServiceObserveRoutersClient) Recv() (*RouterUpdate, error) {
	m := new(RouterUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoutingInformationServiceServer is the server API for RoutingInformationService service.
// All implementations must embed UnimplementedRoutingInformationServiceServer
// for forward compatibility
type RoutingInformationServiceServer interface {
	LPM(context.Context, *LPMRequest) (*LPMResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetRouters(context.Context, *GetRoutersRequest) (*GetRoutersResponse, error)
	GetLonger(context.Context, *GetLongerRequest) (*GetLongerResponse, error)
	ObserveRIB(*ObserveRIBRequest, RoutingInformationService_ObserveRIBServer) error
	DumpRIB(*DumpRIBRequest, RoutingInformationService_DumpRIBServer) error
	ObserveRouters(*ObserveRoutersRequest, RoutingInformationService_ObserveRoutersServer) error
	mustEmbedUnimplementedRoutingInformationServiceServer()
}

// UnimplementedRoutingInformationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoutingInformationServiceServer struct {
}

func (UnimplementedRoutingInformationServiceServer) LPM(context.Context, *LPMRequest) (*LPMResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LPM not implemented")
}
func (UnimplementedRoutingInformationServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRoutingInformationServiceServer) GetRouters(context.Context, *GetRoutersRequest) (*GetRoutersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRouters not implemented")
}
func (UnimplementedRoutingInformationServiceServer) GetLonger(context.Context, *GetLongerRequest) (*GetLongerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLonger not implemented")
}
func (UnimplementedRoutingInformationServiceServer) ObserveRIB(*ObserveRIBRequest, RoutingInformationService_ObserveRIBServer) error {
	return status.Errorf(codes.Unimplemented, "method ObserveRIB not implemented")
}
func (UnimplementedRoutingInformationServiceServer) DumpRIB(*DumpRIBRequest, RoutingInformationService_DumpRIBServer) error {
	return status.Errorf(codes.Unimplemented, "method DumpRIB not implemented")
}
func (UnimplementedRoutingInformationServiceServer) ObserveRouters(*ObserveRoutersRequest, RoutingInformationService_ObserveRoutersServer) error {
	return status.Errorf(codes.Unimplemented, "method ObserveRouters not implemented")
}
func (UnimplementedRoutingInformationServiceServer) mustEmbedUnimplementedRoutingInformationServiceServer() {
}

// UnsafeRoutingInformationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoutingInformationServiceServer will
// result in compilation errors.
type UnsafeRoutingInformationServiceServer interface {
	mustEmbedUnimplementedRoutingInformationServiceServer()
}

func RegisterRoutingInformationServiceServer(s grpc.ServiceRegistrar, srv RoutingInformationServiceServer) {
	s.RegisterService(&RoutingInformationService_ServiceDesc, srv)
}

func _RoutingInformationService_LPM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LPMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingInformationServiceServer).LPM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bio.ris.RoutingInformationService/LPM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingInformationServiceServer).LPM(ctx, req.(*LPMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingInformationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingInformationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bio.ris.RoutingInformationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingInformationServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingInformationService_GetRouters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoutersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingInformationServiceServer).GetRouters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bio.ris.RoutingInformationService/GetRouters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingInformationServiceServer).GetRouters(ctx, req.(*GetRoutersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingInformationService_GetLonger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLongerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutingInformationServiceServer).GetLonger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bio.ris.RoutingInformationService/GetLonger",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutingInformationServiceServer).GetLonger(ctx, req.(*GetLongerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutingInformationService_ObserveRIB_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ObserveRIBRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoutingInformationServiceServer).ObserveRIB(m, &routingInformationServiceObserveRIBServer{stream})
}

type RoutingInformationService_ObserveRIBServer interface {
	Send(*RIBUpdate) error
	grpc.ServerStream
}

type routingInformationServiceObserveRIBServer struct {
	grpc.ServerStream
}

func (x *routingInformationServiceObserveRIBServer) Send(m *RIBUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func _RoutingInformationService_DumpRIB_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DumpRIBRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoutingInformationServiceServer).DumpRIB(m, &routingInformationServiceDumpRIBServer{stream})
}

type RoutingInformationService_DumpRIBServer interface {
	Send(*DumpRIBReply) error
	grpc.ServerStream
}

type routingInformationServiceDumpRIBServer struct {
	grpc.ServerStream
}

func (x *routingInformationServiceDumpRIBServer) Send(m *DumpRIBReply) error {
	return x.ServerStream.SendMsg(m)
}

func _RoutingInformationService_ObserveRouters_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ObserveRoutersRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RoutingInformationServiceServer).ObserveRouters(m, &routingInformationServiceObserveRoutersServer{stream})
}

type RoutingInformationService_ObserveRoutersServer interface {
	Send(*RouterUpdate) error
	grpc.ServerStream
}

type routingInformationServiceObserveRoutersServer struct {
	grpc.ServerStream
}

func (x *routingInformationServiceObserveRoutersServer) Send(m *RouterUpdate) error {
	return x.ServerStream.SendMsg(m)
}

// RoutingInformationService_ServiceDesc is the grpc.ServiceDesc for RoutingInformationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoutingInformationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bio.ris.RoutingInformationService",
	HandlerType: (*RoutingInformationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LPM",
			Handler:    _RoutingInformationService_LPM_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RoutingInformationService_Get_Handler,
		},
		{
			MethodName: "GetRouters",
			Handler:    _RoutingInformationService_GetRouters_Handler,
		},
		{
			MethodName: "GetLonger",
			Handler:    _RoutingInformationService_GetLonger_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ObserveRIB",
			Handler:       _RoutingInformationService_ObserveRIB_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DumpRIB",
			Handler:       _RoutingInformationService_DumpRIB_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ObserveRouters",
			Handler:       _RoutingInformationService_ObserveRouters_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cmd/ris/api/ris.proto",
}
