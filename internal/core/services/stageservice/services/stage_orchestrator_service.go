package services

import (
	"context"
	"log"

	"time"

	stagepb "github.com/ashwin-pf9/DMP2S/services/stageservice/proto"
	"github.com/ashwin-pf9/shared/domain"
	"github.com/ashwin-pf9/shared/ports"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BuildStage implements Stage interface
type StageOrchestratorService struct {
	stagepb.UnimplementedStageServiceServer
	stage ports.Stage
	DB    *gorm.DB
}

func NewStageOrchestratorService(somename ports.Stage, DB *gorm.DB) *StageOrchestratorService {
	return &StageOrchestratorService{
		stage: somename,
		DB:    DB,
	}
}

// GetID returns the stage ID
func (s *StageOrchestratorService) GetStageID() uuid.UUID {
	return s.stage.GetID()
}

// Execute runs the build stage logic
func (s *StageOrchestratorService) ExecuteStage(ctx context.Context, req *stagepb.ExecuteStageRequest) (*stagepb.ExecuteStageResponse, error) {
	// DB := db.InitDatabase()

	log.Printf("stage_orch_service - ExecuteStage called\n")
	executionID, err := uuid.Parse(req.ExecutionId)
	log.Printf("stage_orch_service - ExecutionID : %s\n", executionID)
	if err != nil {
		return &stagepb.ExecuteStageResponse{
			Result:       "",
			ErrorMessage: "Invalid execution_id",
		}, nil
	}

	// Convert proto Stage to domain.Stage
	stageUUID, err := uuid.Parse(req.Stage.Id)
	if err != nil {
		return &stagepb.ExecuteStageResponse{
			ErrorMessage: "Invalid stage ID",
		}, nil
	}
	stage := domain.Stage{
		ID:   stageUUID,
		Name: req.Stage.Name,
	}
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

	_, err = s.stage.Execute(ctx, stage)
	if err != nil {
		endTime := time.Now()
		stageExecution.Status = string(domain.Failed)
		stageExecution.ErrorMessage = err.Error()
		stageExecution.EndedAt = &endTime

		s.DB.Save(&stageExecution)
		return &stagepb.ExecuteStageResponse{
			Result:       "",
			ErrorMessage: "Failed to execute stage",
		}, nil
	}
	// Mark execution as completed
	stageExecution.Status = string(domain.Completed)
	endTime := time.Now()
	stageExecution.EndedAt = &endTime

	s.DB.Save(&stageExecution)

	return &stagepb.ExecuteStageResponse{
		Result:       "Build completed",
		ErrorMessage: "",
	}, nil
}

// HandleError handles any errors during execution
func (s *StageOrchestratorService) HandleError(ctx context.Context, err error) error {
	s.stage.HandleError(ctx, err)
	return err
}

// Rollback rolls back the stage if it fails
func (s *StageOrchestratorService) Rollback(ctx context.Context, input interface{}) error {
	s.stage.Rollback(ctx, "some object")
	return nil
}
