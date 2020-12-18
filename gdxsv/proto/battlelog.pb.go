// Code generated by protoc-gen-go. DO NOT EDIT.
// source: battlelog.proto

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

type BattleLogUser struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	PilotName            string   `protobuf:"bytes,3,opt,name=pilot_name,json=pilotName,proto3" json:"pilot_name,omitempty"`
	GameParam            []byte   `protobuf:"bytes,4,opt,name=game_param,json=gameParam,proto3" json:"game_param,omitempty"`
	BattleCount          int32    `protobuf:"varint,5,opt,name=battle_count,json=battleCount,proto3" json:"battle_count,omitempty"`
	WinCount             int32    `protobuf:"varint,6,opt,name=win_count,json=winCount,proto3" json:"win_count,omitempty"`
	LoseCount            int32    `protobuf:"varint,7,opt,name=lose_count,json=loseCount,proto3" json:"lose_count,omitempty"`
	Platform             string   `protobuf:"bytes,10,opt,name=platform,proto3" json:"platform,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BattleLogUser) Reset()         { *m = BattleLogUser{} }
func (m *BattleLogUser) String() string { return proto.CompactTextString(m) }
func (*BattleLogUser) ProtoMessage()    {}
func (*BattleLogUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bd4546689e62ffe, []int{0}
}

func (m *BattleLogUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleLogUser.Unmarshal(m, b)
}
func (m *BattleLogUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleLogUser.Marshal(b, m, deterministic)
}
func (m *BattleLogUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleLogUser.Merge(m, src)
}
func (m *BattleLogUser) XXX_Size() int {
	return xxx_messageInfo_BattleLogUser.Size(m)
}
func (m *BattleLogUser) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleLogUser.DiscardUnknown(m)
}

var xxx_messageInfo_BattleLogUser proto.InternalMessageInfo

func (m *BattleLogUser) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *BattleLogUser) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *BattleLogUser) GetPilotName() string {
	if m != nil {
		return m.PilotName
	}
	return ""
}

func (m *BattleLogUser) GetGameParam() []byte {
	if m != nil {
		return m.GameParam
	}
	return nil
}

func (m *BattleLogUser) GetBattleCount() int32 {
	if m != nil {
		return m.BattleCount
	}
	return 0
}

func (m *BattleLogUser) GetWinCount() int32 {
	if m != nil {
		return m.WinCount
	}
	return 0
}

func (m *BattleLogUser) GetLoseCount() int32 {
	if m != nil {
		return m.LoseCount
	}
	return 0
}

func (m *BattleLogUser) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

type BattleLogMessage struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Seq                  uint32   `protobuf:"varint,3,opt,name=seq,proto3" json:"seq,omitempty"`
	Timestamp            int64    `protobuf:"varint,10,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BattleLogMessage) Reset()         { *m = BattleLogMessage{} }
func (m *BattleLogMessage) String() string { return proto.CompactTextString(m) }
func (*BattleLogMessage) ProtoMessage()    {}
func (*BattleLogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bd4546689e62ffe, []int{1}
}

func (m *BattleLogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleLogMessage.Unmarshal(m, b)
}
func (m *BattleLogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleLogMessage.Marshal(b, m, deterministic)
}
func (m *BattleLogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleLogMessage.Merge(m, src)
}
func (m *BattleLogMessage) XXX_Size() int {
	return xxx_messageInfo_BattleLogMessage.Size(m)
}
func (m *BattleLogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleLogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BattleLogMessage proto.InternalMessageInfo

func (m *BattleLogMessage) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *BattleLogMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *BattleLogMessage) GetSeq() uint32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *BattleLogMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type BattleLogFile struct {
	GameDisk             string              `protobuf:"bytes,5,opt,name=game_disk,json=gameDisk,proto3" json:"game_disk,omitempty"`
	GdxsvVersion         string              `protobuf:"bytes,2,opt,name=gdxsv_version,json=gdxsvVersion,proto3" json:"gdxsv_version,omitempty"`
	BattleCode           string              `protobuf:"bytes,3,opt,name=battle_code,json=battleCode,proto3" json:"battle_code,omitempty"`
	LogFileVersion       int32               `protobuf:"varint,4,opt,name=log_file_version,json=logFileVersion,proto3" json:"log_file_version,omitempty"`
	RuleBin              []byte              `protobuf:"bytes,10,opt,name=rule_bin,json=ruleBin,proto3" json:"rule_bin,omitempty"`
	Users                []*BattleLogUser    `protobuf:"bytes,11,rep,name=users,proto3" json:"users,omitempty"`
	BattleData           []*BattleLogMessage `protobuf:"bytes,12,rep,name=battle_data,json=battleData,proto3" json:"battle_data,omitempty"`
	StartAt              int64               `protobuf:"varint,20,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	EndAt                int64               `protobuf:"varint,21,opt,name=end_at,json=endAt,proto3" json:"end_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BattleLogFile) Reset()         { *m = BattleLogFile{} }
func (m *BattleLogFile) String() string { return proto.CompactTextString(m) }
func (*BattleLogFile) ProtoMessage()    {}
func (*BattleLogFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bd4546689e62ffe, []int{2}
}

func (m *BattleLogFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleLogFile.Unmarshal(m, b)
}
func (m *BattleLogFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleLogFile.Marshal(b, m, deterministic)
}
func (m *BattleLogFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleLogFile.Merge(m, src)
}
func (m *BattleLogFile) XXX_Size() int {
	return xxx_messageInfo_BattleLogFile.Size(m)
}
func (m *BattleLogFile) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleLogFile.DiscardUnknown(m)
}

var xxx_messageInfo_BattleLogFile proto.InternalMessageInfo

func (m *BattleLogFile) GetGameDisk() string {
	if m != nil {
		return m.GameDisk
	}
	return ""
}

func (m *BattleLogFile) GetGdxsvVersion() string {
	if m != nil {
		return m.GdxsvVersion
	}
	return ""
}

func (m *BattleLogFile) GetBattleCode() string {
	if m != nil {
		return m.BattleCode
	}
	return ""
}

func (m *BattleLogFile) GetLogFileVersion() int32 {
	if m != nil {
		return m.LogFileVersion
	}
	return 0
}

func (m *BattleLogFile) GetRuleBin() []byte {
	if m != nil {
		return m.RuleBin
	}
	return nil
}

func (m *BattleLogFile) GetUsers() []*BattleLogUser {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *BattleLogFile) GetBattleData() []*BattleLogMessage {
	if m != nil {
		return m.BattleData
	}
	return nil
}

func (m *BattleLogFile) GetStartAt() int64 {
	if m != nil {
		return m.StartAt
	}
	return 0
}

func (m *BattleLogFile) GetEndAt() int64 {
	if m != nil {
		return m.EndAt
	}
	return 0
}

func init() {
	proto.RegisterType((*BattleLogUser)(nil), "proto.BattleLogUser")
	proto.RegisterType((*BattleLogMessage)(nil), "proto.BattleLogMessage")
	proto.RegisterType((*BattleLogFile)(nil), "proto.BattleLogFile")
}

func init() {
	proto.RegisterFile("battlelog.proto", fileDescriptor_7bd4546689e62ffe)
}

var fileDescriptor_7bd4546689e62ffe = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x6f, 0xd4, 0x30,
	0x10, 0xc5, 0x95, 0xee, 0x66, 0x37, 0x99, 0xcd, 0xc2, 0xca, 0x6a, 0xd5, 0xf0, 0xa7, 0x62, 0x59,
	0x2e, 0x11, 0x87, 0x1e, 0xe0, 0xc2, 0xb5, 0xa5, 0x42, 0x42, 0x02, 0x84, 0x2c, 0xc1, 0x35, 0xf2,
	0xe2, 0x69, 0x64, 0xd5, 0xb1, 0x83, 0xed, 0x6d, 0xe1, 0x53, 0xf2, 0x8d, 0x10, 0xf2, 0x78, 0xb3,
	0xfc, 0x91, 0x38, 0xc5, 0xf3, 0x7e, 0x99, 0xf1, 0x9b, 0x27, 0xc3, 0xfd, 0xad, 0x08, 0x41, 0xa3,
	0xb6, 0xdd, 0xf9, 0xe0, 0x6c, 0xb0, 0x2c, 0xa7, 0xcf, 0xe6, 0x67, 0x06, 0xcb, 0x4b, 0x42, 0xef,
	0x6c, 0xf7, 0xc9, 0xa3, 0x63, 0xa7, 0x30, 0xdf, 0x79, 0x74, 0xad, 0x92, 0x75, 0xb6, 0xce, 0x9a,
	0x92, 0xcf, 0x62, 0xf9, 0x56, 0xb2, 0x47, 0x50, 0x12, 0x30, 0xa2, 0xc7, 0xfa, 0x88, 0x50, 0x11,
	0x85, 0x0f, 0xa2, 0x47, 0x76, 0x06, 0x30, 0x28, 0x6d, 0x43, 0xa2, 0x13, 0xa2, 0x25, 0x29, 0x23,
	0xee, 0x44, 0x8f, 0xed, 0x20, 0x9c, 0xe8, 0xeb, 0xe9, 0x3a, 0x6b, 0x2a, 0x5e, 0x46, 0xe5, 0x63,
	0x14, 0xd8, 0x53, 0xa8, 0x92, 0xbf, 0xf6, 0x8b, 0xdd, 0x99, 0x50, 0xe7, 0xeb, 0xac, 0xc9, 0xf9,
	0x22, 0x69, 0xaf, 0xa3, 0x14, 0x6f, 0xbf, 0x53, 0x66, 0xcf, 0x67, 0xc4, 0x8b, 0x3b, 0x65, 0x12,
	0x3c, 0x03, 0xd0, 0xd6, 0x8f, 0xdd, 0x73, 0xa2, 0x65, 0x54, 0x12, 0x7e, 0x08, 0xc5, 0xa0, 0x45,
	0xb8, 0xb6, 0xae, 0xaf, 0x21, 0x19, 0x1f, 0xeb, 0x8d, 0x85, 0xd5, 0x61, 0xff, 0xf7, 0xe8, 0xbd,
	0xe8, 0xf0, 0xff, 0x11, 0x30, 0x98, 0x6e, 0xad, 0xfc, 0x4e, 0xdb, 0x57, 0x9c, 0xce, 0x6c, 0x05,
	0x13, 0x8f, 0x5f, 0x69, 0xe5, 0x25, 0x8f, 0x47, 0xf6, 0x18, 0xca, 0xa0, 0x7a, 0xf4, 0x41, 0xf4,
	0x03, 0xdd, 0x37, 0xe1, 0xbf, 0x85, 0xcd, 0x8f, 0xa3, 0x3f, 0x12, 0x7f, 0xa3, 0x34, 0xc6, 0xd5,
	0x28, 0x1c, 0xa9, 0xfc, 0x0d, 0xad, 0x5e, 0xf2, 0x22, 0x0a, 0x57, 0xca, 0xdf, 0xb0, 0x67, 0xb0,
	0xec, 0xe4, 0x37, 0x7f, 0xdb, 0xde, 0xa2, 0xf3, 0xca, 0x9a, 0x7d, 0xf2, 0x15, 0x89, 0x9f, 0x93,
	0xc6, 0x9e, 0xc0, 0xe2, 0x90, 0x9f, 0x1c, 0xe3, 0x87, 0x31, 0x3e, 0x89, 0xac, 0x81, 0x95, 0xb6,
	0x5d, 0x7b, 0xad, 0x34, 0x1e, 0x06, 0x4d, 0x29, 0xa6, 0x7b, 0x3a, 0xb9, 0x18, 0x47, 0x3d, 0x80,
	0xc2, 0xed, 0x34, 0xb6, 0x5b, 0x65, 0xc8, 0x7b, 0xc5, 0xe7, 0xb1, 0xbe, 0x54, 0x86, 0x3d, 0x87,
	0x3c, 0xe6, 0xe0, 0xeb, 0xc5, 0x7a, 0xd2, 0x2c, 0x5e, 0x1c, 0xa7, 0x97, 0x74, 0xfe, 0xd7, 0xf3,
	0xe1, 0xe9, 0x17, 0xf6, 0xea, 0xe0, 0x48, 0x8a, 0x20, 0xea, 0x8a, 0x3a, 0x4e, 0xff, 0xed, 0xd8,
	0x07, 0x3e, 0x5a, 0xbd, 0x12, 0x41, 0x44, 0x03, 0x3e, 0x08, 0x17, 0x5a, 0x11, 0xea, 0x63, 0x0a,
	0x6f, 0x4e, 0xf5, 0x45, 0x60, 0x27, 0x30, 0x43, 0x23, 0x23, 0x38, 0x21, 0x90, 0xa3, 0x91, 0x17,
	0x61, 0x3b, 0xa3, 0xa9, 0x2f, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x65, 0xa9, 0x50, 0x83, 0xe4,
	0x02, 0x00, 0x00,
}