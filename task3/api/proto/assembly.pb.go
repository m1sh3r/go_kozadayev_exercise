package assemblyv1

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

type Engine struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Power         int32                  `protobuf:"varint,3,opt,name=power,proto3" json:"power,omitempty"`
	Volume        float32                `protobuf:"fixed32,4,opt,name=volume,proto3" json:"volume,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Engine) Reset() {
	*x = Engine{}
	mi := &file_assembly_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Engine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Engine) ProtoMessage() {}

func (x *Engine) ProtoReflect() protoreflect.Message {
	mi := &file_assembly_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Engine) Descriptor() ([]byte, []int) {
	return file_assembly_proto_rawDescGZIP(), []int{0}
}

func (x *Engine) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Engine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Engine) GetPower() int32 {
	if x != nil {
		return x.Power
	}
	return 0
}

func (x *Engine) GetVolume() float32 {
	if x != nil {
		return x.Volume
	}
	return 0
}

type Transmission struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type          string                 `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Transmission) Reset() {
	*x = Transmission{}
	mi := &file_assembly_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Transmission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transmission) ProtoMessage() {}

func (x *Transmission) ProtoReflect() protoreflect.Message {
	mi := &file_assembly_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Transmission) Descriptor() ([]byte, []int) {
	return file_assembly_proto_rawDescGZIP(), []int{1}
}

func (x *Transmission) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Transmission) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type CarSpec struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand         string                 `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Model         string                 `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Year          int32                  `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
	Engine        *Engine                `protobuf:"bytes,5,opt,name=engine,proto3" json:"engine,omitempty"`
	Transmission  *Transmission          `protobuf:"bytes,6,opt,name=transmission,proto3" json:"transmission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CarSpec) Reset() {
	*x = CarSpec{}
	mi := &file_assembly_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CarSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarSpec) ProtoMessage() {}

func (x *CarSpec) ProtoReflect() protoreflect.Message {
	mi := &file_assembly_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CarSpec) Descriptor() ([]byte, []int) {
	return file_assembly_proto_rawDescGZIP(), []int{2}
}

func (x *CarSpec) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarSpec) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CarSpec) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CarSpec) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *CarSpec) GetEngine() *Engine {
	if x != nil {
		return x.Engine
	}
	return nil
}

func (x *CarSpec) GetTransmission() *Transmission {
	if x != nil {
		return x.Transmission
	}
	return nil
}

type AssembleCarRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CarId          int64                  `protobuf:"varint,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	EngineId       int64                  `protobuf:"varint,2,opt,name=engine_id,json=engineId,proto3" json:"engine_id,omitempty"`
	TransmissionId int64                  `protobuf:"varint,3,opt,name=transmission_id,json=transmissionId,proto3" json:"transmission_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *AssembleCarRequest) Reset() {
	*x = AssembleCarRequest{}
	mi := &file_assembly_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssembleCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssembleCarRequest) ProtoMessage() {}

func (x *AssembleCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_assembly_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*AssembleCarRequest) Descriptor() ([]byte, []int) {
	return file_assembly_proto_rawDescGZIP(), []int{3}
}

func (x *AssembleCarRequest) GetCarId() int64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *AssembleCarRequest) GetEngineId() int64 {
	if x != nil {
		return x.EngineId
	}
	return 0
}

func (x *AssembleCarRequest) GetTransmissionId() int64 {
	if x != nil {
		return x.TransmissionId
	}
	return 0
}

type AssembleCarResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CarSpec       *CarSpec               `protobuf:"bytes,1,opt,name=car_spec,json=carSpec,proto3" json:"car_spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssembleCarResponse) Reset() {
	*x = AssembleCarResponse{}
	mi := &file_assembly_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssembleCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssembleCarResponse) ProtoMessage() {}

func (x *AssembleCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assembly_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*AssembleCarResponse) Descriptor() ([]byte, []int) {
	return file_assembly_proto_rawDescGZIP(), []int{4}
}

func (x *AssembleCarResponse) GetCarSpec() *CarSpec {
	if x != nil {
		return x.CarSpec
	}
	return nil
}

type GetCarSpecRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CarId         int64                  `protobuf:"varint,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCarSpecRequest) Reset() {
	*x = GetCarSpecRequest{}
	mi := &file_assembly_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCarSpecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarSpecRequest) ProtoMessage() {}

func (x *GetCarSpecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_assembly_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetCarSpecRequest) Descriptor() ([]byte, []int) {
	return file_assembly_proto_rawDescGZIP(), []int{5}
}

func (x *GetCarSpecRequest) GetCarId() int64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

type GetCarSpecResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CarSpec       *CarSpec               `protobuf:"bytes,1,opt,name=car_spec,json=carSpec,proto3" json:"car_spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCarSpecResponse) Reset() {
	*x = GetCarSpecResponse{}
	mi := &file_assembly_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCarSpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarSpecResponse) ProtoMessage() {}

func (x *GetCarSpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_assembly_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetCarSpecResponse) Descriptor() ([]byte, []int) {
	return file_assembly_proto_rawDescGZIP(), []int{6}
}

func (x *GetCarSpecResponse) GetCarSpec() *CarSpec {
	if x != nil {
		return x.CarSpec
	}
	return nil
}

var File_assembly_proto protoreflect.FileDescriptor

const file_assembly_proto_rawDesc = "" +
	"\n" +
	"\x0eassembly.proto\x12\x0fcar_assembly.v1\"Z\n" +
	"\x06Engine\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05power\x18\x03 \x01(\x05R\x05power\x12\x16\n" +
	"\x06volume\x18\x04 \x01(\x02R\x06volume\"2\n" +
	"\fTransmission\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04type\x18\x02 \x01(\tR\x04type\"\xcd\x01\n" +
	"\aCarSpec\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x14\n" +
	"\x05brand\x18\x02 \x01(\tR\x05brand\x12\x14\n" +
	"\x05model\x18\x03 \x01(\tR\x05model\x12\x12\n" +
	"\x04year\x18\x04 \x01(\x05R\x04year\x12/\n" +
	"\x06engine\x18\x05 \x01(\v2\x17.car_assembly.v1.EngineR\x06engine\x12A\n" +
	"\ftransmission\x18\x06 \x01(\v2\x1d.car_assembly.v1.TransmissionR\ftransmission\"q\n" +
	"\x12AssembleCarRequest\x12\x15\n" +
	"\x06car_id\x18\x01 \x01(\x03R\x05carId\x12\x1b\n" +
	"\tengine_id\x18\x02 \x01(\x03R\bengineId\x12'\n" +
	"\x0ftransmission_id\x18\x03 \x01(\x03R\x0etransmissionId\"J\n" +
	"\x13AssembleCarResponse\x123\n" +
	"\bcar_spec\x18\x01 \x01(\v2\x18.car_assembly.v1.CarSpecR\acarSpec\"*\n" +
	"\x11GetCarSpecRequest\x12\x15\n" +
	"\x06car_id\x18\x01 \x01(\x03R\x05carId\"I\n" +
	"\x12GetCarSpecResponse\x123\n" +
	"\bcar_spec\x18\x01 \x01(\v2\x18.car_assembly.v1.CarSpecR\acarSpec2\xc6\x01\n" +
	"\x0fAssemblyService\x12Z\n" +
	"\vAssembleCar\x12#.car_assembly.v1.AssembleCarRequest\x1a$.car_assembly.v1.AssembleCarResponse\"\x00\x12W\n" +
	"\n" +
	"GetCarSpec\x12\".car_assembly.v1.GetCarSpecRequest\x1a#.car_assembly.v1.GetCarSpecResponse\"\x00B2Z0go-kozadayev-exercise/task3/api/proto;assemblyv1b\x06proto3"

var (
	file_assembly_proto_rawDescOnce sync.Once
	file_assembly_proto_rawDescData []byte
)

func file_assembly_proto_rawDescGZIP() []byte {
	file_assembly_proto_rawDescOnce.Do(func() {
		file_assembly_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_assembly_proto_rawDesc), len(file_assembly_proto_rawDesc)))
	})
	return file_assembly_proto_rawDescData
}

var file_assembly_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_assembly_proto_goTypes = []any{
	(*Engine)(nil),
	(*Transmission)(nil),
	(*CarSpec)(nil),
	(*AssembleCarRequest)(nil),
	(*AssembleCarResponse)(nil),
	(*GetCarSpecRequest)(nil),
	(*GetCarSpecResponse)(nil),
}
var file_assembly_proto_depIdxs = []int32{
	0,
	1,
	2,
	2,
	3,
	5,
	4,
	6,
	6,
	4,
	4,
	4,
	0,
}

func init() { file_assembly_proto_init() }
func file_assembly_proto_init() {
	if File_assembly_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_assembly_proto_rawDesc), len(file_assembly_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_assembly_proto_goTypes,
		DependencyIndexes: file_assembly_proto_depIdxs,
		MessageInfos:      file_assembly_proto_msgTypes,
	}.Build()
	File_assembly_proto = out.File
	file_assembly_proto_goTypes = nil
	file_assembly_proto_depIdxs = nil
}
