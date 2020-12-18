// Code generated by protoc-gen-go. DO NOT EDIT.
// source: packet.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageType int32

const (
	MessageType_None        MessageType = 0
	MessageType_HelloServer MessageType = 1
	MessageType_Ping        MessageType = 2
	MessageType_Pong        MessageType = 3
	MessageType_Battle      MessageType = 4
	MessageType_Fin         MessageType = 5
)

var MessageType_name = map[int32]string{
	0: "None",
	1: "HelloServer",
	2: "Ping",
	3: "Pong",
	4: "Battle",
	5: "Fin",
}

var MessageType_value = map[string]int32{
	"None":        0,
	"HelloServer": 1,
	"Ping":        2,
	"Pong":        3,
	"Battle":      4,
	"Fin":         5,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{0}
}

type BattleMessage struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Seq                  uint32   `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BattleMessage) Reset()         { *m = BattleMessage{} }
func (m *BattleMessage) String() string { return proto.CompactTextString(m) }
func (*BattleMessage) ProtoMessage()    {}
func (*BattleMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{0}
}

func (m *BattleMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleMessage.Unmarshal(m, b)
}
func (m *BattleMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleMessage.Marshal(b, m, deterministic)
}
func (m *BattleMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleMessage.Merge(m, src)
}
func (m *BattleMessage) XXX_Size() int {
	return xxx_messageInfo_BattleMessage.Size(m)
}
func (m *BattleMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BattleMessage proto.InternalMessageInfo

func (m *BattleMessage) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *BattleMessage) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *BattleMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type PingMessage struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{1}
}

func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (m *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(m, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PingMessage) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type PongMessage struct {
	Timestamp            int64    `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PublicAddr           string   `protobuf:"bytes,3,opt,name=public_addr,json=publicAddr,proto3" json:"public_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongMessage) Reset()         { *m = PongMessage{} }
func (m *PongMessage) String() string { return proto.CompactTextString(m) }
func (*PongMessage) ProtoMessage()    {}
func (*PongMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{2}
}

func (m *PongMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongMessage.Unmarshal(m, b)
}
func (m *PongMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongMessage.Marshal(b, m, deterministic)
}
func (m *PongMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongMessage.Merge(m, src)
}
func (m *PongMessage) XXX_Size() int {
	return xxx_messageInfo_PongMessage.Size(m)
}
func (m *PongMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PongMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PongMessage proto.InternalMessageInfo

func (m *PongMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *PongMessage) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *PongMessage) GetPublicAddr() string {
	if m != nil {
		return m.PublicAddr
	}
	return ""
}

type HelloServerMessage struct {
	SessionIdDeprecated  string   `protobuf:"bytes,1,opt,name=session_id_deprecated,json=sessionIdDeprecated,proto3" json:"session_id_deprecated,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloServerMessage) Reset()         { *m = HelloServerMessage{} }
func (m *HelloServerMessage) String() string { return proto.CompactTextString(m) }
func (*HelloServerMessage) ProtoMessage()    {}
func (*HelloServerMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{3}
}

func (m *HelloServerMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloServerMessage.Unmarshal(m, b)
}
func (m *HelloServerMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloServerMessage.Marshal(b, m, deterministic)
}
func (m *HelloServerMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloServerMessage.Merge(m, src)
}
func (m *HelloServerMessage) XXX_Size() int {
	return xxx_messageInfo_HelloServerMessage.Size(m)
}
func (m *HelloServerMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloServerMessage.DiscardUnknown(m)
}

var xxx_messageInfo_HelloServerMessage proto.InternalMessageInfo

func (m *HelloServerMessage) GetSessionIdDeprecated() string {
	if m != nil {
		return m.SessionIdDeprecated
	}
	return ""
}

func (m *HelloServerMessage) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *HelloServerMessage) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type FinMessage struct {
	Detail               string   `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FinMessage) Reset()         { *m = FinMessage{} }
func (m *FinMessage) String() string { return proto.CompactTextString(m) }
func (*FinMessage) ProtoMessage()    {}
func (*FinMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{4}
}

func (m *FinMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinMessage.Unmarshal(m, b)
}
func (m *FinMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinMessage.Marshal(b, m, deterministic)
}
func (m *FinMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinMessage.Merge(m, src)
}
func (m *FinMessage) XXX_Size() int {
	return xxx_messageInfo_FinMessage.Size(m)
}
func (m *FinMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_FinMessage.DiscardUnknown(m)
}

var xxx_messageInfo_FinMessage proto.InternalMessageInfo

func (m *FinMessage) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Packet struct {
	Type                 MessageType         `protobuf:"varint,1,opt,name=type,proto3,enum=proto.MessageType" json:"type,omitempty"`
	Seq                  uint32              `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	Ack                  uint32              `protobuf:"varint,3,opt,name=ack,proto3" json:"ack,omitempty"`
	SessionId            string              `protobuf:"bytes,5,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	HelloServerData      *HelloServerMessage `protobuf:"bytes,10,opt,name=hello_server_data,json=helloServerData,proto3" json:"hello_server_data,omitempty"`
	PingData             *PingMessage        `protobuf:"bytes,11,opt,name=ping_data,json=pingData,proto3" json:"ping_data,omitempty"`
	PongData             *PongMessage        `protobuf:"bytes,12,opt,name=pong_data,json=pongData,proto3" json:"pong_data,omitempty"`
	BattleData           []*BattleMessage    `protobuf:"bytes,13,rep,name=battle_data,json=battleData,proto3" json:"battle_data,omitempty"`
	FinData              *FinMessage         `protobuf:"bytes,14,opt,name=fin_data,json=finData,proto3" json:"fin_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_e9ef1a6541f9f9e7, []int{5}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_None
}

func (m *Packet) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Packet) GetAck() uint32 {
	if m != nil {
		return m.Ack
	}
	return 0
}

func (m *Packet) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Packet) GetHelloServerData() *HelloServerMessage {
	if m != nil {
		return m.HelloServerData
	}
	return nil
}

func (m *Packet) GetPingData() *PingMessage {
	if m != nil {
		return m.PingData
	}
	return nil
}

func (m *Packet) GetPongData() *PongMessage {
	if m != nil {
		return m.PongData
	}
	return nil
}

func (m *Packet) GetBattleData() []*BattleMessage {
	if m != nil {
		return m.BattleData
	}
	return nil
}

func (m *Packet) GetFinData() *FinMessage {
	if m != nil {
		return m.FinData
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.MessageType", MessageType_name, MessageType_value)
	proto.RegisterType((*BattleMessage)(nil), "proto.BattleMessage")
	proto.RegisterType((*PingMessage)(nil), "proto.PingMessage")
	proto.RegisterType((*PongMessage)(nil), "proto.PongMessage")
	proto.RegisterType((*HelloServerMessage)(nil), "proto.HelloServerMessage")
	proto.RegisterType((*FinMessage)(nil), "proto.FinMessage")
	proto.RegisterType((*Packet)(nil), "proto.Packet")
}

func init() {
	proto.RegisterFile("packet.proto", fileDescriptor_e9ef1a6541f9f9e7)
}

var fileDescriptor_e9ef1a6541f9f9e7 = []byte{
	// 469 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0x26, 0x49, 0x9b, 0x36, 0x97, 0xb6, 0xcb, 0xcc, 0xaf, 0x20, 0x81, 0x88, 0x22, 0x84, 0x22,
	0x84, 0x86, 0x54, 0xc4, 0x1f, 0x00, 0x2a, 0x13, 0x7b, 0x60, 0x2a, 0x86, 0xf7, 0xc8, 0xad, 0x6f,
	0x9d, 0xd5, 0xcc, 0xce, 0x62, 0x17, 0xa9, 0x8f, 0xfc, 0xe7, 0x28, 0x4e, 0xd2, 0xa6, 0x82, 0xb7,
	0x3d, 0xf9, 0x7c, 0x77, 0xdf, 0x7d, 0x77, 0xdf, 0x1d, 0x4c, 0x4a, 0xb6, 0xde, 0xa2, 0xb9, 0x28,
	0x2b, 0x65, 0x14, 0x19, 0xda, 0x27, 0xbd, 0x86, 0xe9, 0x17, 0x66, 0x4c, 0x81, 0xdf, 0x51, 0x6b,
	0xb6, 0x41, 0xf2, 0x1c, 0x46, 0x3b, 0x8d, 0x55, 0x2e, 0x78, 0xec, 0x24, 0x4e, 0x16, 0x50, 0xbf,
	0xfe, 0x5e, 0x71, 0x12, 0x81, 0xa7, 0xf1, 0x3e, 0x76, 0x13, 0x27, 0x9b, 0xd2, 0xda, 0x24, 0x04,
	0x06, 0x2b, 0xc5, 0xf7, 0xb1, 0x97, 0x38, 0xd9, 0x84, 0x5a, 0x3b, 0x5d, 0x40, 0xb8, 0x14, 0x72,
	0xd3, 0x55, 0x7b, 0x09, 0x81, 0x11, 0x77, 0xa8, 0x0d, 0xbb, 0x2b, 0x6d, 0x3d, 0x8f, 0x1e, 0x1d,
	0x7d, 0x2e, 0xb7, 0xcf, 0x95, 0x22, 0x84, 0x4b, 0xf5, 0xd0, 0x2a, 0xe4, 0x35, 0x84, 0xe5, 0x6e,
	0x55, 0x88, 0x75, 0xce, 0x38, 0xaf, 0x6c, 0x9b, 0x01, 0x85, 0xc6, 0xf5, 0x99, 0xf3, 0x2a, 0xbd,
	0x07, 0xf2, 0x0d, 0x8b, 0x42, 0xfd, 0xc4, 0xea, 0x37, 0x56, 0x1d, 0xdb, 0x1c, 0x9e, 0x6a, 0xd4,
	0x5a, 0x28, 0x99, 0x0b, 0x9e, 0x73, 0x2c, 0x2b, 0x5c, 0x33, 0x83, 0x9d, 0x1e, 0x8f, 0xdb, 0xe0,
	0x15, 0x5f, 0x1c, 0x42, 0x64, 0x06, 0xae, 0xda, 0x5a, 0xfa, 0x31, 0x75, 0xd5, 0xb6, 0xdf, 0x93,
	0x77, 0x32, 0xd9, 0x1b, 0x80, 0x4b, 0x21, 0x3b, 0xaa, 0x67, 0xe0, 0x73, 0x34, 0x4c, 0x14, 0x9d,
	0xd6, 0xcd, 0x2f, 0xfd, 0xe3, 0x81, 0xbf, 0xb4, 0xdb, 0x22, 0x6f, 0x61, 0x60, 0xf6, 0x25, 0xda,
	0x84, 0xd9, 0x9c, 0x34, 0xdb, 0xbb, 0x68, 0x0b, 0xfc, 0xda, 0x97, 0x48, 0x6d, 0xfc, 0x3f, 0xeb,
	0x89, 0xc0, 0x63, 0xeb, 0xad, 0xe5, 0x9f, 0xd2, 0xda, 0x24, 0xaf, 0x00, 0x8e, 0x93, 0xc5, 0x43,
	0x4b, 0x19, 0x1c, 0xc6, 0x21, 0x5f, 0xe1, 0xfc, 0xb6, 0x96, 0x23, 0xd7, 0x56, 0x8f, 0x9c, 0x33,
	0xc3, 0x62, 0x48, 0x9c, 0x2c, 0x9c, 0xbf, 0x68, 0x79, 0xff, 0x95, 0x8b, 0x9e, 0xdd, 0x1e, 0x7d,
	0x0b, 0x66, 0x18, 0xf9, 0x00, 0x41, 0x29, 0xe4, 0xa6, 0x81, 0x87, 0x16, 0xde, 0xb5, 0xdd, 0x3b,
	0x0d, 0x3a, 0xae, 0x93, 0x0e, 0x00, 0xd5, 0x01, 0x26, 0xa7, 0x00, 0xd5, 0x07, 0xa8, 0x16, 0xf0,
	0x09, 0xc2, 0x95, 0x3d, 0xda, 0x06, 0x32, 0x4d, 0xbc, 0x2c, 0x9c, 0x3f, 0x69, 0x21, 0x27, 0xe7,
	0x4c, 0xa1, 0x49, 0xb4, 0xb0, 0xf7, 0x30, 0xbe, 0x11, 0xb2, 0xc1, 0xcc, 0x2c, 0xcd, 0x79, 0x8b,
	0x39, 0xae, 0x84, 0x8e, 0x6e, 0x84, 0xac, 0xb3, 0xdf, 0xfd, 0x80, 0xb0, 0xa7, 0x32, 0x19, 0xc3,
	0xe0, 0x5a, 0x49, 0x8c, 0x1e, 0x91, 0x33, 0x08, 0x7b, 0x32, 0x44, 0x4e, 0x1d, 0xaa, 0x07, 0x8b,
	0x5c, 0x6b, 0x29, 0xb9, 0x89, 0x3c, 0x02, 0xe0, 0x37, 0x8d, 0x44, 0x03, 0x32, 0x02, 0xef, 0x52,
	0xc8, 0x68, 0xb8, 0xf2, 0x2d, 0xdb, 0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0x30, 0x0b,
	0xd0, 0x8a, 0x03, 0x00, 0x00,
}