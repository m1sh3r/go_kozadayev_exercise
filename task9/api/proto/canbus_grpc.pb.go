package canbusv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	CanBusService_TelemetryStream_FullMethodName = "/canbus.v1.CanBusService/TelemetryStream"
	CanBusService_GetAlerts_FullMethodName       = "/canbus.v1.CanBusService/GetAlerts"
)

type CanBusServiceClient interface {
	TelemetryStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Signal, Alert], error)
	GetAlerts(ctx context.Context, in *GetAlertsRequest, opts ...grpc.CallOption) (*GetAlertsResponse, error)
}

type canBusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCanBusServiceClient(cc grpc.ClientConnInterface) CanBusServiceClient {
	return &canBusServiceClient{cc}
}

func (c *canBusServiceClient) TelemetryStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Signal, Alert], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CanBusService_ServiceDesc.Streams[0], CanBusService_TelemetryStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Signal, Alert]{ClientStream: stream}
	return x, nil
}

type CanBusService_TelemetryStreamClient = grpc.BidiStreamingClient[Signal, Alert]

func (c *canBusServiceClient) GetAlerts(ctx context.Context, in *GetAlertsRequest, opts ...grpc.CallOption) (*GetAlertsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAlertsResponse)
	err := c.cc.Invoke(ctx, CanBusService_GetAlerts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type CanBusServiceServer interface {
	TelemetryStream(grpc.BidiStreamingServer[Signal, Alert]) error
	GetAlerts(context.Context, *GetAlertsRequest) (*GetAlertsResponse, error)
	mustEmbedUnimplementedCanBusServiceServer()
}

type UnimplementedCanBusServiceServer struct{}

func (UnimplementedCanBusServiceServer) TelemetryStream(grpc.BidiStreamingServer[Signal, Alert]) error {
	return status.Error(codes.Unimplemented, "method TelemetryStream not implemented")
}
func (UnimplementedCanBusServiceServer) GetAlerts(context.Context, *GetAlertsRequest) (*GetAlertsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetAlerts not implemented")
}
func (UnimplementedCanBusServiceServer) mustEmbedUnimplementedCanBusServiceServer() {}
func (UnimplementedCanBusServiceServer) testEmbeddedByValue()                       {}

type UnsafeCanBusServiceServer interface {
	mustEmbedUnimplementedCanBusServiceServer()
}

func RegisterCanBusServiceServer(s grpc.ServiceRegistrar, srv CanBusServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CanBusService_ServiceDesc, srv)
}

func _CanBusService_TelemetryStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CanBusServiceServer).TelemetryStream(&grpc.GenericServerStream[Signal, Alert]{ServerStream: stream})
}

type CanBusService_TelemetryStreamServer = grpc.BidiStreamingServer[Signal, Alert]

func _CanBusService_GetAlerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CanBusServiceServer).GetAlerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CanBusService_GetAlerts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CanBusServiceServer).GetAlerts(ctx, req.(*GetAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var CanBusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "canbus.v1.CanBusService",
	HandlerType: (*CanBusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAlerts",
			Handler:    _CanBusService_GetAlerts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TelemetryStream",
			Handler:       _CanBusService_TelemetryStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "canbus.proto",
}
