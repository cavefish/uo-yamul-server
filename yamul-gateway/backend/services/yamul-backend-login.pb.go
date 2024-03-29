// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: yamul-backend-login.proto

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

type LoginResponse_LoginResponseValue int32

const (
	LoginResponse_undefined LoginResponse_LoginResponseValue = 0
	LoginResponse_valid     LoginResponse_LoginResponseValue = 1
	LoginResponse_invalid   LoginResponse_LoginResponseValue = 2
)

// Enum value maps for LoginResponse_LoginResponseValue.
var (
	LoginResponse_LoginResponseValue_name = map[int32]string{
		0: "undefined",
		1: "valid",
		2: "invalid",
	}
	LoginResponse_LoginResponseValue_value = map[string]int32{
		"undefined": 0,
		"valid":     1,
		"invalid":   2,
	}
)

func (x LoginResponse_LoginResponseValue) Enum() *LoginResponse_LoginResponseValue {
	p := new(LoginResponse_LoginResponseValue)
	*p = x
	return p
}

func (x LoginResponse_LoginResponseValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginResponse_LoginResponseValue) Descriptor() protoreflect.EnumDescriptor {
	return file_yamul_backend_login_proto_enumTypes[0].Descriptor()
}

func (LoginResponse_LoginResponseValue) Type() protoreflect.EnumType {
	return &file_yamul_backend_login_proto_enumTypes[0]
}

func (x LoginResponse_LoginResponseValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginResponse_LoginResponseValue.Descriptor instead.
func (LoginResponse_LoginResponseValue) EnumDescriptor() ([]byte, []int) {
	return file_yamul_backend_login_proto_rawDescGZIP(), []int{1, 0}
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_yamul_backend_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value LoginResponse_LoginResponseValue `protobuf:"varint,1,opt,name=value,proto3,enum=login.LoginResponse_LoginResponseValue" json:"value,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_yamul_backend_login_proto_rawDescGZIP(), []int{1}
}

func (x *LoginResponse) GetValue() LoginResponse_LoginResponseValue {
	if x != nil {
		return x.Value
	}
	return LoginResponse_undefined
}

var File_yamul_backend_login_proto protoreflect.FileDescriptor

var file_yamul_backend_login_proto_rawDesc = []byte{
	0x0a, 0x19, 0x79, 0x61, 0x6d, 0x75, 0x6c, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x22, 0x46, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x0d, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3b, 0x0a, 0x12, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x0d, 0x0a, 0x09, 0x75, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x69,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x02, 0x32, 0x4a, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x42, 0x0a, 0x24, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x61, 0x76, 0x65,
	0x66, 0x69, 0x73, 0x68, 0x2e, 0x79, 0x61, 0x6d, 0x75, 0x6c, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x50, 0x01, 0x5a, 0x12,
	0x2e, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yamul_backend_login_proto_rawDescOnce sync.Once
	file_yamul_backend_login_proto_rawDescData = file_yamul_backend_login_proto_rawDesc
)

func file_yamul_backend_login_proto_rawDescGZIP() []byte {
	file_yamul_backend_login_proto_rawDescOnce.Do(func() {
		file_yamul_backend_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_yamul_backend_login_proto_rawDescData)
	})
	return file_yamul_backend_login_proto_rawDescData
}

var file_yamul_backend_login_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yamul_backend_login_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yamul_backend_login_proto_goTypes = []interface{}{
	(LoginResponse_LoginResponseValue)(0), // 0: login.LoginResponse.LoginResponseValue
	(*LoginRequest)(nil),                  // 1: login.LoginRequest
	(*LoginResponse)(nil),                 // 2: login.LoginResponse
}
var file_yamul_backend_login_proto_depIdxs = []int32{
	0, // 0: login.LoginResponse.value:type_name -> login.LoginResponse.LoginResponseValue
	1, // 1: login.LoginService.validateLogin:input_type -> login.LoginRequest
	2, // 2: login.LoginService.validateLogin:output_type -> login.LoginResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yamul_backend_login_proto_init() }
func file_yamul_backend_login_proto_init() {
	if File_yamul_backend_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yamul_backend_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_yamul_backend_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResponse); i {
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
			RawDescriptor: file_yamul_backend_login_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yamul_backend_login_proto_goTypes,
		DependencyIndexes: file_yamul_backend_login_proto_depIdxs,
		EnumInfos:         file_yamul_backend_login_proto_enumTypes,
		MessageInfos:      file_yamul_backend_login_proto_msgTypes,
	}.Build()
	File_yamul_backend_login_proto = out.File
	file_yamul_backend_login_proto_rawDesc = nil
	file_yamul_backend_login_proto_goTypes = nil
	file_yamul_backend_login_proto_depIdxs = nil
}
