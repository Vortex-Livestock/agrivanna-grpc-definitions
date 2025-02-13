// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: proto/livestock/v1/weight_record.proto

package generated

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	WeightRecordService_CreateWeightRecord_FullMethodName = "/proto.livestock.v1.WeightRecordService/CreateWeightRecord"
	WeightRecordService_GetWeightRecord_FullMethodName    = "/proto.livestock.v1.WeightRecordService/GetWeightRecord"
	WeightRecordService_UpdateWeightRecord_FullMethodName = "/proto.livestock.v1.WeightRecordService/UpdateWeightRecord"
	WeightRecordService_DeleteWeightRecord_FullMethodName = "/proto.livestock.v1.WeightRecordService/DeleteWeightRecord"
)

// WeightRecordServiceClient is the client API for WeightRecordService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeightRecordServiceClient interface {
	CreateWeightRecord(ctx context.Context, in *CreateWeightRecordRequest, opts ...grpc.CallOption) (*CreateWeightRecordResponse, error)
	GetWeightRecord(ctx context.Context, in *GetWeightRecordRequest, opts ...grpc.CallOption) (*GetWeightRecordResponse, error)
	UpdateWeightRecord(ctx context.Context, in *UpdateWeightRecordRequest, opts ...grpc.CallOption) (*UpdateWeightRecordResponse, error)
	DeleteWeightRecord(ctx context.Context, in *DeleteWeightRecordRequest, opts ...grpc.CallOption) (*DeleteWeightRecordResponse, error)
}

type weightRecordServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeightRecordServiceClient(cc grpc.ClientConnInterface) WeightRecordServiceClient {
	return &weightRecordServiceClient{cc}
}

func (c *weightRecordServiceClient) CreateWeightRecord(ctx context.Context, in *CreateWeightRecordRequest, opts ...grpc.CallOption) (*CreateWeightRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateWeightRecordResponse)
	err := c.cc.Invoke(ctx, WeightRecordService_CreateWeightRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weightRecordServiceClient) GetWeightRecord(ctx context.Context, in *GetWeightRecordRequest, opts ...grpc.CallOption) (*GetWeightRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetWeightRecordResponse)
	err := c.cc.Invoke(ctx, WeightRecordService_GetWeightRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weightRecordServiceClient) UpdateWeightRecord(ctx context.Context, in *UpdateWeightRecordRequest, opts ...grpc.CallOption) (*UpdateWeightRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateWeightRecordResponse)
	err := c.cc.Invoke(ctx, WeightRecordService_UpdateWeightRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weightRecordServiceClient) DeleteWeightRecord(ctx context.Context, in *DeleteWeightRecordRequest, opts ...grpc.CallOption) (*DeleteWeightRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteWeightRecordResponse)
	err := c.cc.Invoke(ctx, WeightRecordService_DeleteWeightRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WeightRecordServiceServer is the server API for WeightRecordService service.
// All implementations must embed UnimplementedWeightRecordServiceServer
// for forward compatibility.
type WeightRecordServiceServer interface {
	CreateWeightRecord(context.Context, *CreateWeightRecordRequest) (*CreateWeightRecordResponse, error)
	GetWeightRecord(context.Context, *GetWeightRecordRequest) (*GetWeightRecordResponse, error)
	UpdateWeightRecord(context.Context, *UpdateWeightRecordRequest) (*UpdateWeightRecordResponse, error)
	DeleteWeightRecord(context.Context, *DeleteWeightRecordRequest) (*DeleteWeightRecordResponse, error)
	mustEmbedUnimplementedWeightRecordServiceServer()
}

// UnimplementedWeightRecordServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWeightRecordServiceServer struct{}

func (UnimplementedWeightRecordServiceServer) CreateWeightRecord(context.Context, *CreateWeightRecordRequest) (*CreateWeightRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWeightRecord not implemented")
}
func (UnimplementedWeightRecordServiceServer) GetWeightRecord(context.Context, *GetWeightRecordRequest) (*GetWeightRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeightRecord not implemented")
}
func (UnimplementedWeightRecordServiceServer) UpdateWeightRecord(context.Context, *UpdateWeightRecordRequest) (*UpdateWeightRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWeightRecord not implemented")
}
func (UnimplementedWeightRecordServiceServer) DeleteWeightRecord(context.Context, *DeleteWeightRecordRequest) (*DeleteWeightRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteWeightRecord not implemented")
}
func (UnimplementedWeightRecordServiceServer) mustEmbedUnimplementedWeightRecordServiceServer() {}
func (UnimplementedWeightRecordServiceServer) testEmbeddedByValue()                             {}

// UnsafeWeightRecordServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeightRecordServiceServer will
// result in compilation errors.
type UnsafeWeightRecordServiceServer interface {
	mustEmbedUnimplementedWeightRecordServiceServer()
}

func RegisterWeightRecordServiceServer(s grpc.ServiceRegistrar, srv WeightRecordServiceServer) {
	// If the following call pancis, it indicates UnimplementedWeightRecordServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WeightRecordService_ServiceDesc, srv)
}

func _WeightRecordService_CreateWeightRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWeightRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeightRecordServiceServer).CreateWeightRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeightRecordService_CreateWeightRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeightRecordServiceServer).CreateWeightRecord(ctx, req.(*CreateWeightRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeightRecordService_GetWeightRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWeightRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeightRecordServiceServer).GetWeightRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeightRecordService_GetWeightRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeightRecordServiceServer).GetWeightRecord(ctx, req.(*GetWeightRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeightRecordService_UpdateWeightRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWeightRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeightRecordServiceServer).UpdateWeightRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeightRecordService_UpdateWeightRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeightRecordServiceServer).UpdateWeightRecord(ctx, req.(*UpdateWeightRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WeightRecordService_DeleteWeightRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteWeightRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WeightRecordServiceServer).DeleteWeightRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WeightRecordService_DeleteWeightRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WeightRecordServiceServer).DeleteWeightRecord(ctx, req.(*DeleteWeightRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WeightRecordService_ServiceDesc is the grpc.ServiceDesc for WeightRecordService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeightRecordService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.livestock.v1.WeightRecordService",
	HandlerType: (*WeightRecordServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWeightRecord",
			Handler:    _WeightRecordService_CreateWeightRecord_Handler,
		},
		{
			MethodName: "GetWeightRecord",
			Handler:    _WeightRecordService_GetWeightRecord_Handler,
		},
		{
			MethodName: "UpdateWeightRecord",
			Handler:    _WeightRecordService_UpdateWeightRecord_Handler,
		},
		{
			MethodName: "DeleteWeightRecord",
			Handler:    _WeightRecordService_DeleteWeightRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/livestock/v1/weight_record.proto",
}
