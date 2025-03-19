package service

import (
	"log"

	"github.com/ashwin-pf9/shared/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PipelineService struct {
	DB *gorm.DB
}

// NewAuthService returns a new instance of AuthService.
func NewPipelineService(DB *gorm.DB) *PipelineService {
	return &PipelineService{
		DB: DB,
	}
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

	// DB := db.InitDatabase()

	// Insert into database
	result := s.DB.Create(&pipeline)
	if result.Error != nil {
		return nil, result.Error
	}

	return pipeline, nil
}

func (s *PipelineService) GetUsersPipelines(userID string) ([]domain.Pipeline, error) {

	// DB := db.InitDatabase()

	var pipelines []domain.Pipeline
	result := s.DB.Where("user_id = ?", userID).Find(&pipelines)
	if result.Error != nil {
		return nil, result.Error
	}
	return pipelines, nil
}

func (s *PipelineService) GetPipelineStages(pipelineID uuid.UUID) []domain.Stage {
	log.Printf("pipeline_service - GetPipelineStages called")

	// DB := db.InitDatabase()

	var stages []domain.Stage
	s.DB.Where("pipeline_id = ?", pipelineID).Find(&stages)
	return stages
}
