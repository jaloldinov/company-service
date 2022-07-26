// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: attribute.proto

package position_service

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

// AttributeServiceClient is the client API for AttributeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttributeServiceClient interface {
	Create(ctx context.Context, in *CreateAttributeRequest, opts ...grpc.CallOption) (*Attribute, error)
	GetAll(ctx context.Context, in *GetAllAttributeRequest, opts ...grpc.CallOption) (*GetAllAttributeResponse, error)
	Get(ctx context.Context, in *AttributeId, opts ...grpc.CallOption) (*Attribute, error)
	Update(ctx context.Context, in *Attribute, opts ...grpc.CallOption) (*AttributeResult, error)
	Delete(ctx context.Context, in *AttributeId, opts ...grpc.CallOption) (*AttributeResult, error)
}

type attributeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAttributeServiceClient(cc grpc.ClientConnInterface) AttributeServiceClient {
	return &attributeServiceClient{cc}
}

func (c *attributeServiceClient) Create(ctx context.Context, in *CreateAttributeRequest, opts ...grpc.CallOption) (*Attribute, error) {
	out := new(Attribute)
	err := c.cc.Invoke(ctx, "/genproto.AttributeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeServiceClient) GetAll(ctx context.Context, in *GetAllAttributeRequest, opts ...grpc.CallOption) (*GetAllAttributeResponse, error) {
	out := new(GetAllAttributeResponse)
	err := c.cc.Invoke(ctx, "/genproto.AttributeService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeServiceClient) Get(ctx context.Context, in *AttributeId, opts ...grpc.CallOption) (*Attribute, error) {
	out := new(Attribute)
	err := c.cc.Invoke(ctx, "/genproto.AttributeService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeServiceClient) Update(ctx context.Context, in *Attribute, opts ...grpc.CallOption) (*AttributeResult, error) {
	out := new(AttributeResult)
	err := c.cc.Invoke(ctx, "/genproto.AttributeService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *attributeServiceClient) Delete(ctx context.Context, in *AttributeId, opts ...grpc.CallOption) (*AttributeResult, error) {
	out := new(AttributeResult)
	err := c.cc.Invoke(ctx, "/genproto.AttributeService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttributeServiceServer is the server API for AttributeService service.
// All implementations must embed UnimplementedAttributeServiceServer
// for forward compatibility
type AttributeServiceServer interface {
	Create(context.Context, *CreateAttributeRequest) (*Attribute, error)
	GetAll(context.Context, *GetAllAttributeRequest) (*GetAllAttributeResponse, error)
	Get(context.Context, *AttributeId) (*Attribute, error)
	Update(context.Context, *Attribute) (*AttributeResult, error)
	Delete(context.Context, *AttributeId) (*AttributeResult, error)
	mustEmbedUnimplementedAttributeServiceServer()
}

// UnimplementedAttributeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAttributeServiceServer struct {
}

func (UnimplementedAttributeServiceServer) Create(context.Context, *CreateAttributeRequest) (*Attribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAttributeServiceServer) GetAll(context.Context, *GetAllAttributeRequest) (*GetAllAttributeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAttributeServiceServer) Get(context.Context, *AttributeId) (*Attribute, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAttributeServiceServer) Update(context.Context, *Attribute) (*AttributeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAttributeServiceServer) Delete(context.Context, *AttributeId) (*AttributeResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAttributeServiceServer) mustEmbedUnimplementedAttributeServiceServer() {}

// UnsafeAttributeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttributeServiceServer will
// result in compilation errors.
type UnsafeAttributeServiceServer interface {
	mustEmbedUnimplementedAttributeServiceServer()
}

func RegisterAttributeServiceServer(s grpc.ServiceRegistrar, srv AttributeServiceServer) {
	s.RegisterService(&AttributeService_ServiceDesc, srv)
}

func _AttributeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).Create(ctx, req.(*CreateAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributeService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAttributeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributeService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).GetAll(ctx, req.(*GetAllAttributeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttributeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributeService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).Get(ctx, req.(*AttributeId))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributeService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Attribute)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributeService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).Update(ctx, req.(*Attribute))
	}
	return interceptor(ctx, in, info, handler)
}

func _AttributeService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttributeId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttributeServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/genproto.AttributeService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttributeServiceServer).Delete(ctx, req.(*AttributeId))
	}
	return interceptor(ctx, in, info, handler)
}

// AttributeService_ServiceDesc is the grpc.ServiceDesc for AttributeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AttributeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "genproto.AttributeService",
	HandlerType: (*AttributeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AttributeService_Create_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AttributeService_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AttributeService_Get_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AttributeService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AttributeService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attribute.proto",
}
