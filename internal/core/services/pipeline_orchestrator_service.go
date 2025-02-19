package services

import (
	"DMP2S/internal/core/domain"
	"DMP2S/internal/core/ports"
	"context"
	"errors"

	"github.com/google/uuid"
)

//Creating A structure with PipelineOrchestratorInterface object, for connecting service to Actual Implementation
//through PipelineOrchestratorInterface interface

type PipelineOrchestratorService struct {
	orchestrator ports.PipelineOrchestratorInterface
}

// Constructor Method : Purpose is to create object of type PipelineOrchestratorService and return its pointer
func NewPipelineOrchestratorService(orchestrator ports.PipelineOrchestratorInterface) *PipelineOrchestratorService {
	return &PipelineOrchestratorService{orchestrator: orchestrator}
}

//There are 4 methods in interface which i want to access from this service via interface
//1. AddState
//2. Execute
//3. GetStatus
//4. Cancel

// Inorder to access interface methods i need to create methods which will belong to struct PipelineOrchestratorService
func (s *PipelineOrchestratorService) AddStageToPipeline(stage domain.Stage) error {
	if s.orchestrator == nil {
		return errors.New("orchestrator is not initialized")
	}

	return s.orchestrator.AddStage(stage)
}

// ExecutePipeline executes the pipeline
func (s *PipelineOrchestratorService) ExecutePipeline(ctx context.Context, input interface{}) (interface{}, error) {
	if s.orchestrator == nil {
		return nil, errors.New("orchestrator is not initialized")
	}
	return s.orchestrator.Execute(ctx, input)
}

// GetPipelineStatus retrieves the execution status of the pipeline
func (s *PipelineOrchestratorService) GetPipelineStatus(pipelineID uuid.UUID) (domain.Status, error) {
	if s.orchestrator == nil {
		return domain.Unknown, errors.New("orchestrator is not initialized")
	}
	return s.orchestrator.GetStatus(pipelineID)
}

// CancelPipeline cancels the pipeline execution
func (s *PipelineOrchestratorService) CancelPipeline(pipelineID uuid.UUID) error {
	if s.orchestrator == nil {
		return errors.New("orchestrator is not initialized")
	}
	return s.orchestrator.Cancel(pipelineID)
}
