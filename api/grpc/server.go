package main

import (
	"grpc/auth"
	grpcauthpb "grpc/auth/proto"
	"grpc/crud"
	grpccrudpb "grpc/crud/proto"
	"grpc/orchestrator"
	pipelinepb "grpc/orchestrator/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	/*Connection to - AuthServiceServer */
	authServer := auth.AuthConnect("authservice-service:50051")
	// authServer := auth.AuthConnect("localhost:50051")
	// defer auth.CONN.Close()

	/*Connection to - CRUDPipelineServiceServer */
	crudServer := crud.CrudConnect("crudservice-service:50054")
	// crudServer := crud.CrudConnect("localhost:50054")
	// defer crud.CONN.Close()

	orchServer := orchestrator.OrchConnect("pipelineservice-service:50053")

	grpcServer := grpc.NewServer()

	/*-- Register CRUDPipelineServiceServer with the gRPC service --*/
	grpccrudpb.RegisterPipelineServiceServer(grpcServer, crudServer)

	/*-- Register AuthServiceServer with the gRPC service --*/
	grpcauthpb.RegisterAuthServiceServer(grpcServer, authServer)

	pipelinepb.RegisterPipelineOrchestratorServiceServer(grpcServer, orchServer)

	/*-- Enable Reflection --*/
	reflection.Register(grpcServer)

	// log.Println("Proxy AuthServiceServer  is running on port 60001...")
	// _ = grpcServer2.Serve(listener2)

	// log.Println("Proxy CRUDPipelineServiceServer  is running on port 60000...")
	// _ = grpcServer1.Serve(listener1)

	// go func() {

	listener, err := net.Listen("tcp", ":60000")
	if err != nil {
		log.Fatalf("Failed to listen on port 60000 %v", err)
	}

	log.Println("Proxy AuthServiceServer is running on port 60000...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve AuthServiceServer: %v", err)
	}

	// }()

}
