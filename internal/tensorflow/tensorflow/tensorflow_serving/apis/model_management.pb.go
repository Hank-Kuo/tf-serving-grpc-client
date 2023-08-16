// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow_serving/apis/model_management.proto

package apis

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	config "grpc-client/internal/tensorflow/tensorflow/tensorflow_serving/config"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReloadConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *config.ModelServerConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *ReloadConfigRequest) Reset() {
	*x = ReloadConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_model_management_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReloadConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadConfigRequest) ProtoMessage() {}

func (x *ReloadConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_model_management_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadConfigRequest.ProtoReflect.Descriptor instead.
func (*ReloadConfigRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_model_management_proto_rawDescGZIP(), []int{0}
}

func (x *ReloadConfigRequest) GetConfig() *config.ModelServerConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

type ReloadConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *StatusProto `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ReloadConfigResponse) Reset() {
	*x = ReloadConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_model_management_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReloadConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReloadConfigResponse) ProtoMessage() {}

func (x *ReloadConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_model_management_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReloadConfigResponse.ProtoReflect.Descriptor instead.
func (*ReloadConfigResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_model_management_proto_rawDescGZIP(), []int{1}
}

func (x *ReloadConfigResponse) GetStatus() *StatusProto {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_tensorflow_serving_apis_model_management_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_model_management_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x1a, 0x24, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x54, 0x0a, 0x13, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x4f, 0x0a, 0x14, 0x52, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0xde, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x42, 0x14, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x12, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0xe2, 0x02, 0x1e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_apis_model_management_proto_rawDescOnce sync.Once
	file_tensorflow_serving_apis_model_management_proto_rawDescData = file_tensorflow_serving_apis_model_management_proto_rawDesc
)

func file_tensorflow_serving_apis_model_management_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_apis_model_management_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_apis_model_management_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_apis_model_management_proto_rawDescData)
	})
	return file_tensorflow_serving_apis_model_management_proto_rawDescData
}

var file_tensorflow_serving_apis_model_management_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_serving_apis_model_management_proto_goTypes = []interface{}{
	(*ReloadConfigRequest)(nil),      // 0: tensorflow.serving.ReloadConfigRequest
	(*ReloadConfigResponse)(nil),     // 1: tensorflow.serving.ReloadConfigResponse
	(*config.ModelServerConfig)(nil), // 2: tensorflow.serving.ModelServerConfig
	(*StatusProto)(nil),              // 3: tensorflow.serving.StatusProto
}
var file_tensorflow_serving_apis_model_management_proto_depIdxs = []int32{
	2, // 0: tensorflow.serving.ReloadConfigRequest.config:type_name -> tensorflow.serving.ModelServerConfig
	3, // 1: tensorflow.serving.ReloadConfigResponse.status:type_name -> tensorflow.serving.StatusProto
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_model_management_proto_init() }
func file_tensorflow_serving_apis_model_management_proto_init() {
	if File_tensorflow_serving_apis_model_management_proto != nil {
		return
	}
	file_tensorflow_serving_apis_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_apis_model_management_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReloadConfigRequest); i {
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
		file_tensorflow_serving_apis_model_management_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReloadConfigResponse); i {
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
			RawDescriptor: file_tensorflow_serving_apis_model_management_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_apis_model_management_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_model_management_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_apis_model_management_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_apis_model_management_proto = out.File
	file_tensorflow_serving_apis_model_management_proto_rawDesc = nil
	file_tensorflow_serving_apis_model_management_proto_goTypes = nil
	file_tensorflow_serving_apis_model_management_proto_depIdxs = nil
}
