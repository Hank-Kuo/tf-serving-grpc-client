// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tensorflow/core/tpu/kernels/tpu_compilation_cache.proto

package kernels

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

// TpuCompilationCacheServiceExternalClient is the client API for TpuCompilationCacheServiceExternal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TpuCompilationCacheServiceExternalClient interface {
	// This method requests the cached proto that the TPU execute op has been
	// instructed to execute.
	GetTpuProgram(ctx context.Context, in *GetTpuProgramRequest, opts ...grpc.CallOption) (*GetTpuProgramResponseExternal, error)
}

type tpuCompilationCacheServiceExternalClient struct {
	cc grpc.ClientConnInterface
}

func NewTpuCompilationCacheServiceExternalClient(cc grpc.ClientConnInterface) TpuCompilationCacheServiceExternalClient {
	return &tpuCompilationCacheServiceExternalClient{cc}
}

func (c *tpuCompilationCacheServiceExternalClient) GetTpuProgram(ctx context.Context, in *GetTpuProgramRequest, opts ...grpc.CallOption) (*GetTpuProgramResponseExternal, error) {
	out := new(GetTpuProgramResponseExternal)
	err := c.cc.Invoke(ctx, "/tensorflow.tpu.TpuCompilationCacheServiceExternal/GetTpuProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TpuCompilationCacheServiceExternalServer is the server API for TpuCompilationCacheServiceExternal service.
// All implementations should embed UnimplementedTpuCompilationCacheServiceExternalServer
// for forward compatibility
type TpuCompilationCacheServiceExternalServer interface {
	// This method requests the cached proto that the TPU execute op has been
	// instructed to execute.
	GetTpuProgram(context.Context, *GetTpuProgramRequest) (*GetTpuProgramResponseExternal, error)
}

// UnimplementedTpuCompilationCacheServiceExternalServer should be embedded to have forward compatible implementations.
type UnimplementedTpuCompilationCacheServiceExternalServer struct {
}

func (UnimplementedTpuCompilationCacheServiceExternalServer) GetTpuProgram(context.Context, *GetTpuProgramRequest) (*GetTpuProgramResponseExternal, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTpuProgram not implemented")
}

// UnsafeTpuCompilationCacheServiceExternalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TpuCompilationCacheServiceExternalServer will
// result in compilation errors.
type UnsafeTpuCompilationCacheServiceExternalServer interface {
	mustEmbedUnimplementedTpuCompilationCacheServiceExternalServer()
}

func RegisterTpuCompilationCacheServiceExternalServer(s grpc.ServiceRegistrar, srv TpuCompilationCacheServiceExternalServer) {
	s.RegisterService(&TpuCompilationCacheServiceExternal_ServiceDesc, srv)
}

func _TpuCompilationCacheServiceExternal_GetTpuProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTpuProgramRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TpuCompilationCacheServiceExternalServer).GetTpuProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.tpu.TpuCompilationCacheServiceExternal/GetTpuProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TpuCompilationCacheServiceExternalServer).GetTpuProgram(ctx, req.(*GetTpuProgramRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TpuCompilationCacheServiceExternal_ServiceDesc is the grpc.ServiceDesc for TpuCompilationCacheServiceExternal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TpuCompilationCacheServiceExternal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.tpu.TpuCompilationCacheServiceExternal",
	HandlerType: (*TpuCompilationCacheServiceExternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTpuProgram",
			Handler:    _TpuCompilationCacheServiceExternal_GetTpuProgram_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow/core/tpu/kernels/tpu_compilation_cache.proto",
}
