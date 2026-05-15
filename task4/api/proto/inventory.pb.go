package inventoryv1

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

type Order struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId     string                 `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Status        string                 `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Order) Reset() {
	*x = Order{}
	mi := &file_inventory_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*Order) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Order) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CreateOrderRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductId     string                 `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	mi := &file_inventory_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type CreateOrderResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Order         *Order                 `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	mi := &file_inventory_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{2}
}

func (x *CreateOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type GetOrderStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int64                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderStatusRequest) Reset() {
	*x = GetOrderStatusRequest{}
	mi := &file_inventory_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderStatusRequest) ProtoMessage() {}

func (x *GetOrderStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetOrderStatusRequest) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderStatusRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type GetOrderStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetOrderStatusResponse) Reset() {
	*x = GetOrderStatusResponse{}
	mi := &file_inventory_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderStatusResponse) ProtoMessage() {}

func (x *GetOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_inventory_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (*GetOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_inventory_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_inventory_proto protoreflect.FileDescriptor

const file_inventory_proto_rawDesc = "" +
	"\n" +
	"\x0finventory.proto\x12\finventory.v1\"N\n" +
	"\x05Order\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1d\n" +
	"\n" +
	"product_id\x18\x02 \x01(\tR\tproductId\x12\x16\n" +
	"\x06status\x18\x03 \x01(\tR\x06status\"3\n" +
	"\x12CreateOrderRequest\x12\x1d\n" +
	"\n" +
	"product_id\x18\x01 \x01(\tR\tproductId\"@\n" +
	"\x13CreateOrderResponse\x12)\n" +
	"\x05order\x18\x01 \x01(\v2\x13.inventory.v1.OrderR\x05order\"2\n" +
	"\x15GetOrderStatusRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\x03R\aorderId\"0\n" +
	"\x16GetOrderStatusResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2\xc7\x01\n" +
	"\x10InventoryService\x12T\n" +
	"\vCreateOrder\x12 .inventory.v1.CreateOrderRequest\x1a!.inventory.v1.CreateOrderResponse\"\x00\x12]\n" +
	"\x0eGetOrderStatus\x12#.inventory.v1.GetOrderStatusRequest\x1a$.inventory.v1.GetOrderStatusResponse\"\x00B3Z1go-kozadayev-exercise/task4/api/proto;inventoryv1b\x06proto3"

var (
	file_inventory_proto_rawDescOnce sync.Once
	file_inventory_proto_rawDescData []byte
)

func file_inventory_proto_rawDescGZIP() []byte {
	file_inventory_proto_rawDescOnce.Do(func() {
		file_inventory_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_inventory_proto_rawDesc), len(file_inventory_proto_rawDesc)))
	})
	return file_inventory_proto_rawDescData
}

var file_inventory_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_inventory_proto_goTypes = []any{
	(*Order)(nil),
	(*CreateOrderRequest)(nil),
	(*CreateOrderResponse)(nil),
	(*GetOrderStatusRequest)(nil),
	(*GetOrderStatusResponse)(nil),
}
var file_inventory_proto_depIdxs = []int32{
	0,
	1,
	3,
	2,
	4,
	3,
	1,
	1,
	1,
	0,
}

func init() { file_inventory_proto_init() }
func file_inventory_proto_init() {
	if File_inventory_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_inventory_proto_rawDesc), len(file_inventory_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_inventory_proto_goTypes,
		DependencyIndexes: file_inventory_proto_depIdxs,
		MessageInfos:      file_inventory_proto_msgTypes,
	}.Build()
	File_inventory_proto = out.File
	file_inventory_proto_goTypes = nil
	file_inventory_proto_depIdxs = nil
}
