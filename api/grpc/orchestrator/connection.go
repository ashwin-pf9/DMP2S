package orchestrator

import (
	pipelinepb "grpc/orchestrator/proto"
	"grpc/orchestrator/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var CONN *grpc.ClientConn

func OrchConnect(url string) *service.PipelineOrchestratorServiceServer {
	CONN, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect %v", err)
	}

	orchClient := pipelinepb.NewPipelineOrchestratorServiceClient(CONN)

	return &service.PipelineOrchestratorServiceServer{
		OrchClient: orchClient,
	}

}
