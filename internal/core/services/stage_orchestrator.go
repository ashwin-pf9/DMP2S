package services

import (
	"context"

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
	return "", nil //TEMPORARY
}

func (s *StageOrchestratorImpl) HandleError(ctx context.Context, err error) error {
	//code for handling error in a stage
	return nil
}

func (s *StageOrchestratorImpl) Rollback(ctx context.Context, input interface{}) error {
	return nil
}
