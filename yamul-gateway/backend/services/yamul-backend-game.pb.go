// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: yamul-backend-game.proto

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

type MsgType int32

const (
	MsgType_TypeUndefined          MsgType = 0 // TODO this is a valid value on UO packet ids
	MsgType_TypeHealthBar          MsgType = 5888
	MsgType_TypeCharacterSelection MsgType = 23808
	MsgType_TypePlayMusic          MsgType = 27904
	MsgType_TypeCreateCharacter    MsgType = 30720
	MsgType_TypeMapChange          MsgType = 48904
	MsgType_TypeApplyWorldPatches  MsgType = 48920
)

// Enum value maps for MsgType.
var (
	MsgType_name = map[int32]string{
		0:     "TypeUndefined",
		5888:  "TypeHealthBar",
		23808: "TypeCharacterSelection",
		27904: "TypePlayMusic",
		30720: "TypeCreateCharacter",
		48904: "TypeMapChange",
		48920: "TypeApplyWorldPatches",
	}
	MsgType_value = map[string]int32{
		"TypeUndefined":          0,
		"TypeHealthBar":          5888,
		"TypeCharacterSelection": 23808,
		"TypePlayMusic":          27904,
		"TypeCreateCharacter":    30720,
		"TypeMapChange":          48904,
		"TypeApplyWorldPatches":  48920,
	}
)

func (x MsgType) Enum() *MsgType {
	p := new(MsgType)
	*p = x
	return p
}

func (x MsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_yamul_backend_game_proto_enumTypes[0].Descriptor()
}

func (MsgType) Type() protoreflect.EnumType {
	return &file_yamul_backend_game_proto_enumTypes[0]
}

func (x MsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgType.Descriptor instead.
func (MsgType) EnumDescriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{0}
}

type MsgHealthBar_Values_Type int32

const (
	MsgHealthBar_Values_GREEN  MsgHealthBar_Values_Type = 0
	MsgHealthBar_Values_YELLOW MsgHealthBar_Values_Type = 1
)

// Enum value maps for MsgHealthBar_Values_Type.
var (
	MsgHealthBar_Values_Type_name = map[int32]string{
		0: "GREEN",
		1: "YELLOW",
	}
	MsgHealthBar_Values_Type_value = map[string]int32{
		"GREEN":  0,
		"YELLOW": 1,
	}
)

func (x MsgHealthBar_Values_Type) Enum() *MsgHealthBar_Values_Type {
	p := new(MsgHealthBar_Values_Type)
	*p = x
	return p
}

func (x MsgHealthBar_Values_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgHealthBar_Values_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_yamul_backend_game_proto_enumTypes[1].Descriptor()
}

func (MsgHealthBar_Values_Type) Type() protoreflect.EnumType {
	return &file_yamul_backend_game_proto_enumTypes[1]
}

func (x MsgHealthBar_Values_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgHealthBar_Values_Type.Descriptor instead.
func (MsgHealthBar_Values_Type) EnumDescriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{3, 0, 0}
}

type MsgApplyWorldPatches struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapCount int32 `protobuf:"varint,1,opt,name=mapCount,proto3" json:"mapCount,omitempty"` // TODO add missing fields
}

func (x *MsgApplyWorldPatches) Reset() {
	*x = MsgApplyWorldPatches{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgApplyWorldPatches) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgApplyWorldPatches) ProtoMessage() {}

func (x *MsgApplyWorldPatches) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgApplyWorldPatches.ProtoReflect.Descriptor instead.
func (*MsgApplyWorldPatches) Descriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{0}
}

func (x *MsgApplyWorldPatches) GetMapCount() int32 {
	if x != nil {
		return x.MapCount
	}
	return 0
}

type MsgCharacterSelection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Slot     int32  `protobuf:"varint,3,opt,name=slot,proto3" json:"slot,omitempty"`
}

func (x *MsgCharacterSelection) Reset() {
	*x = MsgCharacterSelection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCharacterSelection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCharacterSelection) ProtoMessage() {}

func (x *MsgCharacterSelection) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCharacterSelection.ProtoReflect.Descriptor instead.
func (*MsgCharacterSelection) Descriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{1}
}

func (x *MsgCharacterSelection) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MsgCharacterSelection) GetSlot() int32 {
	if x != nil {
		return x.Slot
	}
	return 0
}

type MsgCreateCharacter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *ObjectId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // TODO add missing fields
}

func (x *MsgCreateCharacter) Reset() {
	*x = MsgCreateCharacter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCreateCharacter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCreateCharacter) ProtoMessage() {}

func (x *MsgCreateCharacter) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCreateCharacter.ProtoReflect.Descriptor instead.
func (*MsgCreateCharacter) Descriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{2}
}

func (x *MsgCreateCharacter) GetId() *ObjectId {
	if x != nil {
		return x.Id
	}
	return nil
}

type MsgHealthBar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     *ObjectId              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Values []*MsgHealthBar_Values `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *MsgHealthBar) Reset() {
	*x = MsgHealthBar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgHealthBar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgHealthBar) ProtoMessage() {}

func (x *MsgHealthBar) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgHealthBar.ProtoReflect.Descriptor instead.
func (*MsgHealthBar) Descriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{3}
}

func (x *MsgHealthBar) GetId() *ObjectId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *MsgHealthBar) GetValues() []*MsgHealthBar_Values {
	if x != nil {
		return x.Values
	}
	return nil
}

type MsgMapChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapId int32 `protobuf:"varint,1,opt,name=mapId,proto3" json:"mapId,omitempty"`
}

func (x *MsgMapChange) Reset() {
	*x = MsgMapChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgMapChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgMapChange) ProtoMessage() {}

func (x *MsgMapChange) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgMapChange.ProtoReflect.Descriptor instead.
func (*MsgMapChange) Descriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{4}
}

func (x *MsgMapChange) GetMapId() int32 {
	if x != nil {
		return x.MapId
	}
	return 0
}

type MsgPlayMusic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MusicId int32 `protobuf:"varint,1,opt,name=musicId,proto3" json:"musicId,omitempty"`
}

func (x *MsgPlayMusic) Reset() {
	*x = MsgPlayMusic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgPlayMusic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgPlayMusic) ProtoMessage() {}

func (x *MsgPlayMusic) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgPlayMusic.ProtoReflect.Descriptor instead.
func (*MsgPlayMusic) Descriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{5}
}

func (x *MsgPlayMusic) GetMusicId() int32 {
	if x != nil {
		return x.MusicId
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*Message_HealthBar
	//	*Message_CharacterSelection
	//	*Message_PlayMusic
	//	*Message_CreateCharacter
	//	*Message_MapChange
	//	*Message_ApplyWorldPatches
	Msg isMessage_Msg `protobuf_oneof:"msg"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_game_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_game_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{6}
}

func (m *Message) GetMsg() isMessage_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *Message) GetHealthBar() *MsgHealthBar {
	if x, ok := x.GetMsg().(*Message_HealthBar); ok {
		return x.HealthBar
	}
	return nil
}

func (x *Message) GetCharacterSelection() *MsgCharacterSelection {
	if x, ok := x.GetMsg().(*Message_CharacterSelection); ok {
		return x.CharacterSelection
	}
	return nil
}

func (x *Message) GetPlayMusic() *MsgPlayMusic {
	if x, ok := x.GetMsg().(*Message_PlayMusic); ok {
		return x.PlayMusic
	}
	return nil
}

func (x *Message) GetCreateCharacter() *MsgCreateCharacter {
	if x, ok := x.GetMsg().(*Message_CreateCharacter); ok {
		return x.CreateCharacter
	}
	return nil
}

func (x *Message) GetMapChange() *MsgMapChange {
	if x, ok := x.GetMsg().(*Message_MapChange); ok {
		return x.MapChange
	}
	return nil
}

func (x *Message) GetApplyWorldPatches() *MsgApplyWorldPatches {
	if x, ok := x.GetMsg().(*Message_ApplyWorldPatches); ok {
		return x.ApplyWorldPatches
	}
	return nil
}

type isMessage_Msg interface {
	isMessage_Msg()
}

type Message_HealthBar struct {
	HealthBar *MsgHealthBar `protobuf:"bytes,5888,opt,name=healthBar,proto3,oneof"`
}

type Message_CharacterSelection struct {
	CharacterSelection *MsgCharacterSelection `protobuf:"bytes,23808,opt,name=characterSelection,proto3,oneof"`
}

type Message_PlayMusic struct {
	PlayMusic *MsgPlayMusic `protobuf:"bytes,27904,opt,name=playMusic,proto3,oneof"`
}

type Message_CreateCharacter struct {
	CreateCharacter *MsgCreateCharacter `protobuf:"bytes,30720,opt,name=createCharacter,proto3,oneof"`
}

type Message_MapChange struct {
	MapChange *MsgMapChange `protobuf:"bytes,48904,opt,name=mapChange,proto3,oneof"`
}

type Message_ApplyWorldPatches struct {
	ApplyWorldPatches *MsgApplyWorldPatches `protobuf:"bytes,48920,opt,name=applyWorldPatches,proto3,oneof"`
}

func (*Message_HealthBar) isMessage_Msg() {}

func (*Message_CharacterSelection) isMessage_Msg() {}

func (*Message_PlayMusic) isMessage_Msg() {}

func (*Message_CreateCharacter) isMessage_Msg() {}

func (*Message_MapChange) isMessage_Msg() {}

func (*Message_ApplyWorldPatches) isMessage_Msg() {}

type StreamPackage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type MsgType  `protobuf:"varint,1,opt,name=type,proto3,enum=game.MsgType" json:"type,omitempty"`
	Body *Message `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *StreamPackage) Reset() {
	*x = StreamPackage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_game_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamPackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamPackage) ProtoMessage() {}

func (x *StreamPackage) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_game_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamPackage.ProtoReflect.Descriptor instead.
func (*StreamPackage) Descriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{7}
}

func (x *StreamPackage) GetType() MsgType {
	if x != nil {
		return x.Type
	}
	return MsgType_TypeUndefined
}

func (x *StreamPackage) GetBody() *Message {
	if x != nil {
		return x.Body
	}
	return nil
}

type MsgHealthBar_Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    MsgHealthBar_Values_Type `protobuf:"varint,1,opt,name=type,proto3,enum=game.MsgHealthBar_Values_Type" json:"type,omitempty"`
	Enabled bool                     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *MsgHealthBar_Values) Reset() {
	*x = MsgHealthBar_Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yamul_backend_game_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgHealthBar_Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgHealthBar_Values) ProtoMessage() {}

func (x *MsgHealthBar_Values) ProtoReflect() protoreflect.Message {
	mi := &file_yamul_backend_game_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgHealthBar_Values.ProtoReflect.Descriptor instead.
func (*MsgHealthBar_Values) Descriptor() ([]byte, []int) {
	return file_yamul_backend_game_proto_rawDescGZIP(), []int{3, 0}
}

func (x *MsgHealthBar_Values) GetType() MsgHealthBar_Values_Type {
	if x != nil {
		return x.Type
	}
	return MsgHealthBar_Values_GREEN
}

func (x *MsgHealthBar_Values) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

var File_yamul_backend_game_proto protoreflect.FileDescriptor

var file_yamul_backend_game_proto_rawDesc = []byte{
	0x0a, 0x18, 0x79, 0x61, 0x6d, 0x75, 0x6c, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x61, 0x6d, 0x65,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32,
	0x0a, 0x14, 0x4d, 0x73, 0x67, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x70, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61, 0x70, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0x47, 0x0a, 0x15, 0x4d, 0x73, 0x67, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x22, 0x36, 0x0a, 0x12, 0x4d,
	0x73, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xda, 0x01, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x42, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x73,
	0x67, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x42, 0x61, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x75, 0x0a, 0x06, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x42, 0x61, 0x72, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x22, 0x1d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x52, 0x45,
	0x45, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x59, 0x45, 0x4c, 0x4c, 0x4f, 0x57, 0x10, 0x01,
	0x22, 0x24, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x4d, 0x61, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6d, 0x61, 0x70, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x50, 0x6c, 0x61,
	0x79, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x49, 0x64,
	0x22, 0x98, 0x03, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x33, 0x0a, 0x09,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x42, 0x61, 0x72, 0x18, 0x80, 0x2e, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x42, 0x61, 0x72, 0x48, 0x00, 0x52, 0x09, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x42, 0x61,
	0x72, 0x12, 0x4f, 0x0a, 0x12, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x80, 0xba, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63,
	0x74, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x12,
	0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x18,
	0x80, 0xda, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d,
	0x73, 0x67, 0x50, 0x6c, 0x61, 0x79, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x48, 0x00, 0x52, 0x09, 0x70,
	0x6c, 0x61, 0x79, 0x4d, 0x75, 0x73, 0x69, 0x63, 0x12, 0x46, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x18, 0x80, 0xf0, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x12, 0x34, 0x0a, 0x09, 0x6d, 0x61, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x88, 0xfe,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x73, 0x67,
	0x4d, 0x61, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x48, 0x00, 0x52, 0x09, 0x6d, 0x61, 0x70,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x4c, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x18, 0x98, 0xfe, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x48,
	0x00, 0x52, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x73, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x55, 0x0a, 0x0d, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x67, 0x61, 0x6d,
	0x65, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x67, 0x61, 0x6d, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x2a, 0xb0, 0x01, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11,
	0x0a, 0x0d, 0x54, 0x79, 0x70, 0x65, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0d, 0x54, 0x79, 0x70, 0x65, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x42,
	0x61, 0x72, 0x10, 0x80, 0x2e, 0x12, 0x1c, 0x0a, 0x16, 0x54, 0x79, 0x70, 0x65, 0x43, 0x68, 0x61,
	0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10,
	0x80, 0xba, 0x01, 0x12, 0x13, 0x0a, 0x0d, 0x54, 0x79, 0x70, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x4d,
	0x75, 0x73, 0x69, 0x63, 0x10, 0x80, 0xda, 0x01, 0x12, 0x19, 0x0a, 0x13, 0x54, 0x79, 0x70, 0x65,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x10,
	0x80, 0xf0, 0x01, 0x12, 0x13, 0x0a, 0x0d, 0x54, 0x79, 0x70, 0x65, 0x4d, 0x61, 0x70, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x10, 0x88, 0xfe, 0x02, 0x12, 0x1b, 0x0a, 0x15, 0x54, 0x79, 0x70, 0x65,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x50, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x73, 0x10, 0x98, 0xfe, 0x02, 0x32, 0x4d, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x6e, 0x47, 0x61, 0x6d, 0x65,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x13, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x67, 0x61,
	0x6d, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x41, 0x0a, 0x23, 0x64, 0x65, 0x76, 0x2e, 0x63, 0x61, 0x76, 0x65,
	0x66, 0x69, 0x73, 0x68, 0x2e, 0x79, 0x61, 0x6d, 0x75, 0x6c, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x50, 0x01, 0x5a, 0x12, 0x2e,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yamul_backend_game_proto_rawDescOnce sync.Once
	file_yamul_backend_game_proto_rawDescData = file_yamul_backend_game_proto_rawDesc
)

func file_yamul_backend_game_proto_rawDescGZIP() []byte {
	file_yamul_backend_game_proto_rawDescOnce.Do(func() {
		file_yamul_backend_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_yamul_backend_game_proto_rawDescData)
	})
	return file_yamul_backend_game_proto_rawDescData
}

var file_yamul_backend_game_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yamul_backend_game_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_yamul_backend_game_proto_goTypes = []interface{}{
	(MsgType)(0),                  // 0: game.MsgType
	(MsgHealthBar_Values_Type)(0), // 1: game.MsgHealthBar.Values.Type
	(*MsgApplyWorldPatches)(nil),  // 2: game.MsgApplyWorldPatches
	(*MsgCharacterSelection)(nil), // 3: game.MsgCharacterSelection
	(*MsgCreateCharacter)(nil),    // 4: game.MsgCreateCharacter
	(*MsgHealthBar)(nil),          // 5: game.MsgHealthBar
	(*MsgMapChange)(nil),          // 6: game.MsgMapChange
	(*MsgPlayMusic)(nil),          // 7: game.MsgPlayMusic
	(*Message)(nil),               // 8: game.Message
	(*StreamPackage)(nil),         // 9: game.StreamPackage
	(*MsgHealthBar_Values)(nil),   // 10: game.MsgHealthBar.Values
	(*ObjectId)(nil),              // 11: common.ObjectId
}
var file_yamul_backend_game_proto_depIdxs = []int32{
	11, // 0: game.MsgCreateCharacter.id:type_name -> common.ObjectId
	11, // 1: game.MsgHealthBar.id:type_name -> common.ObjectId
	10, // 2: game.MsgHealthBar.values:type_name -> game.MsgHealthBar.Values
	5,  // 3: game.Message.healthBar:type_name -> game.MsgHealthBar
	3,  // 4: game.Message.characterSelection:type_name -> game.MsgCharacterSelection
	7,  // 5: game.Message.playMusic:type_name -> game.MsgPlayMusic
	4,  // 6: game.Message.createCharacter:type_name -> game.MsgCreateCharacter
	6,  // 7: game.Message.mapChange:type_name -> game.MsgMapChange
	2,  // 8: game.Message.applyWorldPatches:type_name -> game.MsgApplyWorldPatches
	0,  // 9: game.StreamPackage.type:type_name -> game.MsgType
	8,  // 10: game.StreamPackage.body:type_name -> game.Message
	1,  // 11: game.MsgHealthBar.Values.type:type_name -> game.MsgHealthBar.Values.Type
	9,  // 12: game.GameService.openGameStream:input_type -> game.StreamPackage
	9,  // 13: game.GameService.openGameStream:output_type -> game.StreamPackage
	13, // [13:14] is the sub-list for method output_type
	12, // [12:13] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_yamul_backend_game_proto_init() }
func file_yamul_backend_game_proto_init() {
	if File_yamul_backend_game_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yamul_backend_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgApplyWorldPatches); i {
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
		file_yamul_backend_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCharacterSelection); i {
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
		file_yamul_backend_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCreateCharacter); i {
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
		file_yamul_backend_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgHealthBar); i {
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
		file_yamul_backend_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgMapChange); i {
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
		file_yamul_backend_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgPlayMusic); i {
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
		file_yamul_backend_game_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_yamul_backend_game_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamPackage); i {
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
		file_yamul_backend_game_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgHealthBar_Values); i {
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
	file_yamul_backend_game_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Message_HealthBar)(nil),
		(*Message_CharacterSelection)(nil),
		(*Message_PlayMusic)(nil),
		(*Message_CreateCharacter)(nil),
		(*Message_MapChange)(nil),
		(*Message_ApplyWorldPatches)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yamul_backend_game_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yamul_backend_game_proto_goTypes,
		DependencyIndexes: file_yamul_backend_game_proto_depIdxs,
		EnumInfos:         file_yamul_backend_game_proto_enumTypes,
		MessageInfos:      file_yamul_backend_game_proto_msgTypes,
	}.Build()
	File_yamul_backend_game_proto = out.File
	file_yamul_backend_game_proto_rawDesc = nil
	file_yamul_backend_game_proto_goTypes = nil
	file_yamul_backend_game_proto_depIdxs = nil
}