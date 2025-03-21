package main

import (
	"log"
	"net"

	"pipelineservice/events"
	pipelinepb "pipelineservice/proto"
	"pipelineservice/service"
	"pipelineservice/stagepb"

	"github.com/ashwin-pf9/shared/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Printf("pipeline server started\n")
	events.InitNATS() // For publishing

	// Connect to StageService running at stage-service:50052
	conn, err := grpc.Dial("stageservice-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to stage service: %v", err)
	}
	defer conn.Close()

	stageClient := stagepb.NewStageServiceClient(conn)

	// Create orchestrator implementation (if needed, pass stageClient here)
	orchestrator := service.NewPipelineOrchestratorImpl(stageClient)

	// Create Pipeline Orchestrator Service
	pipelineOrchestratorService := service.NewPipelineOrchestratorService(orchestrator, stageClient, db.InitDatabase())

	// Create handler that implements pipelinepb.PipelineServiceServer
	// pipelineHandler := handler.NewPipelineHandler(*pipelineOrchestratorService)

	// Set up and start gRPC server for PipelineService
	listener, err := net.Listen("tcp", ":50053") // PipelineService runs on port 50053
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pipelinepb.RegisterPipelineOrchestratorServiceServer(grpcServer, pipelineOrchestratorService)

	log.Println("PipelineService gRPC server running on port 50053")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
