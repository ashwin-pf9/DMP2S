package services

import (
	"DMP2S/internal/core/domain"
	"DMP2S/internal/core/ports"
	"DMP2S/internal/infrastructure/db"
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

/*
	orchestrator
	-------------
	  interface
	-------------
	implementation
*/

// There is a PipelineOrchestratorInterface contract
// We've to create instance of that contract for each Pipeline

// APPLICATION STRUCT
type PipelineOrchestratorService struct { /*contract instance*/
	orchestrator ports.PipelineOrchestratorInterface
}

// Actual contract instance creation
// CREATE NEW APPLICATION
func NewPipelineOrchestratorService(orchestrator ports.PipelineOrchestratorInterface) *PipelineOrchestratorService {
	//RETURNING INSTANCE OF APPLICATION
	return &PipelineOrchestratorService{
		orchestrator: orchestrator,
	}
}

// Inorder to access interface methods i need to create methods which will belong to struct PipelineOrchestratorService
func (s *PipelineOrchestratorService) AddStageToPipeline(stage domain.Stage) error {
	if stage.Name == "" {
		return errors.New("stage name cannot be empty")
	}

	// Check if pipeline exists
	var pipeline domain.Pipeline
	if err := db.DB.First(&pipeline, "id = ?", stage.PipelineID).Error; err != nil {
		return errors.New("pipeline not found")
	}

	//Call to actual implementation of AddStage(Stage)
	err := s.orchestrator.AddStage(stage)
	if err != nil {
		return errors.New("failed to add stage to pipeline")
	}

	// Add stage to the database
	return db.DB.Create(&stage).Error
}

func (s *PipelineOrchestratorService) ExecutePipeline(ctx context.Context, pipelineID uuid.UUID) (interface{}, error) {
	// Start a new pipeline execution entry
	executionID := uuid.New()
	startTime := time.Now()

	execution := domain.PipelineExecution{
		ID:         executionID,
		PipelineID: pipelineID,
		Status:     string(domain.Running),
		StartedAt:  startTime,
	}

	// Fetch all stages from the database
	var domainStages []domain.Stage
	if err := db.DB.Where("pipeline_id = ?", pipelineID).Find(&domainStages).Error; err != nil || len(domainStages) < 1 {
		return "", errors.New("no stages found")
	}

	//If there is no error in fetching stages and there are more than 0 stages then - Save execution start record in DB
	if err := db.DB.Create(&execution).Error; err != nil {
		return nil, errors.New("failed to start pipeline execution")
	}

	log.Printf("Execution of Pipeline started: %s", executionID)

	//adding "executionID" in the context
	ctx = context.WithValue(ctx, "executionID", executionID)

	// Convert domainStages to interface{} (a slice can be passed as an interface)
	stagesInterface := interface{}(domainStages)

	//CALL TO ACTUAL IMPLEMENTATION
	_, err := s.orchestrator.Execute(ctx, stagesInterface)
	if err != nil {
		execution.Status = string(domain.Failed)
		db.DB.Save(&execution)
		return nil, err
	}

	// Mark execution as complete
	execution.Status = string(domain.Success)
	execution.EndedAt = time.Now()
	db.DB.Save(&execution)

	log.Printf("Pipeline Execution completed: %s", executionID)
	return executionID, nil
}

func (s *PipelineOrchestratorService) DeletePipeline(pipelineID uuid.UUID) error {

	err := db.DB.Where("id = ?", pipelineID).Delete(&domain.Pipeline{}).Error
	if err != nil {
		log.Println("Error deleting pipeline:", err)
		return err
	}

	log.Println("Pipeline deleted successfully")
	return nil
}

func (s *PipelineOrchestratorService) GetStatus(pipelineID uuid.UUID) (domain.Status, error) {
	return domain.Unknown, nil
}

// CancelPipeline cancels the pipeline execution
func (s *PipelineOrchestratorService) Cancel(pipelineID uuid.UUID) error {
	return nil
}
