package services

import (
	"DMP2S/internal/core/domain"
	"context"
	"errors"

	"github.com/google/uuid"
)

// CONCRETE IMPLEMENTATION of the PipelineOrchestratorInterface
type PipelineOrchestrator struct {
	ID     uuid.UUID
	Name   string
	Stages []domain.Stage
}

// AddStage - Adds a new stage to the pipeline
func (p *PipelineOrchestrator) AddStage(stage domain.Stage) error {
	p.Stages = append(p.Stages, stage)
	return nil
}

// Execute - Runs a given pipeline
func (p *PipelineOrchestrator) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	if len(p.Stages) == 0 {
		return nil, errors.New("no stages in pipeline")
	}
	// Simulating execution
	return "Execution completed successfully", nil
}

// GetStatus - Returns the execution status of a pipeline
func (p *PipelineOrchestrator) GetStatus(pipelineID uuid.UUID) (domain.Status, error) {
	if p.ID == pipelineID {
		return domain.Running, nil
	}
	return domain.Unknown, errors.New("pipeline not found")
}

// Cancel - Cancels a pipeline execution
func (p *PipelineOrchestrator) Cancel(pipelineID uuid.UUID) error {
	if p.ID != pipelineID {
		return errors.New("pipeline not found")
	}
	// Simulating cancellation
	return nil
}
