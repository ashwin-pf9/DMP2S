package service

import (
	"context"
	pipelinepb "grpc/orchestrator/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

type PipelineOrchestratorServiceServer struct {
	pipelinepb.UnimplementedPipelineOrchestratorServiceServer
	OrchClient pipelinepb.PipelineOrchestratorServiceClient
}

func (s *PipelineOrchestratorServiceServer) AddStageToPipeline(ctx context.Context, in *pipelinepb.AddStageRequest) (*emptypb.Empty, error) {
	return s.OrchClient.AddStageToPipeline(ctx, in)
}
func (s *PipelineOrchestratorServiceServer) CancelPipeline(ctx context.Context, in *pipelinepb.PipelineIDRequest) (*emptypb.Empty, error) {
	return s.OrchClient.CancelPipeline(ctx, in)
}
func (s *PipelineOrchestratorServiceServer) DeletePipeline(ctx context.Context, in *pipelinepb.PipelineIDRequest) (*emptypb.Empty, error) {
	return s.OrchClient.DeletePipeline(ctx, in)
}
func (s *PipelineOrchestratorServiceServer) ExecutePipeline(ctx context.Context, in *pipelinepb.ExecutePipelineRequest) (*pipelinepb.ExecutionResponse, error) {
	return s.OrchClient.ExecutePipeline(ctx, in)
}
func (s *PipelineOrchestratorServiceServer) GetPipelineStatus(ctx context.Context, in *pipelinepb.PipelineIDRequest) (*pipelinepb.StatusResponse, error) {
	return s.OrchClient.GetPipelineStatus(ctx, in)
}
