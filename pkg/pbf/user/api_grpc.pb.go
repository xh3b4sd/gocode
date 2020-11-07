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
const _ = grpc.SupportPackageIsVersion7

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type APIClient interface {
	Create(ctx context.Context, in *CreateI, opts ...grpc.CallOption) (*CreateO, error)
	Delete(ctx context.Context, in *DeleteI, opts ...grpc.CallOption) (*DeleteO, error)
	Search(ctx context.Context, in *SearchI, opts ...grpc.CallOption) (*SearchO, error)
	Update(ctx context.Context, in *UpdateI, opts ...grpc.CallOption) (*UpdateO, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Create(ctx context.Context, in *CreateI, opts ...grpc.CallOption) (*CreateO, error) {
	out := new(CreateO)
	err := c.cc.Invoke(ctx, "/api.API/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Delete(ctx context.Context, in *DeleteI, opts ...grpc.CallOption) (*DeleteO, error) {
	out := new(DeleteO)
	err := c.cc.Invoke(ctx, "/api.API/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Search(ctx context.Context, in *SearchI, opts ...grpc.CallOption) (*SearchO, error) {
	out := new(SearchO)
	err := c.cc.Invoke(ctx, "/api.API/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Update(ctx context.Context, in *UpdateI, opts ...grpc.CallOption) (*UpdateO, error) {
	out := new(UpdateO)
	err := c.cc.Invoke(ctx, "/api.API/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations must embed UnimplementedAPIServer
// for forward compatibility
type APIServer interface {
	Create(context.Context, *CreateI) (*CreateO, error)
	Delete(context.Context, *DeleteI) (*DeleteO, error)
	Search(context.Context, *SearchI) (*SearchO, error)
	Update(context.Context, *UpdateI) (*UpdateO, error)
	mustEmbedUnimplementedAPIServer()
}

// UnimplementedAPIServer must be embedded to have forward compatible implementations.
type UnimplementedAPIServer struct {
}

func (UnimplementedAPIServer) Create(context.Context, *CreateI) (*CreateO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAPIServer) Delete(context.Context, *DeleteI) (*DeleteO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAPIServer) Search(context.Context, *SearchI) (*SearchO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedAPIServer) Update(context.Context, *UpdateI) (*UpdateO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAPIServer) mustEmbedUnimplementedAPIServer() {}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateI)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Create(ctx, req.(*CreateI))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteI)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Delete(ctx, req.(*DeleteI))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchI)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Search(ctx, req.(*SearchI))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateI)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.API/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Update(ctx, req.(*UpdateI))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _API_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _API_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _API_Search_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _API_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
