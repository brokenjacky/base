// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v5.26.1
// source: msg.proto

package msgBox

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

type MsgBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType string `protobuf:"bytes,1,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	UUID    string `protobuf:"bytes,2,opt,name=UUID,proto3" json:"UUID,omitempty"`
	PayLoad []byte `protobuf:"bytes,3,opt,name=PayLoad,proto3" json:"PayLoad,omitempty"`
}

func (x *MsgBody) Reset() {
	*x = MsgBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgBody) ProtoMessage() {}

func (x *MsgBody) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgBody.ProtoReflect.Descriptor instead.
func (*MsgBody) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *MsgBody) GetMsgType() string {
	if x != nil {
		return x.MsgType
	}
	return ""
}

func (x *MsgBody) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *MsgBody) GetPayLoad() []byte {
	if x != nil {
		return x.PayLoad
	}
	return nil
}

type MailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    int32  `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *MailReq) Reset() {
	*x = MailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailReq) ProtoMessage() {}

func (x *MailReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailReq.ProtoReflect.Descriptor instead.
func (*MailReq) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *MailReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *MailReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x73, 0x67,
	0x42, 0x6f, 0x78, 0x22, 0x51, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x50, 0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50,
	0x61, 0x79, 0x4c, 0x6f, 0x61, 0x64, 0x22, 0x37, 0x0a, 0x07, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32,
	0x3a, 0x0a, 0x0a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x78, 0x12, 0x2c, 0x0a,
	0x06, 0x55, 0x6e, 0x69, 0x4d, 0x73, 0x67, 0x12, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x42, 0x6f, 0x78,
	0x2e, 0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x42, 0x6f,
	0x78, 0x2e, 0x4d, 0x73, 0x67, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e,
	0x2f, 0x6d, 0x73, 0x67, 0x42, 0x6f, 0x78, 0x3b, 0x6d, 0x73, 0x67, 0x42, 0x6f, 0x78, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msg_proto_goTypes = []interface{}{
	(*MsgBody)(nil), // 0: msgBox.MsgBody
	(*MailReq)(nil), // 1: msgBox.MailReq
}
var file_msg_proto_depIdxs = []int32{
	0, // 0: msgBox.MessageBox.UniMsg:input_type -> msgBox.MsgBody
	0, // 1: msgBox.MessageBox.UniMsg:output_type -> msgBox.MsgBody
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgBody); i {
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
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailReq); i {
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
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
