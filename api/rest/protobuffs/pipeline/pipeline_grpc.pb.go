// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: pipeline.proto

package pipelinepb

import (
	"log"
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PipelineService_CreatePipeline_FullMethodName    = "/pipeline.PipelineService/CreatePipeline"
	PipelineService_GetUserPipelines_FullMethodName  = "/pipeline.PipelineService/GetUserPipelines"
	PipelineService_GetPipelineStages_FullMethodName = "/pipeline.PipelineService/GetPipelineStages"
)

// PipelineServiceClient is the client API for PipelineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineServiceClient interface {
	CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*PipelineResponse, error)
	GetUserPipelines(ctx context.Context, in *GetUserPipelinesRequest, opts ...grpc.CallOption) (*PipelinesResponse, error)
	GetPipelineStages(ctx context.Context, in *GetPipelineStagesRequest, opts ...grpc.CallOption) (*StagesResponse, error)
}

type pipelineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineServiceClient(cc grpc.ClientConnInterface) PipelineServiceClient {
	return &pipelineServiceClient{cc}
}

func (c *pipelineServiceClient) CreatePipeline(ctx context.Context, in *CreatePipelineRequest, opts ...grpc.CallOption) (*PipelineResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PipelineResponse)
	err := c.cc.Invoke(ctx, PipelineService_CreatePipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) GetUserPipelines(ctx context.Context, in *GetUserPipelinesRequest, opts ...grpc.CallOption) (*PipelinesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PipelinesResponse)
	err := c.cc.Invoke(ctx, PipelineService_GetUserPipelines_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineServiceClient) GetPipelineStages(ctx context.Context, in *GetPipelineStagesRequest, opts ...grpc.CallOption) (*StagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StagesResponse)
	err := c.cc.Invoke(ctx, PipelineService_GetPipelineStages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineServiceServer is the server API for PipelineService service.
// All implementations must embed UnimplementedPipelineServiceServer
// for forward compatibility.
type PipelineServiceServer interface {
	CreatePipeline(context.Context, *CreatePipelineRequest) (*PipelineResponse, error)
	GetUserPipelines(context.Context, *GetUserPipelinesRequest) (*PipelinesResponse, error)
	GetPipelineStages(context.Context, *GetPipelineStagesRequest) (*StagesResponse, error)
	mustEmbedUnimplementedPipelineServiceServer()
}

// UnimplementedPipelineServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPipelineServiceServer struct{}

func (UnimplementedPipelineServiceServer) CreatePipeline(context.Context, *CreatePipelineRequest) (*PipelineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePipeline not implemented")
}
func (UnimplementedPipelineServiceServer) GetUserPipelines(context.Context, *GetUserPipelinesRequest) (*PipelinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPipelines not implemented")
}
func (UnimplementedPipelineServiceServer) GetPipelineStages(context.Context, *GetPipelineStagesRequest) (*StagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipelineStages not implemented")
}
func (UnimplementedPipelineServiceServer) mustEmbedUnimplementedPipelineServiceServer() {}
func (UnimplementedPipelineServiceServer) testEmbeddedByValue()                         {}

// UnsafePipelineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineServiceServer will
// result in compilation errors.
type UnsafePipelineServiceServer interface {
	mustEmbedUnimplementedPipelineServiceServer()
}

func RegisterPipelineServiceServer(s grpc.ServiceRegistrar, srv PipelineServiceServer) {
	// If the following call pancis, it indicates UnimplementedPipelineServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PipelineService_ServiceDesc, srv)
}

func _PipelineService_CreatePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_CreatePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).CreatePipeline(ctx, req.(*CreatePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_GetUserPipelines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPipelinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).GetUserPipelines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_GetUserPipelines_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).GetUserPipelines(ctx, req.(*GetUserPipelinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineService_GetPipelineStages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPipelineStagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineServiceServer).GetPipelineStages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineService_GetPipelineStages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineServiceServer).GetPipelineStages(ctx, req.(*GetPipelineStagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PipelineService_ServiceDesc is the grpc.ServiceDesc for PipelineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pipeline.PipelineService",
	HandlerType: (*PipelineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipeline",
			Handler:    _PipelineService_CreatePipeline_Handler,
		},
		{
			MethodName: "GetUserPipelines",
			Handler:    _PipelineService_GetUserPipelines_Handler,
		},
		{
			MethodName: "GetPipelineStages",
			Handler:    _PipelineService_GetPipelineStages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pipeline.proto",
}

const (
	PipelineOrchestratorService_ExecutePipeline_FullMethodName    = "/pipeline.PipelineOrchestratorService/ExecutePipeline"
	PipelineOrchestratorService_GetPipelineStatus_FullMethodName  = "/pipeline.PipelineOrchestratorService/GetPipelineStatus"
	PipelineOrchestratorService_CancelPipeline_FullMethodName     = "/pipeline.PipelineOrchestratorService/CancelPipeline"
	PipelineOrchestratorService_AddStageToPipeline_FullMethodName = "/pipeline.PipelineOrchestratorService/AddStageToPipeline"
	PipelineOrchestratorService_DeletePipeline_FullMethodName     = "/pipeline.PipelineOrchestratorService/DeletePipeline"
)

// PipelineOrchestratorServiceClient is the client API for PipelineOrchestratorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PipelineOrchestratorServiceClient interface {
	ExecutePipeline(ctx context.Context, in *ExecutePipelineRequest, opts ...grpc.CallOption) (*ExecutionResponse, error)
	GetPipelineStatus(ctx context.Context, in *PipelineIDRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	CancelPipeline(ctx context.Context, in *PipelineIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AddStageToPipeline(ctx context.Context, in *AddStageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeletePipeline(ctx context.Context, in *PipelineIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pipelineOrchestratorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPipelineOrchestratorServiceClient(cc grpc.ClientConnInterface) PipelineOrchestratorServiceClient {
	return &pipelineOrchestratorServiceClient{cc}
}

func (c *pipelineOrchestratorServiceClient) ExecutePipeline(ctx context.Context, in *ExecutePipelineRequest, opts ...grpc.CallOption) (*ExecutionResponse, error) {
	log.Printf("pipeline_grpc.pb - ExecutePipeline function called\n")
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecutionResponse)
	err := c.cc.Invoke(ctx, PipelineOrchestratorService_ExecutePipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineOrchestratorServiceClient) GetPipelineStatus(ctx context.Context, in *PipelineIDRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, PipelineOrchestratorService_GetPipelineStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineOrchestratorServiceClient) CancelPipeline(ctx context.Context, in *PipelineIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PipelineOrchestratorService_CancelPipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineOrchestratorServiceClient) AddStageToPipeline(ctx context.Context, in *AddStageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PipelineOrchestratorService_AddStageToPipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pipelineOrchestratorServiceClient) DeletePipeline(ctx context.Context, in *PipelineIDRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PipelineOrchestratorService_DeletePipeline_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PipelineOrchestratorServiceServer is the server API for PipelineOrchestratorService service.
// All implementations must embed UnimplementedPipelineOrchestratorServiceServer
// for forward compatibility.
type PipelineOrchestratorServiceServer interface {
	ExecutePipeline(context.Context, *ExecutePipelineRequest) (*ExecutionResponse, error)
	GetPipelineStatus(context.Context, *PipelineIDRequest) (*StatusResponse, error)
	CancelPipeline(context.Context, *PipelineIDRequest) (*emptypb.Empty, error)
	AddStageToPipeline(context.Context, *AddStageRequest) (*emptypb.Empty, error)
	DeletePipeline(context.Context, *PipelineIDRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPipelineOrchestratorServiceServer()
}

// UnimplementedPipelineOrchestratorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPipelineOrchestratorServiceServer struct{}

func (UnimplementedPipelineOrchestratorServiceServer) ExecutePipeline(context.Context, *ExecutePipelineRequest) (*ExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecutePipeline not implemented")
}
func (UnimplementedPipelineOrchestratorServiceServer) GetPipelineStatus(context.Context, *PipelineIDRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPipelineStatus not implemented")
}
func (UnimplementedPipelineOrchestratorServiceServer) CancelPipeline(context.Context, *PipelineIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelPipeline not implemented")
}
func (UnimplementedPipelineOrchestratorServiceServer) AddStageToPipeline(context.Context, *AddStageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStageToPipeline not implemented")
}
func (UnimplementedPipelineOrchestratorServiceServer) DeletePipeline(context.Context, *PipelineIDRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePipeline not implemented")
}
func (UnimplementedPipelineOrchestratorServiceServer) mustEmbedUnimplementedPipelineOrchestratorServiceServer() {
}
func (UnimplementedPipelineOrchestratorServiceServer) testEmbeddedByValue() {}

// UnsafePipelineOrchestratorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PipelineOrchestratorServiceServer will
// result in compilation errors.
type UnsafePipelineOrchestratorServiceServer interface {
	mustEmbedUnimplementedPipelineOrchestratorServiceServer()
}

func RegisterPipelineOrchestratorServiceServer(s grpc.ServiceRegistrar, srv PipelineOrchestratorServiceServer) {
	// If the following call pancis, it indicates UnimplementedPipelineOrchestratorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PipelineOrchestratorService_ServiceDesc, srv)
}

func _PipelineOrchestratorService_ExecutePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecutePipelineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineOrchestratorServiceServer).ExecutePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineOrchestratorService_ExecutePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineOrchestratorServiceServer).ExecutePipeline(ctx, req.(*ExecutePipelineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineOrchestratorService_GetPipelineStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineOrchestratorServiceServer).GetPipelineStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineOrchestratorService_GetPipelineStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineOrchestratorServiceServer).GetPipelineStatus(ctx, req.(*PipelineIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineOrchestratorService_CancelPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineOrchestratorServiceServer).CancelPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineOrchestratorService_CancelPipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineOrchestratorServiceServer).CancelPipeline(ctx, req.(*PipelineIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineOrchestratorService_AddStageToPipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineOrchestratorServiceServer).AddStageToPipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineOrchestratorService_AddStageToPipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineOrchestratorServiceServer).AddStageToPipeline(ctx, req.(*AddStageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PipelineOrchestratorService_DeletePipeline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PipelineOrchestratorServiceServer).DeletePipeline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PipelineOrchestratorService_DeletePipeline_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PipelineOrchestratorServiceServer).DeletePipeline(ctx, req.(*PipelineIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PipelineOrchestratorService_ServiceDesc is the grpc.ServiceDesc for PipelineOrchestratorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PipelineOrchestratorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pipeline.PipelineOrchestratorService",
	HandlerType: (*PipelineOrchestratorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecutePipeline",
			Handler:    _PipelineOrchestratorService_ExecutePipeline_Handler,
		},
		{
			MethodName: "GetPipelineStatus",
			Handler:    _PipelineOrchestratorService_GetPipelineStatus_Handler,
		},
		{
			MethodName: "CancelPipeline",
			Handler:    _PipelineOrchestratorService_CancelPipeline_Handler,
		},
		{
			MethodName: "AddStageToPipeline",
			Handler:    _PipelineOrchestratorService_AddStageToPipeline_Handler,
		},
		{
			MethodName: "DeletePipeline",
			Handler:    _PipelineOrchestratorService_DeletePipeline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pipeline.proto",
}
