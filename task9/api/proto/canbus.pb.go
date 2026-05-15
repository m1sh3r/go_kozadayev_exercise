package canbusv1

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

type Signal struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      string                 `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value         float32                `protobuf:"fixed32,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Signal) Reset() {
	*x = Signal{}
	mi := &file_canbus_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_canbus_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Signal) Descriptor() ([]byte, []int) {
	return file_canbus_proto_rawDescGZIP(), []int{0}
}

func (x *Signal) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Signal) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Signal) GetValue() float32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Alert struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DeviceId      string                 `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Alert) Reset() {
	*x = Alert{}
	mi := &file_canbus_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_canbus_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Alert) Descriptor() ([]byte, []int) {
	return file_canbus_proto_rawDescGZIP(), []int{1}
}

func (x *Alert) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Alert) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Alert) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type GetAlertsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StartTime     int64                  `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       int64                  `protobuf:"varint,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAlertsRequest) Reset() {
	*x = GetAlertsRequest{}
	mi := &file_canbus_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAlertsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertsRequest) ProtoMessage() {}

func (x *GetAlertsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_canbus_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetAlertsRequest) Descriptor() ([]byte, []int) {
	return file_canbus_proto_rawDescGZIP(), []int{2}
}

func (x *GetAlertsRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *GetAlertsRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type GetAlertsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Alerts        []*Alert               `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAlertsResponse) Reset() {
	*x = GetAlertsResponse{}
	mi := &file_canbus_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAlertsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlertsResponse) ProtoMessage() {}

func (x *GetAlertsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_canbus_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetAlertsResponse) Descriptor() ([]byte, []int) {
	return file_canbus_proto_rawDescGZIP(), []int{3}
}

func (x *GetAlertsResponse) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

var File_canbus_proto protoreflect.FileDescriptor

const file_canbus_proto_rawDesc = "" +
	"\n" +
	"\fcanbus.proto\x12\tcanbus.v1\"O\n" +
	"\x06Signal\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\tR\bdeviceId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05value\x18\x03 \x01(\x02R\x05value\"\\\n" +
	"\x05Alert\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\tR\bdeviceId\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\x03R\ttimestamp\"L\n" +
	"\x10GetAlertsRequest\x12\x1d\n" +
	"\n" +
	"start_time\x18\x01 \x01(\x03R\tstartTime\x12\x19\n" +
	"\bend_time\x18\x02 \x01(\x03R\aendTime\"=\n" +
	"\x11GetAlertsResponse\x12(\n" +
	"\x06alerts\x18\x01 \x03(\v2\x10.canbus.v1.AlertR\x06alerts2\x95\x01\n" +
	"\rCanBusService\x12:\n" +
	"\x0fTelemetryStream\x12\x11.canbus.v1.Signal\x1a\x10.canbus.v1.Alert(\x010\x01\x12H\n" +
	"\tGetAlerts\x12\x1b.canbus.v1.GetAlertsRequest\x1a\x1c.canbus.v1.GetAlertsResponse\"\x00B0Z.go-kozadayev-exercise/task9/api/proto;canbusv1b\x06proto3"

var (
	file_canbus_proto_rawDescOnce sync.Once
	file_canbus_proto_rawDescData []byte
)

func file_canbus_proto_rawDescGZIP() []byte {
	file_canbus_proto_rawDescOnce.Do(func() {
		file_canbus_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_canbus_proto_rawDesc), len(file_canbus_proto_rawDesc)))
	})
	return file_canbus_proto_rawDescData
}

var file_canbus_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_canbus_proto_goTypes = []any{
	(*Signal)(nil),
	(*Alert)(nil),
	(*GetAlertsRequest)(nil),
	(*GetAlertsResponse)(nil),
}
var file_canbus_proto_depIdxs = []int32{
	1,
	0,
	2,
	1,
	3,
	3,
	1,
	1,
	1,
	0,
}

func init() { file_canbus_proto_init() }
func file_canbus_proto_init() {
	if File_canbus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_canbus_proto_rawDesc), len(file_canbus_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_canbus_proto_goTypes,
		DependencyIndexes: file_canbus_proto_depIdxs,
		MessageInfos:      file_canbus_proto_msgTypes,
	}.Build()
	File_canbus_proto = out.File
	file_canbus_proto_goTypes = nil
	file_canbus_proto_depIdxs = nil
}
