package smarthomev1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	SmartHomeService_AddReading_FullMethodName      = "/smarthome.v1.SmartHomeService/AddReading"
	SmartHomeService_MonitorReadings_FullMethodName = "/smarthome.v1.SmartHomeService/MonitorReadings"
)

type SmartHomeServiceClient interface {
	AddReading(ctx context.Context, in *AddReadingRequest, opts ...grpc.CallOption) (*AddReadingResponse, error)
	MonitorReadings(ctx context.Context, in *MonitorReadingsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Reading], error)
}

type smartHomeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmartHomeServiceClient(cc grpc.ClientConnInterface) SmartHomeServiceClient {
	return &smartHomeServiceClient{cc}
}

func (c *smartHomeServiceClient) AddReading(ctx context.Context, in *AddReadingRequest, opts ...grpc.CallOption) (*AddReadingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddReadingResponse)
	err := c.cc.Invoke(ctx, SmartHomeService_AddReading_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smartHomeServiceClient) MonitorReadings(ctx context.Context, in *MonitorReadingsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Reading], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SmartHomeService_ServiceDesc.Streams[0], SmartHomeService_MonitorReadings_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MonitorReadingsRequest, Reading]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SmartHomeService_MonitorReadingsClient = grpc.ServerStreamingClient[Reading]

type SmartHomeServiceServer interface {
	AddReading(context.Context, *AddReadingRequest) (*AddReadingResponse, error)
	MonitorReadings(*MonitorReadingsRequest, grpc.ServerStreamingServer[Reading]) error
	mustEmbedUnimplementedSmartHomeServiceServer()
}

type UnimplementedSmartHomeServiceServer struct{}

func (UnimplementedSmartHomeServiceServer) AddReading(context.Context, *AddReadingRequest) (*AddReadingResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method AddReading not implemented")
}
func (UnimplementedSmartHomeServiceServer) MonitorReadings(*MonitorReadingsRequest, grpc.ServerStreamingServer[Reading]) error {
	return status.Error(codes.Unimplemented, "method MonitorReadings not implemented")
}
func (UnimplementedSmartHomeServiceServer) mustEmbedUnimplementedSmartHomeServiceServer() {}
func (UnimplementedSmartHomeServiceServer) testEmbeddedByValue()                          {}

type UnsafeSmartHomeServiceServer interface {
	mustEmbedUnimplementedSmartHomeServiceServer()
}

func RegisterSmartHomeServiceServer(s grpc.ServiceRegistrar, srv SmartHomeServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SmartHomeService_ServiceDesc, srv)
}

func _SmartHomeService_AddReading_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReadingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmartHomeServiceServer).AddReading(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmartHomeService_AddReading_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmartHomeServiceServer).AddReading(ctx, req.(*AddReadingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmartHomeService_MonitorReadings_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorReadingsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SmartHomeServiceServer).MonitorReadings(m, &grpc.GenericServerStream[MonitorReadingsRequest, Reading]{ServerStream: stream})
}

type SmartHomeService_MonitorReadingsServer = grpc.ServerStreamingServer[Reading]

var SmartHomeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "smarthome.v1.SmartHomeService",
	HandlerType: (*SmartHomeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddReading",
			Handler:    _SmartHomeService_AddReading_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorReadings",
			Handler:       _SmartHomeService_MonitorReadings_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "smarthome.proto",
}
