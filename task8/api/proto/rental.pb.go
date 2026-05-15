package rentalv1

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

type Booking struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CarId         int64                  `protobuf:"varint,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	StartDate     int64                  `protobuf:"varint,3,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate       int64                  `protobuf:"varint,4,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Booking) Reset() {
	*x = Booking{}
	mi := &file_rental_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Booking) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{0}
}

func (x *Booking) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Booking) GetCarId() int64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *Booking) GetStartDate() int64 {
	if x != nil {
		return x.StartDate
	}
	return 0
}

func (x *Booking) GetEndDate() int64 {
	if x != nil {
		return x.EndDate
	}
	return 0
}

type CreateBookingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CarId         int64                  `protobuf:"varint,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	StartDate     int64                  `protobuf:"varint,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate       int64                  `protobuf:"varint,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBookingRequest) Reset() {
	*x = CreateBookingRequest{}
	mi := &file_rental_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBookingRequest) ProtoMessage() {}

func (x *CreateBookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateBookingRequest) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBookingRequest) GetCarId() int64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *CreateBookingRequest) GetStartDate() int64 {
	if x != nil {
		return x.StartDate
	}
	return 0
}

func (x *CreateBookingRequest) GetEndDate() int64 {
	if x != nil {
		return x.EndDate
	}
	return 0
}

type CheckAvailabilityRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CarId         int64                  `protobuf:"varint,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	StartDate     int64                  `protobuf:"varint,2,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate       int64                  `protobuf:"varint,3,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckAvailabilityRequest) Reset() {
	*x = CheckAvailabilityRequest{}
	mi := &file_rental_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckAvailabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAvailabilityRequest) ProtoMessage() {}

func (x *CheckAvailabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CheckAvailabilityRequest) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{2}
}

func (x *CheckAvailabilityRequest) GetCarId() int64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *CheckAvailabilityRequest) GetStartDate() int64 {
	if x != nil {
		return x.StartDate
	}
	return 0
}

func (x *CheckAvailabilityRequest) GetEndDate() int64 {
	if x != nil {
		return x.EndDate
	}
	return 0
}

type CheckAvailabilityResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Available     bool                   `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckAvailabilityResponse) Reset() {
	*x = CheckAvailabilityResponse{}
	mi := &file_rental_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckAvailabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckAvailabilityResponse) ProtoMessage() {}

func (x *CheckAvailabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rental_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CheckAvailabilityResponse) Descriptor() ([]byte, []int) {
	return file_rental_proto_rawDescGZIP(), []int{3}
}

func (x *CheckAvailabilityResponse) GetAvailable() bool {
	if x != nil {
		return x.Available
	}
	return false
}

var File_rental_proto protoreflect.FileDescriptor

const file_rental_proto_rawDesc = "" +
	"\n" +
	"\frental.proto\x12\rcar_rental.v1\"j\n" +
	"\aBooking\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x15\n" +
	"\x06car_id\x18\x02 \x01(\x03R\x05carId\x12\x1d\n" +
	"\n" +
	"start_date\x18\x03 \x01(\x03R\tstartDate\x12\x19\n" +
	"\bend_date\x18\x04 \x01(\x03R\aendDate\"g\n" +
	"\x14CreateBookingRequest\x12\x15\n" +
	"\x06car_id\x18\x01 \x01(\x03R\x05carId\x12\x1d\n" +
	"\n" +
	"start_date\x18\x02 \x01(\x03R\tstartDate\x12\x19\n" +
	"\bend_date\x18\x03 \x01(\x03R\aendDate\"k\n" +
	"\x18CheckAvailabilityRequest\x12\x15\n" +
	"\x06car_id\x18\x01 \x01(\x03R\x05carId\x12\x1d\n" +
	"\n" +
	"start_date\x18\x02 \x01(\x03R\tstartDate\x12\x19\n" +
	"\bend_date\x18\x03 \x01(\x03R\aendDate\"9\n" +
	"\x19CheckAvailabilityResponse\x12\x1c\n" +
	"\tavailable\x18\x01 \x01(\bR\tavailable2\xc9\x01\n" +
	"\rRentalService\x12N\n" +
	"\rCreateBooking\x12#.car_rental.v1.CreateBookingRequest\x1a\x16.car_rental.v1.Booking\"\x00\x12h\n" +
	"\x11CheckAvailability\x12'.car_rental.v1.CheckAvailabilityRequest\x1a(.car_rental.v1.CheckAvailabilityResponse\"\x00B0Z.go-kozadayev-exercise/task8/api/proto;rentalv1b\x06proto3"

var (
	file_rental_proto_rawDescOnce sync.Once
	file_rental_proto_rawDescData []byte
)

func file_rental_proto_rawDescGZIP() []byte {
	file_rental_proto_rawDescOnce.Do(func() {
		file_rental_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_rental_proto_rawDesc), len(file_rental_proto_rawDesc)))
	})
	return file_rental_proto_rawDescData
}

var file_rental_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_rental_proto_goTypes = []any{
	(*Booking)(nil),
	(*CreateBookingRequest)(nil),
	(*CheckAvailabilityRequest)(nil),
	(*CheckAvailabilityResponse)(nil),
}
var file_rental_proto_depIdxs = []int32{
	1,
	2,
	0,
	3,
	2,
	0,
	0,
	0,
	0,
}

func init() { file_rental_proto_init() }
func file_rental_proto_init() {
	if File_rental_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_rental_proto_rawDesc), len(file_rental_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rental_proto_goTypes,
		DependencyIndexes: file_rental_proto_depIdxs,
		MessageInfos:      file_rental_proto_msgTypes,
	}.Build()
	File_rental_proto = out.File
	file_rental_proto_goTypes = nil
	file_rental_proto_depIdxs = nil
}
