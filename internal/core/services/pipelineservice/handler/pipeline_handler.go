package handler

// import (
// )

// type PipelineHandler struct {
// 	pipelinepb.UnimplementedPipelineServiceServer // embed to satisfy interface
// 	orchestratorService                           service.PipelineOrchestratorService
// }

// // Constructor
// func NewPipelineOrchestratorHandler(service service.PipelineOrchestratorService) *PipelineHandler {
// 	return &PipelineHandler{
// 		orchestratorService: service,
// 	}
// }

// // Example gRPC method implementation
// func (h *PipelineHandler) ExecutePipeline(ctx context.Context, req *pipelinepb.ExecutePipelineRequest) (*pipelinepb.ExecutionResponse, error) {
// 	log.Panicf("ExecutePipeline in handler called") // -- for debugging
// 	pipelineID, err := uuid.Parse(req.PipelineId)
// 	if err != nil {
// 		return nil, status.Errorf(codes.InvalidArgument, "Invalid Pipeline ID")
// 	}

// 	executionID, err := h.orchestratorService.ExecutePipeline(ctx, pipelineID)
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "Pipeline execution failed: %v", err)
// 	}

// 	return &pipelinepb.ExecutionResponse{
// 		ExecutionId: executionID.(uuid.UUID).String(),
// 	}, nil
// }
