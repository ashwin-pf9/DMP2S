package main

import (
	crudpipelinepb "crudpipelineservice/proto"
	"crudpipelineservice/service"
	"log"
	"net"

	"github.com/ashwin-pf9/shared/db"
	"google.golang.org/grpc"
)

func main() {
	crudService := service.NewPipelineService(db.InitDatabase())

	authServer := service.NewPipelineServer(crudService)

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	crudpipelinepb.RegisterPipelineServiceServer(grpcServer, authServer)

	log.Println("gRPC server running on port 50054...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
