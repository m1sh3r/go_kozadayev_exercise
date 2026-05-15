package maintenancev1

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

type WorkOrder struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CarId         string                 `protobuf:"bytes,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	TotalCost     float32                `protobuf:"fixed32,4,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkOrder) Reset() {
	*x = WorkOrder{}
	mi := &file_maintenance_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkOrder) ProtoMessage() {}

func (x *WorkOrder) ProtoReflect() protoreflect.Message {
	mi := &file_maintenance_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*WorkOrder) Descriptor() ([]byte, []int) {
	return file_maintenance_proto_rawDescGZIP(), []int{0}
}

func (x *WorkOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkOrder) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *WorkOrder) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WorkOrder) GetTotalCost() float32 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

type CreateWorkOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CarId         string                 `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	TotalCost     float32                `protobuf:"fixed32,3,opt,name=total_cost,json=totalCost,proto3" json:"total_cost,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateWorkOrderRequest) Reset() {
	*x = CreateWorkOrderRequest{}
	mi := &file_maintenance_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateWorkOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWorkOrderRequest) ProtoMessage() {}

func (x *CreateWorkOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_maintenance_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateWorkOrderRequest) Descriptor() ([]byte, []int) {
	return file_maintenance_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWorkOrderRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *CreateWorkOrderRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateWorkOrderRequest) GetTotalCost() float32 {
	if x != nil {
		return x.TotalCost
	}
	return 0
}

type GetWorkOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetWorkOrderRequest) Reset() {
	*x = GetWorkOrderRequest{}
	mi := &file_maintenance_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetWorkOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkOrderRequest) ProtoMessage() {}

func (x *GetWorkOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_maintenance_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetWorkOrderRequest) Descriptor() ([]byte, []int) {
	return file_maintenance_proto_rawDescGZIP(), []int{2}
}

func (x *GetWorkOrderRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ExportRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportRequest) Reset() {
	*x = ExportRequest{}
	mi := &file_maintenance_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportRequest) ProtoMessage() {}

func (x *ExportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_maintenance_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*ExportRequest) Descriptor() ([]byte, []int) {
	return file_maintenance_proto_rawDescGZIP(), []int{3}
}

type ExportResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileUrl       string                 `protobuf:"bytes,1,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExportResponse) Reset() {
	*x = ExportResponse{}
	mi := &file_maintenance_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportResponse) ProtoMessage() {}

func (x *ExportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_maintenance_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*ExportResponse) Descriptor() ([]byte, []int) {
	return file_maintenance_proto_rawDescGZIP(), []int{4}
}

func (x *ExportResponse) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

var File_maintenance_proto protoreflect.FileDescriptor

const file_maintenance_proto_rawDesc = "" +
	"\n" +
	"\x11maintenance.proto\x12\x0emaintenance.v1\"s\n" +
	"\tWorkOrder\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x15\n" +
	"\x06car_id\x18\x02 \x01(\tR\x05carId\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"total_cost\x18\x04 \x01(\x02R\ttotalCost\"p\n" +
	"\x16CreateWorkOrderRequest\x12\x15\n" +
	"\x06car_id\x18\x01 \x01(\tR\x05carId\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"total_cost\x18\x03 \x01(\x02R\ttotalCost\"%\n" +
	"\x13GetWorkOrderRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"\x0f\n" +
	"\rExportRequest\"+\n" +
	"\x0eExportResponse\x12\x19\n" +
	"\bfile_url\x18\x01 \x01(\tR\afileUrl2\x90\x02\n" +
	"\x12MaintenanceService\x12V\n" +
	"\x0fCreateWorkOrder\x12&.maintenance.v1.CreateWorkOrderRequest\x1a\x19.maintenance.v1.WorkOrder\"\x00\x12P\n" +
	"\fGetWorkOrder\x12#.maintenance.v1.GetWorkOrderRequest\x1a\x19.maintenance.v1.WorkOrder\"\x00\x12P\n" +
	"\rExportToExcel\x12\x1d.maintenance.v1.ExportRequest\x1a\x1e.maintenance.v1.ExportResponse\"\x00B6Z4go-kozadayev-exercise/task10/api/proto;maintenancev1b\x06proto3"

var (
	file_maintenance_proto_rawDescOnce sync.Once
	file_maintenance_proto_rawDescData []byte
)

func file_maintenance_proto_rawDescGZIP() []byte {
	file_maintenance_proto_rawDescOnce.Do(func() {
		file_maintenance_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_maintenance_proto_rawDesc), len(file_maintenance_proto_rawDesc)))
	})
	return file_maintenance_proto_rawDescData
}

var file_maintenance_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_maintenance_proto_goTypes = []any{
	(*WorkOrder)(nil),
	(*CreateWorkOrderRequest)(nil),
	(*GetWorkOrderRequest)(nil),
	(*ExportRequest)(nil),
	(*ExportResponse)(nil),
}
var file_maintenance_proto_depIdxs = []int32{
	1,
	2,
	3,
	0,
	0,
	4,
	3,
	0,
	0,
	0,
	0,
}

func init() { file_maintenance_proto_init() }
func file_maintenance_proto_init() {
	if File_maintenance_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_maintenance_proto_rawDesc), len(file_maintenance_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_maintenance_proto_goTypes,
		DependencyIndexes: file_maintenance_proto_depIdxs,
		MessageInfos:      file_maintenance_proto_msgTypes,
	}.Build()
	File_maintenance_proto = out.File
	file_maintenance_proto_goTypes = nil
	file_maintenance_proto_depIdxs = nil
}
