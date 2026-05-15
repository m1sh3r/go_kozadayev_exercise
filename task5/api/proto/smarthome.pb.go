package smarthomev1

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

type Reading struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      string                 `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Value         float32                `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Reading) Reset() {
	*x = Reading{}
	mi := &file_smarthome_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Reading) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reading) ProtoMessage() {}

func (x *Reading) ProtoReflect() protoreflect.Message {
	mi := &file_smarthome_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Reading) Descriptor() ([]byte, []int) {
	return file_smarthome_proto_rawDescGZIP(), []int{0}
}

func (x *Reading) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Reading) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Reading) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type AddReadingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      string                 `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Value         float32                `protobuf:"fixed32,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddReadingRequest) Reset() {
	*x = AddReadingRequest{}
	mi := &file_smarthome_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddReadingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReadingRequest) ProtoMessage() {}

func (x *AddReadingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smarthome_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*AddReadingRequest) Descriptor() ([]byte, []int) {
	return file_smarthome_proto_rawDescGZIP(), []int{1}
}

func (x *AddReadingRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *AddReadingRequest) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type AddReadingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddReadingResponse) Reset() {
	*x = AddReadingResponse{}
	mi := &file_smarthome_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddReadingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReadingResponse) ProtoMessage() {}

func (x *AddReadingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_smarthome_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*AddReadingResponse) Descriptor() ([]byte, []int) {
	return file_smarthome_proto_rawDescGZIP(), []int{2}
}

func (x *AddReadingResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type MonitorReadingsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      string                 `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MonitorReadingsRequest) Reset() {
	*x = MonitorReadingsRequest{}
	mi := &file_smarthome_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MonitorReadingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonitorReadingsRequest) ProtoMessage() {}

func (x *MonitorReadingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_smarthome_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*MonitorReadingsRequest) Descriptor() ([]byte, []int) {
	return file_smarthome_proto_rawDescGZIP(), []int{3}
}

func (x *MonitorReadingsRequest) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

var File_smarthome_proto protoreflect.FileDescriptor

const file_smarthome_proto_rawDesc = "" +
	"\n" +
	"\x0fsmarthome.proto\x12\fsmarthome.v1\"Z\n" +
	"\aReading\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\tR\bdeviceId\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x02R\x05value\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\x03R\ttimestamp\"F\n" +
	"\x11AddReadingRequest\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\tR\bdeviceId\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x02R\x05value\".\n" +
	"\x12AddReadingResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\"5\n" +
	"\x16MonitorReadingsRequest\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\tR\bdeviceId2\xb9\x01\n" +
	"\x10SmartHomeService\x12Q\n" +
	"\n" +
	"AddReading\x12\x1f.smarthome.v1.AddReadingRequest\x1a .smarthome.v1.AddReadingResponse\"\x00\x12R\n" +
	"\x0fMonitorReadings\x12$.smarthome.v1.MonitorReadingsRequest\x1a\x15.smarthome.v1.Reading\"\x000\x01B3Z1go-kozadayev-exercise/task5/api/proto;smarthomev1b\x06proto3"

var (
	file_smarthome_proto_rawDescOnce sync.Once
	file_smarthome_proto_rawDescData []byte
)

func file_smarthome_proto_rawDescGZIP() []byte {
	file_smarthome_proto_rawDescOnce.Do(func() {
		file_smarthome_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_smarthome_proto_rawDesc), len(file_smarthome_proto_rawDesc)))
	})
	return file_smarthome_proto_rawDescData
}

var file_smarthome_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_smarthome_proto_goTypes = []any{
	(*Reading)(nil),
	(*AddReadingRequest)(nil),
	(*AddReadingResponse)(nil),
	(*MonitorReadingsRequest)(nil),
}
var file_smarthome_proto_depIdxs = []int32{
	1,
	3,
	2,
	0,
	2,
	0,
	0,
	0,
	0,
}

func init() { file_smarthome_proto_init() }
func file_smarthome_proto_init() {
	if File_smarthome_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_smarthome_proto_rawDesc), len(file_smarthome_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_smarthome_proto_goTypes,
		DependencyIndexes: file_smarthome_proto_depIdxs,
		MessageInfos:      file_smarthome_proto_msgTypes,
	}.Build()
	File_smarthome_proto = out.File
	file_smarthome_proto_goTypes = nil
	file_smarthome_proto_depIdxs = nil
}
