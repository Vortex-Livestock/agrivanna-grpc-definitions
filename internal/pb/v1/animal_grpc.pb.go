// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: v1/animal.proto

package v1

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
	AnimalService_GetAnimalById_FullMethodName = "/proto.v1.AnimalService/GetAnimalById"
	AnimalService_GetAnimals_FullMethodName    = "/proto.v1.AnimalService/GetAnimals"
	AnimalService_CreateAnimal_FullMethodName  = "/proto.v1.AnimalService/CreateAnimal"
	AnimalService_UpdateAnimal_FullMethodName  = "/proto.v1.AnimalService/UpdateAnimal"
	AnimalService_DeleteAnimal_FullMethodName  = "/proto.v1.AnimalService/DeleteAnimal"
)

// AnimalServiceClient is the client API for AnimalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnimalServiceClient interface {
	GetAnimalById(ctx context.Context, in *GetAnimalByIdRequest, opts ...grpc.CallOption) (*GetAnimalByIdResponse, error)
	GetAnimals(ctx context.Context, in *GetAnimalsRequest, opts ...grpc.CallOption) (*GetAnimalsResponse, error)
	CreateAnimal(ctx context.Context, in *CreateAnimalRequest, opts ...grpc.CallOption) (*CreateAnimalResponse, error)
	UpdateAnimal(ctx context.Context, in *UpdateAnimalRequest, opts ...grpc.CallOption) (*UpdateAnimalResponse, error)
	DeleteAnimal(ctx context.Context, in *DeleteAnimalRequest, opts ...grpc.CallOption) (*DeleteAnimalResponse, error)
}

type animalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnimalServiceClient(cc grpc.ClientConnInterface) AnimalServiceClient {
	return &animalServiceClient{cc}
}

func (c *animalServiceClient) GetAnimalById(ctx context.Context, in *GetAnimalByIdRequest, opts ...grpc.CallOption) (*GetAnimalByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAnimalByIdResponse)
	err := c.cc.Invoke(ctx, AnimalService_GetAnimalById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalServiceClient) GetAnimals(ctx context.Context, in *GetAnimalsRequest, opts ...grpc.CallOption) (*GetAnimalsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAnimalsResponse)
	err := c.cc.Invoke(ctx, AnimalService_GetAnimals_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalServiceClient) CreateAnimal(ctx context.Context, in *CreateAnimalRequest, opts ...grpc.CallOption) (*CreateAnimalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAnimalResponse)
	err := c.cc.Invoke(ctx, AnimalService_CreateAnimal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalServiceClient) UpdateAnimal(ctx context.Context, in *UpdateAnimalRequest, opts ...grpc.CallOption) (*UpdateAnimalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAnimalResponse)
	err := c.cc.Invoke(ctx, AnimalService_UpdateAnimal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *animalServiceClient) DeleteAnimal(ctx context.Context, in *DeleteAnimalRequest, opts ...grpc.CallOption) (*DeleteAnimalResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAnimalResponse)
	err := c.cc.Invoke(ctx, AnimalService_DeleteAnimal_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnimalServiceServer is the server API for AnimalService service.
// All implementations must embed UnimplementedAnimalServiceServer
// for forward compatibility.
type AnimalServiceServer interface {
	GetAnimalById(context.Context, *GetAnimalByIdRequest) (*GetAnimalByIdResponse, error)
	GetAnimals(context.Context, *GetAnimalsRequest) (*GetAnimalsResponse, error)
	CreateAnimal(context.Context, *CreateAnimalRequest) (*CreateAnimalResponse, error)
	UpdateAnimal(context.Context, *UpdateAnimalRequest) (*UpdateAnimalResponse, error)
	DeleteAnimal(context.Context, *DeleteAnimalRequest) (*DeleteAnimalResponse, error)
	mustEmbedUnimplementedAnimalServiceServer()
}

// UnimplementedAnimalServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAnimalServiceServer struct{}

func (UnimplementedAnimalServiceServer) GetAnimalById(context.Context, *GetAnimalByIdRequest) (*GetAnimalByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimalById not implemented")
}
func (UnimplementedAnimalServiceServer) GetAnimals(context.Context, *GetAnimalsRequest) (*GetAnimalsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnimals not implemented")
}
func (UnimplementedAnimalServiceServer) CreateAnimal(context.Context, *CreateAnimalRequest) (*CreateAnimalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAnimal not implemented")
}
func (UnimplementedAnimalServiceServer) UpdateAnimal(context.Context, *UpdateAnimalRequest) (*UpdateAnimalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAnimal not implemented")
}
func (UnimplementedAnimalServiceServer) DeleteAnimal(context.Context, *DeleteAnimalRequest) (*DeleteAnimalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnimal not implemented")
}
func (UnimplementedAnimalServiceServer) mustEmbedUnimplementedAnimalServiceServer() {}
func (UnimplementedAnimalServiceServer) testEmbeddedByValue()                       {}

// UnsafeAnimalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnimalServiceServer will
// result in compilation errors.
type UnsafeAnimalServiceServer interface {
	mustEmbedUnimplementedAnimalServiceServer()
}

func RegisterAnimalServiceServer(s grpc.ServiceRegistrar, srv AnimalServiceServer) {
	// If the following call pancis, it indicates UnimplementedAnimalServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AnimalService_ServiceDesc, srv)
}

func _AnimalService_GetAnimalById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnimalByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServiceServer).GetAnimalById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnimalService_GetAnimalById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServiceServer).GetAnimalById(ctx, req.(*GetAnimalByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnimalService_GetAnimals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnimalsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServiceServer).GetAnimals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnimalService_GetAnimals_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServiceServer).GetAnimals(ctx, req.(*GetAnimalsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnimalService_CreateAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServiceServer).CreateAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnimalService_CreateAnimal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServiceServer).CreateAnimal(ctx, req.(*CreateAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnimalService_UpdateAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServiceServer).UpdateAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnimalService_UpdateAnimal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServiceServer).UpdateAnimal(ctx, req.(*UpdateAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AnimalService_DeleteAnimal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAnimalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnimalServiceServer).DeleteAnimal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AnimalService_DeleteAnimal_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnimalServiceServer).DeleteAnimal(ctx, req.(*DeleteAnimalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AnimalService_ServiceDesc is the grpc.ServiceDesc for AnimalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AnimalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.v1.AnimalService",
	HandlerType: (*AnimalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAnimalById",
			Handler:    _AnimalService_GetAnimalById_Handler,
		},
		{
			MethodName: "GetAnimals",
			Handler:    _AnimalService_GetAnimals_Handler,
		},
		{
			MethodName: "CreateAnimal",
			Handler:    _AnimalService_CreateAnimal_Handler,
		},
		{
			MethodName: "UpdateAnimal",
			Handler:    _AnimalService_UpdateAnimal_Handler,
		},
		{
			MethodName: "DeleteAnimal",
			Handler:    _AnimalService_DeleteAnimal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/animal.proto",
}
