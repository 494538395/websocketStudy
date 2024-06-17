// 多人在线游戏引擎交互协议定义

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.3
// source: engine.proto

package roompb

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

// 消息类型
type MessageType int32

const (
	MessageType_NONE      MessageType = 0
	MessageType_STATE     MessageType = 1 // 状态消息
	MessageType_OPERATION MessageType = 2 // 操作消息, c=>s
	MessageType_REQUEST   MessageType = 3 // 请求消息
	MessageType_ACK       MessageType = 4 // 确认消息
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "NONE",
		1: "STATE",
		2: "OPERATION",
		3: "REQUEST",
		4: "ACK",
	}
	MessageType_value = map[string]int32{
		"NONE":      0,
		"STATE":     1,
		"OPERATION": 2,
		"REQUEST":   3,
		"ACK":       4,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_engine_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_engine_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{0}
}

// 当前帧对应消息列表
type FrameMsgPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameId  uint32     `protobuf:"varint,1,opt,name=frameId,proto3" json:"frameId,omitempty"`   // 服务器当前帧id
	MsgList  []*Message `protobuf:"bytes,2,rep,name=msgList,proto3" json:"msgList,omitempty"`    // 消息列表
	SendTime int64      `protobuf:"varint,3,opt,name=sendTime,proto3" json:"sendTime,omitempty"` // 发送时间, 单位=ms
}

func (x *FrameMsgPacket) Reset() {
	*x = FrameMsgPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrameMsgPacket) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrameMsgPacket) ProtoMessage() {}

func (x *FrameMsgPacket) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrameMsgPacket.ProtoReflect.Descriptor instead.
func (*FrameMsgPacket) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{0}
}

func (x *FrameMsgPacket) GetFrameId() uint32 {
	if x != nil {
		return x.FrameId
	}
	return 0
}

func (x *FrameMsgPacket) GetMsgList() []*Message {
	if x != nil {
		return x.MsgList
	}
	return nil
}

func (x *FrameMsgPacket) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

// 消息体
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType MessageType `protobuf:"varint,1,opt,name=msgType,proto3,enum=engine.MessageType" json:"msgType,omitempty"` // 消息类型
	Id      string      `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`                                    // 消息id
	Event   string      `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`                              // 事件名
	Data    []byte      `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`                                // 消息体
	UserId  string      `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`                            // 用户id
	FrameId uint32      `protobuf:"varint,6,opt,name=frameId,proto3" json:"frameId,omitempty"`                         // 消息帧id
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[1]
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
	return file_engine_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetMsgType() MessageType {
	if x != nil {
		return x.MsgType
	}
	return MessageType_NONE
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Message) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Message) GetFrameId() uint32 {
	if x != nil {
		return x.FrameId
	}
	return 0
}

// 系统设置
type LaunchInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fps             int32 `protobuf:"varint,1,opt,name=fps,proto3" json:"fps,omitempty"`                         // 帧率
	StartTime       int64 `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`             // 启动时间, 单位=ms
	FrameWindowSize int32 `protobuf:"varint,3,opt,name=frameWindowSize,proto3" json:"frameWindowSize,omitempty"` // 帧窗口大小
	MaxPredictNum   int32 `protobuf:"varint,4,opt,name=maxPredictNum,proto3" json:"maxPredictNum,omitempty"`     // 最大预测帧数
	MsgDelayOffset  int32 `protobuf:"varint,5,opt,name=msgDelayOffset,proto3" json:"msgDelayOffset,omitempty"`   // 消息延迟偏移, 单位帧数
}

func (x *LaunchInfo) Reset() {
	*x = LaunchInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_engine_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LaunchInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaunchInfo) ProtoMessage() {}

func (x *LaunchInfo) ProtoReflect() protoreflect.Message {
	mi := &file_engine_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaunchInfo.ProtoReflect.Descriptor instead.
func (*LaunchInfo) Descriptor() ([]byte, []int) {
	return file_engine_proto_rawDescGZIP(), []int{2}
}

func (x *LaunchInfo) GetFps() int32 {
	if x != nil {
		return x.Fps
	}
	return 0
}

func (x *LaunchInfo) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *LaunchInfo) GetFrameWindowSize() int32 {
	if x != nil {
		return x.FrameWindowSize
	}
	return 0
}

func (x *LaunchInfo) GetMaxPredictNum() int32 {
	if x != nil {
		return x.MaxPredictNum
	}
	return 0
}

func (x *LaunchInfo) GetMsgDelayOffset() int32 {
	if x != nil {
		return x.MsgDelayOffset
	}
	return 0
}

var File_engine_proto protoreflect.FileDescriptor

var file_engine_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x22, 0x71, 0x0a, 0x0e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x4d,
	0x73, 0x67, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x61, 0x6d,
	0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x49, 0x64, 0x12, 0x29, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xa4, 0x01, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x22, 0xb4, 0x01, 0x0a, 0x0a, 0x4c, 0x61, 0x75, 0x6e, 0x63, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x10, 0x0a, 0x03, 0x66, 0x70, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x66, 0x70,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x53, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x57,
	0x69, 0x6e, 0x64, 0x6f, 0x77, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x61, 0x78,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0d, 0x6d, 0x61, 0x78, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x12,
	0x26, 0x0a, 0x0e, 0x6d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x61, 0x79, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x73, 0x67, 0x44, 0x65, 0x6c, 0x61,
	0x79, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x2a, 0x47, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45,
	0x51, 0x55, 0x45, 0x53, 0x54, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x43, 0x4b, 0x10, 0x04,
	0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x72, 0x6f, 0x6f, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_engine_proto_rawDescOnce sync.Once
	file_engine_proto_rawDescData = file_engine_proto_rawDesc
)

func file_engine_proto_rawDescGZIP() []byte {
	file_engine_proto_rawDescOnce.Do(func() {
		file_engine_proto_rawDescData = protoimpl.X.CompressGZIP(file_engine_proto_rawDescData)
	})
	return file_engine_proto_rawDescData
}

var file_engine_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_engine_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_engine_proto_goTypes = []interface{}{
	(MessageType)(0),       // 0: engine.MessageType
	(*FrameMsgPacket)(nil), // 1: engine.FrameMsgPacket
	(*Message)(nil),        // 2: engine.Message
	(*LaunchInfo)(nil),     // 3: engine.LaunchInfo
}
var file_engine_proto_depIdxs = []int32{
	2, // 0: engine.FrameMsgPacket.msgList:type_name -> engine.Message
	0, // 1: engine.Message.msgType:type_name -> engine.MessageType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_engine_proto_init() }
func file_engine_proto_init() {
	if File_engine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_engine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrameMsgPacket); i {
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
		file_engine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_engine_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LaunchInfo); i {
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
			RawDescriptor: file_engine_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_engine_proto_goTypes,
		DependencyIndexes: file_engine_proto_depIdxs,
		EnumInfos:         file_engine_proto_enumTypes,
		MessageInfos:      file_engine_proto_msgTypes,
	}.Build()
	File_engine_proto = out.File
	file_engine_proto_rawDesc = nil
	file_engine_proto_goTypes = nil
	file_engine_proto_depIdxs = nil
}