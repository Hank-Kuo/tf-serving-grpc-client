// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/core/protobuf/data_service.proto

package protobuf

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

// tf.data service deployment mode.
type DeploymentMode int32

const (
	DeploymentMode_DEPLOYMENT_MODE_UNSPECIFIED DeploymentMode = 0
	// tf.data service workers colocate with TF workers.
	DeploymentMode_DEPLOYMENT_MODE_COLOCATED DeploymentMode = 1
	// tf.data service workers run in dedicated tf.data hosts.
	DeploymentMode_DEPLOYMENT_MODE_REMOTE DeploymentMode = 2
	// tf.data service workers run in colocated TF hosts and dedicated tf.data
	// hosts.
	DeploymentMode_DEPLOYMENT_MODE_HYBRID DeploymentMode = 3
)

// Enum value maps for DeploymentMode.
var (
	DeploymentMode_name = map[int32]string{
		0: "DEPLOYMENT_MODE_UNSPECIFIED",
		1: "DEPLOYMENT_MODE_COLOCATED",
		2: "DEPLOYMENT_MODE_REMOTE",
		3: "DEPLOYMENT_MODE_HYBRID",
	}
	DeploymentMode_value = map[string]int32{
		"DEPLOYMENT_MODE_UNSPECIFIED": 0,
		"DEPLOYMENT_MODE_COLOCATED":   1,
		"DEPLOYMENT_MODE_REMOTE":      2,
		"DEPLOYMENT_MODE_HYBRID":      3,
	}
)

func (x DeploymentMode) Enum() *DeploymentMode {
	p := new(DeploymentMode)
	*p = x
	return p
}

func (x DeploymentMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentMode) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_protobuf_data_service_proto_enumTypes[0].Descriptor()
}

func (DeploymentMode) Type() protoreflect.EnumType {
	return &file_tensorflow_core_protobuf_data_service_proto_enumTypes[0]
}

func (x DeploymentMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentMode.Descriptor instead.
func (DeploymentMode) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_data_service_proto_rawDescGZIP(), []int{0}
}

// Specifies how data is sharded among tf.data service workers.
type ProcessingModeDef_ShardingPolicy int32

const (
	// No sharding will be performed. Each worker produces the entire dataset
	// without any sharding. With this mode, the best practice is to shuffle the
	// dataset nondeterministically so that workers process the dataset in
	// different orders.
	ProcessingModeDef_OFF ProcessingModeDef_ShardingPolicy = 0
	// The input dataset is dynamically split among workers at runtime. Each
	// worker gets the next split when it reads data from the dispatcher. There
	// is no fixed sharding with this mode.
	ProcessingModeDef_DYNAMIC ProcessingModeDef_ShardingPolicy = 1
	// The following are static sharding policies. The semantics are similar to
	// `tf.data.experimental.AutoShardPolicy`. These policies require:
	//   - The tf.data service cluster has a fixed size, and you need to specify
	//     the workers in DispatcherConfig.
	//   - Each client only reads from the local tf.data service worker.
	//
	// Shards by input files (each worker will get a set of files to process).
	// When this option is selected, make sure that there is at least as many
	// files as workers. If there are fewer input files than workers, a runtime
	// error will be raised.
	ProcessingModeDef_FILE ProcessingModeDef_ShardingPolicy = 2
	// Shards by elements produced by the dataset. Each worker will process the
	// whole dataset and discard the portion that is not for itself. Note that
	// for this mode to correctly partitions the dataset elements, the dataset
	// needs to produce elements in a deterministic order.
	ProcessingModeDef_DATA ProcessingModeDef_ShardingPolicy = 3
	// Attempts FILE-based sharding, falling back to DATA-based sharding on
	// failures.
	ProcessingModeDef_FILE_OR_DATA ProcessingModeDef_ShardingPolicy = 4
	// Looks for the presence of `shard(SHARD_HINT, ...)` which is treated as a
	// placeholder to replace with `shard(num_workers, worker_index)`.
	ProcessingModeDef_HINT ProcessingModeDef_ShardingPolicy = 5
)

// Enum value maps for ProcessingModeDef_ShardingPolicy.
var (
	ProcessingModeDef_ShardingPolicy_name = map[int32]string{
		0: "OFF",
		1: "DYNAMIC",
		2: "FILE",
		3: "DATA",
		4: "FILE_OR_DATA",
		5: "HINT",
	}
	ProcessingModeDef_ShardingPolicy_value = map[string]int32{
		"OFF":          0,
		"DYNAMIC":      1,
		"FILE":         2,
		"DATA":         3,
		"FILE_OR_DATA": 4,
		"HINT":         5,
	}
)

func (x ProcessingModeDef_ShardingPolicy) Enum() *ProcessingModeDef_ShardingPolicy {
	p := new(ProcessingModeDef_ShardingPolicy)
	*p = x
	return p
}

func (x ProcessingModeDef_ShardingPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProcessingModeDef_ShardingPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_protobuf_data_service_proto_enumTypes[1].Descriptor()
}

func (ProcessingModeDef_ShardingPolicy) Type() protoreflect.EnumType {
	return &file_tensorflow_core_protobuf_data_service_proto_enumTypes[1]
}

func (x ProcessingModeDef_ShardingPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProcessingModeDef_ShardingPolicy.Descriptor instead.
func (ProcessingModeDef_ShardingPolicy) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_data_service_proto_rawDescGZIP(), []int{0, 0}
}

type DataServiceMetadata_Compression int32

const (
	DataServiceMetadata_COMPRESSION_UNSPECIFIED DataServiceMetadata_Compression = 0
	// No compression.
	DataServiceMetadata_COMPRESSION_OFF DataServiceMetadata_Compression = 1
	// Snappy compression as defined in tensorflow/core/platform/snappy.h.
	DataServiceMetadata_COMPRESSION_SNAPPY DataServiceMetadata_Compression = 2
)

// Enum value maps for DataServiceMetadata_Compression.
var (
	DataServiceMetadata_Compression_name = map[int32]string{
		0: "COMPRESSION_UNSPECIFIED",
		1: "COMPRESSION_OFF",
		2: "COMPRESSION_SNAPPY",
	}
	DataServiceMetadata_Compression_value = map[string]int32{
		"COMPRESSION_UNSPECIFIED": 0,
		"COMPRESSION_OFF":         1,
		"COMPRESSION_SNAPPY":      2,
	}
)

func (x DataServiceMetadata_Compression) Enum() *DataServiceMetadata_Compression {
	p := new(DataServiceMetadata_Compression)
	*p = x
	return p
}

func (x DataServiceMetadata_Compression) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataServiceMetadata_Compression) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_protobuf_data_service_proto_enumTypes[2].Descriptor()
}

func (DataServiceMetadata_Compression) Type() protoreflect.EnumType {
	return &file_tensorflow_core_protobuf_data_service_proto_enumTypes[2]
}

func (x DataServiceMetadata_Compression) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataServiceMetadata_Compression.Descriptor instead.
func (DataServiceMetadata_Compression) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_data_service_proto_rawDescGZIP(), []int{1, 0}
}

// Next tag: 2
type ProcessingModeDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShardingPolicy ProcessingModeDef_ShardingPolicy `protobuf:"varint,1,opt,name=sharding_policy,json=shardingPolicy,proto3,enum=tensorflow.data.ProcessingModeDef_ShardingPolicy" json:"sharding_policy,omitempty"`
}

func (x *ProcessingModeDef) Reset() {
	*x = ProcessingModeDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_data_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessingModeDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessingModeDef) ProtoMessage() {}

func (x *ProcessingModeDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_data_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessingModeDef.ProtoReflect.Descriptor instead.
func (*ProcessingModeDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_data_service_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessingModeDef) GetShardingPolicy() ProcessingModeDef_ShardingPolicy {
	if x != nil {
		return x.ShardingPolicy
	}
	return ProcessingModeDef_OFF
}

// Metadata related to tf.data service datasets.
// Next tag: 4
type DataServiceMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to OptionalElementSpec:
	//
	//	*DataServiceMetadata_ElementSpec
	OptionalElementSpec isDataServiceMetadata_OptionalElementSpec `protobuf_oneof:"optional_element_spec"`
	Compression         DataServiceMetadata_Compression           `protobuf:"varint,2,opt,name=compression,proto3,enum=tensorflow.data.DataServiceMetadata_Compression" json:"compression,omitempty"`
	// Cardinality of the dataset.
	Cardinality int64 `protobuf:"varint,3,opt,name=cardinality,proto3" json:"cardinality,omitempty"`
}

func (x *DataServiceMetadata) Reset() {
	*x = DataServiceMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_data_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataServiceMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataServiceMetadata) ProtoMessage() {}

func (x *DataServiceMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_data_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataServiceMetadata.ProtoReflect.Descriptor instead.
func (*DataServiceMetadata) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_data_service_proto_rawDescGZIP(), []int{1}
}

func (m *DataServiceMetadata) GetOptionalElementSpec() isDataServiceMetadata_OptionalElementSpec {
	if m != nil {
		return m.OptionalElementSpec
	}
	return nil
}

func (x *DataServiceMetadata) GetElementSpec() []byte {
	if x, ok := x.GetOptionalElementSpec().(*DataServiceMetadata_ElementSpec); ok {
		return x.ElementSpec
	}
	return nil
}

func (x *DataServiceMetadata) GetCompression() DataServiceMetadata_Compression {
	if x != nil {
		return x.Compression
	}
	return DataServiceMetadata_COMPRESSION_UNSPECIFIED
}

func (x *DataServiceMetadata) GetCardinality() int64 {
	if x != nil {
		return x.Cardinality
	}
	return 0
}

type isDataServiceMetadata_OptionalElementSpec interface {
	isDataServiceMetadata_OptionalElementSpec()
}

type DataServiceMetadata_ElementSpec struct {
	// Serialized element spec.
	ElementSpec []byte `protobuf:"bytes,1,opt,name=element_spec,json=elementSpec,proto3,oneof"`
}

func (*DataServiceMetadata_ElementSpec) isDataServiceMetadata_OptionalElementSpec() {}

type CrossTrainerCacheOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrainerId string `protobuf:"bytes,1,opt,name=trainer_id,json=trainerId,proto3" json:"trainer_id,omitempty"`
}

func (x *CrossTrainerCacheOptions) Reset() {
	*x = CrossTrainerCacheOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_data_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrossTrainerCacheOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrossTrainerCacheOptions) ProtoMessage() {}

func (x *CrossTrainerCacheOptions) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_data_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrossTrainerCacheOptions.ProtoReflect.Descriptor instead.
func (*CrossTrainerCacheOptions) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_data_service_proto_rawDescGZIP(), []int{2}
}

func (x *CrossTrainerCacheOptions) GetTrainerId() string {
	if x != nil {
		return x.TrainerId
	}
	return ""
}

// Data service config available to the client through GetDataServiceConfig RPC.
// Next tag: 2
type DataServiceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeploymentMode DeploymentMode `protobuf:"varint,1,opt,name=deployment_mode,json=deploymentMode,proto3,enum=tensorflow.data.DeploymentMode" json:"deployment_mode,omitempty"`
}

func (x *DataServiceConfig) Reset() {
	*x = DataServiceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_data_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataServiceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataServiceConfig) ProtoMessage() {}

func (x *DataServiceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_data_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataServiceConfig.ProtoReflect.Descriptor instead.
func (*DataServiceConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_data_service_proto_rawDescGZIP(), []int{3}
}

func (x *DataServiceConfig) GetDeploymentMode() DeploymentMode {
	if x != nil {
		return x.DeploymentMode
	}
	return DeploymentMode_DEPLOYMENT_MODE_UNSPECIFIED
}

var File_tensorflow_core_protobuf_data_service_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_data_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x22, 0xc7,
	0x01, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64,
	0x65, 0x44, 0x65, 0x66, 0x12, 0x5a, 0x0a, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x44, 0x65,
	0x66, 0x2e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x0e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x22, 0x56, 0x0a, 0x0e, 0x53, 0x68, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x12, 0x07, 0x0a, 0x03, 0x4f, 0x46, 0x46, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44,
	0x59, 0x4e, 0x41, 0x4d, 0x49, 0x43, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x41, 0x54, 0x41, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c,
	0x46, 0x49, 0x4c, 0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x04, 0x12, 0x08,
	0x0a, 0x04, 0x48, 0x49, 0x4e, 0x54, 0x10, 0x05, 0x22, 0xa2, 0x02, 0x0a, 0x13, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x23, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x52, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x63, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x63, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x57, 0x0a, 0x0b, 0x43,
	0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f,
	0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4d, 0x50, 0x52,
	0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x4e, 0x41, 0x50,
	0x50, 0x59, 0x10, 0x02, 0x42, 0x17, 0x0a, 0x15, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x22, 0x39, 0x0a,
	0x18, 0x43, 0x72, 0x6f, 0x73, 0x73, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x48, 0x0a,
	0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x2a, 0x88, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x44,
	0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x43,
	0x4f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45,
	0x4d, 0x4f, 0x54, 0x45, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59,
	0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x48, 0x59, 0x42, 0x52, 0x49, 0x44,
	0x10, 0x03, 0x42, 0xc9, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x42, 0x10, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x43,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0xa2, 0x02, 0x03, 0x54, 0x44, 0x58, 0xaa, 0x02, 0x0f, 0x54, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61, 0xca, 0x02, 0x0f, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x44, 0x61, 0x74, 0x61, 0xe2, 0x02, 0x1b,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x54, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_data_service_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_data_service_proto_rawDescData = file_tensorflow_core_protobuf_data_service_proto_rawDesc
)

func file_tensorflow_core_protobuf_data_service_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_data_service_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_data_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_data_service_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_data_service_proto_rawDescData
}

var file_tensorflow_core_protobuf_data_service_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_tensorflow_core_protobuf_data_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_core_protobuf_data_service_proto_goTypes = []interface{}{
	(DeploymentMode)(0),                   // 0: tensorflow.data.DeploymentMode
	(ProcessingModeDef_ShardingPolicy)(0), // 1: tensorflow.data.ProcessingModeDef.ShardingPolicy
	(DataServiceMetadata_Compression)(0),  // 2: tensorflow.data.DataServiceMetadata.Compression
	(*ProcessingModeDef)(nil),             // 3: tensorflow.data.ProcessingModeDef
	(*DataServiceMetadata)(nil),           // 4: tensorflow.data.DataServiceMetadata
	(*CrossTrainerCacheOptions)(nil),      // 5: tensorflow.data.CrossTrainerCacheOptions
	(*DataServiceConfig)(nil),             // 6: tensorflow.data.DataServiceConfig
}
var file_tensorflow_core_protobuf_data_service_proto_depIdxs = []int32{
	1, // 0: tensorflow.data.ProcessingModeDef.sharding_policy:type_name -> tensorflow.data.ProcessingModeDef.ShardingPolicy
	2, // 1: tensorflow.data.DataServiceMetadata.compression:type_name -> tensorflow.data.DataServiceMetadata.Compression
	0, // 2: tensorflow.data.DataServiceConfig.deployment_mode:type_name -> tensorflow.data.DeploymentMode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_data_service_proto_init() }
func file_tensorflow_core_protobuf_data_service_proto_init() {
	if File_tensorflow_core_protobuf_data_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_data_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProcessingModeDef); i {
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
		file_tensorflow_core_protobuf_data_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataServiceMetadata); i {
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
		file_tensorflow_core_protobuf_data_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrossTrainerCacheOptions); i {
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
		file_tensorflow_core_protobuf_data_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataServiceConfig); i {
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
	file_tensorflow_core_protobuf_data_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*DataServiceMetadata_ElementSpec)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_protobuf_data_service_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_data_service_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_data_service_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_protobuf_data_service_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_protobuf_data_service_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_data_service_proto = out.File
	file_tensorflow_core_protobuf_data_service_proto_rawDesc = nil
	file_tensorflow_core_protobuf_data_service_proto_goTypes = nil
	file_tensorflow_core_protobuf_data_service_proto_depIdxs = nil
}
