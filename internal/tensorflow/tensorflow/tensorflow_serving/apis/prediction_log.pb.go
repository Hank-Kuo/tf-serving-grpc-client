// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: tensorflow_serving/apis/prediction_log.proto

package apis

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

type ClassifyLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *ClassificationRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response *ClassificationResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *ClassifyLog) Reset() {
	*x = ClassifyLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassifyLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassifyLog) ProtoMessage() {}

func (x *ClassifyLog) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassifyLog.ProtoReflect.Descriptor instead.
func (*ClassifyLog) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_prediction_log_proto_rawDescGZIP(), []int{0}
}

func (x *ClassifyLog) GetRequest() *ClassificationRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ClassifyLog) GetResponse() *ClassificationResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type RegressLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *RegressionRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response *RegressionResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *RegressLog) Reset() {
	*x = RegressLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegressLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegressLog) ProtoMessage() {}

func (x *RegressLog) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegressLog.ProtoReflect.Descriptor instead.
func (*RegressLog) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_prediction_log_proto_rawDescGZIP(), []int{1}
}

func (x *RegressLog) GetRequest() *RegressionRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *RegressLog) GetResponse() *RegressionResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type PredictLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *PredictRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response *PredictResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *PredictLog) Reset() {
	*x = PredictLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictLog) ProtoMessage() {}

func (x *PredictLog) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictLog.ProtoReflect.Descriptor instead.
func (*PredictLog) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_prediction_log_proto_rawDescGZIP(), []int{2}
}

func (x *PredictLog) GetRequest() *PredictRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *PredictLog) GetResponse() *PredictResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type PredictStreamedLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  []*PredictRequest  `protobuf:"bytes,1,rep,name=request,proto3" json:"request,omitempty"`
	Response []*PredictResponse `protobuf:"bytes,2,rep,name=response,proto3" json:"response,omitempty"`
}

func (x *PredictStreamedLog) Reset() {
	*x = PredictStreamedLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictStreamedLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictStreamedLog) ProtoMessage() {}

func (x *PredictStreamedLog) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictStreamedLog.ProtoReflect.Descriptor instead.
func (*PredictStreamedLog) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_prediction_log_proto_rawDescGZIP(), []int{3}
}

func (x *PredictStreamedLog) GetRequest() []*PredictRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *PredictStreamedLog) GetResponse() []*PredictResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type MultiInferenceLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *MultiInferenceRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response *MultiInferenceResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *MultiInferenceLog) Reset() {
	*x = MultiInferenceLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiInferenceLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiInferenceLog) ProtoMessage() {}

func (x *MultiInferenceLog) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiInferenceLog.ProtoReflect.Descriptor instead.
func (*MultiInferenceLog) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_prediction_log_proto_rawDescGZIP(), []int{4}
}

func (x *MultiInferenceLog) GetRequest() *MultiInferenceRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *MultiInferenceLog) GetResponse() *MultiInferenceResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

type SessionRunLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request  *SessionRunRequest  `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Response *SessionRunResponse `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SessionRunLog) Reset() {
	*x = SessionRunLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionRunLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionRunLog) ProtoMessage() {}

func (x *SessionRunLog) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionRunLog.ProtoReflect.Descriptor instead.
func (*SessionRunLog) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_prediction_log_proto_rawDescGZIP(), []int{5}
}

func (x *SessionRunLog) GetRequest() *SessionRunRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *SessionRunLog) GetResponse() *SessionRunResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

// Logged model inference request.
type PredictionLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogMetadata *LogMetadata `protobuf:"bytes,1,opt,name=log_metadata,json=logMetadata,proto3" json:"log_metadata,omitempty"`
	// Types that are assignable to LogType:
	//
	//	*PredictionLog_ClassifyLog
	//	*PredictionLog_RegressLog
	//	*PredictionLog_PredictLog
	//	*PredictionLog_PredictStreamedLog
	//	*PredictionLog_MultiInferenceLog
	//	*PredictionLog_SessionRunLog
	LogType isPredictionLog_LogType `protobuf_oneof:"log_type"`
}

func (x *PredictionLog) Reset() {
	*x = PredictionLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PredictionLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PredictionLog) ProtoMessage() {}

func (x *PredictionLog) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_apis_prediction_log_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PredictionLog.ProtoReflect.Descriptor instead.
func (*PredictionLog) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_apis_prediction_log_proto_rawDescGZIP(), []int{6}
}

func (x *PredictionLog) GetLogMetadata() *LogMetadata {
	if x != nil {
		return x.LogMetadata
	}
	return nil
}

func (m *PredictionLog) GetLogType() isPredictionLog_LogType {
	if m != nil {
		return m.LogType
	}
	return nil
}

func (x *PredictionLog) GetClassifyLog() *ClassifyLog {
	if x, ok := x.GetLogType().(*PredictionLog_ClassifyLog); ok {
		return x.ClassifyLog
	}
	return nil
}

func (x *PredictionLog) GetRegressLog() *RegressLog {
	if x, ok := x.GetLogType().(*PredictionLog_RegressLog); ok {
		return x.RegressLog
	}
	return nil
}

func (x *PredictionLog) GetPredictLog() *PredictLog {
	if x, ok := x.GetLogType().(*PredictionLog_PredictLog); ok {
		return x.PredictLog
	}
	return nil
}

func (x *PredictionLog) GetPredictStreamedLog() *PredictStreamedLog {
	if x, ok := x.GetLogType().(*PredictionLog_PredictStreamedLog); ok {
		return x.PredictStreamedLog
	}
	return nil
}

func (x *PredictionLog) GetMultiInferenceLog() *MultiInferenceLog {
	if x, ok := x.GetLogType().(*PredictionLog_MultiInferenceLog); ok {
		return x.MultiInferenceLog
	}
	return nil
}

func (x *PredictionLog) GetSessionRunLog() *SessionRunLog {
	if x, ok := x.GetLogType().(*PredictionLog_SessionRunLog); ok {
		return x.SessionRunLog
	}
	return nil
}

type isPredictionLog_LogType interface {
	isPredictionLog_LogType()
}

type PredictionLog_ClassifyLog struct {
	ClassifyLog *ClassifyLog `protobuf:"bytes,2,opt,name=classify_log,json=classifyLog,proto3,oneof"`
}

type PredictionLog_RegressLog struct {
	RegressLog *RegressLog `protobuf:"bytes,3,opt,name=regress_log,json=regressLog,proto3,oneof"`
}

type PredictionLog_PredictLog struct {
	PredictLog *PredictLog `protobuf:"bytes,6,opt,name=predict_log,json=predictLog,proto3,oneof"`
}

type PredictionLog_PredictStreamedLog struct {
	PredictStreamedLog *PredictStreamedLog `protobuf:"bytes,7,opt,name=predict_streamed_log,json=predictStreamedLog,proto3,oneof"`
}

type PredictionLog_MultiInferenceLog struct {
	MultiInferenceLog *MultiInferenceLog `protobuf:"bytes,4,opt,name=multi_inference_log,json=multiInferenceLog,proto3,oneof"`
}

type PredictionLog_SessionRunLog struct {
	SessionRunLog *SessionRunLog `protobuf:"bytes,5,opt,name=session_run_log,json=sessionRunLog,proto3,oneof"`
}

func (*PredictionLog_ClassifyLog) isPredictionLog_LogType() {}

func (*PredictionLog_RegressLog) isPredictionLog_LogType() {}

func (*PredictionLog_PredictLog) isPredictionLog_LogType() {}

func (*PredictionLog_PredictStreamedLog) isPredictionLog_LogType() {}

func (*PredictionLog_MultiInferenceLog) isPredictionLog_LogType() {}

func (*PredictionLog_SessionRunLog) isPredictionLog_LogType() {}

var File_tensorflow_serving_apis_prediction_log_proto protoreflect.FileDescriptor

var file_tensorflow_serving_apis_prediction_log_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x1a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x25, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2d, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9a, 0x01, 0x0a, 0x0b, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x4c, 0x6f, 0x67,
	0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x91, 0x01,
	0x0a, 0x0a, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x3f, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x8b, 0x01, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x4c, 0x6f, 0x67,
	0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x93, 0x01, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x3c, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65,
	0x64, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x11, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49,
	0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x43, 0x0a, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x46, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49, 0x6e, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x3f, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xad, 0x04, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f,
	0x67, 0x12, 0x42, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x6f, 0x67,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x44, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66,
	0x79, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x4c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x0b,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x79, 0x4c, 0x6f, 0x67, 0x12, 0x41, 0x0a, 0x0b, 0x72,
	0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67,
	0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x41,
	0x0a, 0x0b, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74,
	0x4c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x4c, 0x6f,
	0x67, 0x12, 0x5a, 0x0a, 0x14, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x12, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x64, 0x4c, 0x6f, 0x67, 0x12, 0x57, 0x0a,
	0x13, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x69, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x4c, 0x6f,
	0x67, 0x48, 0x00, 0x52, 0x11, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x49, 0x6e, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x4b, 0x0a, 0x0f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x75, 0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e, 0x4c,
	0x6f, 0x67, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6e,
	0x4c, 0x6f, 0x67, 0x42, 0x0a, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42,
	0xdc, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x42, 0x12, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x42, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x54, 0x53, 0x58, 0xaa, 0x02, 0x12,
	0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0xca, 0x02, 0x12, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0xe2, 0x02, 0x1e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x54, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_apis_prediction_log_proto_rawDescOnce sync.Once
	file_tensorflow_serving_apis_prediction_log_proto_rawDescData = file_tensorflow_serving_apis_prediction_log_proto_rawDesc
)

func file_tensorflow_serving_apis_prediction_log_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_apis_prediction_log_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_apis_prediction_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_apis_prediction_log_proto_rawDescData)
	})
	return file_tensorflow_serving_apis_prediction_log_proto_rawDescData
}

var file_tensorflow_serving_apis_prediction_log_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tensorflow_serving_apis_prediction_log_proto_goTypes = []interface{}{
	(*ClassifyLog)(nil),            // 0: tensorflow.serving.ClassifyLog
	(*RegressLog)(nil),             // 1: tensorflow.serving.RegressLog
	(*PredictLog)(nil),             // 2: tensorflow.serving.PredictLog
	(*PredictStreamedLog)(nil),     // 3: tensorflow.serving.PredictStreamedLog
	(*MultiInferenceLog)(nil),      // 4: tensorflow.serving.MultiInferenceLog
	(*SessionRunLog)(nil),          // 5: tensorflow.serving.SessionRunLog
	(*PredictionLog)(nil),          // 6: tensorflow.serving.PredictionLog
	(*ClassificationRequest)(nil),  // 7: tensorflow.serving.ClassificationRequest
	(*ClassificationResponse)(nil), // 8: tensorflow.serving.ClassificationResponse
	(*RegressionRequest)(nil),      // 9: tensorflow.serving.RegressionRequest
	(*RegressionResponse)(nil),     // 10: tensorflow.serving.RegressionResponse
	(*PredictRequest)(nil),         // 11: tensorflow.serving.PredictRequest
	(*PredictResponse)(nil),        // 12: tensorflow.serving.PredictResponse
	(*MultiInferenceRequest)(nil),  // 13: tensorflow.serving.MultiInferenceRequest
	(*MultiInferenceResponse)(nil), // 14: tensorflow.serving.MultiInferenceResponse
	(*SessionRunRequest)(nil),      // 15: tensorflow.serving.SessionRunRequest
	(*SessionRunResponse)(nil),     // 16: tensorflow.serving.SessionRunResponse
	(*LogMetadata)(nil),            // 17: tensorflow.serving.LogMetadata
}
var file_tensorflow_serving_apis_prediction_log_proto_depIdxs = []int32{
	7,  // 0: tensorflow.serving.ClassifyLog.request:type_name -> tensorflow.serving.ClassificationRequest
	8,  // 1: tensorflow.serving.ClassifyLog.response:type_name -> tensorflow.serving.ClassificationResponse
	9,  // 2: tensorflow.serving.RegressLog.request:type_name -> tensorflow.serving.RegressionRequest
	10, // 3: tensorflow.serving.RegressLog.response:type_name -> tensorflow.serving.RegressionResponse
	11, // 4: tensorflow.serving.PredictLog.request:type_name -> tensorflow.serving.PredictRequest
	12, // 5: tensorflow.serving.PredictLog.response:type_name -> tensorflow.serving.PredictResponse
	11, // 6: tensorflow.serving.PredictStreamedLog.request:type_name -> tensorflow.serving.PredictRequest
	12, // 7: tensorflow.serving.PredictStreamedLog.response:type_name -> tensorflow.serving.PredictResponse
	13, // 8: tensorflow.serving.MultiInferenceLog.request:type_name -> tensorflow.serving.MultiInferenceRequest
	14, // 9: tensorflow.serving.MultiInferenceLog.response:type_name -> tensorflow.serving.MultiInferenceResponse
	15, // 10: tensorflow.serving.SessionRunLog.request:type_name -> tensorflow.serving.SessionRunRequest
	16, // 11: tensorflow.serving.SessionRunLog.response:type_name -> tensorflow.serving.SessionRunResponse
	17, // 12: tensorflow.serving.PredictionLog.log_metadata:type_name -> tensorflow.serving.LogMetadata
	0,  // 13: tensorflow.serving.PredictionLog.classify_log:type_name -> tensorflow.serving.ClassifyLog
	1,  // 14: tensorflow.serving.PredictionLog.regress_log:type_name -> tensorflow.serving.RegressLog
	2,  // 15: tensorflow.serving.PredictionLog.predict_log:type_name -> tensorflow.serving.PredictLog
	3,  // 16: tensorflow.serving.PredictionLog.predict_streamed_log:type_name -> tensorflow.serving.PredictStreamedLog
	4,  // 17: tensorflow.serving.PredictionLog.multi_inference_log:type_name -> tensorflow.serving.MultiInferenceLog
	5,  // 18: tensorflow.serving.PredictionLog.session_run_log:type_name -> tensorflow.serving.SessionRunLog
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_apis_prediction_log_proto_init() }
func file_tensorflow_serving_apis_prediction_log_proto_init() {
	if File_tensorflow_serving_apis_prediction_log_proto != nil {
		return
	}
	file_tensorflow_serving_apis_classification_proto_init()
	file_tensorflow_serving_apis_inference_proto_init()
	file_tensorflow_serving_apis_logging_proto_init()
	file_tensorflow_serving_apis_predict_proto_init()
	file_tensorflow_serving_apis_regression_proto_init()
	file_tensorflow_serving_apis_session_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_apis_prediction_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassifyLog); i {
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
		file_tensorflow_serving_apis_prediction_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegressLog); i {
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
		file_tensorflow_serving_apis_prediction_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictLog); i {
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
		file_tensorflow_serving_apis_prediction_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictStreamedLog); i {
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
		file_tensorflow_serving_apis_prediction_log_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiInferenceLog); i {
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
		file_tensorflow_serving_apis_prediction_log_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionRunLog); i {
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
		file_tensorflow_serving_apis_prediction_log_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PredictionLog); i {
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
	file_tensorflow_serving_apis_prediction_log_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*PredictionLog_ClassifyLog)(nil),
		(*PredictionLog_RegressLog)(nil),
		(*PredictionLog_PredictLog)(nil),
		(*PredictionLog_PredictStreamedLog)(nil),
		(*PredictionLog_MultiInferenceLog)(nil),
		(*PredictionLog_SessionRunLog)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_serving_apis_prediction_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_apis_prediction_log_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_apis_prediction_log_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_apis_prediction_log_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_apis_prediction_log_proto = out.File
	file_tensorflow_serving_apis_prediction_log_proto_rawDesc = nil
	file_tensorflow_serving_apis_prediction_log_proto_goTypes = nil
	file_tensorflow_serving_apis_prediction_log_proto_depIdxs = nil
}
