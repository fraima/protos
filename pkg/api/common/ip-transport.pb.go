// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: common/ip-transport.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//IP net transport
type Networks_NetIP_Transport int32

const (
	Networks_NetIP_TCP Networks_NetIP_Transport = 0
	Networks_NetIP_UDP Networks_NetIP_Transport = 1
)

// Enum value maps for Networks_NetIP_Transport.
var (
	Networks_NetIP_Transport_name = map[int32]string{
		0: "TCP",
		1: "UDP",
	}
	Networks_NetIP_Transport_value = map[string]int32{
		"TCP": 0,
		"UDP": 1,
	}
)

func (x Networks_NetIP_Transport) Enum() *Networks_NetIP_Transport {
	p := new(Networks_NetIP_Transport)
	*p = x
	return p
}

func (x Networks_NetIP_Transport) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Networks_NetIP_Transport) Descriptor() protoreflect.EnumDescriptor {
	return file_common_ip_transport_proto_enumTypes[0].Descriptor()
}

func (Networks_NetIP_Transport) Type() protoreflect.EnumType {
	return &file_common_ip_transport_proto_enumTypes[0]
}

func (x Networks_NetIP_Transport) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Networks_NetIP_Transport.Descriptor instead.
func (Networks_NetIP_Transport) EnumDescriptor() ([]byte, []int) {
	return file_common_ip_transport_proto_rawDescGZIP(), []int{0, 0, 0}
}

//Networks
type Networks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Networks) Reset() {
	*x = Networks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ip_transport_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Networks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Networks) ProtoMessage() {}

func (x *Networks) ProtoReflect() protoreflect.Message {
	mi := &file_common_ip_transport_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Networks.ProtoReflect.Descriptor instead.
func (*Networks) Descriptor() ([]byte, []int) {
	return file_common_ip_transport_proto_rawDescGZIP(), []int{0}
}

//IP network
type Networks_NetIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//IP net address range
	CIDR string `protobuf:"bytes,1,opt,name=CIDR,proto3" json:"CIDR,omitempty"`
}

func (x *Networks_NetIP) Reset() {
	*x = Networks_NetIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ip_transport_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Networks_NetIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Networks_NetIP) ProtoMessage() {}

func (x *Networks_NetIP) ProtoReflect() protoreflect.Message {
	mi := &file_common_ip_transport_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Networks_NetIP.ProtoReflect.Descriptor instead.
func (*Networks_NetIP) Descriptor() ([]byte, []int) {
	return file_common_ip_transport_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Networks_NetIP) GetCIDR() string {
	if x != nil {
		return x.CIDR
	}
	return ""
}

//IP net port range
type Networks_NetIP_PortRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From uint32 `protobuf:"varint,1,opt,name=from,proto3" json:"from,omitempty"`
	To   uint32 `protobuf:"varint,2,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *Networks_NetIP_PortRange) Reset() {
	*x = Networks_NetIP_PortRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_ip_transport_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Networks_NetIP_PortRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Networks_NetIP_PortRange) ProtoMessage() {}

func (x *Networks_NetIP_PortRange) ProtoReflect() protoreflect.Message {
	mi := &file_common_ip_transport_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Networks_NetIP_PortRange.ProtoReflect.Descriptor instead.
func (*Networks_NetIP_PortRange) Descriptor() ([]byte, []int) {
	return file_common_ip_transport_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Networks_NetIP_PortRange) GetFrom() uint32 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *Networks_NetIP_PortRange) GetTo() uint32 {
	if x != nil {
		return x.To
	}
	return 0
}

var File_common_ip_transport_proto protoreflect.FileDescriptor

var file_common_ip_transport_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x69, 0x70, 0x2d, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x08, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x1a, 0x6b, 0x0a, 0x05, 0x4e, 0x65, 0x74, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x49,
	0x44, 0x52, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x49, 0x44, 0x52, 0x1a, 0x2f,
	0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x74, 0x6f, 0x22,
	0x1d, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x07, 0x0a, 0x03,
	0x54, 0x43, 0x50, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x50, 0x10, 0x01, 0x42, 0x2e,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x48, 0x2d, 0x42,
	0x46, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_ip_transport_proto_rawDescOnce sync.Once
	file_common_ip_transport_proto_rawDescData = file_common_ip_transport_proto_rawDesc
)

func file_common_ip_transport_proto_rawDescGZIP() []byte {
	file_common_ip_transport_proto_rawDescOnce.Do(func() {
		file_common_ip_transport_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_ip_transport_proto_rawDescData)
	})
	return file_common_ip_transport_proto_rawDescData
}

var file_common_ip_transport_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_ip_transport_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_ip_transport_proto_goTypes = []interface{}{
	(Networks_NetIP_Transport)(0),    // 0: common.Networks.NetIP.Transport
	(*Networks)(nil),                 // 1: common.Networks
	(*Networks_NetIP)(nil),           // 2: common.Networks.NetIP
	(*Networks_NetIP_PortRange)(nil), // 3: common.Networks.NetIP.PortRange
}
var file_common_ip_transport_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_ip_transport_proto_init() }
func file_common_ip_transport_proto_init() {
	if File_common_ip_transport_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_ip_transport_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Networks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_ip_transport_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Networks_NetIP); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_ip_transport_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Networks_NetIP_PortRange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_ip_transport_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_ip_transport_proto_goTypes,
		DependencyIndexes: file_common_ip_transport_proto_depIdxs,
		EnumInfos:         file_common_ip_transport_proto_enumTypes,
		MessageInfos:      file_common_ip_transport_proto_msgTypes,
	}.Build()
	File_common_ip_transport_proto = out.File
	file_common_ip_transport_proto_rawDesc = nil
	file_common_ip_transport_proto_goTypes = nil
	file_common_ip_transport_proto_depIdxs = nil
}
