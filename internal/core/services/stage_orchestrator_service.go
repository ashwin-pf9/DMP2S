package services

import (
	"DMP2S/internal/core/domain"
	"DMP2S/internal/core/ports"
	"DMP2S/internal/infrastructure/db"
	"context"
	"errors"
	"time"

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
	//Extracting executionID from context
	executionID, ok := ctx.Value("executionID").(uuid.UUID)
	if !ok {
		return nil, errors.New("executionID not found in context")
	}

	// Simulated execution logic
	stage := input.(domain.Stage)
	startTime := time.Now()

	stageExecution := domain.StageExecution{
		ID:           uuid.New(),
		StageID:      stage.ID,
		ExecutionID:  executionID,
		Status:       string(domain.Running),
		StartedAt:    startTime,
		EndedAt:      &time.Time{},
		ErrorMessage: "no error",
	}

	_, err := b.stage.Execute(ctx, input)
	if err != nil {
		endTime := time.Now()
		stageExecution.Status = string(domain.Failed)
		stageExecution.ErrorMessage = err.Error()
		stageExecution.EndedAt = &endTime

		db.DB.Save(&stageExecution)
		return nil, err
	}

	// Mark execution as completed
	stageExecution.Status = string(domain.Completed)
	endTime := time.Now()
	stageExecution.EndedAt = &endTime

	db.DB.Save(&stageExecution)

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
