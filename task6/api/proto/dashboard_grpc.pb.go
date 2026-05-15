package dashboardv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	DashboardService_GetDashboard_FullMethodName        = "/dashboard.v1.DashboardService/GetDashboard"
	DashboardService_CreateOwner_FullMethodName         = "/dashboard.v1.DashboardService/CreateOwner"
	DashboardService_CreateCar_FullMethodName           = "/dashboard.v1.DashboardService/CreateCar"
	DashboardService_CreateServiceRecord_FullMethodName = "/dashboard.v1.DashboardService/CreateServiceRecord"
)

type DashboardServiceClient interface {
	GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...grpc.CallOption) (*GetDashboardResponse, error)

	CreateOwner(ctx context.Context, in *CreateOwnerRequest, opts ...grpc.CallOption) (*Owner, error)
	CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*Car, error)
	CreateServiceRecord(ctx context.Context, in *CreateServiceRecordRequest, opts ...grpc.CallOption) (*ServiceRecord, error)
}

type dashboardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDashboardServiceClient(cc grpc.ClientConnInterface) DashboardServiceClient {
	return &dashboardServiceClient{cc}
}

func (c *dashboardServiceClient) GetDashboard(ctx context.Context, in *GetDashboardRequest, opts ...grpc.CallOption) (*GetDashboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDashboardResponse)
	err := c.cc.Invoke(ctx, DashboardService_GetDashboard_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) CreateOwner(ctx context.Context, in *CreateOwnerRequest, opts ...grpc.CallOption) (*Owner, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Owner)
	err := c.cc.Invoke(ctx, DashboardService_CreateOwner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*Car, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Car)
	err := c.cc.Invoke(ctx, DashboardService_CreateCar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardServiceClient) CreateServiceRecord(ctx context.Context, in *CreateServiceRecordRequest, opts ...grpc.CallOption) (*ServiceRecord, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServiceRecord)
	err := c.cc.Invoke(ctx, DashboardService_CreateServiceRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type DashboardServiceServer interface {
	GetDashboard(context.Context, *GetDashboardRequest) (*GetDashboardResponse, error)

	CreateOwner(context.Context, *CreateOwnerRequest) (*Owner, error)
	CreateCar(context.Context, *CreateCarRequest) (*Car, error)
	CreateServiceRecord(context.Context, *CreateServiceRecordRequest) (*ServiceRecord, error)
	mustEmbedUnimplementedDashboardServiceServer()
}

type UnimplementedDashboardServiceServer struct{}

func (UnimplementedDashboardServiceServer) GetDashboard(context.Context, *GetDashboardRequest) (*GetDashboardResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetDashboard not implemented")
}
func (UnimplementedDashboardServiceServer) CreateOwner(context.Context, *CreateOwnerRequest) (*Owner, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateOwner not implemented")
}
func (UnimplementedDashboardServiceServer) CreateCar(context.Context, *CreateCarRequest) (*Car, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateCar not implemented")
}
func (UnimplementedDashboardServiceServer) CreateServiceRecord(context.Context, *CreateServiceRecordRequest) (*ServiceRecord, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateServiceRecord not implemented")
}
func (UnimplementedDashboardServiceServer) mustEmbedUnimplementedDashboardServiceServer() {}
func (UnimplementedDashboardServiceServer) testEmbeddedByValue()                          {}

type UnsafeDashboardServiceServer interface {
	mustEmbedUnimplementedDashboardServiceServer()
}

func RegisterDashboardServiceServer(s grpc.ServiceRegistrar, srv DashboardServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DashboardService_ServiceDesc, srv)
}

func _DashboardService_GetDashboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDashboardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).GetDashboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardService_GetDashboard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).GetDashboard(ctx, req.(*GetDashboardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_CreateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).CreateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardService_CreateOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).CreateOwner(ctx, req.(*CreateOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_CreateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).CreateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardService_CreateCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).CreateCar(ctx, req.(*CreateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DashboardService_CreateServiceRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServiceRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServiceServer).CreateServiceRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DashboardService_CreateServiceRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServiceServer).CreateServiceRecord(ctx, req.(*CreateServiceRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var DashboardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dashboard.v1.DashboardService",
	HandlerType: (*DashboardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDashboard",
			Handler:    _DashboardService_GetDashboard_Handler,
		},
		{
			MethodName: "CreateOwner",
			Handler:    _DashboardService_CreateOwner_Handler,
		},
		{
			MethodName: "CreateCar",
			Handler:    _DashboardService_CreateCar_Handler,
		},
		{
			MethodName: "CreateServiceRecord",
			Handler:    _DashboardService_CreateServiceRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dashboard.proto",
}
