// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: comms/comms.proto

package comms

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

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_comms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_comms_comms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_comms_comms_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comms_comms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_comms_comms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_comms_comms_proto_rawDescGZIP(), []int{1}
}

func (x *Ack) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_comms_comms_proto protoreflect.FileDescriptor

var file_comms_comms_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6e, 0x73, 0x22, 0x1f, 0x0a, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x1f, 0x0a, 0x03,
	0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x31, 0x0a,
	0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x20,
	0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6e, 0x73, 0x2e,
	0x4d, 0x73, 0x67, 0x1a, 0x0b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6e, 0x73, 0x2e, 0x41, 0x63, 0x6b,
	0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x77, 0x64, 0x6f, 0x77, 0x69, 0x63, 0x7a, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x73, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comms_comms_proto_rawDescOnce sync.Once
	file_comms_comms_proto_rawDescData = file_comms_comms_proto_rawDesc
)

func file_comms_comms_proto_rawDescGZIP() []byte {
	file_comms_comms_proto_rawDescOnce.Do(func() {
		file_comms_comms_proto_rawDescData = protoimpl.X.CompressGZIP(file_comms_comms_proto_rawDescData)
	})
	return file_comms_comms_proto_rawDescData
}

var file_comms_comms_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_comms_comms_proto_goTypes = []any{
	(*Msg)(nil), // 0: commns.Msg
	(*Ack)(nil), // 1: commns.Ack
}
var file_comms_comms_proto_depIdxs = []int32{
	0, // 0: commns.PersonService.Post:input_type -> commns.Msg
	1, // 1: commns.PersonService.Post:output_type -> commns.Ack
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_comms_comms_proto_init() }
func file_comms_comms_proto_init() {
	if File_comms_comms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comms_comms_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Msg); i {
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
		file_comms_comms_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Ack); i {
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
			RawDescriptor: file_comms_comms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comms_comms_proto_goTypes,
		DependencyIndexes: file_comms_comms_proto_depIdxs,
		MessageInfos:      file_comms_comms_proto_msgTypes,
	}.Build()
	File_comms_comms_proto = out.File
	file_comms_comms_proto_rawDesc = nil
	file_comms_comms_proto_goTypes = nil
	file_comms_comms_proto_depIdxs = nil
}
