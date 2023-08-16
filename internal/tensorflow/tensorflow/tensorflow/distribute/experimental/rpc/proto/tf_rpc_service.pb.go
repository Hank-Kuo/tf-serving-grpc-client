// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/distribute/experimental/rpc/proto/tf_rpc_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	framework "grpc-client/internal/tensorflow/tensorflow/tensorflow/core/framework"
	protobuf "grpc-client/internal/tensorflow/tensorflow/tensorflow/core/protobuf"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method       string                   `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	InputTensors []*framework.TensorProto `protobuf:"bytes,2,rep,name=input_tensors,json=inputTensors,proto3" json:"input_tensors,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *CallRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *CallRequest) GetInputTensors() []*framework.TensorProto {
	if x != nil {
		return x.InputTensors
	}
	return nil
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OutputTensors []*framework.TensorProto `protobuf:"bytes,1,rep,name=output_tensors,json=outputTensors,proto3" json:"output_tensors,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *CallResponse) GetOutputTensors() []*framework.TensorProto {
	if x != nil {
		return x.OutputTensors
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescGZIP(), []int{2}
}

type RegisteredMethod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method      string                    `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	InputSpecs  *protobuf.StructuredValue `protobuf:"bytes,2,opt,name=input_specs,json=inputSpecs,proto3" json:"input_specs,omitempty"`
	OutputSpecs *protobuf.StructuredValue `protobuf:"bytes,3,opt,name=output_specs,json=outputSpecs,proto3" json:"output_specs,omitempty"`
}

func (x *RegisteredMethod) Reset() {
	*x = RegisteredMethod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisteredMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisteredMethod) ProtoMessage() {}

func (x *RegisteredMethod) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisteredMethod.ProtoReflect.Descriptor instead.
func (*RegisteredMethod) Descriptor() ([]byte, []int) {
	return file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescGZIP(), []int{3}
}

func (x *RegisteredMethod) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RegisteredMethod) GetInputSpecs() *protobuf.StructuredValue {
	if x != nil {
		return x.InputSpecs
	}
	return nil
}

func (x *RegisteredMethod) GetOutputSpecs() *protobuf.StructuredValue {
	if x != nil {
		return x.OutputSpecs
	}
	return nil
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegisteredMethods []*RegisteredMethod `protobuf:"bytes,1,rep,name=registered_methods,json=registeredMethods,proto3" json:"registered_methods,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponse) GetRegisteredMethods() []*RegisteredMethod {
	if x != nil {
		return x.RegisteredMethods
	}
	return nil
}

var File_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto protoreflect.FileDescriptor

var file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x41, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65,
	0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74,
	0x66, 0x5f, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x72, 0x70, 0x63, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x63, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3c, 0x0a, 0x0d, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0c, 0x69, 0x6e, 0x70, 0x75, 0x74,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x22, 0x4e, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x0e, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x0d, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x3c, 0x0a, 0x0b, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x73, 0x12, 0x3e, 0x0a, 0x0c, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x53, 0x70, 0x65, 0x63,
	0x73, 0x22, 0x5f, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4f, 0x0a, 0x12, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x73, 0x32, 0x96, 0x01, 0x0a, 0x0a, 0x52, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x43, 0x0a, 0x04, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xd9, 0x01, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x72,
	0x70, 0x63, 0x42, 0x11, 0x54, 0x66, 0x52, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x57, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xa2, 0x02, 0x03, 0x54, 0x52, 0x58, 0xaa, 0x02, 0x0e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x52, 0x70, 0x63, 0xca, 0x02, 0x0e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x52, 0x70, 0x63, 0xe2, 0x02, 0x1a, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x52, 0x70, 0x63, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x3a, 0x3a, 0x52, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescOnce sync.Once
	file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescData = file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDesc
)

func file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescGZIP() []byte {
	file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescOnce.Do(func() {
		file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescData)
	})
	return file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDescData
}

var file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_goTypes = []interface{}{
	(*CallRequest)(nil),              // 0: tensorflow.rpc.CallRequest
	(*CallResponse)(nil),             // 1: tensorflow.rpc.CallResponse
	(*ListRequest)(nil),              // 2: tensorflow.rpc.ListRequest
	(*RegisteredMethod)(nil),         // 3: tensorflow.rpc.RegisteredMethod
	(*ListResponse)(nil),             // 4: tensorflow.rpc.ListResponse
	(*framework.TensorProto)(nil),    // 5: tensorflow.TensorProto
	(*protobuf.StructuredValue)(nil), // 6: tensorflow.StructuredValue
}
var file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_depIdxs = []int32{
	5, // 0: tensorflow.rpc.CallRequest.input_tensors:type_name -> tensorflow.TensorProto
	5, // 1: tensorflow.rpc.CallResponse.output_tensors:type_name -> tensorflow.TensorProto
	6, // 2: tensorflow.rpc.RegisteredMethod.input_specs:type_name -> tensorflow.StructuredValue
	6, // 3: tensorflow.rpc.RegisteredMethod.output_specs:type_name -> tensorflow.StructuredValue
	3, // 4: tensorflow.rpc.ListResponse.registered_methods:type_name -> tensorflow.rpc.RegisteredMethod
	0, // 5: tensorflow.rpc.RpcService.Call:input_type -> tensorflow.rpc.CallRequest
	2, // 6: tensorflow.rpc.RpcService.List:input_type -> tensorflow.rpc.ListRequest
	1, // 7: tensorflow.rpc.RpcService.Call:output_type -> tensorflow.rpc.CallResponse
	4, // 8: tensorflow.rpc.RpcService.List:output_type -> tensorflow.rpc.ListResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_init() }
func file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_init() {
	if File_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
		file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisteredMethod); i {
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
		file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
			RawDescriptor: file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_goTypes,
		DependencyIndexes: file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_depIdxs,
		MessageInfos:      file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_msgTypes,
	}.Build()
	File_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto = out.File
	file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_rawDesc = nil
	file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_goTypes = nil
	file_tensorflow_distribute_experimental_rpc_proto_tf_rpc_service_proto_depIdxs = nil
}
