package service

import (
	"context"
	crudpipelinepb "crudpipelineservice/proto"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PipelineServer struct {
	crudpipelinepb.UnimplementedPipelineServiceServer
	pipelineService *PipelineService
}

func NewPipelineServer(pipelineImpl *PipelineService) *PipelineServer {
	return &PipelineServer{pipelineService: pipelineImpl}
}

func (s *PipelineServer) CreatePipeline(ctx context.Context, req *crudpipelinepb.CreatePipelineRequest) (*crudpipelinepb.PipelineResponse, error) {
	userID := req.GetUserId()
	name := req.GetName()

	pipeline, err := s.pipelineService.CreatePipeline(userID, name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create pipeline: %v", err)
	}

	resp := &crudpipelinepb.PipelineResponse{
		Pipeline: &crudpipelinepb.Pipeline{
			Id:     pipeline.ID.String(),
			UserId: pipeline.UserID.String(),
			Name:   pipeline.Name,
		},
	}
	return resp, nil
}

func (s *PipelineServer) GetUserPipelines(ctx context.Context, req *crudpipelinepb.GetUserPipelinesRequest) (*crudpipelinepb.PipelinesResponse, error) {
	userID := req.GetUserId()
	pipelines := s.pipelineService.GetUsersPipelines(userID)

	var protoPipelines []*crudpipelinepb.Pipeline
	for _, p := range pipelines {
		protoPipelines = append(protoPipelines, &crudpipelinepb.Pipeline{
			Id:     p.ID.String(),
			UserId: p.UserID.String(),
			Name:   p.Name,
		})
	}

	return &crudpipelinepb.PipelinesResponse{
		Pipelines: protoPipelines,
	}, nil
}

func (s *PipelineServer) GetPipelineStages(ctx context.Context, req *crudpipelinepb.GetPipelineStagesRequest) (*crudpipelinepb.StagesResponse, error) {
	pipelineID, err := uuid.Parse(req.GetPipelineId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid pipeline_id: %v", err)
	}

	log.Printf("crud_bridge - GetPipelineStages called")

	stages := s.pipelineService.GetPipelineStages(pipelineID)

	var protoStages []*crudpipelinepb.Stage
	for _, s := range stages {
		protoStages = append(protoStages, &crudpipelinepb.Stage{
			Id:         s.ID.String(),
			PipelineId: s.PipelineID.String(),
			Name:       s.Name,
		})
	}

	return &crudpipelinepb.StagesResponse{
		Stages: protoStages,
	}, nil
}
