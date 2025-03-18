package service

import (
	"context"
	"errors"
	"log"

	"time"

	"github.com/ashwin-pf9/DMP2S/services/pipelineservice/events"
	"github.com/ashwin-pf9/DMP2S/services/pipelineservice/stagepb"
	"github.com/ashwin-pf9/shared/domain"
	"github.com/google/uuid"
)

type PipelineOrchestratorImpl struct {
	stageClient stagepb.StageServiceClient
}

func NewPipelineOrchestratorImpl(client stagepb.StageServiceClient) *PipelineOrchestratorImpl {
	return &PipelineOrchestratorImpl{stageClient: client}
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
	//   CTX
	/*-- ctx parameter contains "executionID" --*/

	// Type assertion to convert interface{}
	// Type assertion for domain.Stage slice
	stageList, ok := input.([]domain.Stage)
	if !ok {
		return nil, errors.New("failed to assert input to []domain.Stage")
	}

	// Iterate through each domain.Stage and wrap it into StageOrchestratorService
	for _, stage := range stageList {
		// // Create StageOrchestratorService for each stage
		// stageImpl := NewStageOrchestratorImpl()                 // Actual implementation
		// stageService := NewStageOrchestratorService(&stageImpl) // Service layer

		// Call methods on the service layer
		log.Printf("Executing stage: %s", stage.Name)
		// Broadcast that the stage has started
		events.SendUpdate(stage.ID.String(), "Running")

		executionIDValue := ctx.Value("executionID")

		executionID, ok := executionIDValue.(uuid.UUID)
		if !ok {
			return nil, errors.New("executionID not found or invalid type in context")
		}

		req := &stagepb.ExecuteStageRequest{
			Stage: &stagepb.Stage{
				Id:         stage.ID.String(),
				Name:       stage.Name,
				PipelineId: stage.PipelineID.String(),
			},
			ExecutionId: executionID.String(), // Now it's valid
		}

		// Call gRPC method on stage service
		res, err := imp.stageClient.ExecuteStage(ctx, req)
		if err != nil {
			log.Printf("Stage %s failed: %v", stage.Name, err)
			events.SendUpdate(stage.ID.String(), "Failed")
			return nil, err
		}

		log.Printf("Stage %s succeeded: %v\n", stage.Name, res)
		// Broadcast success status
		events.SendUpdate(stage.ID.String(), "Completed")

		//pausing for a second
		time.Sleep(2 * time.Second)
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
