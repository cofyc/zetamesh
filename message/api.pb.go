// Copyright 2020 ZetaMesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: api.proto

package message

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PacketType int32

const (
	PacketType_Heartbeat     PacketType = 0
	PacketType_Relay         PacketType = 1
	PacketType_OpenTunnel    PacketType = 2
	PacketType_OpenTunnelAck PacketType = 3
	PacketType_Ping          PacketType = 4
	PacketType_Pong          PacketType = 5
	PacketType_Data          PacketType = 6
)

// Enum value maps for PacketType.
var (
	PacketType_name = map[int32]string{
		0: "Heartbeat",
		1: "Relay",
		2: "OpenTunnel",
		3: "OpenTunnelAck",
		4: "Ping",
		5: "Pong",
		6: "Data",
	}
	PacketType_value = map[string]int32{
		"Heartbeat":     0,
		"Relay":         1,
		"OpenTunnel":    2,
		"OpenTunnelAck": 3,
		"Ping":          4,
		"Pong":          5,
		"Data":          6,
	}
)

func (x PacketType) Enum() *PacketType {
	p := new(PacketType)
	*p = x
	return p
}

func (x PacketType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PacketType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[0].Descriptor()
}

func (PacketType) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[0]
}

func (x PacketType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PacketType.Descriptor instead.
func (PacketType) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

type StatusCode int32

const (
	StatusCode_Success        StatusCode = 0
	StatusCode_ServerInternal StatusCode = 1
	StatusCode_InvalidVersion StatusCode = 2
	StatusCode_AddConflicted  StatusCode = 3
	StatusCode_VersionTooOld  StatusCode = 4
	StatusCode_KeyNotMatched  StatusCode = 5
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "Success",
		1: "ServerInternal",
		2: "InvalidVersion",
		3: "AddConflicted",
		4: "VersionTooOld",
		5: "KeyNotMatched",
	}
	StatusCode_value = map[string]int32{
		"Success":        0,
		"ServerInternal": 1,
		"InvalidVersion": 2,
		"AddConflicted":  3,
		"VersionTooOld":  4,
		"KeyNotMatched":  5,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_proto_enumTypes[1].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_api_proto_enumTypes[1]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

type CtrlHeartbeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtAddress string `protobuf:"bytes,1,opt,name=virtAddress,proto3" json:"virtAddress,omitempty"`
}

func (x *CtrlHeartbeat) Reset() {
	*x = CtrlHeartbeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtrlHeartbeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtrlHeartbeat) ProtoMessage() {}

func (x *CtrlHeartbeat) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtrlHeartbeat.ProtoReflect.Descriptor instead.
func (*CtrlHeartbeat) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *CtrlHeartbeat) GetVirtAddress() string {
	if x != nil {
		return x.VirtAddress
	}
	return ""
}

type CtrlPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtAddress string `protobuf:"bytes,1,opt,name=virtAddress,proto3" json:"virtAddress,omitempty"`
	Nonce       string `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *CtrlPing) Reset() {
	*x = CtrlPing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtrlPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtrlPing) ProtoMessage() {}

func (x *CtrlPing) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtrlPing.ProtoReflect.Descriptor instead.
func (*CtrlPing) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *CtrlPing) GetVirtAddress() string {
	if x != nil {
		return x.VirtAddress
	}
	return ""
}

func (x *CtrlPing) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type CtrlPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtAddress string `protobuf:"bytes,1,opt,name=virtAddress,proto3" json:"virtAddress,omitempty"`
	Nonce       string `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (x *CtrlPong) Reset() {
	*x = CtrlPong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtrlPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtrlPong) ProtoMessage() {}

func (x *CtrlPong) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtrlPong.ProtoReflect.Descriptor instead.
func (*CtrlPong) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *CtrlPong) GetVirtAddress() string {
	if x != nil {
		return x.VirtAddress
	}
	return ""
}

func (x *CtrlPong) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

type CtrlOpenTunnel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AckId       int64  `protobuf:"varint,1,opt,name=ackId,proto3" json:"ackId,omitempty"`
	VirtAddress string `protobuf:"bytes,2,opt,name=virtAddress,proto3" json:"virtAddress,omitempty"`
	UdpAddress  string `protobuf:"bytes,3,opt,name=udpAddress,proto3" json:"udpAddress,omitempty"`
}

func (x *CtrlOpenTunnel) Reset() {
	*x = CtrlOpenTunnel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtrlOpenTunnel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtrlOpenTunnel) ProtoMessage() {}

func (x *CtrlOpenTunnel) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtrlOpenTunnel.ProtoReflect.Descriptor instead.
func (*CtrlOpenTunnel) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *CtrlOpenTunnel) GetAckId() int64 {
	if x != nil {
		return x.AckId
	}
	return 0
}

func (x *CtrlOpenTunnel) GetVirtAddress() string {
	if x != nil {
		return x.VirtAddress
	}
	return ""
}

func (x *CtrlOpenTunnel) GetUdpAddress() string {
	if x != nil {
		return x.UdpAddress
	}
	return ""
}

type CtrlOpenTunnelAck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AckId int64 `protobuf:"varint,1,opt,name=ackId,proto3" json:"ackId,omitempty"`
}

func (x *CtrlOpenTunnelAck) Reset() {
	*x = CtrlOpenTunnelAck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtrlOpenTunnelAck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtrlOpenTunnelAck) ProtoMessage() {}

func (x *CtrlOpenTunnelAck) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtrlOpenTunnelAck.ProtoReflect.Descriptor instead.
func (*CtrlOpenTunnelAck) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *CtrlOpenTunnelAck) GetAckId() int64 {
	if x != nil {
		return x.AckId
	}
	return 0
}

type CtrlRelay struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtAddress string `protobuf:"bytes,1,opt,name=virtAddress,proto3" json:"virtAddress,omitempty"`
	Data        []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CtrlRelay) Reset() {
	*x = CtrlRelay{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CtrlRelay) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CtrlRelay) ProtoMessage() {}

func (x *CtrlRelay) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CtrlRelay.ProtoReflect.Descriptor instead.
func (*CtrlRelay) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *CtrlRelay) GetVirtAddress() string {
	if x != nil {
		return x.VirtAddress
	}
	return ""
}

func (x *CtrlRelay) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a, 0x0d, 0x43,
	0x74, 0x72, 0x6c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x76, 0x69, 0x72, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x42,
	0x0a, 0x08, 0x43, 0x74, 0x72, 0x6c, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x69,
	0x72, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x76, 0x69, 0x72, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x22, 0x42, 0x0a, 0x08, 0x43, 0x74, 0x72, 0x6c, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x20,
	0x0a, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x22, 0x68, 0x0a, 0x0e, 0x43, 0x74, 0x72, 0x6c, 0x4f, 0x70,
	0x65, 0x6e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x63, 0x6b, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x69, 0x72, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x64, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x64, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x29, 0x0a, 0x11, 0x43, 0x74, 0x72, 0x6c, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x41, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x09, 0x43,
	0x74, 0x72, 0x6c, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x76, 0x69, 0x72, 0x74,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76,
	0x69, 0x72, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x67,
	0x0a, 0x0a, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x52,
	0x65, 0x6c, 0x61, 0x79, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x75,
	0x6e, 0x6e, 0x65, 0x6c, 0x41, 0x63, 0x6b, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x10, 0x05, 0x12, 0x08, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x10, 0x06, 0x2a, 0x7a, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x64,
	0x64, 0x43, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x65, 0x64, 0x10, 0x03, 0x12, 0x11, 0x0a,
	0x0d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6f, 0x4f, 0x6c, 0x64, 0x10, 0x04,
	0x12, 0x11, 0x0a, 0x0d, 0x4b, 0x65, 0x79, 0x4e, 0x6f, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x64, 0x10, 0x05, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_proto_goTypes = []interface{}{
	(PacketType)(0),           // 0: PacketType
	(StatusCode)(0),           // 1: StatusCode
	(*CtrlHeartbeat)(nil),     // 2: CtrlHeartbeat
	(*CtrlPing)(nil),          // 3: CtrlPing
	(*CtrlPong)(nil),          // 4: CtrlPong
	(*CtrlOpenTunnel)(nil),    // 5: CtrlOpenTunnel
	(*CtrlOpenTunnelAck)(nil), // 6: CtrlOpenTunnelAck
	(*CtrlRelay)(nil),         // 7: CtrlRelay
}
var file_api_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtrlHeartbeat); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtrlPing); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtrlPong); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtrlOpenTunnel); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtrlOpenTunnelAck); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CtrlRelay); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		EnumInfos:         file_api_proto_enumTypes,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
