package services

import (
	"DMP2S/internal/core/domain"
	"context"
	"log"
	"time"

	"github.com/google/uuid"
)

/*ACTUAL IMPLEMENTATION*/
type StageOrchestratorImpl struct {
}

func NewStageOrchestratorImpl() StageOrchestratorImpl {
	return StageOrchestratorImpl{}
}

func (s *StageOrchestratorImpl) GetID() uuid.UUID {
	//code to get stage id
	id := uuid.New() //TEMPORARY
	return id
}

func (s *StageOrchestratorImpl) Execute(ctx context.Context, input interface{}) (interface{}, error) {
	//code for executing a stage
	stage := input.(domain.Stage)

	log.Printf("execution of stage \"%s\" started", stage.Name)

	// Inside ExecutePipeline function, before executing each stage:
	time.Sleep(3 * time.Second) // Pause for 3 seconds

	log.Printf("execution of stage \"%s\" ended", stage.Name)
	return "", nil //TEMPORARY
}

func (s *StageOrchestratorImpl) HandleError(ctx context.Context, err error) error {
	//code for handling error in a stage
	return nil
}

func (s *StageOrchestratorImpl) Rollback(ctx context.Context, input interface{}) error {
	return nil
}
