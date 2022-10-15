// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.2
// source: message.proto

package message

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

// cd common/idl; protoc -I message  --go_out=plugins=grpc:message  message/message.proto
type CmdType int32

const (
	CmdType_Login     CmdType = 0
	CmdType_Heartbeat CmdType = 1
	CmdType_ReConn    CmdType = 2
	CmdType_ACK       CmdType = 3
)

// Enum value maps for CmdType.
var (
	CmdType_name = map[int32]string{
		0: "Login",
		1: "Heartbeat",
		2: "ReConn",
		3: "ACK",
	}
	CmdType_value = map[string]int32{
		"Login":     0,
		"Heartbeat": 1,
		"ReConn":    2,
		"ACK":       3,
	}
)

func (x CmdType) Enum() *CmdType {
	p := new(CmdType)
	*p = x
	return p
}

func (x CmdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CmdType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (CmdType) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x CmdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CmdType.Descriptor instead.
func (CmdType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

// 顶层cmd pb结构
type MsgCmd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    CmdType `protobuf:"varint,1,opt,name=Type,proto3,enum=message.CmdType" json:"Type,omitempty"`
	Payload []byte  `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (x *MsgCmd) Reset() {
	*x = MsgCmd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgCmd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgCmd) ProtoMessage() {}

func (x *MsgCmd) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgCmd.ProtoReflect.Descriptor instead.
func (*MsgCmd) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *MsgCmd) GetType() CmdType {
	if x != nil {
		return x.Type
	}
	return CmdType_Login
}

func (x *MsgCmd) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// ACK 消息
type ACKMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   uint32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Msg    string  `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	Type   CmdType `protobuf:"varint,3,opt,name=Type,proto3,enum=message.CmdType" json:"Type,omitempty"`
	ConnID uint64  `protobuf:"varint,4,opt,name=ConnID,proto3" json:"ConnID,omitempty"`
}

func (x *ACKMsg) Reset() {
	*x = ACKMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ACKMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ACKMsg) ProtoMessage() {}

func (x *ACKMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ACKMsg.ProtoReflect.Descriptor instead.
func (*ACKMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *ACKMsg) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ACKMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ACKMsg) GetType() CmdType {
	if x != nil {
		return x.Type
	}
	return CmdType_Login
}

func (x *ACKMsg) GetConnID() uint64 {
	if x != nil {
		return x.ConnID
	}
	return 0
}

// 登陆消息
type LoginMsgHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceID uint64 `protobuf:"varint,1,opt,name=DeviceID,proto3" json:"DeviceID,omitempty"`
}

func (x *LoginMsgHead) Reset() {
	*x = LoginMsgHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMsgHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMsgHead) ProtoMessage() {}

func (x *LoginMsgHead) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMsgHead.ProtoReflect.Descriptor instead.
func (*LoginMsgHead) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *LoginMsgHead) GetDeviceID() uint64 {
	if x != nil {
		return x.DeviceID
	}
	return 0
}

type LoginMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head         *LoginMsgHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	LoginMsgBody []byte        `protobuf:"bytes,2,opt,name=LoginMsgBody,proto3" json:"LoginMsgBody,omitempty"`
}

func (x *LoginMsg) Reset() {
	*x = LoginMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginMsg) ProtoMessage() {}

func (x *LoginMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginMsg.ProtoReflect.Descriptor instead.
func (*LoginMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *LoginMsg) GetHead() *LoginMsgHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *LoginMsg) GetLoginMsgBody() []byte {
	if x != nil {
		return x.LoginMsgBody
	}
	return nil
}

// 心跳消息
type HeartbeatMsgHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HeartbeatMsgHead) Reset() {
	*x = HeartbeatMsgHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatMsgHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatMsgHead) ProtoMessage() {}

func (x *HeartbeatMsgHead) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatMsgHead.ProtoReflect.Descriptor instead.
func (*HeartbeatMsgHead) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

type HeartbeatMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head             *HeartbeatMsgHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	HeartbeatMsgBody []byte            `protobuf:"bytes,2,opt,name=HeartbeatMsgBody,proto3" json:"HeartbeatMsgBody,omitempty"`
}

func (x *HeartbeatMsg) Reset() {
	*x = HeartbeatMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatMsg) ProtoMessage() {}

func (x *HeartbeatMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatMsg.ProtoReflect.Descriptor instead.
func (*HeartbeatMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *HeartbeatMsg) GetHead() *HeartbeatMsgHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *HeartbeatMsg) GetHeartbeatMsgBody() []byte {
	if x != nil {
		return x.HeartbeatMsgBody
	}
	return nil
}

// 重连消息
type ReConnMsgHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnID uint64 `protobuf:"varint,1,opt,name=ConnID,proto3" json:"ConnID,omitempty"`
}

func (x *ReConnMsgHead) Reset() {
	*x = ReConnMsgHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReConnMsgHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReConnMsgHead) ProtoMessage() {}

func (x *ReConnMsgHead) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReConnMsgHead.ProtoReflect.Descriptor instead.
func (*ReConnMsgHead) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (x *ReConnMsgHead) GetConnID() uint64 {
	if x != nil {
		return x.ConnID
	}
	return 0
}

type ReConnMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Head          *ReConnMsgHead `protobuf:"bytes,1,opt,name=Head,proto3" json:"Head,omitempty"`
	ReConnMsgBody []byte         `protobuf:"bytes,2,opt,name=ReConnMsgBody,proto3" json:"ReConnMsgBody,omitempty"`
}

func (x *ReConnMsg) Reset() {
	*x = ReConnMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReConnMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReConnMsg) ProtoMessage() {}

func (x *ReConnMsg) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReConnMsg.ProtoReflect.Descriptor instead.
func (*ReConnMsg) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *ReConnMsg) GetHead() *ReConnMsgHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *ReConnMsg) GetReConnMsgBody() []byte {
	if x != nil {
		return x.ReConnMsgBody
	}
	return nil
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x48, 0x0a, 0x06, 0x4d, 0x73, 0x67, 0x43,
	0x6d, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x6c, 0x0a, 0x06, 0x41, 0x43, 0x4b, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x24, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x6e,
	0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x44,
	0x22, 0x2a, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x22, 0x59, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x29, 0x0a, 0x04, 0x48, 0x65, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x48,
	0x65, 0x61, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x4d, 0x73, 0x67, 0x42,
	0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x12, 0x0a, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x22, 0x69, 0x0a, 0x0c, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x2d, 0x0a, 0x04, 0x48,
	0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x73, 0x67,
	0x48, 0x65, 0x61, 0x64, 0x52, 0x04, 0x48, 0x65, 0x61, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x4d,
	0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x27, 0x0a, 0x0d, 0x52, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x4d, 0x73, 0x67, 0x48, 0x65, 0x61, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x6e, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x43, 0x6f, 0x6e, 0x6e, 0x49, 0x44, 0x22,
	0x5d, 0x0a, 0x09, 0x52, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x04,
	0x48, 0x65, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x73, 0x67, 0x48, 0x65,
	0x61, 0x64, 0x52, 0x04, 0x48, 0x65, 0x61, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x52, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0d, 0x52, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x2a, 0x38,
	0x0a, 0x07, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61,
	0x74, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x10, 0x02, 0x12,
	0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b, 0x10, 0x03, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_message_proto_goTypes = []interface{}{
	(CmdType)(0),             // 0: message.CmdType
	(*MsgCmd)(nil),           // 1: message.MsgCmd
	(*ACKMsg)(nil),           // 2: message.ACKMsg
	(*LoginMsgHead)(nil),     // 3: message.LoginMsgHead
	(*LoginMsg)(nil),         // 4: message.LoginMsg
	(*HeartbeatMsgHead)(nil), // 5: message.HeartbeatMsgHead
	(*HeartbeatMsg)(nil),     // 6: message.HeartbeatMsg
	(*ReConnMsgHead)(nil),    // 7: message.ReConnMsgHead
	(*ReConnMsg)(nil),        // 8: message.ReConnMsg
}
var file_message_proto_depIdxs = []int32{
	0, // 0: message.MsgCmd.Type:type_name -> message.CmdType
	0, // 1: message.ACKMsg.Type:type_name -> message.CmdType
	3, // 2: message.LoginMsg.Head:type_name -> message.LoginMsgHead
	5, // 3: message.HeartbeatMsg.Head:type_name -> message.HeartbeatMsgHead
	7, // 4: message.ReConnMsg.Head:type_name -> message.ReConnMsgHead
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgCmd); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ACKMsg); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginMsgHead); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginMsg); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatMsgHead); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatMsg); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReConnMsgHead); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReConnMsg); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
