package services

import (
	"DMP2S/internal/core/domain"
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
)

type PipelineOrchestratorImpl struct {
}

func NewPipelineOrchestratorImpl() *PipelineOrchestratorImpl {
	return &PipelineOrchestratorImpl{}
}

func (imp *PipelineOrchestratorImpl) AddStage(stage domain.Stage) error {

	if stage.Name == "" {
		return errors.New("stage name cannot be empty")
	}
	//appending Stage object to the in memory list of stages
	//p.pipeline.Stages = append(p.pipeline.Stages, stage)
	log.Printf("stage \"%s\"added successfully\n", stage.Name)
	//Have to figure out how to make this persistent by making database insert query

	return nil
}

func (imp *PipelineOrchestratorImpl) Execute(ctx context.Context, input interface{}) (interface{}, error) {

	// Type assertion to convert interface{}
	// Type assertion for domain.Stage slice
	stageList, ok := input.([]domain.Stage)
	if !ok {
		return nil, errors.New("failed to assert input to []domain.Stage")
	}

	// Iterate through each domain.Stage and wrap it into StageOrchestratorService
	for _, stage := range stageList {
		// Create StageOrchestratorService for each stage
		stageImpl := NewStageOrchestratorImpl()                 // Actual implementation
		stageService := NewStageOrchestratorService(&stageImpl) // Service layer

		// Call methods on the service layer
		log.Printf("Executing stage: %s", stage.Name)

		// Execute stage
		result, err := stageService.ExecuteStage(ctx, stage)
		if err != nil {
			log.Printf("Stage %s failed: %v", stage.Name, err)
			stageService.HandleError(ctx, err)
			return nil, err
		}

		log.Printf("Stage %s succeeded: %v\n", stage.Name, result)
	}

	return "an object", nil
}

func (imp *PipelineOrchestratorImpl) GetStatus(pipelineID uuid.UUID) (domain.Status, error) {
	//code ....

	return domain.Unknown, nil
}

func (imp *PipelineOrchestratorImpl) Cancel(pipelineID uuid.UUID) error {
	//code....

	return nil
}

// // CONCRETE IMPLEMENTATION of the PipelineOrchestratorInterface
// type PipelineOrchestrator struct {
// 	ID     uuid.UUID
// 	Name   string
// 	Stages []domain.Stage
// }

// // AddStage - Adds a new stage to the pipeline
// func (p *PipelineOrchestrator) AddStage(stage domain.Stage) error {
// 	if stage.Name == "" {
// 		return errors.New("stage name cannot be empty")
// 	}
// 	p.Stages = append(p.Stages, stage)
// 	return nil
// }

// // Execute - Runs a given pipeline
// func (p *PipelineOrchestrator) Execute(ctx context.Context, input interface{}) (interface{}, error) {
// 	if len(p.Stages) == 0 {
// 		return nil, errors.New("no stages in pipeline")
// 	}

// 	for _, stage := range p.Stages {
// 		result, err := stage.Execute(ctx, input)
// 		if err != nil {
// 			return nil, err
// 		}
// 		input = result
// 	}
// 	return input, nil
// }

// // GetStatus - Returns the execution status of a pipeline
// func (p *PipelineOrchestrator) GetStatus(pipelineID uuid.UUID) (domain.Status, error) {
// 	if p.ID == pipelineID {
// 		return domain.Running, nil
// 	}
// 	return domain.Unknown, errors.New("pipeline not found")
// }

// // Cancel - Cancels a pipeline execution
// func (p *PipelineOrchestrator) Cancel(pipelineID uuid.UUID) error {
// 	if p.ID != pipelineID {
// 		return errors.New("pipeline not found")
// 	}
// 	// Simulating cancellation
// 	return nil
// }
