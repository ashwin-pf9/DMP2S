package service

import (
	"log"

	"github.com/ashwin-pf9/shared/db"
	"github.com/ashwin-pf9/shared/domain"
	"github.com/google/uuid"
)

type PipelineService struct{}

// NewAuthService returns a new instance of AuthService.
func NewPipelineService() *PipelineService {
	return &PipelineService{}
}

func (s *PipelineService) CreatePipeline(userID string, name string) (*domain.Pipeline, error) {
	if name == "" {
		name = "Untitled Pipeline"
	}

	pipeline := &domain.Pipeline{
		ID:     uuid.New(), // Generate UUID
		UserID: uuid.MustParse(userID),
		Name:   name,
	}

	DB := db.InitDatabase()

	// Insert into database
	result := DB.Create(&pipeline)
	if result.Error != nil {
		return nil, result.Error
	}

	return pipeline, nil
}

func (s *PipelineService) GetUsersPipelines(userID string) []domain.Pipeline {

	DB := db.InitDatabase()

	var pipelines []domain.Pipeline
	//db.DB - Using this global DB instance to communicate with database
	DB.Where("user_id = ?", userID).Find(&pipelines)
	return pipelines
}

func (s *PipelineService) GetPipelineStages(pipelineID uuid.UUID) []domain.Stage {
	log.Printf("pipeline_service - GetPipelineStages called")
	DB := db.InitDatabase()

	var stages []domain.Stage
	DB.Where("pipeline_id = ?", pipelineID).Find(&stages)
	return stages
}
