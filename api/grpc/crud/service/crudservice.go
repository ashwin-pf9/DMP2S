package service

import (
	"context"
	"fmt"
	grpccrudpb "grpc/crud/proto"
	// grpccrudpb "grpc/crud/proto"
)

type CrudServiceServer struct {
	grpccrudpb.UnimplementedPipelineServiceServer
	CrudClient grpccrudpb.PipelineServiceClient
}

func (s *CrudServiceServer) CreatePipeline(ctx context.Context, in *grpccrudpb.CreatePipelineRequest) (*grpccrudpb.CreatePipelineResponse, error) {
	fmt.Println("Create pipelines function called - crudservice.go")
	return s.CrudClient.CreatePipeline(ctx, in)
}
func (s *CrudServiceServer) GetUserPipelines(ctx context.Context, in *grpccrudpb.GetPipelinesRequest) (*grpccrudpb.GetPipelinesResponse, error) {
	fmt.Println("Get user pipelines function Locally running - crudservice.go")
	return s.CrudClient.GetUserPipelines(ctx, in)
}
func (s *CrudServiceServer) GetPipelineStages(ctx context.Context, in *grpccrudpb.GetPipelineStagesRequest) (*grpccrudpb.GetPipelineStagesResponse, error) {
	fmt.Println("Get pipeline stages function called - crudservice.go")
	return s.CrudClient.GetPipelineStages(ctx, in)
}
