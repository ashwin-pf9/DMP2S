package services

import (
	"DMP2S/internal/core/ports"
	"context"

	"github.com/google/uuid"
)

// BuildStage implements Stage interface
type StageOrchestratorService struct {
	stage ports.Stage
}

func NewStageOrchestratorService(somename ports.Stage) *StageOrchestratorService {
	return &StageOrchestratorService{
		stage: somename,
	}
}

// GetID returns the stage ID
func (b *StageOrchestratorService) GetStageID() uuid.UUID {
	return b.stage.GetID()
}

// Execute runs the build stage logic
func (b *StageOrchestratorService) ExecuteStage(ctx context.Context, input interface{}) (interface{}, error) {
	// Simulated execution logic
	b.stage.Execute(ctx, "some object for single stage execution")
	return "Build completed", nil
}

// HandleError handles any errors during execution
func (b *StageOrchestratorService) HandleError(ctx context.Context, err error) error {
	b.stage.HandleError(ctx, err)
	return err
}

// Rollback rolls back the stage if it fails
func (b *StageOrchestratorService) Rollback(ctx context.Context, input interface{}) error {
	b.stage.Rollback(ctx, "some object")
	return nil
}
