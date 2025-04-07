package gateways

import (
	"context"
	clipb "grpc/proto"
)

type OrchestratorServiceGateway struct {
	clipb.UnimplementedPipelineOrchestratorServiceServer
	OrchestratorClient clipb.PipelineOrchestratorServiceClient
}

func (a *OrchestratorServiceGateway) Execute(ctx context.Context, req *clipb.ExecutePipelineRequest) (*clipb.ExecutePipelineResponse, error) {
	return a.OrchestratorClient.ExecutePipeline(ctx, req)
}
