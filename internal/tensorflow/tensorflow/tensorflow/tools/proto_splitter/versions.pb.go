// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/tools/proto_splitter/versions.proto

package proto_splitter

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

// Version information for Chunked protos.
type VersionDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Information about the Splitter used to split and write this data.
	SplitterVersion int32 `protobuf:"varint,1,opt,name=splitter_version,json=splitterVersion,proto3" json:"splitter_version,omitempty"`
	// The version of the Join implemention required to consume this data.
	JoinVersion int32 `protobuf:"varint,2,opt,name=join_version,json=joinVersion,proto3" json:"join_version,omitempty"`
	// Specific consumer versions which are disallowed (e.g. due to bugs).
	BadConsumers []int32 `protobuf:"varint,3,rep,packed,name=bad_consumers,json=badConsumers,proto3" json:"bad_consumers,omitempty"`
}

func (x *VersionDef) Reset() {
	*x = VersionDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tools_proto_splitter_versions_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionDef) ProtoMessage() {}

func (x *VersionDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tools_proto_splitter_versions_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionDef.ProtoReflect.Descriptor instead.
func (*VersionDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_tools_proto_splitter_versions_proto_rawDescGZIP(), []int{0}
}

func (x *VersionDef) GetSplitterVersion() int32 {
	if x != nil {
		return x.SplitterVersion
	}
	return 0
}

func (x *VersionDef) GetJoinVersion() int32 {
	if x != nil {
		return x.JoinVersion
	}
	return 0
}

func (x *VersionDef) GetBadConsumers() []int32 {
	if x != nil {
		return x.BadConsumers
	}
	return nil
}

var File_tensorflow_tools_proto_splitter_versions_proto protoreflect.FileDescriptor

var file_tensorflow_tools_proto_splitter_versions_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x6f, 0x6f,
	0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72,
	0x22, 0x7f, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x12, 0x29,
	0x0a, 0x10, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x6a, 0x6f, 0x69,
	0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x6a, 0x6f, 0x69, 0x6e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x62, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x0c, 0x62, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x73, 0x42, 0xc6, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x73, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x42, 0x0d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x70, 0x6c,
	0x69, 0x74, 0x74, 0x65, 0x72, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02,
	0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0xca, 0x02,
	0x0d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0xe2, 0x02,
	0x19, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tensorflow_tools_proto_splitter_versions_proto_rawDescOnce sync.Once
	file_tensorflow_tools_proto_splitter_versions_proto_rawDescData = file_tensorflow_tools_proto_splitter_versions_proto_rawDesc
)

func file_tensorflow_tools_proto_splitter_versions_proto_rawDescGZIP() []byte {
	file_tensorflow_tools_proto_splitter_versions_proto_rawDescOnce.Do(func() {
		file_tensorflow_tools_proto_splitter_versions_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tools_proto_splitter_versions_proto_rawDescData)
	})
	return file_tensorflow_tools_proto_splitter_versions_proto_rawDescData
}

var file_tensorflow_tools_proto_splitter_versions_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_tools_proto_splitter_versions_proto_goTypes = []interface{}{
	(*VersionDef)(nil), // 0: proto_splitter.VersionDef
}
var file_tensorflow_tools_proto_splitter_versions_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_tools_proto_splitter_versions_proto_init() }
func file_tensorflow_tools_proto_splitter_versions_proto_init() {
	if File_tensorflow_tools_proto_splitter_versions_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tools_proto_splitter_versions_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionDef); i {
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
			RawDescriptor: file_tensorflow_tools_proto_splitter_versions_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tools_proto_splitter_versions_proto_goTypes,
		DependencyIndexes: file_tensorflow_tools_proto_splitter_versions_proto_depIdxs,
		MessageInfos:      file_tensorflow_tools_proto_splitter_versions_proto_msgTypes,
	}.Build()
	File_tensorflow_tools_proto_splitter_versions_proto = out.File
	file_tensorflow_tools_proto_splitter_versions_proto_rawDesc = nil
	file_tensorflow_tools_proto_splitter_versions_proto_goTypes = nil
	file_tensorflow_tools_proto_splitter_versions_proto_depIdxs = nil
}
