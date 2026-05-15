package dashboardv1

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

type Owner struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Owner) Reset() {
	*x = Owner{}
	mi := &file_dashboard_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Owner) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *Owner) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Owner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Car struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand         string                 `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	OwnerId       int64                  `protobuf:"varint,3,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Car) Reset() {
	*x = Car{}
	mi := &file_dashboard_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Car) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *Car) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Car) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Car) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type ServiceRecord struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CarId         int64                  `protobuf:"varint,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Date          int64                  `protobuf:"varint,4,opt,name=date,proto3" json:"date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceRecord) Reset() {
	*x = ServiceRecord{}
	mi := &file_dashboard_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRecord) ProtoMessage() {}

func (x *ServiceRecord) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*ServiceRecord) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceRecord) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ServiceRecord) GetCarId() int64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *ServiceRecord) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ServiceRecord) GetDate() int64 {
	if x != nil {
		return x.Date
	}
	return 0
}

type GetDashboardRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OwnerId       int64                  `protobuf:"varint,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDashboardRequest) Reset() {
	*x = GetDashboardRequest{}
	mi := &file_dashboard_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDashboardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardRequest) ProtoMessage() {}

func (x *GetDashboardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetDashboardRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{3}
}

func (x *GetDashboardRequest) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type CarDashboardEntry struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Car           *Car                   `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
	LastService   *ServiceRecord         `protobuf:"bytes,2,opt,name=last_service,json=lastService,proto3" json:"last_service,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CarDashboardEntry) Reset() {
	*x = CarDashboardEntry{}
	mi := &file_dashboard_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CarDashboardEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarDashboardEntry) ProtoMessage() {}

func (x *CarDashboardEntry) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CarDashboardEntry) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{4}
}

func (x *CarDashboardEntry) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

func (x *CarDashboardEntry) GetLastService() *ServiceRecord {
	if x != nil {
		return x.LastService
	}
	return nil
}

type GetDashboardResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         *Owner                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Cars          []*CarDashboardEntry   `protobuf:"bytes,2,rep,name=cars,proto3" json:"cars,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDashboardResponse) Reset() {
	*x = GetDashboardResponse{}
	mi := &file_dashboard_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDashboardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDashboardResponse) ProtoMessage() {}

func (x *GetDashboardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetDashboardResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{5}
}

func (x *GetDashboardResponse) GetOwner() *Owner {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *GetDashboardResponse) GetCars() []*CarDashboardEntry {
	if x != nil {
		return x.Cars
	}
	return nil
}

type CreateOwnerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOwnerRequest) Reset() {
	*x = CreateOwnerRequest{}
	mi := &file_dashboard_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOwnerRequest) ProtoMessage() {}

func (x *CreateOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateOwnerRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{6}
}

func (x *CreateOwnerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateCarRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Brand         string                 `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	OwnerId       int64                  `protobuf:"varint,2,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCarRequest) Reset() {
	*x = CreateCarRequest{}
	mi := &file_dashboard_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarRequest) ProtoMessage() {}

func (x *CreateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateCarRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{7}
}

func (x *CreateCarRequest) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CreateCarRequest) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

type CreateServiceRecordRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CarId         int64                  `protobuf:"varint,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateServiceRecordRequest) Reset() {
	*x = CreateServiceRecordRequest{}
	mi := &file_dashboard_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateServiceRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateServiceRecordRequest) ProtoMessage() {}

func (x *CreateServiceRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateServiceRecordRequest) Descriptor() ([]byte, []int) {
	return file_dashboard_proto_rawDescGZIP(), []int{8}
}

func (x *CreateServiceRecordRequest) GetCarId() int64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *CreateServiceRecordRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_dashboard_proto protoreflect.FileDescriptor

const file_dashboard_proto_rawDesc = "" +
	"\n" +
	"\x0fdashboard.proto\x12\fdashboard.v1\"+\n" +
	"\x05Owner\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\"F\n" +
	"\x03Car\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x14\n" +
	"\x05brand\x18\x02 \x01(\tR\x05brand\x12\x19\n" +
	"\bowner_id\x18\x03 \x01(\x03R\aownerId\"l\n" +
	"\rServiceRecord\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x15\n" +
	"\x06car_id\x18\x02 \x01(\x03R\x05carId\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x12\n" +
	"\x04date\x18\x04 \x01(\x03R\x04date\"0\n" +
	"\x13GetDashboardRequest\x12\x19\n" +
	"\bowner_id\x18\x01 \x01(\x03R\aownerId\"x\n" +
	"\x11CarDashboardEntry\x12#\n" +
	"\x03car\x18\x01 \x01(\v2\x11.dashboard.v1.CarR\x03car\x12>\n" +
	"\flast_service\x18\x02 \x01(\v2\x1b.dashboard.v1.ServiceRecordR\vlastService\"v\n" +
	"\x14GetDashboardResponse\x12)\n" +
	"\x05owner\x18\x01 \x01(\v2\x13.dashboard.v1.OwnerR\x05owner\x123\n" +
	"\x04cars\x18\x02 \x03(\v2\x1f.dashboard.v1.CarDashboardEntryR\x04cars\"(\n" +
	"\x12CreateOwnerRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"C\n" +
	"\x10CreateCarRequest\x12\x14\n" +
	"\x05brand\x18\x01 \x01(\tR\x05brand\x12\x19\n" +
	"\bowner_id\x18\x02 \x01(\x03R\aownerId\"U\n" +
	"\x1aCreateServiceRecordRequest\x12\x15\n" +
	"\x06car_id\x18\x01 \x01(\x03R\x05carId\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription2\xcf\x02\n" +
	"\x10DashboardService\x12W\n" +
	"\fGetDashboard\x12!.dashboard.v1.GetDashboardRequest\x1a\".dashboard.v1.GetDashboardResponse\"\x00\x12D\n" +
	"\vCreateOwner\x12 .dashboard.v1.CreateOwnerRequest\x1a\x13.dashboard.v1.Owner\x12>\n" +
	"\tCreateCar\x12\x1e.dashboard.v1.CreateCarRequest\x1a\x11.dashboard.v1.Car\x12\\\n" +
	"\x13CreateServiceRecord\x12(.dashboard.v1.CreateServiceRecordRequest\x1a\x1b.dashboard.v1.ServiceRecordB3Z1go-kozadayev-exercise/task6/api/proto;dashboardv1b\x06proto3"

var (
	file_dashboard_proto_rawDescOnce sync.Once
	file_dashboard_proto_rawDescData []byte
)

func file_dashboard_proto_rawDescGZIP() []byte {
	file_dashboard_proto_rawDescOnce.Do(func() {
		file_dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_dashboard_proto_rawDesc), len(file_dashboard_proto_rawDesc)))
	})
	return file_dashboard_proto_rawDescData
}

var file_dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_dashboard_proto_goTypes = []any{
	(*Owner)(nil),
	(*Car)(nil),
	(*ServiceRecord)(nil),
	(*GetDashboardRequest)(nil),
	(*CarDashboardEntry)(nil),
	(*GetDashboardResponse)(nil),
	(*CreateOwnerRequest)(nil),
	(*CreateCarRequest)(nil),
	(*CreateServiceRecordRequest)(nil),
}
var file_dashboard_proto_depIdxs = []int32{
	1,
	2,
	0,
	4,
	3,
	6,
	7,
	8,
	5,
	0,
	1,
	2,
	8,
	4,
	4,
	4,
	0,
}

func init() { file_dashboard_proto_init() }
func file_dashboard_proto_init() {
	if File_dashboard_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_dashboard_proto_rawDesc), len(file_dashboard_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dashboard_proto_goTypes,
		DependencyIndexes: file_dashboard_proto_depIdxs,
		MessageInfos:      file_dashboard_proto_msgTypes,
	}.Build()
	File_dashboard_proto = out.File
	file_dashboard_proto_goTypes = nil
	file_dashboard_proto_depIdxs = nil
}
