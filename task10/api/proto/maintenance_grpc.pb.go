package maintenancev1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	MaintenanceService_CreateWorkOrder_FullMethodName = "/maintenance.v1.MaintenanceService/CreateWorkOrder"
	MaintenanceService_GetWorkOrder_FullMethodName    = "/maintenance.v1.MaintenanceService/GetWorkOrder"
	MaintenanceService_ExportToExcel_FullMethodName   = "/maintenance.v1.MaintenanceService/ExportToExcel"
)

type MaintenanceServiceClient interface {
	CreateWorkOrder(ctx context.Context, in *CreateWorkOrderRequest, opts ...grpc.CallOption) (*WorkOrder, error)
	GetWorkOrder(ctx context.Context, in *GetWorkOrderRequest, opts ...grpc.CallOption) (*WorkOrder, error)
	ExportToExcel(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error)
}

type maintenanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaintenanceServiceClient(cc grpc.ClientConnInterface) MaintenanceServiceClient {
	return &maintenanceServiceClient{cc}
}

func (c *maintenanceServiceClient) CreateWorkOrder(ctx context.Context, in *CreateWorkOrderRequest, opts ...grpc.CallOption) (*WorkOrder, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WorkOrder)
	err := c.cc.Invoke(ctx, MaintenanceService_CreateWorkOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintenanceServiceClient) GetWorkOrder(ctx context.Context, in *GetWorkOrderRequest, opts ...grpc.CallOption) (*WorkOrder, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WorkOrder)
	err := c.cc.Invoke(ctx, MaintenanceService_GetWorkOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *maintenanceServiceClient) ExportToExcel(ctx context.Context, in *ExportRequest, opts ...grpc.CallOption) (*ExportResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExportResponse)
	err := c.cc.Invoke(ctx, MaintenanceService_ExportToExcel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type MaintenanceServiceServer interface {
	CreateWorkOrder(context.Context, *CreateWorkOrderRequest) (*WorkOrder, error)
	GetWorkOrder(context.Context, *GetWorkOrderRequest) (*WorkOrder, error)
	ExportToExcel(context.Context, *ExportRequest) (*ExportResponse, error)
	mustEmbedUnimplementedMaintenanceServiceServer()
}

type UnimplementedMaintenanceServiceServer struct{}

func (UnimplementedMaintenanceServiceServer) CreateWorkOrder(context.Context, *CreateWorkOrderRequest) (*WorkOrder, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateWorkOrder not implemented")
}
func (UnimplementedMaintenanceServiceServer) GetWorkOrder(context.Context, *GetWorkOrderRequest) (*WorkOrder, error) {
	return nil, status.Error(codes.Unimplemented, "method GetWorkOrder not implemented")
}
func (UnimplementedMaintenanceServiceServer) ExportToExcel(context.Context, *ExportRequest) (*ExportResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method ExportToExcel not implemented")
}
func (UnimplementedMaintenanceServiceServer) mustEmbedUnimplementedMaintenanceServiceServer() {}
func (UnimplementedMaintenanceServiceServer) testEmbeddedByValue()                            {}

type UnsafeMaintenanceServiceServer interface {
	mustEmbedUnimplementedMaintenanceServiceServer()
}

func RegisterMaintenanceServiceServer(s grpc.ServiceRegistrar, srv MaintenanceServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MaintenanceService_ServiceDesc, srv)
}

func _MaintenanceService_CreateWorkOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorkOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServiceServer).CreateWorkOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaintenanceService_CreateWorkOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServiceServer).CreateWorkOrder(ctx, req.(*CreateWorkOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintenanceService_GetWorkOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServiceServer).GetWorkOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaintenanceService_GetWorkOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServiceServer).GetWorkOrder(ctx, req.(*GetWorkOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MaintenanceService_ExportToExcel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaintenanceServiceServer).ExportToExcel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MaintenanceService_ExportToExcel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaintenanceServiceServer).ExportToExcel(ctx, req.(*ExportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var MaintenanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "maintenance.v1.MaintenanceService",
	HandlerType: (*MaintenanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorkOrder",
			Handler:    _MaintenanceService_CreateWorkOrder_Handler,
		},
		{
			MethodName: "GetWorkOrder",
			Handler:    _MaintenanceService_GetWorkOrder_Handler,
		},
		{
			MethodName: "ExportToExcel",
			Handler:    _MaintenanceService_ExportToExcel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "maintenance.proto",
}
