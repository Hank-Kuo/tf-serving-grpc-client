// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/profiler/profiler_options.proto

package profiler

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	protobuf "grpc-client/internal/tensorflow/tensorflow/tensorflow/tsl/profiler/protobuf"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of tensorflow/tsl/profiler/protobuf/profiler_options.proto.

type ProfileOptions_DeviceType = protobuf.ProfileOptions_DeviceType

const ProfileOptions_UNSPECIFIED = protobuf.ProfileOptions_UNSPECIFIED
const ProfileOptions_CPU = protobuf.ProfileOptions_CPU
const ProfileOptions_GPU = protobuf.ProfileOptions_GPU
const ProfileOptions_TPU = protobuf.ProfileOptions_TPU
const ProfileOptions_PLUGGABLE_DEVICE = protobuf.ProfileOptions_PLUGGABLE_DEVICE

var ProfileOptions_DeviceType_name = protobuf.ProfileOptions_DeviceType_name
var ProfileOptions_DeviceType_value = protobuf.ProfileOptions_DeviceType_value

type ProfileOptions = protobuf.ProfileOptions
type RemoteProfilerSessionManagerOptions = protobuf.RemoteProfilerSessionManagerOptions

var File_tensorflow_core_profiler_profiler_options_proto protoreflect.FileDescriptor

var file_tensorflow_core_profiler_profiler_options_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x75,
	0x6d, 0x6d, 0x79, 0x1a, 0x37, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0xd2, 0x01, 0x0a,
	0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x64, 0x75, 0x6d, 0x6d, 0x79, 0x42, 0x14, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x72, 0xa2, 0x02, 0x03, 0x54, 0x44, 0x58, 0xaa, 0x02, 0x10, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0xca, 0x02, 0x10, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0xe2, 0x02,
	0x1c, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x44, 0x75, 0x6d, 0x6d,
	0x79, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tensorflow_core_profiler_profiler_options_proto_goTypes = []interface{}{}
var file_tensorflow_core_profiler_profiler_options_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_core_profiler_profiler_options_proto_init() }
func file_tensorflow_core_profiler_profiler_options_proto_init() {
	if File_tensorflow_core_profiler_profiler_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_profiler_profiler_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_profiler_profiler_options_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_profiler_profiler_options_proto_depIdxs,
	}.Build()
	File_tensorflow_core_profiler_profiler_options_proto = out.File
	file_tensorflow_core_profiler_profiler_options_proto_rawDesc = nil
	file_tensorflow_core_profiler_profiler_options_proto_goTypes = nil
	file_tensorflow_core_profiler_profiler_options_proto_depIdxs = nil
}
