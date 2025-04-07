package main

import (
	"grpc/gateways"
	clipb "grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Enable server reflection
	reflection.Register(grpcServer)

	/* CONNECTING TO AUTH SERVICE */
	// Dial to actual AuthService microservice (running on a different host)
	// auth_conn, err := grpc.Dial("authservice-service:50051", grpc.WithInsecure())
	auth_conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to AuthService: %v", err)
	}
	defer auth_conn.Close()

	// Create AuthService gRPC client
	authClient := clipb.NewAuthServiceClient(auth_conn)

	clipb.RegisterAuthServiceServer(grpcServer, &gateways.AuthServiceGateway{
		AuthClient: authClient,
	})

	/* AUTH SERVICE END */

	/* CONNECTING TO CRUD SERVICE */
	crud_conn, err := grpc.Dial("crudservice-service:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to CRUD Service: %v", err)
	}
	defer crud_conn.Close()

	crudClient := clipb.NewCRUDPipelineServiceClient(crud_conn)

	clipb.RegisterCRUDPipelineServiceServer(grpcServer, &gateways.CRUDServiceGateway{
		CRUDClient: crudClient,
	})

	/* CRUD SERVICE END */

	/* CONNECTING TO PIPELINE ORCHESTRATOR SERVICE */
	orch_conn, err := grpc.Dial("pipelineservice-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Pipeline Orchestrator Service: %v", err)
	}
	defer orch_conn.Close()

	orchClient := clipb.NewPipelineOrchestratorServiceClient(orch_conn)

	clipb.RegisterPipelineOrchestratorServiceServer(grpcServer, &gateways.OrchestratorServiceGateway{
		OrchestratorClient: orchClient,
	})

	/* PIPELINE ORCHESTRATOR SERVICE END */
	log.Printf("gRPC server is running on port 50055")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
