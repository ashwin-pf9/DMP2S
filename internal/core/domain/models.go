package domain

/* Data Representation */

import (
	"time"

	"github.com/google/uuid"
)

// // User represents a user in the system
// type User struct {
// 	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
// 	Username string    `gorm:"unique;not null" json:"username"`
// 	Password string    `json:"-"` // Exclude from JSON responses
// }

// type Pipeline struct {
// 	ID     uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
// 	Name   string    `gorm:"type:text;not null" json:"name"`
// 	UserID uuid.UUID `gorm:"type:uuid;not null;constraint:foreignKey:UserID;references:users(ID)" json:"user_id"` //foreign key for connecting pipeline with user
// 	//Not sure about how this Multi-Valued attribute will be translated: [This should be removed and a foreign key should be added in the Stages table referencing to Pipelines tables primary key]
// 	Stages []Stage `json:"stages"`
// }

// Pipeline represents a user's pipeline
type Pipeline struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name   string    `gorm:"not null" json:"name"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	// Stages []Stage   `json:"stages"` // Relationship
}

// type Stage struct {
// 	ID         uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
// 	Name       string    `gorm:"type:text" json:"name"`
// 	PipelineID uuid.UUID `gorm:"type:uuid;not null;constraint:foreignKey:PipelineID:pipelines(ID)" json:"pipeline_id"` //Foreign key for connecting user with stage
// }

// Stage represents a stage in a pipeline
type Stage struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name       string    `gorm:"not null" json:"name"`
	PipelineID uuid.UUID `gorm:"type:uuid;not null" json:"pipeline_id"`
	Pipeline   Pipeline  `gorm:"foreignKey:PipelineID" json:"-"`
}

// PipelineExecution tracks a pipeline run
type PipelineExecution struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	PipelineID uuid.UUID `gorm:"type:uuid;not null" json:"pipeline_id"`
	Pipeline   Pipeline  `gorm:"foreignKey:PipelineID" json:"-"`
	Status     string    `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	StartedAt  time.Time `gorm:"default:now()" json:"started_at"`
	EndedAt    time.Time `json:"ended_at,omitempty"`
}

// StageExecution tracks the execution of a stage
type StageExecution struct {
	ID           uuid.UUID         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StageID      uuid.UUID         `gorm:"type:uuid;not null" json:"stage_id"`
	ExecutionID  uuid.UUID         `gorm:"type:uuid;not null" json:"execution_id"`
	Stage        Stage             `gorm:"foreignKey:StageID" json:"-"`
	Execution    PipelineExecution `gorm:"foreignKey:ExecutionID" json:"-"`
	Status       string            `gorm:"type:varchar(50);not null;default:'pending'" json:"status"`
	StartedAt    time.Time         `gorm:"default:now()" json:"started_at"`
	EndedAt      *time.Time        `json:"ended_at,omitempty"`
	ErrorMessage string            `json:"error_message,omitempty"`
}

// Status Type for telling the status of pipeline //
type Status string

const (
	Pending Status = "Pending"
	Running Status = "Running"
	Failed  Status = "Failed"
	Success Status = "Success"

	Unknown Status = "Unknown" //Error occured before fetching status
)
