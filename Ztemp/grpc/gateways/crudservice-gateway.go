package gateways

import (
	"context"
	clipb "grpc/proto"
)

type CRUDServiceGateway struct {
	clipb.UnimplementedCRUDPipelineServiceServer
	CRUDClient clipb.CRUDPipelineServiceClient
}

func (a *CRUDServiceGateway) CreatePipeline(ctx context.Context, req *clipb.CreatePipelineRequest) (*clipb.CreatePipelineResponse, error) {
	//Forwarding request to actual microservice
	return a.CRUDClient.CreatePipeline(ctx, req)
}

func (a *CRUDServiceGateway) GetUserPipelines(ctx context.Context, req *clipb.GetPipelinesRequest) (*clipb.GetPipelinesResponse, error) {
	return a.CRUDClient.GetUserPipelines(ctx, req)

}

func (a *CRUDServiceGateway) GetPipelineStages(ctx context.Context, req *clipb.GetPipelineStagesRequest) (*clipb.GetPipelineStagesResponse, error) {
	return a.CRUDClient.GetPipelineStages(ctx, req)
}
