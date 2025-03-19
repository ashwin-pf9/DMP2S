package service

import (
	"context"
	"errors"
	"log"

	"time"

	pipelinepb "pipelineservice/proto"

	"pipelineservice/stagepb"

	"github.com/ashwin-pf9/shared/domain"
	"github.com/ashwin-pf9/shared/ports"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
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
	orchestrator                                              ports.PipelineOrchestratorInterface
	stageClient                                               stagepb.StageServiceClient // gRPC client interface
	DB                                                        *gorm.DB
	pipelinepb.UnimplementedPipelineOrchestratorServiceServer // embed to satisfy interface
}

// Actual contract instance creation
// CREATE NEW APPLICATION
func NewPipelineOrchestratorService(orchestrator ports.PipelineOrchestratorInterface, stageClient stagepb.StageServiceClient, DB *gorm.DB) *PipelineOrchestratorService {
	//RETURNING INSTANCE OF APPLICATION
	return &PipelineOrchestratorService{
		orchestrator: orchestrator,
		stageClient:  stageClient,
		DB:           DB,
	}
}

/*
	ExecutePipeline(context.Context, *ExecutePipelineRequest) (*ExecutionResponse, error)
	GetPipelineStatus(context.Context, *PipelineIDRequest) (*StatusResponse, error)
	CancelPipeline(context.Context, *PipelineIDRequest) (*emptypb.Empty, error)
	AddStageToPipeline(context.Context, *AddStageRequest) (*emptypb.Empty, error)
	DeletePipeline(context.Context, *PipelineIDRequest) (*emptypb.Empty, error)
*/

// Inorder to access interface methods i need to create methods which will belong to struct PipelineOrchestratorService
// AddStageToPipeline RPC
func (s *PipelineOrchestratorService) AddStageToPipeline(ctx context.Context, req *pipelinepb.AddStageRequest) (*emptypb.Empty, error) {
	// DB := db.InitDatabase()
	// Validate input
	if req.GetStage().GetName() == "" {
		return nil, errors.New("stage name cannot be empty")
	}

	// Check if pipeline exists
	var pipeline domain.Pipeline
	if err := s.DB.First(&pipeline, "id = ?", req.GetStage().GetPipelineId()).Error; err != nil {
		return nil, errors.New("pipeline not found")
	}

	// Convert protobuf Stage to domain Stage
	stage := domain.Stage{
		ID:         uuid.New(),
		PipelineID: uuid.MustParse(req.GetStage().GetPipelineId()),
		Name:       req.GetStage().GetName(),
		// Add other necessary fields here
	}

	// Call orchestrator to add stage
	err := s.orchestrator.AddStage(stage)
	if err != nil {
		return nil, errors.New("failed to add stage to pipeline")
	}

	// Save stage to DB
	if err := s.DB.Create(&stage).Error; err != nil {
		return nil, errors.New("failed to save stage to database")
	}

	return &emptypb.Empty{}, nil
}

// ExecutePipeline(context.Context, *ExecutePipelineRequest) (*ExecutionResponse, error)
func (s *PipelineOrchestratorService) ExecutePipeline(ctx context.Context, req *pipelinepb.ExecutePipelineRequest) (*pipelinepb.ExecutionResponse, error) {
	// DB := db.InitDatabase()
	log.Printf("pipeline_orch_service - ExecutePipeline called")
	pipelineID, err := uuid.Parse(req.GetPipelineId())
	if err != nil {
		return nil, errors.New("invalid pipeline ID format")
	}

	// Create a new pipeline execution entry
	executionID := uuid.New()
	startTime := time.Now()

	execution := domain.PipelineExecution{
		ID:         executionID,
		PipelineID: pipelineID,
		Status:     string(domain.Running),
		StartedAt:  startTime,
	}
	log.Printf("Execution of Pipeline started New: %s", executionID)

	// Save execution start to DB
	if err := s.DB.Create(&execution).Error; err != nil {
		return nil, errors.New("failed to start pipeline execution")
	}

	// Fetch stages
	var domainStages []domain.Stage
	if err := s.DB.Where("pipeline_id = ?", pipelineID).Find(&domainStages).Error; err != nil || len(domainStages) < 1 {
		// Mark execution as failed
		execution.Status = string(domain.Failed)
		s.DB.Save(&execution)
		return nil, errors.New("no stages found for pipeline")
	}

	// Add execution ID to context
	ctx = context.WithValue(ctx, "executionID", executionID)

	// Call orchestrator to execute
	_, err = s.orchestrator.Execute(ctx, domainStages)
	if err != nil {
		execution.Status = string(domain.Failed)
		s.DB.Save(&execution)
		return nil, errors.New("pipeline execution failed")
	}

	// Mark as successful
	execution.Status = string(domain.Success)
	execution.EndedAt = time.Now()
	s.DB.Save(&execution)

	log.Printf("Pipeline Execution completed: %s", executionID)

	// Return gRPC response
	return &pipelinepb.ExecutionResponse{
		ExecutionId: executionID.String(),
	}, nil
}

// DeletePipeline(context.Context, *PipelineIDRequest) (*emptypb.Empty, error)
func (s *PipelineOrchestratorService) DeletePipeline(ctx context.Context, req *pipelinepb.PipelineIDRequest) (*emptypb.Empty, error) {
	// DB := db.InitDatabase()
	// Parse pipeline ID from request
	pipelineID, err := uuid.Parse(req.GetPipelineId())
	if err != nil {
		log.Println("Invalid pipeline ID:", err)
		return nil, errors.New("invalid pipeline ID format")
	}

	// Delete pipeline
	if err := s.DB.Where("id = ?", pipelineID).Delete(&domain.Pipeline{}).Error; err != nil {
		log.Println("Error deleting pipeline:", err)
		return nil, errors.New("failed to delete pipeline")
	}

	log.Printf("Pipeline %s deleted successfully", pipelineID)
	return &emptypb.Empty{}, nil
}

// GetPipelineStatus(context.Context, *PipelineIDRequest) (*StatusResponse, error)
func (s *PipelineOrchestratorService) GetPipelineStatus(ctx context.Context, req *pipelinepb.PipelineIDRequest) (*pipelinepb.StatusResponse, error) {
	return &pipelinepb.StatusResponse{}, nil
}

// CancelPipeline(context.Context, *PipelineIDRequest) (*emptypb.Empty, error)
// CancelPipeline cancels the pipeline execution
func (s *PipelineOrchestratorService) CancelPipeline(ctx context.Context, req *pipelinepb.PipelineIDRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
