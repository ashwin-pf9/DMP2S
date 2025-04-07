package service

import (
	pipelinepb "DMP2S/global-protos/orch"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Global variable to hold the gRPC client
var orchClient pipelinepb.PipelineOrchestratorServiceClient

// initGRPCClient initializes the connection to the gRPC server
func InitOrchClient(url string) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	orchClient = pipelinepb.NewPipelineOrchestratorServiceClient(conn)
}

func ExecutePipeline(pipelineID string) {
	req := &pipelinepb.ExecutePipelineRequest{
		PipelineId: pipelineID,
	}

	resp, err := orchClient.ExecutePipeline(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to Execute Pipeline %v", err)
	}

	fmt.Printf("Pipeline Execution ID : %s\n", resp.ExecutionId)
}

func DeletePipeline(pipelineID string) {
	req := &pipelinepb.PipelineIDRequest{
		PipelineId: pipelineID,
	}

	_, err := orchClient.DeletePipeline(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to delete pipeline %v\n", err)
	}

	fmt.Printf("Pipeline deleted successfully\n")
}
