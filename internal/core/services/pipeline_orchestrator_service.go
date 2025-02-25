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

	// Save execution start record in DB
	if err := db.DB.Create(&execution).Error; err != nil {
		return nil, errors.New("failed to start pipeline execution")
	}

	log.Printf("Pipeline Execution started: %s", executionID)

	// Fetch all stages from the database
	var domainStages []domain.Stage
	if err := db.DB.Where("pipeline_id = ?", pipelineID).Order("\"order\" ASC").Find(&domainStages).Error; err != nil {
		return "", errors.New("no stages found")
	}

	// Convert domainStages to interface{} (a slice can be passed as an interface)
	stagesInterface := interface{}(domainStages)

	//CALL TO ACTUAL IMPLEMENTATION
	_, err := s.orchestrator.Execute(ctx, stagesInterface)
	if err != nil {
		execution.Status = string(domain.Failed)
		db.DB.Save(execution)
		return nil, err
	}

	// Mark execution as complete
	execution.Status = string(domain.Success)
	execution.EndedAt = &startTime
	db.DB.Save(&execution)

	log.Printf("Pipeline Execution completed: %s", executionID)
	return executionID, nil
}

// // ExecutePipeline executes the pipeline
// func (s *PipelineOrchestratorService) Execute(ctx context.Context, pipelineID uuid.UUID) (interface{}, error) {

// 	// Start a new pipeline execution entry
// 	executionID := uuid.New()
// 	startTime := time.Now()

// 	execution := domain.PipelineExecution{
// 		ID:         executionID,
// 		PipelineID: pipelineID,
// 		Status:     string(domain.Running),
// 		StartedAt:  startTime,
// 	}

// 	// Save execution start record in DB
// 	if err := db.DB.Create(&execution).Error; err != nil {
// 		return nil, errors.New("failed to start pipeline execution")
// 	}

// 	log.Printf("Pipeline Execution started: %s", executionID)

// 	// Fetch all stages in order
// 	var stages []services.StageOrchestratorService
// 	if err := db.DB.Where("pipeline_id = ?", pipelineID).Order("\"order\" ASC").Find(&stages).Error; err != nil {
// 		return "", errors.New("no stages found")
// 	}

// 	//Execute each stage sequentially
// 	for _, stage := range stages {
// 		log.Printf("Executing stage: %s", stage.GetID())

// 		// Execute stage logic (replace with actual logic)
// 		result, err := stage.Execute(ctx, nil)
// 	}

// 	return "", nil
// }

// GetPipelineStatus retrieves the execution status of the pipeline
func (s *PipelineOrchestratorService) GetStatus(pipelineID uuid.UUID) (domain.Status, error) {
	return domain.Unknown, nil
}

// CancelPipeline cancels the pipeline execution
func (s *PipelineOrchestratorService) Cancel(pipelineID uuid.UUID) error {
	return nil
}
