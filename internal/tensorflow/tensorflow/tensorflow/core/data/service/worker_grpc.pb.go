// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: tensorflow/core/data/service/worker.proto

package service

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

// WorkerServiceClient is the client API for WorkerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerServiceClient interface {
	// Processes a task for a dataset, making elements available to clients.
	ProcessTask(ctx context.Context, in *ProcessTaskRequest, opts ...grpc.CallOption) (*ProcessTaskResponse, error)
	// Gets the next dataset element.
	GetElement(ctx context.Context, in *GetElementRequest, opts ...grpc.CallOption) (*GetElementResponse, error)
	// Gets the tasks currently being executed by the worker.
	GetWorkerTasks(ctx context.Context, in *GetWorkerTasksRequest, opts ...grpc.CallOption) (*GetWorkerTasksResponse, error)
	// Gets the progresses of the snapshot tasks currently being executed by the
	// worker.
	GetSnapshotTaskProgresses(ctx context.Context, in *GetSnapshotTaskProgressesRequest, opts ...grpc.CallOption) (*GetSnapshotTaskProgressesResponse, error)
}

type workerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerServiceClient(cc grpc.ClientConnInterface) WorkerServiceClient {
	return &workerServiceClient{cc}
}

func (c *workerServiceClient) ProcessTask(ctx context.Context, in *ProcessTaskRequest, opts ...grpc.CallOption) (*ProcessTaskResponse, error) {
	out := new(ProcessTaskResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.WorkerService/ProcessTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetElement(ctx context.Context, in *GetElementRequest, opts ...grpc.CallOption) (*GetElementResponse, error) {
	out := new(GetElementResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.WorkerService/GetElement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetWorkerTasks(ctx context.Context, in *GetWorkerTasksRequest, opts ...grpc.CallOption) (*GetWorkerTasksResponse, error) {
	out := new(GetWorkerTasksResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.WorkerService/GetWorkerTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workerServiceClient) GetSnapshotTaskProgresses(ctx context.Context, in *GetSnapshotTaskProgressesRequest, opts ...grpc.CallOption) (*GetSnapshotTaskProgressesResponse, error) {
	out := new(GetSnapshotTaskProgressesResponse)
	err := c.cc.Invoke(ctx, "/tensorflow.data.WorkerService/GetSnapshotTaskProgresses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkerServiceServer is the server API for WorkerService service.
// All implementations should embed UnimplementedWorkerServiceServer
// for forward compatibility
type WorkerServiceServer interface {
	// Processes a task for a dataset, making elements available to clients.
	ProcessTask(context.Context, *ProcessTaskRequest) (*ProcessTaskResponse, error)
	// Gets the next dataset element.
	GetElement(context.Context, *GetElementRequest) (*GetElementResponse, error)
	// Gets the tasks currently being executed by the worker.
	GetWorkerTasks(context.Context, *GetWorkerTasksRequest) (*GetWorkerTasksResponse, error)
	// Gets the progresses of the snapshot tasks currently being executed by the
	// worker.
	GetSnapshotTaskProgresses(context.Context, *GetSnapshotTaskProgressesRequest) (*GetSnapshotTaskProgressesResponse, error)
}

// UnimplementedWorkerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedWorkerServiceServer struct {
}

func (UnimplementedWorkerServiceServer) ProcessTask(context.Context, *ProcessTaskRequest) (*ProcessTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessTask not implemented")
}
func (UnimplementedWorkerServiceServer) GetElement(context.Context, *GetElementRequest) (*GetElementResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetElement not implemented")
}
func (UnimplementedWorkerServiceServer) GetWorkerTasks(context.Context, *GetWorkerTasksRequest) (*GetWorkerTasksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkerTasks not implemented")
}
func (UnimplementedWorkerServiceServer) GetSnapshotTaskProgresses(context.Context, *GetSnapshotTaskProgressesRequest) (*GetSnapshotTaskProgressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapshotTaskProgresses not implemented")
}

// UnsafeWorkerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WorkerServiceServer will
// result in compilation errors.
type UnsafeWorkerServiceServer interface {
	mustEmbedUnimplementedWorkerServiceServer()
}

func RegisterWorkerServiceServer(s grpc.ServiceRegistrar, srv WorkerServiceServer) {
	s.RegisterService(&WorkerService_ServiceDesc, srv)
}

func _WorkerService_ProcessTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).ProcessTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.WorkerService/ProcessTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).ProcessTask(ctx, req.(*ProcessTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetElement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetElementRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetElement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.WorkerService/GetElement",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetElement(ctx, req.(*GetElementRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetWorkerTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkerTasksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetWorkerTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.WorkerService/GetWorkerTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetWorkerTasks(ctx, req.(*GetWorkerTasksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkerService_GetSnapshotTaskProgresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotTaskProgressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkerServiceServer).GetSnapshotTaskProgresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tensorflow.data.WorkerService/GetSnapshotTaskProgresses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkerServiceServer).GetSnapshotTaskProgresses(ctx, req.(*GetSnapshotTaskProgressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WorkerService_ServiceDesc is the grpc.ServiceDesc for WorkerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WorkerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tensorflow.data.WorkerService",
	HandlerType: (*WorkerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProcessTask",
			Handler:    _WorkerService_ProcessTask_Handler,
		},
		{
			MethodName: "GetElement",
			Handler:    _WorkerService_GetElement_Handler,
		},
		{
			MethodName: "GetWorkerTasks",
			Handler:    _WorkerService_GetWorkerTasks_Handler,
		},
		{
			MethodName: "GetSnapshotTaskProgresses",
			Handler:    _WorkerService_GetSnapshotTaskProgresses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tensorflow/core/data/service/worker.proto",
}
