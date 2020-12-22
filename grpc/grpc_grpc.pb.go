// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpcpb

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CRUDClient is the client API for CRUD service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CRUDClient interface {
	Create(ctx context.Context, in *Structmsg, opts ...grpc.CallOption) (*Errmsg, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Errmsg, error)
	Update(ctx context.Context, in *Structmsg, opts ...grpc.CallOption) (*Errmsg, error)
	Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Structmsg, error)
	CreateUser(ctx context.Context, in *Userstruct, opts ...grpc.CallOption) (*Errmsg, error)
	DeleteUser(ctx context.Context, in *Userstruct, opts ...grpc.CallOption) (*Errmsg, error)
	UpdateUser(ctx context.Context, in *Userstruct, opts ...grpc.CallOption) (*Errmsg, error)
	Login(ctx context.Context, in *Userstruct, opts ...grpc.CallOption) (*Jwtoken, error)
}

type cRUDClient struct {
	cc grpc.ClientConnInterface
}

func NewCRUDClient(cc grpc.ClientConnInterface) CRUDClient {
	return &cRUDClient{cc}
}

func (c *cRUDClient) Create(ctx context.Context, in *Structmsg, opts ...grpc.CallOption) (*Errmsg, error) {
	out := new(Errmsg)
	err := c.cc.Invoke(ctx, "/grpcpb.CRUD/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Errmsg, error) {
	out := new(Errmsg)
	err := c.cc.Invoke(ctx, "/grpcpb.CRUD/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Update(ctx context.Context, in *Structmsg, opts ...grpc.CallOption) (*Errmsg, error) {
	out := new(Errmsg)
	err := c.cc.Invoke(ctx, "/grpcpb.CRUD/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Structmsg, error) {
	out := new(Structmsg)
	err := c.cc.Invoke(ctx, "/grpcpb.CRUD/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) CreateUser(ctx context.Context, in *Userstruct, opts ...grpc.CallOption) (*Errmsg, error) {
	out := new(Errmsg)
	err := c.cc.Invoke(ctx, "/grpcpb.CRUD/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) DeleteUser(ctx context.Context, in *Userstruct, opts ...grpc.CallOption) (*Errmsg, error) {
	out := new(Errmsg)
	err := c.cc.Invoke(ctx, "/grpcpb.CRUD/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) UpdateUser(ctx context.Context, in *Userstruct, opts ...grpc.CallOption) (*Errmsg, error) {
	out := new(Errmsg)
	err := c.cc.Invoke(ctx, "/grpcpb.CRUD/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cRUDClient) Login(ctx context.Context, in *Userstruct, opts ...grpc.CallOption) (*Jwtoken, error) {
	out := new(Jwtoken)
	err := c.cc.Invoke(ctx, "/grpcpb.CRUD/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CRUDServer is the server API for CRUD service.
// All implementations must embed UnimplementedCRUDServer
// for forward compatibility
type CRUDServer interface {
	Create(context.Context, *Structmsg) (*Errmsg, error)
	Delete(context.Context, *Id) (*Errmsg, error)
	Update(context.Context, *Structmsg) (*Errmsg, error)
	Read(context.Context, *Id) (*Structmsg, error)
	CreateUser(context.Context, *Userstruct) (*Errmsg, error)
	DeleteUser(context.Context, *Userstruct) (*Errmsg, error)
	UpdateUser(context.Context, *Userstruct) (*Errmsg, error)
	Login(context.Context, *Userstruct) (*Jwtoken, error)
}

// UnimplementedCRUDServer must be embedded to have forward compatible implementations.
type UnimplementedCRUDServer struct {
}

func (UnimplementedCRUDServer) Create(context.Context, *Structmsg) (*Errmsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (UnimplementedCRUDServer) Delete(context.Context, *Id) (*Errmsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func (UnimplementedCRUDServer) Update(context.Context, *Structmsg) (*Errmsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func (UnimplementedCRUDServer) Read(context.Context, *Id) (*Structmsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func (UnimplementedCRUDServer) CreateUser(context.Context, *Userstruct) (*Errmsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

func (UnimplementedCRUDServer) DeleteUser(context.Context, *Userstruct) (*Errmsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}

func (UnimplementedCRUDServer) UpdateUser(context.Context, *Userstruct) (*Errmsg, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func (UnimplementedCRUDServer) Login(context.Context, *Userstruct) (*Jwtoken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCRUDServer) mustEmbedUnimplementedCRUDServer() {}

// UnsafeCRUDServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CRUDServer will
// result in compilation errors.
type UnsafeCRUDServer interface {
	mustEmbedUnimplementedCRUDServer()
}

func RegisterCRUDServer(s grpc.ServiceRegistrar, srv CRUDServer) {
	s.RegisterService(&_CRUD_serviceDesc, srv)
}

func _CRUD_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Structmsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpb.CRUD/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Create(ctx, req.(*Structmsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpb.CRUD/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Structmsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpb.CRUD/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Update(ctx, req.(*Structmsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpb.CRUD/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Read(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Userstruct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpb.CRUD/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).CreateUser(ctx, req.(*Userstruct))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Userstruct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpb.CRUD/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).DeleteUser(ctx, req.(*Userstruct))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Userstruct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpb.CRUD/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).UpdateUser(ctx, req.(*Userstruct))
	}
	return interceptor(ctx, in, info, handler)
}

func _CRUD_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Userstruct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CRUDServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpb.CRUD/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CRUDServer).Login(ctx, req.(*Userstruct))
	}
	return interceptor(ctx, in, info, handler)
}

var _CRUD_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcpb.CRUD",
	HandlerType: (*CRUDServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CRUD_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CRUD_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CRUD_Update_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _CRUD_Read_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _CRUD_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _CRUD_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _CRUD_UpdateUser_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _CRUD_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/grpc.proto",
}