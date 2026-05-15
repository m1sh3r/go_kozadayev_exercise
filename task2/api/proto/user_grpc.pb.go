package userv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	UserService_CreateUser_FullMethodName        = "/user.v1.UserService/CreateUser"
	UserService_AddDocumentToUser_FullMethodName = "/user.v1.UserService/AddDocumentToUser"
	UserService_GetUserDocuments_FullMethodName  = "/user.v1.UserService/GetUserDocuments"
)

type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	AddDocumentToUser(ctx context.Context, in *AddDocumentToUserRequest, opts ...grpc.CallOption) (*AddDocumentToUserResponse, error)
	GetUserDocuments(ctx context.Context, in *GetUserDocumentsRequest, opts ...grpc.CallOption) (*GetUserDocumentsResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddDocumentToUser(ctx context.Context, in *AddDocumentToUserRequest, opts ...grpc.CallOption) (*AddDocumentToUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddDocumentToUserResponse)
	err := c.cc.Invoke(ctx, UserService_AddDocumentToUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserDocuments(ctx context.Context, in *GetUserDocumentsRequest, opts ...grpc.CallOption) (*GetUserDocumentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserDocumentsResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserDocuments_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	AddDocumentToUser(context.Context, *AddDocumentToUserRequest) (*AddDocumentToUserResponse, error)
	GetUserDocuments(context.Context, *GetUserDocumentsRequest) (*GetUserDocumentsResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) AddDocumentToUser(context.Context, *AddDocumentToUserRequest) (*AddDocumentToUserResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method AddDocumentToUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserDocuments(context.Context, *GetUserDocumentsRequest) (*GetUserDocumentsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetUserDocuments not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddDocumentToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDocumentToUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddDocumentToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddDocumentToUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddDocumentToUser(ctx, req.(*AddDocumentToUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserDocuments_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserDocuments(ctx, req.(*GetUserDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "AddDocumentToUser",
			Handler:    _UserService_AddDocumentToUser_Handler,
		},
		{
			MethodName: "GetUserDocuments",
			Handler:    _UserService_GetUserDocuments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
