package searchenginev1

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)

	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Article struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId      string                 `protobuf:"bytes,4,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Article) Reset() {
	*x = Article{}
	mi := &file_search_engine_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_search_engine_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Article) Descriptor() ([]byte, []int) {
	return file_search_engine_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type CreateArticleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	AuthorId      string                 `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateArticleRequest) Reset() {
	*x = CreateArticleRequest{}
	mi := &file_search_engine_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateArticleRequest) ProtoMessage() {}

func (x *CreateArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_engine_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateArticleRequest) Descriptor() ([]byte, []int) {
	return file_search_engine_proto_rawDescGZIP(), []int{1}
}

func (x *CreateArticleRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateArticleRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateArticleRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

type GetArticleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetArticleRequest) Reset() {
	*x = GetArticleRequest{}
	mi := &file_search_engine_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleRequest) ProtoMessage() {}

func (x *GetArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_engine_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetArticleRequest) Descriptor() ([]byte, []int) {
	return file_search_engine_proto_rawDescGZIP(), []int{2}
}

func (x *GetArticleRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type IndexArticleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Article       *Article               `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IndexArticleRequest) Reset() {
	*x = IndexArticleRequest{}
	mi := &file_search_engine_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndexArticleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexArticleRequest) ProtoMessage() {}

func (x *IndexArticleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_engine_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*IndexArticleRequest) Descriptor() ([]byte, []int) {
	return file_search_engine_proto_rawDescGZIP(), []int{3}
}

func (x *IndexArticleRequest) GetArticle() *Article {
	if x != nil {
		return x.Article
	}
	return nil
}

type IndexArticleResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IndexArticleResponse) Reset() {
	*x = IndexArticleResponse{}
	mi := &file_search_engine_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IndexArticleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexArticleResponse) ProtoMessage() {}

func (x *IndexArticleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_engine_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*IndexArticleResponse) Descriptor() ([]byte, []int) {
	return file_search_engine_proto_rawDescGZIP(), []int{4}
}

func (x *IndexArticleResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_search_engine_proto protoreflect.FileDescriptor

const file_search_engine_proto_rawDesc = "" +
	"\n" +
	"\x13search_engine.proto\x12\x10search_engine.v1\"f\n" +
	"\aArticle\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05title\x18\x02 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x03 \x01(\tR\acontent\x12\x1b\n" +
	"\tauthor_id\x18\x04 \x01(\tR\bauthorId\"c\n" +
	"\x14CreateArticleRequest\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x12\x1b\n" +
	"\tauthor_id\x18\x03 \x01(\tR\bauthorId\"#\n" +
	"\x11GetArticleRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"J\n" +
	"\x13IndexArticleRequest\x123\n" +
	"\aarticle\x18\x01 \x01(\v2\x19.search_engine.v1.ArticleR\aarticle\"0\n" +
	"\x14IndexArticleResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess2\xb6\x01\n" +
	"\x0eContentService\x12T\n" +
	"\rCreateArticle\x12&.search_engine.v1.CreateArticleRequest\x1a\x19.search_engine.v1.Article\"\x00\x12N\n" +
	"\n" +
	"GetArticle\x12#.search_engine.v1.GetArticleRequest\x1a\x19.search_engine.v1.Article\"\x002n\n" +
	"\rSearchService\x12]\n" +
	"\fIndexArticle\x12%.search_engine.v1.IndexArticleRequest\x1a&.search_engine.v1.IndexArticleResponseB6Z4go-kozadayev-exercise/task7/api/proto;searchenginev1b\x06proto3"

var (
	file_search_engine_proto_rawDescOnce sync.Once
	file_search_engine_proto_rawDescData []byte
)

func file_search_engine_proto_rawDescGZIP() []byte {
	file_search_engine_proto_rawDescOnce.Do(func() {
		file_search_engine_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_search_engine_proto_rawDesc), len(file_search_engine_proto_rawDesc)))
	})
	return file_search_engine_proto_rawDescData
}

var file_search_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_search_engine_proto_goTypes = []any{
	(*Article)(nil),
	(*CreateArticleRequest)(nil),
	(*GetArticleRequest)(nil),
	(*IndexArticleRequest)(nil),
	(*IndexArticleResponse)(nil),
}
var file_search_engine_proto_depIdxs = []int32{
	0,
	1,
	2,
	3,
	0,
	0,
	4,
	4,
	1,
	1,
	1,
	0,
}

func init() { file_search_engine_proto_init() }
func file_search_engine_proto_init() {
	if File_search_engine_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_search_engine_proto_rawDesc), len(file_search_engine_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_search_engine_proto_goTypes,
		DependencyIndexes: file_search_engine_proto_depIdxs,
		MessageInfos:      file_search_engine_proto_msgTypes,
	}.Build()
	File_search_engine_proto = out.File
	file_search_engine_proto_goTypes = nil
	file_search_engine_proto_depIdxs = nil
}
