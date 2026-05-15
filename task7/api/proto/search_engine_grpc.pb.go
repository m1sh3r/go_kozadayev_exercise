package searchenginev1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

const _ = grpc.SupportPackageIsVersion9

const (
	ContentService_CreateArticle_FullMethodName = "/search_engine.v1.ContentService/CreateArticle"
	ContentService_GetArticle_FullMethodName    = "/search_engine.v1.ContentService/GetArticle"
)

type ContentServiceClient interface {
	CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*Article, error)
	GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*Article, error)
}

type contentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContentServiceClient(cc grpc.ClientConnInterface) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...grpc.CallOption) (*Article, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Article)
	err := c.cc.Invoke(ctx, ContentService_CreateArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...grpc.CallOption) (*Article, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Article)
	err := c.cc.Invoke(ctx, ContentService_GetArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type ContentServiceServer interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*Article, error)
	GetArticle(context.Context, *GetArticleRequest) (*Article, error)
	mustEmbedUnimplementedContentServiceServer()
}

type UnimplementedContentServiceServer struct{}

func (UnimplementedContentServiceServer) CreateArticle(context.Context, *CreateArticleRequest) (*Article, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateArticle not implemented")
}
func (UnimplementedContentServiceServer) GetArticle(context.Context, *GetArticleRequest) (*Article, error) {
	return nil, status.Error(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedContentServiceServer) mustEmbedUnimplementedContentServiceServer() {}
func (UnimplementedContentServiceServer) testEmbeddedByValue()                        {}

type UnsafeContentServiceServer interface {
	mustEmbedUnimplementedContentServiceServer()
}

func RegisterContentServiceServer(s grpc.ServiceRegistrar, srv ContentServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ContentService_ServiceDesc, srv)
}

func _ContentService_CreateArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).CreateArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_CreateArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).CreateArticle(ctx, req.(*CreateArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ContentService_GetArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).GetArticle(ctx, req.(*GetArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var ContentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search_engine.v1.ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateArticle",
			Handler:    _ContentService_CreateArticle_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _ContentService_GetArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search_engine.proto",
}

const (
	SearchService_IndexArticle_FullMethodName = "/search_engine.v1.SearchService/IndexArticle"
)

type SearchServiceClient interface {
	IndexArticle(ctx context.Context, in *IndexArticleRequest, opts ...grpc.CallOption) (*IndexArticleResponse, error)
}

type searchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchServiceClient(cc grpc.ClientConnInterface) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) IndexArticle(ctx context.Context, in *IndexArticleRequest, opts ...grpc.CallOption) (*IndexArticleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IndexArticleResponse)
	err := c.cc.Invoke(ctx, SearchService_IndexArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type SearchServiceServer interface {
	IndexArticle(context.Context, *IndexArticleRequest) (*IndexArticleResponse, error)
	mustEmbedUnimplementedSearchServiceServer()
}

type UnimplementedSearchServiceServer struct{}

func (UnimplementedSearchServiceServer) IndexArticle(context.Context, *IndexArticleRequest) (*IndexArticleResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method IndexArticle not implemented")
}
func (UnimplementedSearchServiceServer) mustEmbedUnimplementedSearchServiceServer() {}
func (UnimplementedSearchServiceServer) testEmbeddedByValue()                       {}

type UnsafeSearchServiceServer interface {
	mustEmbedUnimplementedSearchServiceServer()
}

func RegisterSearchServiceServer(s grpc.ServiceRegistrar, srv SearchServiceServer) {

	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SearchService_ServiceDesc, srv)
}

func _SearchService_IndexArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IndexArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).IndexArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SearchService_IndexArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).IndexArticle(ctx, req.(*IndexArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var SearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "search_engine.v1.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IndexArticle",
			Handler:    _SearchService_IndexArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search_engine.proto",
}
