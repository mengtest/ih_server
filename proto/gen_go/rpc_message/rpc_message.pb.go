// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc_message.proto

/*
Package msg_rpc_message is a generated protocol buffer package.

It is generated from these files:
	rpc_message.proto

It has these top-level messages:
	G2GPlayerInfoRequest
	G2GPlayerInfoResponse
	G2GPlayerInfoNotify
	G2GPlayerBattleInfoRequest
	G2GPlayerBattleInfoResponse
	G2GPlayerMultiInfoRequest
	PlayerInfo
	G2GPlayerMultiInfoResponse
	G2GFriendInfoRequest
	G2GFriendInfoResponse
	G2GFriendAskRequest
	G2GFriendAskResponse
	G2GCarnivalBeInvitedRequest
	G2GCarnivalBeInvitedResponse
*/
package msg_rpc_message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MSGID int32

const (
	MSGID_NONE                            MSGID = 0
	MSGID_G2G_PLAYER_INFO_REQUEST         MSGID = 1
	MSGID_G2G_PLAYER_INFO_RESPONSE        MSGID = 2
	MSGID_G2G_PLAYER_INFO_NOTIFY          MSGID = 3
	MSGID_G2G_PLAYER_BATTLE_INFO_REQUEST  MSGID = 4
	MSGID_G2G_PLAYER_BATTLE_INFO_RESPONSE MSGID = 5
	MSGID_G2G_PLAYER_MULTI_INFO_REQUEST   MSGID = 10
	MSGID_G2G_PLAYER_MULTI_INFO_RESPONSE  MSGID = 11
	MSGID_G2G_FRIEND_INFO_REQUEST         MSGID = 100
	MSGID_G2G_FRIEND_INFO_RESPONSE        MSGID = 101
	MSGID_G2G_FRIEND_ASK_REQUEST          MSGID = 102
	MSGID_G2G_FRIEND_ASK_RESPONSE         MSGID = 103
	MSGID_G2G_FRIEND_AGREE_REQUEST        MSGID = 104
	MSGID_G2G_FRIEND_AGREE_RESPONSE       MSGID = 105
	MSGID_G2G_FRIEND_REFUSE_REQUEST       MSGID = 106
	MSGID_G2G_FRIEND_REFUSE_RESPONSE      MSGID = 107
	MSGID_G2G_FRIEND_REMOVE_REQUEST       MSGID = 108
	MSGID_G2G_FRIEND_REMOVE_RESPONSE      MSGID = 109
	MSGID_G2G_FRIEND_REMOVE_NOTIFY        MSGID = 110
	MSGID_G2G_FRIEND_GIVE_POINTS_REQUEST  MSGID = 111
	MSGID_G2G_FRIEND_GIVE_POINTS_RESPONSE MSGID = 112
	MSGID_G2G_FRIEND_GET_POINTS_REQUEST   MSGID = 113
	MSGID_G2G_FRIEND_GET_POINTS_RESPONSE  MSGID = 114
	// 嘉年華
	MSGID_G2G_CARNIVAL_BE_INVITED_REQUEST  MSGID = 1000
	MSGID_G2G_CARNIVAL_BE_INVITED_RESPONSE MSGID = 1001
)

var MSGID_name = map[int32]string{
	0:    "NONE",
	1:    "G2G_PLAYER_INFO_REQUEST",
	2:    "G2G_PLAYER_INFO_RESPONSE",
	3:    "G2G_PLAYER_INFO_NOTIFY",
	4:    "G2G_PLAYER_BATTLE_INFO_REQUEST",
	5:    "G2G_PLAYER_BATTLE_INFO_RESPONSE",
	10:   "G2G_PLAYER_MULTI_INFO_REQUEST",
	11:   "G2G_PLAYER_MULTI_INFO_RESPONSE",
	100:  "G2G_FRIEND_INFO_REQUEST",
	101:  "G2G_FRIEND_INFO_RESPONSE",
	102:  "G2G_FRIEND_ASK_REQUEST",
	103:  "G2G_FRIEND_ASK_RESPONSE",
	104:  "G2G_FRIEND_AGREE_REQUEST",
	105:  "G2G_FRIEND_AGREE_RESPONSE",
	106:  "G2G_FRIEND_REFUSE_REQUEST",
	107:  "G2G_FRIEND_REFUSE_RESPONSE",
	108:  "G2G_FRIEND_REMOVE_REQUEST",
	109:  "G2G_FRIEND_REMOVE_RESPONSE",
	110:  "G2G_FRIEND_REMOVE_NOTIFY",
	111:  "G2G_FRIEND_GIVE_POINTS_REQUEST",
	112:  "G2G_FRIEND_GIVE_POINTS_RESPONSE",
	113:  "G2G_FRIEND_GET_POINTS_REQUEST",
	114:  "G2G_FRIEND_GET_POINTS_RESPONSE",
	1000: "G2G_CARNIVAL_BE_INVITED_REQUEST",
	1001: "G2G_CARNIVAL_BE_INVITED_RESPONSE",
}
var MSGID_value = map[string]int32{
	"NONE": 0,
	"G2G_PLAYER_INFO_REQUEST":          1,
	"G2G_PLAYER_INFO_RESPONSE":         2,
	"G2G_PLAYER_INFO_NOTIFY":           3,
	"G2G_PLAYER_BATTLE_INFO_REQUEST":   4,
	"G2G_PLAYER_BATTLE_INFO_RESPONSE":  5,
	"G2G_PLAYER_MULTI_INFO_REQUEST":    10,
	"G2G_PLAYER_MULTI_INFO_RESPONSE":   11,
	"G2G_FRIEND_INFO_REQUEST":          100,
	"G2G_FRIEND_INFO_RESPONSE":         101,
	"G2G_FRIEND_ASK_REQUEST":           102,
	"G2G_FRIEND_ASK_RESPONSE":          103,
	"G2G_FRIEND_AGREE_REQUEST":         104,
	"G2G_FRIEND_AGREE_RESPONSE":        105,
	"G2G_FRIEND_REFUSE_REQUEST":        106,
	"G2G_FRIEND_REFUSE_RESPONSE":       107,
	"G2G_FRIEND_REMOVE_REQUEST":        108,
	"G2G_FRIEND_REMOVE_RESPONSE":       109,
	"G2G_FRIEND_REMOVE_NOTIFY":         110,
	"G2G_FRIEND_GIVE_POINTS_REQUEST":   111,
	"G2G_FRIEND_GIVE_POINTS_RESPONSE":  112,
	"G2G_FRIEND_GET_POINTS_REQUEST":    113,
	"G2G_FRIEND_GET_POINTS_RESPONSE":   114,
	"G2G_CARNIVAL_BE_INVITED_REQUEST":  1000,
	"G2G_CARNIVAL_BE_INVITED_RESPONSE": 1001,
}

func (x MSGID) String() string {
	return proto.EnumName(MSGID_name, int32(x))
}
func (MSGID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 玩家信息请求
type G2GPlayerInfoRequest struct {
}

func (m *G2GPlayerInfoRequest) Reset()                    { *m = G2GPlayerInfoRequest{} }
func (m *G2GPlayerInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*G2GPlayerInfoRequest) ProtoMessage()               {}
func (*G2GPlayerInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// 玩家信息返回
type G2GPlayerInfoResponse struct {
	UniqueId string `protobuf:"bytes,2,opt,name=UniqueId" json:"UniqueId,omitempty"`
	Account  string `protobuf:"bytes,3,opt,name=Account" json:"Account,omitempty"`
	Level    int32  `protobuf:"varint,4,opt,name=Level" json:"Level,omitempty"`
	Head     int32  `protobuf:"varint,5,opt,name=Head" json:"Head,omitempty"`
}

func (m *G2GPlayerInfoResponse) Reset()                    { *m = G2GPlayerInfoResponse{} }
func (m *G2GPlayerInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*G2GPlayerInfoResponse) ProtoMessage()               {}
func (*G2GPlayerInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *G2GPlayerInfoResponse) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *G2GPlayerInfoResponse) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *G2GPlayerInfoResponse) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *G2GPlayerInfoResponse) GetHead() int32 {
	if m != nil {
		return m.Head
	}
	return 0
}

// 玩家信息通知
type G2GPlayerInfoNotify struct {
}

func (m *G2GPlayerInfoNotify) Reset()                    { *m = G2GPlayerInfoNotify{} }
func (m *G2GPlayerInfoNotify) String() string            { return proto.CompactTextString(m) }
func (*G2GPlayerInfoNotify) ProtoMessage()               {}
func (*G2GPlayerInfoNotify) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// 玩家战斗信息请求
type G2GPlayerBattleInfoRequest struct {
}

func (m *G2GPlayerBattleInfoRequest) Reset()                    { *m = G2GPlayerBattleInfoRequest{} }
func (m *G2GPlayerBattleInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*G2GPlayerBattleInfoRequest) ProtoMessage()               {}
func (*G2GPlayerBattleInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// 玩家战斗信息返回
type G2GPlayerBattleInfoResponse struct {
}

func (m *G2GPlayerBattleInfoResponse) Reset()                    { *m = G2GPlayerBattleInfoResponse{} }
func (m *G2GPlayerBattleInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*G2GPlayerBattleInfoResponse) ProtoMessage()               {}
func (*G2GPlayerBattleInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// 多个玩家信息请求
type G2GPlayerMultiInfoRequest struct {
}

func (m *G2GPlayerMultiInfoRequest) Reset()                    { *m = G2GPlayerMultiInfoRequest{} }
func (m *G2GPlayerMultiInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*G2GPlayerMultiInfoRequest) ProtoMessage()               {}
func (*G2GPlayerMultiInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type PlayerInfo struct {
	PlayerId int32  `protobuf:"varint,1,opt,name=PlayerId" json:"PlayerId,omitempty"`
	UniqueId string `protobuf:"bytes,2,opt,name=UniqueId" json:"UniqueId,omitempty"`
	Account  string `protobuf:"bytes,3,opt,name=Account" json:"Account,omitempty"`
	Level    int32  `protobuf:"varint,4,opt,name=Level" json:"Level,omitempty"`
	Head     int32  `protobuf:"varint,5,opt,name=Head" json:"Head,omitempty"`
}

func (m *PlayerInfo) Reset()                    { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()               {}
func (*PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PlayerInfo) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *PlayerInfo) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *PlayerInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *PlayerInfo) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *PlayerInfo) GetHead() int32 {
	if m != nil {
		return m.Head
	}
	return 0
}

// 多个玩家信息返回
type G2GPlayerMultiInfoResponse struct {
	PlayerInfos []*PlayerInfo `protobuf:"bytes,1,rep,name=PlayerInfos" json:"PlayerInfos,omitempty"`
}

func (m *G2GPlayerMultiInfoResponse) Reset()                    { *m = G2GPlayerMultiInfoResponse{} }
func (m *G2GPlayerMultiInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*G2GPlayerMultiInfoResponse) ProtoMessage()               {}
func (*G2GPlayerMultiInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *G2GPlayerMultiInfoResponse) GetPlayerInfos() []*PlayerInfo {
	if m != nil {
		return m.PlayerInfos
	}
	return nil
}

// 好友信息请求
type G2GFriendInfoRequest struct {
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId" json:"FriendId,omitempty"`
}

func (m *G2GFriendInfoRequest) Reset()                    { *m = G2GFriendInfoRequest{} }
func (m *G2GFriendInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*G2GFriendInfoRequest) ProtoMessage()               {}
func (*G2GFriendInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *G2GFriendInfoRequest) GetFriendId() int32 {
	if m != nil {
		return m.FriendId
	}
	return 0
}

// 好友信息返回
type G2GFriendInfoResponse struct {
}

func (m *G2GFriendInfoResponse) Reset()                    { *m = G2GFriendInfoResponse{} }
func (m *G2GFriendInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*G2GFriendInfoResponse) ProtoMessage()               {}
func (*G2GFriendInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

// 添加好友请求
type G2GFriendAskRequest struct {
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId" json:"FriendId,omitempty"`
}

func (m *G2GFriendAskRequest) Reset()                    { *m = G2GFriendAskRequest{} }
func (m *G2GFriendAskRequest) String() string            { return proto.CompactTextString(m) }
func (*G2GFriendAskRequest) ProtoMessage()               {}
func (*G2GFriendAskRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *G2GFriendAskRequest) GetFriendId() int32 {
	if m != nil {
		return m.FriendId
	}
	return 0
}

// 添加好友返回
type G2GFriendAskResponse struct {
}

func (m *G2GFriendAskResponse) Reset()                    { *m = G2GFriendAskResponse{} }
func (m *G2GFriendAskResponse) String() string            { return proto.CompactTextString(m) }
func (*G2GFriendAskResponse) ProtoMessage()               {}
func (*G2GFriendAskResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

// ----------------------------------- 嘉年華 ----------------------------------
// 被邀請請求
type G2GCarnivalBeInvitedRequest struct {
	InviteCode string `protobuf:"bytes,1,opt,name=InviteCode" json:"InviteCode,omitempty"`
}

func (m *G2GCarnivalBeInvitedRequest) Reset()                    { *m = G2GCarnivalBeInvitedRequest{} }
func (m *G2GCarnivalBeInvitedRequest) String() string            { return proto.CompactTextString(m) }
func (*G2GCarnivalBeInvitedRequest) ProtoMessage()               {}
func (*G2GCarnivalBeInvitedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *G2GCarnivalBeInvitedRequest) GetInviteCode() string {
	if m != nil {
		return m.InviteCode
	}
	return ""
}

type G2GCarnivalBeInvitedResponse struct {
	InviteCode string `protobuf:"bytes,2,opt,name=InviteCode" json:"InviteCode,omitempty"`
}

func (m *G2GCarnivalBeInvitedResponse) Reset()                    { *m = G2GCarnivalBeInvitedResponse{} }
func (m *G2GCarnivalBeInvitedResponse) String() string            { return proto.CompactTextString(m) }
func (*G2GCarnivalBeInvitedResponse) ProtoMessage()               {}
func (*G2GCarnivalBeInvitedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *G2GCarnivalBeInvitedResponse) GetInviteCode() string {
	if m != nil {
		return m.InviteCode
	}
	return ""
}

func init() {
	proto.RegisterType((*G2GPlayerInfoRequest)(nil), "msg.rpc_message.G2GPlayerInfoRequest")
	proto.RegisterType((*G2GPlayerInfoResponse)(nil), "msg.rpc_message.G2GPlayerInfoResponse")
	proto.RegisterType((*G2GPlayerInfoNotify)(nil), "msg.rpc_message.G2GPlayerInfoNotify")
	proto.RegisterType((*G2GPlayerBattleInfoRequest)(nil), "msg.rpc_message.G2GPlayerBattleInfoRequest")
	proto.RegisterType((*G2GPlayerBattleInfoResponse)(nil), "msg.rpc_message.G2GPlayerBattleInfoResponse")
	proto.RegisterType((*G2GPlayerMultiInfoRequest)(nil), "msg.rpc_message.G2GPlayerMultiInfoRequest")
	proto.RegisterType((*PlayerInfo)(nil), "msg.rpc_message.PlayerInfo")
	proto.RegisterType((*G2GPlayerMultiInfoResponse)(nil), "msg.rpc_message.G2GPlayerMultiInfoResponse")
	proto.RegisterType((*G2GFriendInfoRequest)(nil), "msg.rpc_message.G2GFriendInfoRequest")
	proto.RegisterType((*G2GFriendInfoResponse)(nil), "msg.rpc_message.G2GFriendInfoResponse")
	proto.RegisterType((*G2GFriendAskRequest)(nil), "msg.rpc_message.G2GFriendAskRequest")
	proto.RegisterType((*G2GFriendAskResponse)(nil), "msg.rpc_message.G2GFriendAskResponse")
	proto.RegisterType((*G2GCarnivalBeInvitedRequest)(nil), "msg.rpc_message.G2GCarnivalBeInvitedRequest")
	proto.RegisterType((*G2GCarnivalBeInvitedResponse)(nil), "msg.rpc_message.G2GCarnivalBeInvitedResponse")
	proto.RegisterEnum("msg.rpc_message.MSGID", MSGID_name, MSGID_value)
}

func init() { proto.RegisterFile("rpc_message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 620 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0xd2, 0x26, 0xfd, 0x99, 0x5e, 0x7c, 0x61, 0xe9, 0x8f, 0x69, 0xda, 0x12, 0x0c, 0x48,
	0x15, 0x17, 0x91, 0x08, 0xd7, 0x45, 0x72, 0xda, 0x8d, 0x59, 0x91, 0xd8, 0xc1, 0x76, 0x22, 0x55,
	0x5c, 0x58, 0xa6, 0xde, 0x16, 0x53, 0xd7, 0x4e, 0x6d, 0xa7, 0xa8, 0x6f, 0xc0, 0x73, 0xf0, 0x74,
	0xf0, 0x16, 0xa8, 0xeb, 0xb5, 0xb3, 0xde, 0xa4, 0x12, 0x37, 0xdc, 0x65, 0xe6, 0xcc, 0x39, 0x33,
	0x67, 0xc6, 0xab, 0xc0, 0x93, 0x64, 0x7a, 0xe1, 0xde, 0xd0, 0x34, 0xf5, 0xae, 0x68, 0x67, 0x9a,
	0xc4, 0x59, 0x8c, 0xfe, 0xbf, 0x49, 0xaf, 0x3a, 0x42, 0x5a, 0xdd, 0x85, 0x6d, 0xbd, 0xab, 0x8f,
	0x42, 0xef, 0x9e, 0x26, 0x24, 0xba, 0x8c, 0x2d, 0x7a, 0x3b, 0xa3, 0x69, 0xa6, 0x7e, 0x87, 0x1d,
	0x29, 0x9f, 0x4e, 0xe3, 0x28, 0xa5, 0x68, 0x1f, 0x36, 0xc6, 0x51, 0x70, 0x3b, 0xa3, 0xc4, 0x57,
	0x56, 0xda, 0xb5, 0xe3, 0x4d, 0xab, 0x8c, 0x91, 0x02, 0xeb, 0xda, 0xc5, 0x45, 0x3c, 0x8b, 0x32,
	0x65, 0x95, 0x41, 0x45, 0x88, 0xb6, 0xa1, 0x31, 0xa0, 0x77, 0x34, 0x54, 0xea, 0xed, 0xda, 0x71,
	0xc3, 0xca, 0x03, 0x84, 0xa0, 0xfe, 0x81, 0x7a, 0xbe, 0xd2, 0x60, 0x49, 0xf6, 0x5b, 0xdd, 0x81,
	0xa7, 0x95, 0xc6, 0x46, 0x9c, 0x05, 0x97, 0xf7, 0xea, 0x01, 0xec, 0x97, 0xe9, 0x9e, 0x97, 0x65,
	0x21, 0x15, 0xa7, 0x3d, 0x84, 0xd6, 0x52, 0x34, 0x9f, 0x59, 0x6d, 0xc1, 0xb3, 0x12, 0x1e, 0xce,
	0xc2, 0x2c, 0x10, 0xb9, 0x3f, 0x6a, 0x00, 0xf3, 0x76, 0x0f, 0xfe, 0x78, 0xe4, 0x2b, 0x35, 0x36,
	0x57, 0x19, 0xff, 0x73, 0xef, 0x9f, 0x05, 0x93, 0xc2, 0x9c, 0x7c, 0xf3, 0x27, 0xb0, 0x35, 0x9f,
	0x33, 0x55, 0x6a, 0xed, 0xd5, 0xe3, 0xad, 0x6e, 0xab, 0x23, 0x5d, 0xb4, 0x23, 0xdc, 0x4c, 0xac,
	0x57, 0xbb, 0xec, 0xd2, 0xfd, 0x24, 0xa0, 0x91, 0x2f, 0xf8, 0x7f, 0x30, 0xc5, 0x93, 0xa5, 0xe1,
	0x22, 0x56, 0xf7, 0xd8, 0x57, 0x20, 0x72, 0xf8, 0x46, 0xdf, 0xb2, 0x2b, 0xe5, 0x80, 0x96, 0x5e,
	0xff, 0x8d, 0xd6, 0xae, 0xd0, 0x9f, 0x51, 0xb8, 0xd4, 0x09, 0xbb, 0xdd, 0xa9, 0x97, 0x44, 0xc1,
	0x9d, 0x17, 0xf6, 0x28, 0x89, 0xee, 0x82, 0x8c, 0xfa, 0x85, 0xe4, 0x11, 0x40, 0x9e, 0x39, 0x8d,
	0x7d, 0xca, 0x44, 0x37, 0x2d, 0x21, 0xa3, 0xbe, 0x87, 0x83, 0xe5, 0x74, 0xbe, 0xb5, 0x2a, 0x7f,
	0x45, 0xe6, 0xbf, 0xf9, 0xb9, 0x06, 0x8d, 0xa1, 0xad, 0x93, 0x33, 0xb4, 0x01, 0x75, 0xc3, 0x34,
	0x70, 0xf3, 0x3f, 0xd4, 0x82, 0x3d, 0xbd, 0xab, 0xbb, 0xa3, 0x81, 0x76, 0x8e, 0x2d, 0x97, 0x18,
	0x7d, 0xd3, 0xb5, 0xf0, 0xa7, 0x31, 0xb6, 0x9d, 0x66, 0x0d, 0x1d, 0x80, 0xb2, 0x08, 0xda, 0x23,
	0xd3, 0xb0, 0x71, 0x73, 0x05, 0xed, 0xc3, 0xae, 0x8c, 0x1a, 0xa6, 0x43, 0xfa, 0xe7, 0xcd, 0x55,
	0xa4, 0xc2, 0x91, 0x80, 0xf5, 0x34, 0xc7, 0x19, 0xe0, 0xaa, 0x7a, 0x1d, 0xbd, 0x84, 0xe7, 0x8f,
	0xd6, 0xf0, 0x26, 0x0d, 0xf4, 0x02, 0x0e, 0x85, 0xa2, 0xe1, 0x78, 0xe0, 0x90, 0xaa, 0x0e, 0x48,
	0xbd, 0x2a, 0x25, 0x5c, 0x66, 0xab, 0xb0, 0xd9, 0xb7, 0x08, 0x36, 0xce, 0xaa, 0x02, 0x7e, 0x61,
	0xb3, 0x0a, 0x72, 0x2a, 0x2d, 0x6c, 0x72, 0x54, 0xb3, 0x3f, 0x96, 0xcc, 0x4b, 0x49, 0x36, 0xc7,
	0x38, 0xf1, 0x4a, 0x92, 0xd5, 0x74, 0x0b, 0xe3, 0x92, 0xfa, 0x15, 0x1d, 0xb2, 0x87, 0x2a, 0xa3,
	0x9c, 0x1c, 0x48, 0xb0, 0x85, 0xfb, 0x63, 0x7b, 0xce, 0xfe, 0x86, 0x8e, 0xd8, 0xf3, 0x59, 0x80,
	0x39, 0xfd, 0x7a, 0x81, 0x3e, 0x34, 0x27, 0x73, 0x7a, 0xb8, 0x40, 0xe7, 0x30, 0xa7, 0xdf, 0x48,
	0xa3, 0x73, 0x9c, 0x1f, 0x37, 0x2a, 0x16, 0xce, 0x51, 0x9d, 0x4c, 0xb0, 0x3b, 0x32, 0x89, 0xe1,
	0xd8, 0x65, 0x87, 0xb8, 0x38, 0xee, 0xd2, 0x1a, 0xde, 0x66, 0x5a, 0x1c, 0xb7, 0x28, 0xc2, 0x8e,
	0xac, 0x73, 0x2b, 0xf7, 0x12, 0x4b, 0xb8, 0x4c, 0x82, 0x5e, 0xe5, 0xbd, 0x4e, 0x35, 0xcb, 0x20,
	0x13, 0x6d, 0xe0, 0xf6, 0x1e, 0x3e, 0xa3, 0x09, 0x71, 0xf0, 0x59, 0x29, 0xf4, 0x6b, 0x1d, 0xbd,
	0x86, 0xf6, 0xe3, 0x55, 0x5c, 0xeb, 0xf7, 0xfa, 0x97, 0x35, 0xf6, 0xef, 0xf1, 0xee, 0x4f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x05, 0x40, 0x99, 0x77, 0x52, 0x06, 0x00, 0x00,
}
