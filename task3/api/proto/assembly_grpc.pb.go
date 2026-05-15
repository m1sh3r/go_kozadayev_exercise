package assemblyv1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	AssemblyService_AssembleCar_FullMethodName = "/car_assembly.v1.AssemblyService/AssembleCar"
	AssemblyService_GetCarSpec_FullMethodName  = "/car_assembly.v1.AssemblyService/GetCarSpec"
)

type AssemblyServiceClient interface {
	AssembleCar(ctx context.Context, in *AssembleCarRequest, opts ...grpc.CallOption) (*AssembleCarResponse, error)
	GetCarSpec(ctx context.Context, in *GetCarSpecRequest, opts ...grpc.CallOption) (*GetCarSpecResponse, error)
}

type assemblyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAssemblyServiceClient(cc grpc.ClientConnInterface) AssemblyServiceClient {
	return &assemblyServiceClient{cc}
}

func (c *assemblyServiceClient) AssembleCar(ctx context.Context, in *AssembleCarRequest, opts ...grpc.CallOption) (*AssembleCarResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AssembleCarResponse)
	err := c.cc.Invoke(ctx, AssemblyService_AssembleCar_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assemblyServiceClient) GetCarSpec(ctx context.Context, in *GetCarSpecRequest, opts ...grpc.CallOption) (*GetCarSpecResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCarSpecResponse)
	err := c.cc.Invoke(ctx, AssemblyService_GetCarSpec_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type AssemblyServiceServer interface {
	AssembleCar(context.Context, *AssembleCarRequest) (*AssembleCarResponse, error)
	GetCarSpec(context.Context, *GetCarSpecRequest) (*GetCarSpecResponse, error)
	mustEmbedUnimplementedAssemblyServiceServer()
}

type UnimplementedAssemblyServiceServer struct{}

func (UnimplementedAssemblyServiceServer) AssembleCar(context.Context, *AssembleCarRequest) (*AssembleCarResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method AssembleCar not implemented")
}
func (UnimplementedAssemblyServiceServer) GetCarSpec(context.Context, *GetCarSpecRequest) (*GetCarSpecResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetCarSpec not implemented")
}
func (UnimplementedAssemblyServiceServer) mustEmbedUnimplementedAssemblyServiceServer() {}
func (UnimplementedAssemblyServiceServer) testEmbeddedByValue()                         {}

type UnsafeAssemblyServiceServer interface {
	mustEmbedUnimplementedAssemblyServiceServer()
}

func RegisterAssemblyServiceServer(s grpc.ServiceRegistrar, srv AssemblyServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AssemblyService_ServiceDesc, srv)
}

func _AssemblyService_AssembleCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssembleCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssemblyServiceServer).AssembleCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssemblyService_AssembleCar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssemblyServiceServer).AssembleCar(ctx, req.(*AssembleCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssemblyService_GetCarSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCarSpecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssemblyServiceServer).GetCarSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssemblyService_GetCarSpec_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssemblyServiceServer).GetCarSpec(ctx, req.(*GetCarSpecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var AssemblyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "car_assembly.v1.AssemblyService",
	HandlerType: (*AssemblyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AssembleCar",
			Handler:    _AssemblyService_AssembleCar_Handler,
		},
		{
			MethodName: "GetCarSpec",
			Handler:    _AssemblyService_GetCarSpec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "assembly.proto",
}
