// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: auditpb/message.proto

package auditpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Audit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Action    string                 `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	Actor     string                 `protobuf:"bytes,3,opt,name=actor,proto3" json:"actor,omitempty"`
	Entity    *Entity                `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	SourceIp  string                 `protobuf:"bytes,5,opt,name=sourceIp,proto3" json:"sourceIp,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Audit) Reset() {
	*x = Audit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auditpb_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audit) ProtoMessage() {}

func (x *Audit) ProtoReflect() protoreflect.Message {
	mi := &file_auditpb_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audit.ProtoReflect.Descriptor instead.
func (*Audit) Descriptor() ([]byte, []int) {
	return file_auditpb_message_proto_rawDescGZIP(), []int{0}
}

func (x *Audit) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Audit) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Audit) GetActor() string {
	if x != nil {
		return x.Actor
	}
	return ""
}

func (x *Audit) GetEntity() *Entity {
	if x != nil {
		return x.Entity
	}
	return nil
}

func (x *Audit) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

func (x *Audit) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Entity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Entity) Reset() {
	*x = Entity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auditpb_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entity) ProtoMessage() {}

func (x *Entity) ProtoReflect() protoreflect.Message {
	mi := &file_auditpb_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entity.ProtoReflect.Descriptor instead.
func (*Entity) Descriptor() ([]byte, []int) {
	return file_auditpb_message_proto_rawDescGZIP(), []int{1}
}

func (x *Entity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Entity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_auditpb_message_proto protoreflect.FileDescriptor

var file_auditpb_message_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x64, 0x69, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x75, 0x64, 0x69, 0x74, 0x70, 0x62,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x70, 0x12, 0x38,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2c, 0x0a, 0x06, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x8d, 0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x70, 0x62, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x34, 0x76, 0x6f, 0x63, 0x61, 0x62, 0x2d, 0x62, 0x6f,
	0x6f, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x70, 0x62, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x41,
	0x58, 0x58, 0xaa, 0x02, 0x07, 0x41, 0x75, 0x64, 0x69, 0x74, 0x70, 0x62, 0xca, 0x02, 0x07, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x70, 0x62, 0xe2, 0x02, 0x13, 0x41, 0x75, 0x64, 0x69, 0x74, 0x70, 0x62,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auditpb_message_proto_rawDescOnce sync.Once
	file_auditpb_message_proto_rawDescData = file_auditpb_message_proto_rawDesc
)

func file_auditpb_message_proto_rawDescGZIP() []byte {
	file_auditpb_message_proto_rawDescOnce.Do(func() {
		file_auditpb_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_auditpb_message_proto_rawDescData)
	})
	return file_auditpb_message_proto_rawDescData
}

var file_auditpb_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auditpb_message_proto_goTypes = []interface{}{
	(*Audit)(nil),                 // 0: auditpb.Audit
	(*Entity)(nil),                // 1: auditpb.Entity
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_auditpb_message_proto_depIdxs = []int32{
	1, // 0: auditpb.Audit.entity:type_name -> auditpb.Entity
	2, // 1: auditpb.Audit.createdAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auditpb_message_proto_init() }
func file_auditpb_message_proto_init() {
	if File_auditpb_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auditpb_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audit); i {
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
		file_auditpb_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entity); i {
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
			RawDescriptor: file_auditpb_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auditpb_message_proto_goTypes,
		DependencyIndexes: file_auditpb_message_proto_depIdxs,
		MessageInfos:      file_auditpb_message_proto_msgTypes,
	}.Build()
	File_auditpb_message_proto = out.File
	file_auditpb_message_proto_rawDesc = nil
	file_auditpb_message_proto_goTypes = nil
	file_auditpb_message_proto_depIdxs = nil
}
