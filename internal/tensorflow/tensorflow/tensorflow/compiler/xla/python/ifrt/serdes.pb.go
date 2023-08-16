// Copyright 2023 The TensorFlow Authors. All Rights Reserved.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//==============================================================================

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/xla/python/ifrt/serdes.proto

package ifrt

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

// Wire format for objects serialized from `Serializable`.
type Serialized struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeName string `protobuf:"bytes,1,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	Data     []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Serialized) Reset() {
	*x = Serialized{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_xla_python_ifrt_serdes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Serialized) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Serialized) ProtoMessage() {}

func (x *Serialized) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_xla_python_ifrt_serdes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Serialized.ProtoReflect.Descriptor instead.
func (*Serialized) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDescGZIP(), []int{0}
}

func (x *Serialized) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *Serialized) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_tensorflow_compiler_xla_python_ifrt_serdes_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x2f, 0x69, 0x66, 0x72, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x64, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x78, 0x6c, 0x61, 0x2e, 0x69, 0x66, 0x72, 0x74, 0x22, 0x3d, 0x0a, 0x0a,
	0x53, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0xac, 0x01, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x6c, 0x61, 0x2e, 0x69, 0x66, 0x72, 0x74, 0x42, 0x0b, 0x53, 0x65,
	0x72, 0x64, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4e, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x78, 0x6c, 0x61, 0x2f,
	0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x2f, 0x69, 0x66, 0x72, 0x74, 0xa2, 0x02, 0x03, 0x58, 0x49,
	0x58, 0xaa, 0x02, 0x08, 0x58, 0x6c, 0x61, 0x2e, 0x49, 0x66, 0x72, 0x74, 0xca, 0x02, 0x08, 0x58,
	0x6c, 0x61, 0x5c, 0x49, 0x66, 0x72, 0x74, 0xe2, 0x02, 0x14, 0x58, 0x6c, 0x61, 0x5c, 0x49, 0x66,
	0x72, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x09, 0x58, 0x6c, 0x61, 0x3a, 0x3a, 0x49, 0x66, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDescData = file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDesc
)

func file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDescData)
	})
	return file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDescData
}

var file_tensorflow_compiler_xla_python_ifrt_serdes_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_compiler_xla_python_ifrt_serdes_proto_goTypes = []interface{}{
	(*Serialized)(nil), // 0: xla.ifrt.Serialized
}
var file_tensorflow_compiler_xla_python_ifrt_serdes_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_xla_python_ifrt_serdes_proto_init() }
func file_tensorflow_compiler_xla_python_ifrt_serdes_proto_init() {
	if File_tensorflow_compiler_xla_python_ifrt_serdes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_xla_python_ifrt_serdes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Serialized); i {
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
			RawDescriptor: file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_xla_python_ifrt_serdes_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_xla_python_ifrt_serdes_proto_depIdxs,
		MessageInfos:      file_tensorflow_compiler_xla_python_ifrt_serdes_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_xla_python_ifrt_serdes_proto = out.File
	file_tensorflow_compiler_xla_python_ifrt_serdes_proto_rawDesc = nil
	file_tensorflow_compiler_xla_python_ifrt_serdes_proto_goTypes = nil
	file_tensorflow_compiler_xla_python_ifrt_serdes_proto_depIdxs = nil
}
