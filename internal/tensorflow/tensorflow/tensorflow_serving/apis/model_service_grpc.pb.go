// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tensorflow_serving/apis/model_service.proto

package apis

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ModelServiceClient is the client API for ModelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelServiceClient interface {
	// Gets status of model. If the ModelSpec in the request does not specify
	// version, information about all versions of the model will be returned. If
	// the ModelSpec in the request does specify a version, the status of only
	// that version will be returned.
	GetModelStatus(ctx context.Context, in *GetModelStatusRequest, opts ...grpc.CallOption) (*GetModelStatusResponse, error)
	// Reloads the set of served models. The new config supersedes the old one,
	// so if a model is omitted from the new config it will be unloaded and no
	// longer served.
	HandleReloadConfigRequest(ctx context.Context, in *ReloadConfigRequest, opts ...grpc.CallOption) (*ReloadConfigResponse, error)
}

type modelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelServiceClient(cc grpc.ClientConnInterface) ModelServiceClient {
	return &modelServiceClient{cc}
}

func (c *modelServiceClient) GetModelStatus(ctx context.Context, in *GetModelStatusRequest, opts ...grpc.CallOption) (*GetModelStatusResponse, error) {
	out := new(GetModelStatusResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.ModelService/GetModelStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) HandleReloadConfigRequest(ctx context.Context, in *ReloadConfigRequest, opts ...grpc.CallOption) (*ReloadConfigResponse, error) {
	out := new(ReloadConfigResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.serving.ModelService/HandleReloadConfigRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ModelServiceServer is the server API for ModelService service.
// All implementations should embed UnimplementedModelServiceServer
// for forward compatibility
type ModelServiceServer interface {
	// Gets status of model. If the ModelSpec in the request does not specify
	// version, information about all versions of the model will be returned. If
	// the ModelSpec in the request does specify a version, the status of only
	// that version will be returned.
	GetModelStatus(context.Context, *GetModelStatusRequest) (*GetModelStatusResponse, error)
	// Reloads the set of served models. The new config supersedes the old one,
	// so if a model is omitted from the new config it will be unloaded and no
	// longer served.
	HandleReloadConfigRequest(context.Context, *ReloadConfigRequest) (*ReloadConfigResponse, error)
}

// UnimplementedModelServiceServer should be embedded to have forward compatible implementations.
type UnimplementedModelServiceServer struct {
}

func (UnimplementedModelServiceServer) GetModelStatus(context.Context, *GetModelStatusRequest) (*GetModelStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetModelStatus not implemented")
}
func (UnimplementedModelServiceServer) HandleReloadConfigRequest(context.Context, *ReloadConfigRequest) (*ReloadConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleReloadConfigRequest not implemented")
}

// UnsafeModelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelServiceServer will
// result in compilation errors.
type UnsafeModelServiceServer interface {
	mustEmbedUnimplementedModelServiceServer()
}

func RegisterModelServiceServer(s grpc.ServiceRegistrar, srv ModelServiceServer) {
	s.RegisterService(&ModelService_ServiceDesc, srv)
}

func _ModelService_GetModelStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetModelStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).GetModelStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.ModelService/GetModelStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).GetModelStatus(ctx, req.(*GetModelStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_HandleReloadConfigRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).HandleReloadConfigRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.serving.ModelService/HandleReloadConfigRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).HandleReloadConfigRequest(ctx, req.(*ReloadConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ModelService_ServiceDesc is the grpc.ServiceDesc for ModelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.serving.ModelService",
	HandlerType: (*ModelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetModelStatus",
			Handler:    _ModelService_GetModelStatus_Handler,
		},
		{
			MethodName: "HandleReloadConfigRequest",
			Handler:    _ModelService_HandleReloadConfigRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow_serving/apis/model_service.proto",
}
