// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: staffpb/message.proto

package staffpb

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

type Staff struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	IsAdmin bool   `protobuf:"varint,4,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
}

func (x *Staff) Reset() {
	*x = Staff{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staffpb_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Staff) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Staff) ProtoMessage() {}

func (x *Staff) ProtoReflect() protoreflect.Message {
	mi := &file_staffpb_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Staff.ProtoReflect.Descriptor instead.
func (*Staff) Descriptor() ([]byte, []int) {
	return file_staffpb_message_proto_rawDescGZIP(), []int{0}
}

func (x *Staff) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Staff) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Staff) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Staff) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

var File_staffpb_message_proto protoreflect.FileDescriptor

var file_staffpb_message_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x61, 0x66, 0x66, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x61, 0x66, 0x66, 0x70, 0x62,
	0x22, 0x5b, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x42, 0x8d, 0x01,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x70, 0x62, 0x42, 0x0c, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x76,
	0x6f, 0x63, 0x61, 0x62, 0x2d, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2d, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x74,
	0x61, 0x66, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x61, 0x66,
	0x66, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x53, 0x74, 0x61, 0x66,
	0x66, 0x70, 0x62, 0xca, 0x02, 0x07, 0x53, 0x74, 0x61, 0x66, 0x66, 0x70, 0x62, 0xe2, 0x02, 0x13,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x53, 0x74, 0x61, 0x66, 0x66, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staffpb_message_proto_rawDescOnce sync.Once
	file_staffpb_message_proto_rawDescData = file_staffpb_message_proto_rawDesc
)

func file_staffpb_message_proto_rawDescGZIP() []byte {
	file_staffpb_message_proto_rawDescOnce.Do(func() {
		file_staffpb_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_staffpb_message_proto_rawDescData)
	})
	return file_staffpb_message_proto_rawDescData
}

var file_staffpb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_staffpb_message_proto_goTypes = []interface{}{
	(*Staff)(nil), // 0: staffpb.Staff
}
var file_staffpb_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_staffpb_message_proto_init() }
func file_staffpb_message_proto_init() {
	if File_staffpb_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_staffpb_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Staff); i {
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
			RawDescriptor: file_staffpb_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_staffpb_message_proto_goTypes,
		DependencyIndexes: file_staffpb_message_proto_depIdxs,
		MessageInfos:      file_staffpb_message_proto_msgTypes,
	}.Build()
	File_staffpb_message_proto = out.File
	file_staffpb_message_proto_rawDesc = nil
	file_staffpb_message_proto_goTypes = nil
	file_staffpb_message_proto_depIdxs = nil
}
