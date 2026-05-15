package rentalv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	RentalService_CreateBooking_FullMethodName     = "/car_rental.v1.RentalService/CreateBooking"
	RentalService_CheckAvailability_FullMethodName = "/car_rental.v1.RentalService/CheckAvailability"
)

type RentalServiceClient interface {
	CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*Booking, error)
	CheckAvailability(ctx context.Context, in *CheckAvailabilityRequest, opts ...grpc.CallOption) (*CheckAvailabilityResponse, error)
}

type rentalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRentalServiceClient(cc grpc.ClientConnInterface) RentalServiceClient {
	return &rentalServiceClient{cc}
}

func (c *rentalServiceClient) CreateBooking(ctx context.Context, in *CreateBookingRequest, opts ...grpc.CallOption) (*Booking, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Booking)
	err := c.cc.Invoke(ctx, RentalService_CreateBooking_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rentalServiceClient) CheckAvailability(ctx context.Context, in *CheckAvailabilityRequest, opts ...grpc.CallOption) (*CheckAvailabilityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckAvailabilityResponse)
	err := c.cc.Invoke(ctx, RentalService_CheckAvailability_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type RentalServiceServer interface {
	CreateBooking(context.Context, *CreateBookingRequest) (*Booking, error)
	CheckAvailability(context.Context, *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error)
	mustEmbedUnimplementedRentalServiceServer()
}

type UnimplementedRentalServiceServer struct{}

func (UnimplementedRentalServiceServer) CreateBooking(context.Context, *CreateBookingRequest) (*Booking, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateBooking not implemented")
}
func (UnimplementedRentalServiceServer) CheckAvailability(context.Context, *CheckAvailabilityRequest) (*CheckAvailabilityResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CheckAvailability not implemented")
}
func (UnimplementedRentalServiceServer) mustEmbedUnimplementedRentalServiceServer() {}
func (UnimplementedRentalServiceServer) testEmbeddedByValue()                       {}

type UnsafeRentalServiceServer interface {
	mustEmbedUnimplementedRentalServiceServer()
}

func RegisterRentalServiceServer(s grpc.ServiceRegistrar, srv RentalServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RentalService_ServiceDesc, srv)
}

func _RentalService_CreateBooking_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).CreateBooking(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RentalService_CreateBooking_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).CreateBooking(ctx, req.(*CreateBookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RentalService_CheckAvailability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckAvailabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RentalServiceServer).CheckAvailability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RentalService_CheckAvailability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RentalServiceServer).CheckAvailability(ctx, req.(*CheckAvailabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var RentalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "car_rental.v1.RentalService",
	HandlerType: (*RentalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBooking",
			Handler:    _RentalService_CreateBooking_Handler,
		},
		{
			MethodName: "CheckAvailability",
			Handler:    _RentalService_CheckAvailability_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rental.proto",
}
