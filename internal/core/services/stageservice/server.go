package main

import (
	"log"
	"net"

	stagepb "github.com/ashwin-pf9/DMP2S/services/stageservice/proto"
	"github.com/ashwin-pf9/DMP2S/services/stageservice/services"
	"github.com/ashwin-pf9/shared/db"
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

	stageServer := services.NewStageOrchestratorService(&stageImpl, db.InitDatabase())

	// Register gRPC server with implementation
	stagepb.RegisterStageServiceServer(grpcServer, stageServer)

	log.Println("StageService gRPC server started on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
