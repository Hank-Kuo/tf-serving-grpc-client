// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/protobuf/tpu/compilation_result.proto

package tpu

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	service "grpc-client/internal/tensorflow/tensorflow/tensorflow/compiler/xla/service"
	_ "grpc-client/internal/tensorflow/tensorflow/tensorflow/core/protobuf"
	protobuf "grpc-client/internal/tensorflow/tensorflow/tensorflow/tsl/protobuf"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CompilationResultProto_ErrorCode int32

const (
	CompilationResultProto_UNKNOWN       CompilationResultProto_ErrorCode = 0
	CompilationResultProto_OUT_OF_MEMORY CompilationResultProto_ErrorCode = 1
)

// Enum value maps for CompilationResultProto_ErrorCode.
var (
	CompilationResultProto_ErrorCode_name = map[int32]string{
		0: "UNKNOWN",
		1: "OUT_OF_MEMORY",
	}
	CompilationResultProto_ErrorCode_value = map[string]int32{
		"UNKNOWN":       0,
		"OUT_OF_MEMORY": 1,
	}
)

func (x CompilationResultProto_ErrorCode) Enum() *CompilationResultProto_ErrorCode {
	p := new(CompilationResultProto_ErrorCode)
	*p = x
	return p
}

func (x CompilationResultProto_ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompilationResultProto_ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_protobuf_tpu_compilation_result_proto_enumTypes[0].Descriptor()
}

func (CompilationResultProto_ErrorCode) Type() protoreflect.EnumType {
	return &file_tensorflow_core_protobuf_tpu_compilation_result_proto_enumTypes[0]
}

func (x CompilationResultProto_ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompilationResultProto_ErrorCode.Descriptor instead.
func (CompilationResultProto_ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDescGZIP(), []int{0, 0}
}

// Describes the result of a TPU compilation. This is also used as TPU
// compilation result status payload.
// URI: "type.googleapis.com/tensorflow.tpu.CompilationResultProto"
type CompilationResultProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The error message, if any, returned during compilation.
	StatusCode         protobuf.Code `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3,enum=tensorflow.error.Code" json:"status_code,omitempty"`
	StatusErrorMessage string        `protobuf:"bytes,2,opt,name=status_error_message,json=statusErrorMessage,proto3" json:"status_error_message,omitempty"`
	// HLO proto.
	HloProtos []*service.HloProto              `protobuf:"bytes,3,rep,name=hlo_protos,json=hloProtos,proto3" json:"hlo_protos,omitempty"`
	ErrorCode CompilationResultProto_ErrorCode `protobuf:"varint,4,opt,name=error_code,json=errorCode,proto3,enum=tensorflow.tpu.CompilationResultProto_ErrorCode" json:"error_code,omitempty"`
}

func (x *CompilationResultProto) Reset() {
	*x = CompilationResultProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_tpu_compilation_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompilationResultProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompilationResultProto) ProtoMessage() {}

func (x *CompilationResultProto) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_tpu_compilation_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompilationResultProto.ProtoReflect.Descriptor instead.
func (*CompilationResultProto) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDescGZIP(), []int{0}
}

func (x *CompilationResultProto) GetStatusCode() protobuf.Code {
	if x != nil {
		return x.StatusCode
	}
	return protobuf.Code(0)
}

func (x *CompilationResultProto) GetStatusErrorMessage() string {
	if x != nil {
		return x.StatusErrorMessage
	}
	return ""
}

func (x *CompilationResultProto) GetHloProtos() []*service.HloProto {
	if x != nil {
		return x.HloProtos
	}
	return nil
}

func (x *CompilationResultProto) GetErrorCode() CompilationResultProto_ErrorCode {
	if x != nil {
		return x.ErrorCode
	}
	return CompilationResultProto_UNKNOWN
}

var File_tensorflow_core_protobuf_tpu_compilation_result_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDesc = []byte{
	0x0a, 0x35, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x70, 0x75, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x70, 0x75, 0x1a, 0x29, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x68, 0x6c, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf,
	0x02, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x68, 0x6c, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x48,
	0x6c, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x09, 0x68, 0x6c, 0x6f, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x12, 0x4f, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x70, 0x75, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x2b, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x4d, 0x45, 0x4d, 0x4f, 0x52, 0x59, 0x10, 0x01,
	0x42, 0xd1, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x74, 0x70, 0x75, 0x42, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x47, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x70, 0x75, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03,
	0x54, 0x54, 0x58, 0xaa, 0x02, 0x0e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x54, 0x70, 0x75, 0xca, 0x02, 0x0e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5c, 0x54, 0x70, 0x75, 0xe2, 0x02, 0x1a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x5c, 0x54, 0x70, 0x75, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a,
	0x3a, 0x54, 0x70, 0x75, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDescData = file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDesc
)

func file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDescData
}

var file_tensorflow_core_protobuf_tpu_compilation_result_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_protobuf_tpu_compilation_result_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_tpu_compilation_result_proto_goTypes = []interface{}{
	(CompilationResultProto_ErrorCode)(0), // 0: tensorflow.tpu.CompilationResultProto.ErrorCode
	(*CompilationResultProto)(nil),        // 1: tensorflow.tpu.CompilationResultProto
	(protobuf.Code)(0),                    // 2: tensorflow.error.Code
	(*service.HloProto)(nil),              // 3: xla.HloProto
}
var file_tensorflow_core_protobuf_tpu_compilation_result_proto_depIdxs = []int32{
	2, // 0: tensorflow.tpu.CompilationResultProto.status_code:type_name -> tensorflow.error.Code
	3, // 1: tensorflow.tpu.CompilationResultProto.hlo_protos:type_name -> xla.HloProto
	0, // 2: tensorflow.tpu.CompilationResultProto.error_code:type_name -> tensorflow.tpu.CompilationResultProto.ErrorCode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_tpu_compilation_result_proto_init() }
func file_tensorflow_core_protobuf_tpu_compilation_result_proto_init() {
	if File_tensorflow_core_protobuf_tpu_compilation_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_tpu_compilation_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompilationResultProto); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_tpu_compilation_result_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_tpu_compilation_result_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_protobuf_tpu_compilation_result_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_protobuf_tpu_compilation_result_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_tpu_compilation_result_proto = out.File
	file_tensorflow_core_protobuf_tpu_compilation_result_proto_rawDesc = nil
	file_tensorflow_core_protobuf_tpu_compilation_result_proto_goTypes = nil
	file_tensorflow_core_protobuf_tpu_compilation_result_proto_depIdxs = nil
}
