package main

import (
	"log"
	"net"
	stagepb "stageservice/proto"
	"stageservice/services"
	"stageservice/shared/db"

	"google.golang.org/grpc"
)

// Start gRPC server
func main() {
	db.InitDatabase()
	// Listen on port 50052 (or change as needed)
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Create actual implementation of StageOrchestrator
	stageImpl := services.NewStageOrchestratorImpl()

	stageServer := services.NewStageOrchestratorService(&stageImpl)

	// Register gRPC server with implementation
	stagepb.RegisterStageServiceServer(grpcServer, stageServer)

	log.Println("StageService gRPC server started on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
