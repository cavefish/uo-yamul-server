// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: common.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ObjectDirection int32

const (
	ObjectDirection_north   ObjectDirection = 0
	ObjectDirection_right   ObjectDirection = 1
	ObjectDirection_east    ObjectDirection = 2
	ObjectDirection_down    ObjectDirection = 3
	ObjectDirection_south   ObjectDirection = 4
	ObjectDirection_left    ObjectDirection = 5
	ObjectDirection_west    ObjectDirection = 6
	ObjectDirection_up      ObjectDirection = 7
	ObjectDirection_mask    ObjectDirection = 112
	ObjectDirection_running ObjectDirection = 128
	ObjectDirection_none    ObjectDirection = 237
)

// Enum value maps for ObjectDirection.
var (
	ObjectDirection_name = map[int32]string{
		0:   "north",
		1:   "right",
		2:   "east",
		3:   "down",
		4:   "south",
		5:   "left",
		6:   "west",
		7:   "up",
		112: "mask",
		128: "running",
		237: "none",
	}
	ObjectDirection_value = map[string]int32{
		"north":   0,
		"right":   1,
		"east":    2,
		"down":    3,
		"south":   4,
		"left":    5,
		"west":    6,
		"up":      7,
		"mask":    112,
		"running": 128,
		"none":    237,
	}
)

func (x ObjectDirection) Enum() *ObjectDirection {
	p := new(ObjectDirection)
	*p = x
	return p
}

func (x ObjectDirection) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectDirection) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (ObjectDirection) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x ObjectDirection) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectDirection.Descriptor instead.
func (ObjectDirection) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ObjectId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ObjectId) Reset() {
	*x = ObjectId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectId) ProtoMessage() {}

func (x *ObjectId) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectId.ProtoReflect.Descriptor instead.
func (*ObjectId) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{2}
}

func (x *ObjectId) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Coordinate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	XLoc uint32 `protobuf:"varint,5,opt,name=xLoc,proto3" json:"xLoc,omitempty"`
	YLoc uint32 `protobuf:"varint,6,opt,name=yLoc,proto3" json:"yLoc,omitempty"`
	ZLoc int32  `protobuf:"varint,9,opt,name=zLoc,proto3" json:"zLoc,omitempty"`
}

func (x *Coordinate) Reset() {
	*x = Coordinate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Coordinate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Coordinate) ProtoMessage() {}

func (x *Coordinate) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Coordinate.ProtoReflect.Descriptor instead.
func (*Coordinate) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{3}
}

func (x *Coordinate) GetXLoc() uint32 {
	if x != nil {
		return x.XLoc
	}
	return 0
}

func (x *Coordinate) GetYLoc() uint32 {
	if x != nil {
		return x.YLoc
	}
	return 0
}

func (x *Coordinate) GetZLoc() int32 {
	if x != nil {
		return x.ZLoc
	}
	return 0
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x1c, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20, 0x0a,
	0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x48, 0x0a, 0x0a, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x78, 0x4c, 0x6f, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x78, 0x4c, 0x6f,
	0x63, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x4c, 0x6f, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x79, 0x4c, 0x6f, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x7a, 0x4c, 0x6f, 0x63, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x7a, 0x4c, 0x6f, 0x63, 0x2a, 0x85, 0x01, 0x0a, 0x0f, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x09, 0x0a,
	0x05, 0x6e, 0x6f, 0x72, 0x74, 0x68, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x65, 0x61, 0x73, 0x74, 0x10, 0x02, 0x12, 0x08, 0x0a,
	0x04, 0x64, 0x6f, 0x77, 0x6e, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x74, 0x68,
	0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04,
	0x77, 0x65, 0x73, 0x74, 0x10, 0x06, 0x12, 0x06, 0x0a, 0x02, 0x75, 0x70, 0x10, 0x07, 0x12, 0x08,
	0x0a, 0x04, 0x6d, 0x61, 0x73, 0x6b, 0x10, 0x70, 0x12, 0x0c, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x10, 0x80, 0x01, 0x12, 0x09, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x10, 0xed,
	0x01, 0x42, 0x43, 0x0a, 0x25, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x61, 0x76, 0x65, 0x66, 0x69, 0x73,
	0x68, 0x2e, 0x79, 0x61, 0x6d, 0x75, 0x6c, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x50, 0x01, 0x5a, 0x12, 0x2e, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x88, 0x01, 0x01, 0xa0, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_proto_goTypes = []interface{}{
	(ObjectDirection)(0), // 0: common.ObjectDirection
	(*Empty)(nil),        // 1: common.Empty
	(*UUID)(nil),         // 2: common.UUID
	(*ObjectId)(nil),     // 3: common.ObjectId
	(*Coordinate)(nil),   // 4: common.Coordinate
}
var file_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
		file_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectId); i {
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
		file_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Coordinate); i {
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
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
