// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow/lite/experimental/acceleration/configuration/configuration.proto

package configuration

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	configuration "grpc-client/internal/tensorflow/tensorflow/tensorflow/lite/acceleration/configuration"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of tensorflow/lite/acceleration/configuration/configuration.proto.

type ExecutionPreference = configuration.ExecutionPreference

const ExecutionPreference_ANY = configuration.ExecutionPreference_ANY
const ExecutionPreference_LOW_LATENCY = configuration.ExecutionPreference_LOW_LATENCY
const ExecutionPreference_LOW_POWER = configuration.ExecutionPreference_LOW_POWER
const ExecutionPreference_FORCE_CPU = configuration.ExecutionPreference_FORCE_CPU

var ExecutionPreference_name = configuration.ExecutionPreference_name
var ExecutionPreference_value = configuration.ExecutionPreference_value

type Delegate = configuration.Delegate

const Delegate_NONE = configuration.Delegate_NONE
const Delegate_NNAPI = configuration.Delegate_NNAPI
const Delegate_GPU = configuration.Delegate_GPU
const Delegate_HEXAGON = configuration.Delegate_HEXAGON
const Delegate_XNNPACK = configuration.Delegate_XNNPACK
const Delegate_EDGETPU = configuration.Delegate_EDGETPU
const Delegate_EDGETPU_CORAL = configuration.Delegate_EDGETPU_CORAL
const Delegate_CORE_ML = configuration.Delegate_CORE_ML
const Delegate_ARMNN = configuration.Delegate_ARMNN

var Delegate_name = configuration.Delegate_name
var Delegate_value = configuration.Delegate_value

type NNAPIExecutionPreference = configuration.NNAPIExecutionPreference

const NNAPIExecutionPreference_UNDEFINED = configuration.NNAPIExecutionPreference_UNDEFINED
const NNAPIExecutionPreference_NNAPI_LOW_POWER = configuration.NNAPIExecutionPreference_NNAPI_LOW_POWER
const NNAPIExecutionPreference_NNAPI_FAST_SINGLE_ANSWER = configuration.NNAPIExecutionPreference_NNAPI_FAST_SINGLE_ANSWER
const NNAPIExecutionPreference_NNAPI_SUSTAINED_SPEED = configuration.NNAPIExecutionPreference_NNAPI_SUSTAINED_SPEED

var NNAPIExecutionPreference_name = configuration.NNAPIExecutionPreference_name
var NNAPIExecutionPreference_value = configuration.NNAPIExecutionPreference_value

type NNAPIExecutionPriority = configuration.NNAPIExecutionPriority

const NNAPIExecutionPriority_NNAPI_PRIORITY_UNDEFINED = configuration.NNAPIExecutionPriority_NNAPI_PRIORITY_UNDEFINED
const NNAPIExecutionPriority_NNAPI_PRIORITY_LOW = configuration.NNAPIExecutionPriority_NNAPI_PRIORITY_LOW
const NNAPIExecutionPriority_NNAPI_PRIORITY_MEDIUM = configuration.NNAPIExecutionPriority_NNAPI_PRIORITY_MEDIUM
const NNAPIExecutionPriority_NNAPI_PRIORITY_HIGH = configuration.NNAPIExecutionPriority_NNAPI_PRIORITY_HIGH

var NNAPIExecutionPriority_name = configuration.NNAPIExecutionPriority_name
var NNAPIExecutionPriority_value = configuration.NNAPIExecutionPriority_value

type GPUBackend = configuration.GPUBackend

const GPUBackend_UNSET = configuration.GPUBackend_UNSET
const GPUBackend_OPENCL = configuration.GPUBackend_OPENCL
const GPUBackend_OPENGL = configuration.GPUBackend_OPENGL

var GPUBackend_name = configuration.GPUBackend_name
var GPUBackend_value = configuration.GPUBackend_value

type GPUInferencePriority = configuration.GPUInferencePriority

const GPUInferencePriority_GPU_PRIORITY_AUTO = configuration.GPUInferencePriority_GPU_PRIORITY_AUTO
const GPUInferencePriority_GPU_PRIORITY_MAX_PRECISION = configuration.GPUInferencePriority_GPU_PRIORITY_MAX_PRECISION
const GPUInferencePriority_GPU_PRIORITY_MIN_LATENCY = configuration.GPUInferencePriority_GPU_PRIORITY_MIN_LATENCY
const GPUInferencePriority_GPU_PRIORITY_MIN_MEMORY_USAGE = configuration.GPUInferencePriority_GPU_PRIORITY_MIN_MEMORY_USAGE

var GPUInferencePriority_name = configuration.GPUInferencePriority_name
var GPUInferencePriority_value = configuration.GPUInferencePriority_value

type GPUInferenceUsage = configuration.GPUInferenceUsage

const GPUInferenceUsage_GPU_INFERENCE_PREFERENCE_FAST_SINGLE_ANSWER = configuration.GPUInferenceUsage_GPU_INFERENCE_PREFERENCE_FAST_SINGLE_ANSWER
const GPUInferenceUsage_GPU_INFERENCE_PREFERENCE_SUSTAINED_SPEED = configuration.GPUInferenceUsage_GPU_INFERENCE_PREFERENCE_SUSTAINED_SPEED

var GPUInferenceUsage_name = configuration.GPUInferenceUsage_name
var GPUInferenceUsage_value = configuration.GPUInferenceUsage_value

type XNNPackFlags = configuration.XNNPackFlags

const XNNPackFlags_TFLITE_XNNPACK_DELEGATE_NO_FLAGS = configuration.XNNPackFlags_TFLITE_XNNPACK_DELEGATE_NO_FLAGS
const XNNPackFlags_TFLITE_XNNPACK_DELEGATE_FLAG_QS8 = configuration.XNNPackFlags_TFLITE_XNNPACK_DELEGATE_FLAG_QS8
const XNNPackFlags_TFLITE_XNNPACK_DELEGATE_FLAG_QU8 = configuration.XNNPackFlags_TFLITE_XNNPACK_DELEGATE_FLAG_QU8
const XNNPackFlags_TFLITE_XNNPACK_DELEGATE_FLAG_QS8_QU8 = configuration.XNNPackFlags_TFLITE_XNNPACK_DELEGATE_FLAG_QS8_QU8
const XNNPackFlags_TFLITE_XNNPACK_DELEGATE_FLAG_FORCE_FP16 = configuration.XNNPackFlags_TFLITE_XNNPACK_DELEGATE_FLAG_FORCE_FP16

var XNNPackFlags_name = configuration.XNNPackFlags_name
var XNNPackFlags_value = configuration.XNNPackFlags_value

type EdgeTpuPowerState = configuration.EdgeTpuPowerState

const EdgeTpuPowerState_UNDEFINED_POWERSTATE = configuration.EdgeTpuPowerState_UNDEFINED_POWERSTATE
const EdgeTpuPowerState_TPU_CORE_OFF = configuration.EdgeTpuPowerState_TPU_CORE_OFF
const EdgeTpuPowerState_READY = configuration.EdgeTpuPowerState_READY
const EdgeTpuPowerState_ACTIVE_MIN_POWER = configuration.EdgeTpuPowerState_ACTIVE_MIN_POWER
const EdgeTpuPowerState_ACTIVE_VERY_LOW_POWER = configuration.EdgeTpuPowerState_ACTIVE_VERY_LOW_POWER
const EdgeTpuPowerState_ACTIVE_LOW_POWER = configuration.EdgeTpuPowerState_ACTIVE_LOW_POWER
const EdgeTpuPowerState_ACTIVE = configuration.EdgeTpuPowerState_ACTIVE
const EdgeTpuPowerState_OVER_DRIVE = configuration.EdgeTpuPowerState_OVER_DRIVE

var EdgeTpuPowerState_name = configuration.EdgeTpuPowerState_name
var EdgeTpuPowerState_value = configuration.EdgeTpuPowerState_value

type BenchmarkEventType = configuration.BenchmarkEventType

const BenchmarkEventType_UNDEFINED_BENCHMARK_EVENT_TYPE = configuration.BenchmarkEventType_UNDEFINED_BENCHMARK_EVENT_TYPE
const BenchmarkEventType_START = configuration.BenchmarkEventType_START
const BenchmarkEventType_END = configuration.BenchmarkEventType_END
const BenchmarkEventType_ERROR = configuration.BenchmarkEventType_ERROR
const BenchmarkEventType_LOGGED = configuration.BenchmarkEventType_LOGGED
const BenchmarkEventType_RECOVERED_ERROR = configuration.BenchmarkEventType_RECOVERED_ERROR

var BenchmarkEventType_name = configuration.BenchmarkEventType_name
var BenchmarkEventType_value = configuration.BenchmarkEventType_value

type BenchmarkStage = configuration.BenchmarkStage

const BenchmarkStage_UNKNOWN = configuration.BenchmarkStage_UNKNOWN
const BenchmarkStage_INITIALIZATION = configuration.BenchmarkStage_INITIALIZATION
const BenchmarkStage_INFERENCE = configuration.BenchmarkStage_INFERENCE

var BenchmarkStage_name = configuration.BenchmarkStage_name
var BenchmarkStage_value = configuration.BenchmarkStage_value

type CoreMLSettings_EnabledDevices = configuration.CoreMLSettings_EnabledDevices

const CoreMLSettings_DEVICES_ALL = configuration.CoreMLSettings_DEVICES_ALL
const CoreMLSettings_DEVICES_WITH_NEURAL_ENGINE = configuration.CoreMLSettings_DEVICES_WITH_NEURAL_ENGINE

var CoreMLSettings_EnabledDevices_name = configuration.CoreMLSettings_EnabledDevices_name
var CoreMLSettings_EnabledDevices_value = configuration.CoreMLSettings_EnabledDevices_value

type EdgeTpuDeviceSpec_PlatformType = configuration.EdgeTpuDeviceSpec_PlatformType

const EdgeTpuDeviceSpec_MMIO = configuration.EdgeTpuDeviceSpec_MMIO
const EdgeTpuDeviceSpec_REFERENCE = configuration.EdgeTpuDeviceSpec_REFERENCE
const EdgeTpuDeviceSpec_SIMULATOR = configuration.EdgeTpuDeviceSpec_SIMULATOR
const EdgeTpuDeviceSpec_REMOTE_SIMULATOR = configuration.EdgeTpuDeviceSpec_REMOTE_SIMULATOR

var EdgeTpuDeviceSpec_PlatformType_name = configuration.EdgeTpuDeviceSpec_PlatformType_name
var EdgeTpuDeviceSpec_PlatformType_value = configuration.EdgeTpuDeviceSpec_PlatformType_value

type EdgeTpuSettings_FloatTruncationType = configuration.EdgeTpuSettings_FloatTruncationType

const EdgeTpuSettings_UNSPECIFIED = configuration.EdgeTpuSettings_UNSPECIFIED
const EdgeTpuSettings_NO_TRUNCATION = configuration.EdgeTpuSettings_NO_TRUNCATION
const EdgeTpuSettings_BFLOAT16 = configuration.EdgeTpuSettings_BFLOAT16
const EdgeTpuSettings_HALF = configuration.EdgeTpuSettings_HALF

var EdgeTpuSettings_FloatTruncationType_name = configuration.EdgeTpuSettings_FloatTruncationType_name
var EdgeTpuSettings_FloatTruncationType_value = configuration.EdgeTpuSettings_FloatTruncationType_value

type EdgeTpuSettings_QosClass = configuration.EdgeTpuSettings_QosClass

const EdgeTpuSettings_QOS_UNDEFINED = configuration.EdgeTpuSettings_QOS_UNDEFINED
const EdgeTpuSettings_BEST_EFFORT = configuration.EdgeTpuSettings_BEST_EFFORT
const EdgeTpuSettings_REALTIME = configuration.EdgeTpuSettings_REALTIME

var EdgeTpuSettings_QosClass_name = configuration.EdgeTpuSettings_QosClass_name
var EdgeTpuSettings_QosClass_value = configuration.EdgeTpuSettings_QosClass_value

type GoogleEdgeTpuSettings_Priority = configuration.GoogleEdgeTpuSettings_Priority

const GoogleEdgeTpuSettings_PRIORITY_UNDEFINED = configuration.GoogleEdgeTpuSettings_PRIORITY_UNDEFINED
const GoogleEdgeTpuSettings_PRIORITY_LOW = configuration.GoogleEdgeTpuSettings_PRIORITY_LOW
const GoogleEdgeTpuSettings_PRIORITY_MEDIUM = configuration.GoogleEdgeTpuSettings_PRIORITY_MEDIUM
const GoogleEdgeTpuSettings_PRIORITY_HIGH = configuration.GoogleEdgeTpuSettings_PRIORITY_HIGH

var GoogleEdgeTpuSettings_Priority_name = configuration.GoogleEdgeTpuSettings_Priority_name
var GoogleEdgeTpuSettings_Priority_value = configuration.GoogleEdgeTpuSettings_Priority_value

type GoogleEdgeTpuSettings_TriState = configuration.GoogleEdgeTpuSettings_TriState

const GoogleEdgeTpuSettings_TRISTATE_UNDEFINED = configuration.GoogleEdgeTpuSettings_TRISTATE_UNDEFINED
const GoogleEdgeTpuSettings_TRISTATE_FALSE = configuration.GoogleEdgeTpuSettings_TRISTATE_FALSE
const GoogleEdgeTpuSettings_TRISTATE_TRUE = configuration.GoogleEdgeTpuSettings_TRISTATE_TRUE

var GoogleEdgeTpuSettings_TriState_name = configuration.GoogleEdgeTpuSettings_TriState_name
var GoogleEdgeTpuSettings_TriState_value = configuration.GoogleEdgeTpuSettings_TriState_value

type CoralSettings_Performance = configuration.CoralSettings_Performance

const CoralSettings_UNDEFINED = configuration.CoralSettings_UNDEFINED
const CoralSettings_MAXIMUM = configuration.CoralSettings_MAXIMUM
const CoralSettings_HIGH = configuration.CoralSettings_HIGH
const CoralSettings_MEDIUM = configuration.CoralSettings_MEDIUM
const CoralSettings_LOW = configuration.CoralSettings_LOW

var CoralSettings_Performance_name = configuration.CoralSettings_Performance_name
var CoralSettings_Performance_value = configuration.CoralSettings_Performance_value

type ComputeSettings = configuration.ComputeSettings
type NNAPISettings = configuration.NNAPISettings
type GPUSettings = configuration.GPUSettings

const Default_GPUSettings_EnableQuantizedInference = configuration.Default_GPUSettings_EnableQuantizedInference
const Default_GPUSettings_InferencePriority1 = configuration.Default_GPUSettings_InferencePriority1
const Default_GPUSettings_InferencePriority2 = configuration.Default_GPUSettings_InferencePriority2
const Default_GPUSettings_InferencePriority3 = configuration.Default_GPUSettings_InferencePriority3

type HexagonSettings = configuration.HexagonSettings
type XNNPackSettings = configuration.XNNPackSettings

const Default_XNNPackSettings_Flags = configuration.Default_XNNPackSettings_Flags

type CoreMLSettings = configuration.CoreMLSettings

const Default_CoreMLSettings_MaxDelegatedPartitions = configuration.Default_CoreMLSettings_MaxDelegatedPartitions
const Default_CoreMLSettings_MinNodesPerPartition = configuration.Default_CoreMLSettings_MinNodesPerPartition

type StableDelegateLoaderSettings = configuration.StableDelegateLoaderSettings
type CompilationCachingSettings = configuration.CompilationCachingSettings
type EdgeTpuDeviceSpec = configuration.EdgeTpuDeviceSpec
type EdgeTpuInactivePowerConfig = configuration.EdgeTpuInactivePowerConfig
type EdgeTpuSettings = configuration.EdgeTpuSettings

const Default_EdgeTpuSettings_InferencePriority = configuration.Default_EdgeTpuSettings_InferencePriority
const Default_EdgeTpuSettings_QosClass = configuration.Default_EdgeTpuSettings_QosClass

type GoogleEdgeTpuSettings = configuration.GoogleEdgeTpuSettings

const Default_GoogleEdgeTpuSettings_LogVerbosity = configuration.Default_GoogleEdgeTpuSettings_LogVerbosity
const Default_GoogleEdgeTpuSettings_EnableTracing = configuration.Default_GoogleEdgeTpuSettings_EnableTracing
const Default_GoogleEdgeTpuSettings_ModelIdentifier = configuration.Default_GoogleEdgeTpuSettings_ModelIdentifier
const Default_GoogleEdgeTpuSettings_UseAsyncApi = configuration.Default_GoogleEdgeTpuSettings_UseAsyncApi
const Default_GoogleEdgeTpuSettings_DelegateShouldManageCacheForInputs = configuration.Default_GoogleEdgeTpuSettings_DelegateShouldManageCacheForInputs
const Default_GoogleEdgeTpuSettings_DelegateShouldManageCacheForOutputs = configuration.Default_GoogleEdgeTpuSettings_DelegateShouldManageCacheForOutputs
const Default_GoogleEdgeTpuSettings_AllowFp16PrecisionForFp32 = configuration.Default_GoogleEdgeTpuSettings_AllowFp16PrecisionForFp32

type CoralSettings = configuration.CoralSettings

const Default_CoralSettings_Performance = configuration.Default_CoralSettings_Performance

type CPUSettings = configuration.CPUSettings

const Default_CPUSettings_NumThreads = configuration.Default_CPUSettings_NumThreads

type ArmNNSettings = configuration.ArmNNSettings
type TFLiteSettings = configuration.TFLiteSettings
type FallbackSettings = configuration.FallbackSettings
type BenchmarkMetric = configuration.BenchmarkMetric
type BenchmarkResult = configuration.BenchmarkResult
type ErrorCode = configuration.ErrorCode
type BenchmarkError = configuration.BenchmarkError
type BenchmarkEvent = configuration.BenchmarkEvent
type BestAccelerationDecision = configuration.BestAccelerationDecision
type BenchmarkInitializationFailure = configuration.BenchmarkInitializationFailure
type MiniBenchmarkEvent = configuration.MiniBenchmarkEvent
type ModelFile = configuration.ModelFile
type ModelIdGroup = configuration.ModelIdGroup
type BenchmarkStoragePaths = configuration.BenchmarkStoragePaths
type ValidationSettings = configuration.ValidationSettings
type MinibenchmarkSettings = configuration.MinibenchmarkSettings
type BenchmarkEventStorage = configuration.BenchmarkEventStorage
type BenchmarkResult_InferenceOutput = configuration.BenchmarkResult_InferenceOutput

var File_tensorflow_lite_experimental_acceleration_configuration_configuration_proto protoreflect.FileDescriptor

var file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x74,
	0x65, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x74,
	0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69, 0x74, 0x65, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0xdb, 0x01, 0x0a, 0x10,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x42, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x62, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x6c, 0x69,
	0x74, 0x65, 0x2f, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x6c, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x54, 0x50, 0x58,
	0xaa, 0x02, 0x0c, 0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xca,
	0x02, 0x0c, 0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0xe2, 0x02,
	0x18, 0x54, 0x66, 0x6c, 0x69, 0x74, 0x65, 0x5c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x54, 0x66, 0x6c, 0x69,
	0x74, 0x65, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x00,
}

var file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_goTypes = []interface{}{}
var file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_init() }
func file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_init() {
	if File_tensorflow_lite_experimental_acceleration_configuration_configuration_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_goTypes,
		DependencyIndexes: file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_depIdxs,
	}.Build()
	File_tensorflow_lite_experimental_acceleration_configuration_configuration_proto = out.File
	file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_rawDesc = nil
	file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_goTypes = nil
	file_tensorflow_lite_experimental_acceleration_configuration_configuration_proto_depIdxs = nil
}