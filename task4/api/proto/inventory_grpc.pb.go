package inventoryv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	InventoryService_CreateOrder_FullMethodName    = "/inventory.v1.InventoryService/CreateOrder"
	InventoryService_GetOrderStatus_FullMethodName = "/inventory.v1.InventoryService/GetOrderStatus"
)

type InventoryServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	GetOrderStatus(ctx context.Context, in *GetOrderStatusRequest, opts ...grpc.CallOption) (*GetOrderStatusResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, InventoryService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetOrderStatus(ctx context.Context, in *GetOrderStatusRequest, opts ...grpc.CallOption) (*GetOrderStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderStatusResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetOrderStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type InventoryServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	GetOrderStatus(context.Context, *GetOrderStatusRequest) (*GetOrderStatusResponse, error)
	mustEmbedUnimplementedInventoryServiceServer()
}

type UnimplementedInventoryServiceServer struct{}

func (UnimplementedInventoryServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedInventoryServiceServer) GetOrderStatus(context.Context, *GetOrderStatusRequest) (*GetOrderStatusResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetOrderStatus not implemented")
}
func (UnimplementedInventoryServiceServer) mustEmbedUnimplementedInventoryServiceServer() {}
func (UnimplementedInventoryServiceServer) testEmbeddedByValue()                          {}

type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetOrderStatus(ctx, req.(*GetOrderStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.v1.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _InventoryService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderStatus",
			Handler:    _InventoryService_GetOrderStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory.proto",
}
