// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/tsl/profiler/protobuf/profiler_options.proto

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

type ProfileOptions_DeviceType int32

const (
	ProfileOptions_UNSPECIFIED      ProfileOptions_DeviceType = 0
	ProfileOptions_CPU              ProfileOptions_DeviceType = 1
	ProfileOptions_GPU              ProfileOptions_DeviceType = 2
	ProfileOptions_TPU              ProfileOptions_DeviceType = 3
	ProfileOptions_PLUGGABLE_DEVICE ProfileOptions_DeviceType = 4
)

// Enum value maps for ProfileOptions_DeviceType.
var (
	ProfileOptions_DeviceType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "CPU",
		2: "GPU",
		3: "TPU",
		4: "PLUGGABLE_DEVICE",
	}
	ProfileOptions_DeviceType_value = map[string]int32{
		"UNSPECIFIED":      0,
		"CPU":              1,
		"GPU":              2,
		"TPU":              3,
		"PLUGGABLE_DEVICE": 4,
	}
)

func (x ProfileOptions_DeviceType) Enum() *ProfileOptions_DeviceType {
	p := new(ProfileOptions_DeviceType)
	*p = x
	return p
}

func (x ProfileOptions_DeviceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProfileOptions_DeviceType) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_enumTypes[0].Descriptor()
}

func (ProfileOptions_DeviceType) Type() protoreflect.EnumType {
	return &file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_enumTypes[0]
}

func (x ProfileOptions_DeviceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProfileOptions_DeviceType.Descriptor instead.
func (ProfileOptions_DeviceType) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDescGZIP(), []int{0, 0}
}

// Next ID: 11
type ProfileOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Some default value of option are not proto3 default value. Use this version
	// to determine if we should use default option value instead of proto3
	// default value.
	Version uint32 `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	// Device type to profile/trace: (version >= 1)
	// DeviceType::UNSPECIFIED: All registered device profiler will be enabled.
	// DeviceType::CPU: only CPU will be profiled.
	// DeviceType::GPU: only CPU/GPU will be profiled.
	// DeviceType::TPU: only CPU/TPU will be profiled.
	// DeviceType::PLUGGABLE_DEVICE: only CPU/pluggable devices with profilers
	// will be profiled.
	DeviceType ProfileOptions_DeviceType `protobuf:"varint,6,opt,name=device_type,json=deviceType,proto3,enum=tensorflow.ProfileOptions_DeviceType" json:"device_type,omitempty"`
	// We don't collect the dataset ops by default for better trace-viewer
	// scalability. The caller can manually set this field to include the ops.
	IncludeDatasetOps bool `protobuf:"varint,1,opt,name=include_dataset_ops,json=includeDatasetOps,proto3" json:"include_dataset_ops,omitempty"`
	// Levels of host tracing: (version >= 1)
	//   - Level 0 is used to disable host traces.
	//   - Level 1 enables tracing of only user instrumented (or default) TraceMe.
	//   - Level 2 enables tracing of all level 1 TraceMe(s) and instrumented high
	//     level program execution details (expensive TF ops, XLA ops, etc).
	//     This is the default.
	//   - Level 3 enables tracing of all level 2 TraceMe(s) and more verbose
	//     (low-level) program execution details (cheap TF ops, etc).
	HostTracerLevel uint32 `protobuf:"varint,2,opt,name=host_tracer_level,json=hostTracerLevel,proto3" json:"host_tracer_level,omitempty"`
	// Levels of device tracing: (version >= 1)
	//   - Level 0 is used to disable device traces.
	//   - Level 1 is used to enable device traces.
	//   - More levels might be defined for specific device for controlling the
	//     verbosity of the trace.
	DeviceTracerLevel uint32 `protobuf:"varint,3,opt,name=device_tracer_level,json=deviceTracerLevel,proto3" json:"device_tracer_level,omitempty"`
	// Whether enable python function calls tracing. Runtime overhead ensues if
	// enabled. Default off. (version >= 1)
	PythonTracerLevel uint32 `protobuf:"varint,4,opt,name=python_tracer_level,json=pythonTracerLevel,proto3" json:"python_tracer_level,omitempty"`
	// Whether serialize hlo_proto when XLA is used. (version >= 1)
	EnableHloProto bool `protobuf:"varint,7,opt,name=enable_hlo_proto,json=enableHloProto,proto3" json:"enable_hlo_proto,omitempty"`
	// The local profiler starts profiling at this Unix timestamp in nanoseconds.
	StartTimestampNs uint64 `protobuf:"varint,8,opt,name=start_timestamp_ns,json=startTimestampNs,proto3" json:"start_timestamp_ns,omitempty"`
	// The local profiler collects `duration_ms` milliseconds of data. If the
	// value is 0, profiling continues until interrupted.
	DurationMs uint64 `protobuf:"varint,9,opt,name=duration_ms,json=durationMs,proto3" json:"duration_ms,omitempty"`
	// Directory to save profile data to. No-op when empty.
	RepositoryPath string `protobuf:"bytes,10,opt,name=repository_path,json=repositoryPath,proto3" json:"repository_path,omitempty"`
}

func (x *ProfileOptions) Reset() {
	*x = ProfileOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileOptions) ProtoMessage() {}

func (x *ProfileOptions) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileOptions.ProtoReflect.Descriptor instead.
func (*ProfileOptions) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileOptions) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ProfileOptions) GetDeviceType() ProfileOptions_DeviceType {
	if x != nil {
		return x.DeviceType
	}
	return ProfileOptions_UNSPECIFIED
}

func (x *ProfileOptions) GetIncludeDatasetOps() bool {
	if x != nil {
		return x.IncludeDatasetOps
	}
	return false
}

func (x *ProfileOptions) GetHostTracerLevel() uint32 {
	if x != nil {
		return x.HostTracerLevel
	}
	return 0
}

func (x *ProfileOptions) GetDeviceTracerLevel() uint32 {
	if x != nil {
		return x.DeviceTracerLevel
	}
	return 0
}

func (x *ProfileOptions) GetPythonTracerLevel() uint32 {
	if x != nil {
		return x.PythonTracerLevel
	}
	return 0
}

func (x *ProfileOptions) GetEnableHloProto() bool {
	if x != nil {
		return x.EnableHloProto
	}
	return false
}

func (x *ProfileOptions) GetStartTimestampNs() uint64 {
	if x != nil {
		return x.StartTimestampNs
	}
	return 0
}

func (x *ProfileOptions) GetDurationMs() uint64 {
	if x != nil {
		return x.DurationMs
	}
	return 0
}

func (x *ProfileOptions) GetRepositoryPath() string {
	if x != nil {
		return x.RepositoryPath
	}
	return ""
}

// Options for remote profiler session manager.
// Next ID: 6
type RemoteProfilerSessionManagerOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Options for each local profiler.
	ProfilerOptions *ProfileOptions `protobuf:"bytes,1,opt,name=profiler_options,json=profilerOptions,proto3" json:"profiler_options,omitempty"`
	// List of servers to profile. Supported formats: host:port.
	ServiceAddresses []string `protobuf:"bytes,2,rep,name=service_addresses,json=serviceAddresses,proto3" json:"service_addresses,omitempty"`
	// Unix timestamp of when the session was started.
	SessionCreationTimestampNs uint64 `protobuf:"varint,3,opt,name=session_creation_timestamp_ns,json=sessionCreationTimestampNs,proto3" json:"session_creation_timestamp_ns,omitempty"`
	// Maximum time (in milliseconds) a profiling session manager waits for all
	// profilers to finish after issuing gRPC request. If value is 0, session
	// continues until interrupted. Otherwise, value must be greater than
	// profiler_options.duration_ms.
	MaxSessionDurationMs uint64 `protobuf:"varint,4,opt,name=max_session_duration_ms,json=maxSessionDurationMs,proto3" json:"max_session_duration_ms,omitempty"`
	// Start of profiling is delayed by this much (in milliseconds).
	DelayMs uint64 `protobuf:"varint,5,opt,name=delay_ms,json=delayMs,proto3" json:"delay_ms,omitempty"`
}

func (x *RemoteProfilerSessionManagerOptions) Reset() {
	*x = RemoteProfilerSessionManagerOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoteProfilerSessionManagerOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoteProfilerSessionManagerOptions) ProtoMessage() {}

func (x *RemoteProfilerSessionManagerOptions) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoteProfilerSessionManagerOptions.ProtoReflect.Descriptor instead.
func (*RemoteProfilerSessionManagerOptions) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDescGZIP(), []int{1}
}

func (x *RemoteProfilerSessionManagerOptions) GetProfilerOptions() *ProfileOptions {
	if x != nil {
		return x.ProfilerOptions
	}
	return nil
}

func (x *RemoteProfilerSessionManagerOptions) GetServiceAddresses() []string {
	if x != nil {
		return x.ServiceAddresses
	}
	return nil
}

func (x *RemoteProfilerSessionManagerOptions) GetSessionCreationTimestampNs() uint64 {
	if x != nil {
		return x.SessionCreationTimestampNs
	}
	return 0
}

func (x *RemoteProfilerSessionManagerOptions) GetMaxSessionDurationMs() uint64 {
	if x != nil {
		return x.MaxSessionDurationMs
	}
	return 0
}

func (x *RemoteProfilerSessionManagerOptions) GetDelayMs() uint64 {
	if x != nil {
		return x.DelayMs
	}
	return 0
}

var File_tensorflow_tsl_profiler_protobuf_profiler_options_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDesc = []byte{
	0x0a, 0x37, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0xa0, 0x04, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x5f, 0x6f, 0x70,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x4f, 0x70, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x68, 0x6f,
	0x73, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x68, 0x6f, 0x73, 0x74, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x11, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x28, 0x0a, 0x10, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x68, 0x6c, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x6c, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x2c, 0x0a, 0x12, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x5f, 0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x4e, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12,
	0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x74, 0x68, 0x22, 0x4e, 0x0a, 0x0a, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x10, 0x01,
	0x12, 0x07, 0x0a, 0x03, 0x47, 0x50, 0x55, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x50, 0x55,
	0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x4c, 0x55, 0x47, 0x47, 0x41, 0x42, 0x4c, 0x45, 0x5f,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x04, 0x22, 0xae, 0x02, 0x0a, 0x23, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x45, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x1d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x5f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x1a, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x4e, 0x73, 0x12, 0x35, 0x0a, 0x17, 0x6d, 0x61, 0x78, 0x5f, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6d, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x4d, 0x73, 0x42, 0xbb, 0x01, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x42, 0x14, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0xca, 0x02, 0x0a, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0xe2, 0x02, 0x16, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDescData = file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDesc
)

func file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDescData)
	})
	return file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDescData
}

var file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_goTypes = []interface{}{
	(ProfileOptions_DeviceType)(0),              // 0: tensorflow.ProfileOptions.DeviceType
	(*ProfileOptions)(nil),                      // 1: tensorflow.ProfileOptions
	(*RemoteProfilerSessionManagerOptions)(nil), // 2: tensorflow.RemoteProfilerSessionManagerOptions
}
var file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_depIdxs = []int32{
	0, // 0: tensorflow.ProfileOptions.device_type:type_name -> tensorflow.ProfileOptions.DeviceType
	1, // 1: tensorflow.RemoteProfilerSessionManagerOptions.profiler_options:type_name -> tensorflow.ProfileOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_init() }
func file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_init() {
	if File_tensorflow_tsl_profiler_protobuf_profiler_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileOptions); i {
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
		file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoteProfilerSessionManagerOptions); i {
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
			RawDescriptor: file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_depIdxs,
		EnumInfos:         file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_enumTypes,
		MessageInfos:      file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_profiler_protobuf_profiler_options_proto = out.File
	file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_rawDesc = nil
	file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_goTypes = nil
	file_tensorflow_tsl_profiler_protobuf_profiler_options_proto_depIdxs = nil
}