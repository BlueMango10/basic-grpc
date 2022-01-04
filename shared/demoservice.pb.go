// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: demoservice.proto

package shared

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

type Finger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Finger) Reset() {
	*x = Finger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demoservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Finger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Finger) ProtoMessage() {}

func (x *Finger) ProtoReflect() protoreflect.Message {
	mi := &file_demoservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Finger.ProtoReflect.Descriptor instead.
func (*Finger) Descriptor() ([]byte, []int) {
	return file_demoservice_proto_rawDescGZIP(), []int{0}
}

func (x *Finger) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Reaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Insults []string `protobuf:"bytes,1,rep,name=insults,proto3" json:"insults,omitempty"`
}

func (x *Reaction) Reset() {
	*x = Reaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demoservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reaction) ProtoMessage() {}

func (x *Reaction) ProtoReflect() protoreflect.Message {
	mi := &file_demoservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reaction.ProtoReflect.Descriptor instead.
func (*Reaction) Descriptor() ([]byte, []int) {
	return file_demoservice_proto_rawDescGZIP(), []int{1}
}

func (x *Reaction) GetInsults() []string {
	if x != nil {
		return x.Insults
	}
	return nil
}

var File_demoservice_proto protoreflect.FileDescriptor

var file_demoservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x64, 0x65, 0x6d, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x06, 0x46, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x24, 0x0a, 0x08, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x69,
	0x6e, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6e,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x32, 0x29, 0x0a, 0x0b, 0x44, 0x65, 0x6d, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x04, 0x50, 0x6f, 0x6b, 0x65, 0x12, 0x07, 0x2e, 0x46,
	0x69, 0x6e, 0x67, 0x65, 0x72, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x1d, 0x5a, 0x1b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demoservice_proto_rawDescOnce sync.Once
	file_demoservice_proto_rawDescData = file_demoservice_proto_rawDesc
)

func file_demoservice_proto_rawDescGZIP() []byte {
	file_demoservice_proto_rawDescOnce.Do(func() {
		file_demoservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_demoservice_proto_rawDescData)
	})
	return file_demoservice_proto_rawDescData
}

var file_demoservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_demoservice_proto_goTypes = []interface{}{
	(*Finger)(nil),   // 0: Finger
	(*Reaction)(nil), // 1: Reaction
}
var file_demoservice_proto_depIdxs = []int32{
	0, // 0: DemoService.Poke:input_type -> Finger
	1, // 1: DemoService.Poke:output_type -> Reaction
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demoservice_proto_init() }
func file_demoservice_proto_init() {
	if File_demoservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demoservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Finger); i {
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
		file_demoservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reaction); i {
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
			RawDescriptor: file_demoservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demoservice_proto_goTypes,
		DependencyIndexes: file_demoservice_proto_depIdxs,
		MessageInfos:      file_demoservice_proto_msgTypes,
	}.Build()
	File_demoservice_proto = out.File
	file_demoservice_proto_rawDesc = nil
	file_demoservice_proto_goTypes = nil
	file_demoservice_proto_depIdxs = nil
}
