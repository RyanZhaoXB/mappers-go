// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: v1/mapperapi.proto

package mapperv1

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

// DeviceManagerServiceClient is the client API for DeviceManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceManagerServiceClient interface {
	// 定义MapperRegister方法，接受MapperRegister消息， 并返回MapperRegister消息
	MapperRegister(ctx context.Context, in *MapperRegisterRequest, opts ...grpc.CallOption) (*MapperRegisterResponse, error)
	ReportDeviceStatus(ctx context.Context, in *ReportDeviceStatusRequest, opts ...grpc.CallOption) (*ReportDeviceStatusResponse, error)
}

type deviceManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceManagerServiceClient(cc grpc.ClientConnInterface) DeviceManagerServiceClient {
	return &deviceManagerServiceClient{cc}
}

func (c *deviceManagerServiceClient) MapperRegister(ctx context.Context, in *MapperRegisterRequest, opts ...grpc.CallOption) (*MapperRegisterResponse, error) {
	out := new(MapperRegisterResponse)
	err := c.cc.Invoke(ctx, "/mapperv1.DeviceManagerService/MapperRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceManagerServiceClient) ReportDeviceStatus(ctx context.Context, in *ReportDeviceStatusRequest, opts ...grpc.CallOption) (*ReportDeviceStatusResponse, error) {
	out := new(ReportDeviceStatusResponse)
	err := c.cc.Invoke(ctx, "/mapperv1.DeviceManagerService/ReportDeviceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceManagerServiceServer is the server API for DeviceManagerService service.
// All implementations must embed UnimplementedDeviceManagerServiceServer
// for forward compatibility
type DeviceManagerServiceServer interface {
	// 定义MapperRegister方法，接受MapperRegister消息， 并返回MapperRegister消息
	MapperRegister(context.Context, *MapperRegisterRequest) (*MapperRegisterResponse, error)
	ReportDeviceStatus(context.Context, *ReportDeviceStatusRequest) (*ReportDeviceStatusResponse, error)
	mustEmbedUnimplementedDeviceManagerServiceServer()
}

// UnimplementedDeviceManagerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceManagerServiceServer struct {
}

func (UnimplementedDeviceManagerServiceServer) MapperRegister(context.Context, *MapperRegisterRequest) (*MapperRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MapperRegister not implemented")
}
func (UnimplementedDeviceManagerServiceServer) ReportDeviceStatus(context.Context, *ReportDeviceStatusRequest) (*ReportDeviceStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportDeviceStatus not implemented")
}
func (UnimplementedDeviceManagerServiceServer) mustEmbedUnimplementedDeviceManagerServiceServer() {}

// UnsafeDeviceManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceManagerServiceServer will
// result in compilation errors.
type UnsafeDeviceManagerServiceServer interface {
	mustEmbedUnimplementedDeviceManagerServiceServer()
}

func RegisterDeviceManagerServiceServer(s grpc.ServiceRegistrar, srv DeviceManagerServiceServer) {
	s.RegisterService(&DeviceManagerService_ServiceDesc, srv)
}

func _DeviceManagerService_MapperRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapperRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManagerServiceServer).MapperRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapperv1.DeviceManagerService/MapperRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManagerServiceServer).MapperRegister(ctx, req.(*MapperRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceManagerService_ReportDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportDeviceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceManagerServiceServer).ReportDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mapperv1.DeviceManagerService/ReportDeviceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceManagerServiceServer).ReportDeviceStatus(ctx, req.(*ReportDeviceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceManagerService_ServiceDesc is the grpc.ServiceDesc for DeviceManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mapperv1.DeviceManagerService",
	HandlerType: (*DeviceManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MapperRegister",
			Handler:    _DeviceManagerService_MapperRegister_Handler,
		},
		{
			MethodName: "ReportDeviceStatus",
			Handler:    _DeviceManagerService_ReportDeviceStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/mapperapi.proto",
}