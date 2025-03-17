package service

import (
	"pipelineservice/shared/db"
	"pipelineservice/shared/domain"

	"github.com/google/uuid"
)

func CreatePipeline(userID string, name string) (*domain.Pipeline, error) {
	if name == "" {
		name = "Untitled Pipeline"
	}

	pipeline := &domain.Pipeline{
		ID:     uuid.New(), // Generate UUID
		UserID: uuid.MustParse(userID),
		Name:   name,
	}

	// Insert into database
	result := db.DB.Create(&pipeline)
	if result.Error != nil {
		return nil, result.Error
	}

	return pipeline, nil
}

func GetUsersPipelines(userID string) []domain.Pipeline {
	var pipelines []domain.Pipeline
	//db.DB - Using this global DB instance to communicate with database
	db.DB.Where("user_id = ?", userID).Find(&pipelines)
	return pipelines
}

func GetPipelineStages(pipelineID uuid.UUID) []domain.Stage {
	var stages []domain.Stage
	db.DB.Where("pipeline_id = ?", pipelineID).Find(&stages)
	return stages
}
