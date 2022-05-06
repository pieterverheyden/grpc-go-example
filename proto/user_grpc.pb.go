// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: proto/user.proto

package grpc_go_example

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserManagementClient is the client API for UserManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagementClient interface {
	CreateOrUpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserByEmail(ctx context.Context, in *UserEmail, opts ...grpc.CallOption) (*User, error)
	GetAllUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (UserManagement_GetAllUsersClient, error)
	RemoveUserByEmail(ctx context.Context, in *UserEmail, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementClient(cc grpc.ClientConnInterface) UserManagementClient {
	return &userManagementClient{cc}
}

func (c *userManagementClient) CreateOrUpdateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserManagement/CreateOrUpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) GetUserByEmail(ctx context.Context, in *UserEmail, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserManagement/GetUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) GetAllUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (UserManagement_GetAllUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &UserManagement_ServiceDesc.Streams[0], "/user.UserManagement/GetAllUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &userManagementGetAllUsersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type UserManagement_GetAllUsersClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type userManagementGetAllUsersClient struct {
	grpc.ClientStream
}

func (x *userManagementGetAllUsersClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userManagementClient) RemoveUserByEmail(ctx context.Context, in *UserEmail, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.UserManagement/RemoveUserByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServer is the server API for UserManagement service.
// All implementations must embed UnimplementedUserManagementServer
// for forward compatibility
type UserManagementServer interface {
	CreateOrUpdateUser(context.Context, *User) (*emptypb.Empty, error)
	GetUserByEmail(context.Context, *UserEmail) (*User, error)
	GetAllUsers(*emptypb.Empty, UserManagement_GetAllUsersServer) error
	RemoveUserByEmail(context.Context, *UserEmail) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserManagementServer()
}

// UnimplementedUserManagementServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagementServer struct {
}

func (UnimplementedUserManagementServer) CreateOrUpdateUser(context.Context, *User) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateUser not implemented")
}
func (UnimplementedUserManagementServer) GetUserByEmail(context.Context, *UserEmail) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByEmail not implemented")
}
func (UnimplementedUserManagementServer) GetAllUsers(*emptypb.Empty, UserManagement_GetAllUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllUsers not implemented")
}
func (UnimplementedUserManagementServer) RemoveUserByEmail(context.Context, *UserEmail) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserByEmail not implemented")
}
func (UnimplementedUserManagementServer) mustEmbedUnimplementedUserManagementServer() {}

// UnsafeUserManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServer will
// result in compilation errors.
type UnsafeUserManagementServer interface {
	mustEmbedUnimplementedUserManagementServer()
}

func RegisterUserManagementServer(s grpc.ServiceRegistrar, srv UserManagementServer) {
	s.RegisterService(&UserManagement_ServiceDesc, srv)
}

func _UserManagement_CreateOrUpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).CreateOrUpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserManagement/CreateOrUpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).CreateOrUpdateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_GetUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserManagement/GetUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetUserByEmail(ctx, req.(*UserEmail))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_GetAllUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(UserManagementServer).GetAllUsers(m, &userManagementGetAllUsersServer{stream})
}

type UserManagement_GetAllUsersServer interface {
	Send(*User) error
	grpc.ServerStream
}

type userManagementGetAllUsersServer struct {
	grpc.ServerStream
}

func (x *userManagementGetAllUsersServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _UserManagement_RemoveUserByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserEmail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).RemoveUserByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserManagement/RemoveUserByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).RemoveUserByEmail(ctx, req.(*UserEmail))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagement_ServiceDesc is the grpc.ServiceDesc for UserManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserManagement",
	HandlerType: (*UserManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdateUser",
			Handler:    _UserManagement_CreateOrUpdateUser_Handler,
		},
		{
			MethodName: "GetUserByEmail",
			Handler:    _UserManagement_GetUserByEmail_Handler,
		},
		{
			MethodName: "RemoveUserByEmail",
			Handler:    _UserManagement_RemoveUserByEmail_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllUsers",
			Handler:       _UserManagement_GetAllUsers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/user.proto",
}
