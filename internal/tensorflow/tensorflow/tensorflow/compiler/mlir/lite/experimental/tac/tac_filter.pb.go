// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/compiler/mlir/lite/experimental/tac/tac_filter.proto

package tac

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

// Function filter types that are supported. If one function is matched for
// two rules with conflict, INCLUDE_TARGET_ANNOTATION has higher priority.
type FunctionFilter_FunctionFilterType int32

const (
	// To skip this function in the target annotation pass. This means all ops
	// in this function run on TFLite CPU.
	FunctionFilter_SKIP_TARGET_ANNOTATION FunctionFilter_FunctionFilterType = 0
	// To include this function in the target annotation pass. This has higher
	// priority than `SKIP_TARGET_ANNOTATION`.
	FunctionFilter_INCLUDE_TARGET_ANNOTATION FunctionFilter_FunctionFilterType = 1
)

// Enum value maps for FunctionFilter_FunctionFilterType.
var (
	FunctionFilter_FunctionFilterType_name = map[int32]string{
		0: "SKIP_TARGET_ANNOTATION",
		1: "INCLUDE_TARGET_ANNOTATION",
	}
	FunctionFilter_FunctionFilterType_value = map[string]int32{
		"SKIP_TARGET_ANNOTATION":    0,
		"INCLUDE_TARGET_ANNOTATION": 1,
	}
)

func (x FunctionFilter_FunctionFilterType) Enum() *FunctionFilter_FunctionFilterType {
	p := new(FunctionFilter_FunctionFilterType)
	*p = x
	return p
}

func (x FunctionFilter_FunctionFilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FunctionFilter_FunctionFilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_enumTypes[0].Descriptor()
}

func (FunctionFilter_FunctionFilterType) Type() protoreflect.EnumType {
	return &file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_enumTypes[0]
}

func (x FunctionFilter_FunctionFilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FunctionFilter_FunctionFilterType.Descriptor instead.
func (FunctionFilter_FunctionFilterType) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescGZIP(), []int{2, 0}
}

// A list of filters for TAC users to run ops/functions on ML hardwares. The
// intuition is that, for ops/functions that can be run on ML hardware (e.g.
// EdgeTPU) and TFLite CPU, TAC users give a hint that they're more performant
// to run on TFLite CPU. These filters give the TAC users freedom to specify the
// parts that they want to use other hardware to accelerate.
type TacFilters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of filters/rules to specify the parts that user wants to run on
	// other hardware.
	TacFilters []*TacFilter `protobuf:"bytes,1,rep,name=tac_filters,json=tacFilters,proto3" json:"tac_filters,omitempty"`
}

func (x *TacFilters) Reset() {
	*x = TacFilters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TacFilters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TacFilters) ProtoMessage() {}

func (x *TacFilters) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TacFilters.ProtoReflect.Descriptor instead.
func (*TacFilters) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescGZIP(), []int{0}
}

func (x *TacFilters) GetTacFilters() []*TacFilter {
	if x != nil {
		return x.TacFilters
	}
	return nil
}

// A filter can be used for an op or function.
type TacFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Filter:
	//
	//	*TacFilter_OpFilter
	//	*TacFilter_FunctionFilter
	Filter isTacFilter_Filter `protobuf_oneof:"filter"`
}

func (x *TacFilter) Reset() {
	*x = TacFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TacFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TacFilter) ProtoMessage() {}

func (x *TacFilter) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TacFilter.ProtoReflect.Descriptor instead.
func (*TacFilter) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescGZIP(), []int{1}
}

func (m *TacFilter) GetFilter() isTacFilter_Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (x *TacFilter) GetOpFilter() *OpFilter {
	if x, ok := x.GetFilter().(*TacFilter_OpFilter); ok {
		return x.OpFilter
	}
	return nil
}

func (x *TacFilter) GetFunctionFilter() *FunctionFilter {
	if x, ok := x.GetFilter().(*TacFilter_FunctionFilter); ok {
		return x.FunctionFilter
	}
	return nil
}

type isTacFilter_Filter interface {
	isTacFilter_Filter()
}

type TacFilter_OpFilter struct {
	OpFilter *OpFilter `protobuf:"bytes,1,opt,name=op_filter,json=opFilter,proto3,oneof"`
}

type TacFilter_FunctionFilter struct {
	FunctionFilter *FunctionFilter `protobuf:"bytes,2,opt,name=function_filter,json=functionFilter,proto3,oneof"`
}

func (*TacFilter_OpFilter) isTacFilter_Filter() {}

func (*TacFilter_FunctionFilter) isTacFilter_Filter() {}

// Function filter is to include/exclude a function in the target annotation
// pass in the TAC tool pipeline.
type FunctionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This name corresponds to the TFLite subgraph name in the flatbuffer.
	// `function_name_pattern` supports regex matching.
	FunctionNamePattern string                            `protobuf:"bytes,1,opt,name=function_name_pattern,json=functionNamePattern,proto3" json:"function_name_pattern,omitempty"`
	FilterType          FunctionFilter_FunctionFilterType `protobuf:"varint,2,opt,name=filter_type,json=filterType,proto3,enum=third_party.tensorflow.compiler.mlir.lite.experimental.tac.FunctionFilter_FunctionFilterType" json:"filter_type,omitempty"`
}

func (x *FunctionFilter) Reset() {
	*x = FunctionFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FunctionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FunctionFilter) ProtoMessage() {}

func (x *FunctionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FunctionFilter.ProtoReflect.Descriptor instead.
func (*FunctionFilter) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescGZIP(), []int{2}
}

func (x *FunctionFilter) GetFunctionNamePattern() string {
	if x != nil {
		return x.FunctionNamePattern
	}
	return ""
}

func (x *FunctionFilter) GetFilterType() FunctionFilter_FunctionFilterType {
	if x != nil {
		return x.FilterType
	}
	return FunctionFilter_SKIP_TARGET_ANNOTATION
}

// Op filter is to filter out ops that user wants to run. Ops with this filter
// run on TFLite CPU.
type OpFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This name corresponds to the mlir::Location of the tensor.
	// `op_name_pattern` supports regex matching.
	OpNamePattern string `protobuf:"bytes,1,opt,name=op_name_pattern,json=opNamePattern,proto3" json:"op_name_pattern,omitempty"`
}

func (x *OpFilter) Reset() {
	*x = OpFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpFilter) ProtoMessage() {}

func (x *OpFilter) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpFilter.ProtoReflect.Descriptor instead.
func (*OpFilter) Descriptor() ([]byte, []int) {
	return file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescGZIP(), []int{3}
}

func (x *OpFilter) GetOpNamePattern() string {
	if x != nil {
		return x.OpNamePattern
	}
	return ""
}

var File_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto protoreflect.FileDescriptor

var file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x6c, 0x69, 0x72, 0x2f, 0x6c, 0x69, 0x74, 0x65, 0x2f,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x74, 0x61, 0x63,
	0x2f, 0x74, 0x61, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x3a, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x63, 0x22, 0x74, 0x0a,
	0x0a, 0x54, 0x61, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x66, 0x0a, 0x0b, 0x74,
	0x61, 0x63, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x45, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c,
	0x65, 0x72, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x70,
	0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x63, 0x2e, 0x54, 0x61,
	0x63, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x74, 0x61, 0x63, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x22, 0xf1, 0x01, 0x0a, 0x09, 0x54, 0x61, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x63, 0x0a, 0x09, 0x6f, 0x70, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x6c, 0x69, 0x74, 0x65,
	0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x74, 0x61,
	0x63, 0x2e, 0x4f, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x70,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x75, 0x0a, 0x0f, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x4a, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65,
	0x72, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x70, 0x65,
	0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x63, 0x2e, 0x46, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x66,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x08, 0x0a,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x95, 0x02, 0x0a, 0x0e, 0x46, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x15, 0x66, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x7e,
	0x0a, 0x0b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x5d, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x6c, 0x69, 0x72, 0x2e, 0x6c, 0x69, 0x74, 0x65, 0x2e,
	0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x74, 0x61, 0x63,
	0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x4f,
	0x0a, 0x12, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x4b, 0x49, 0x50, 0x5f, 0x54, 0x41, 0x52,
	0x47, 0x45, 0x54, 0x5f, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x00,
	0x12, 0x1d, 0x0a, 0x19, 0x49, 0x4e, 0x43, 0x4c, 0x55, 0x44, 0x45, 0x5f, 0x54, 0x41, 0x52, 0x47,
	0x45, 0x54, 0x5f, 0x41, 0x4e, 0x4e, 0x4f, 0x54, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x22,
	0x32, 0x0a, 0x08, 0x4f, 0x70, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0f, 0x6f,
	0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x74, 0x74,
	0x65, 0x72, 0x6e, 0x42, 0xb9, 0x03, 0x0a, 0x3e, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x6d, 0x6c, 0x69, 0x72,
	0x2e, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x6c, 0x2e, 0x74, 0x61, 0x63, 0x42, 0x0e, 0x54, 0x61, 0x63, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x6c, 0x69, 0x72, 0x2f, 0x6c, 0x69,
	0x74, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f,
	0x74, 0x61, 0x63, 0xa2, 0x02, 0x07, 0x54, 0x54, 0x43, 0x4d, 0x4c, 0x45, 0x54, 0xaa, 0x02, 0x39,
	0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x4d,
	0x6c, 0x69, 0x72, 0x2e, 0x4c, 0x69, 0x74, 0x65, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2e, 0x54, 0x61, 0x63, 0xca, 0x02, 0x39, 0x54, 0x68, 0x69, 0x72,
	0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x5c, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x5c, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x5c, 0x4d, 0x6c, 0x69, 0x72, 0x5c,
	0x4c, 0x69, 0x74, 0x65, 0x5c, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61,
	0x6c, 0x5c, 0x54, 0x61, 0x63, 0xe2, 0x02, 0x45, 0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x5c, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x43, 0x6f,
	0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72, 0x5c, 0x4d, 0x6c, 0x69, 0x72, 0x5c, 0x4c, 0x69, 0x74, 0x65,
	0x5c, 0x45, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x5c, 0x54, 0x61,
	0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3f,
	0x54, 0x68, 0x69, 0x72, 0x64, 0x50, 0x61, 0x72, 0x74, 0x79, 0x3a, 0x3a, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x72,
	0x3a, 0x3a, 0x4d, 0x6c, 0x69, 0x72, 0x3a, 0x3a, 0x4c, 0x69, 0x74, 0x65, 0x3a, 0x3a, 0x45, 0x78,
	0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x3a, 0x3a, 0x54, 0x61, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescOnce sync.Once
	file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescData = file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDesc
)

func file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescGZIP() []byte {
	file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescOnce.Do(func() {
		file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescData)
	})
	return file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDescData
}

var file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_goTypes = []interface{}{
	(FunctionFilter_FunctionFilterType)(0), // 0: third_party.tensorflow.compiler.mlir.lite.experimental.tac.FunctionFilter.FunctionFilterType
	(*TacFilters)(nil),                     // 1: third_party.tensorflow.compiler.mlir.lite.experimental.tac.TacFilters
	(*TacFilter)(nil),                      // 2: third_party.tensorflow.compiler.mlir.lite.experimental.tac.TacFilter
	(*FunctionFilter)(nil),                 // 3: third_party.tensorflow.compiler.mlir.lite.experimental.tac.FunctionFilter
	(*OpFilter)(nil),                       // 4: third_party.tensorflow.compiler.mlir.lite.experimental.tac.OpFilter
}
var file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_depIdxs = []int32{
	2, // 0: third_party.tensorflow.compiler.mlir.lite.experimental.tac.TacFilters.tac_filters:type_name -> third_party.tensorflow.compiler.mlir.lite.experimental.tac.TacFilter
	4, // 1: third_party.tensorflow.compiler.mlir.lite.experimental.tac.TacFilter.op_filter:type_name -> third_party.tensorflow.compiler.mlir.lite.experimental.tac.OpFilter
	3, // 2: third_party.tensorflow.compiler.mlir.lite.experimental.tac.TacFilter.function_filter:type_name -> third_party.tensorflow.compiler.mlir.lite.experimental.tac.FunctionFilter
	0, // 3: third_party.tensorflow.compiler.mlir.lite.experimental.tac.FunctionFilter.filter_type:type_name -> third_party.tensorflow.compiler.mlir.lite.experimental.tac.FunctionFilter.FunctionFilterType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_init() }
func file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_init() {
	if File_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TacFilters); i {
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
		file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TacFilter); i {
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
		file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FunctionFilter); i {
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
		file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpFilter); i {
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
	file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TacFilter_OpFilter)(nil),
		(*TacFilter_FunctionFilter)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_goTypes,
		DependencyIndexes: file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_depIdxs,
		EnumInfos:         file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_enumTypes,
		MessageInfos:      file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_msgTypes,
	}.Build()
	File_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto = out.File
	file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_rawDesc = nil
	file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_goTypes = nil
	file_tensorflow_compiler_mlir_lite_experimental_tac_tac_filter_proto_depIdxs = nil
}
