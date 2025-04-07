package service

import (
	grpccrudpb "DMP2S/global-protos/crud"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Global variable to hold the gRPC client
var crudClient grpccrudpb.PipelineServiceClient

// initGRPCClient initializes the connection to the gRPC server
func InitCrudClient(url string) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	crudClient = grpccrudpb.NewPipelineServiceClient(conn)
}

func CreatePipeline(userID, name string) {
	req := &grpccrudpb.CreatePipelineRequest{
		UserId: userID,
		Name:   name,
	}

	resp, err := crudClient.CreatePipeline(context.Background(), req)
	if err != nil {
		log.Fatalf("Error creating pipeline: %v", err)
	}
	fmt.Printf("Pipeline created! pipelineID: %s\n", resp.Pipeline.Id)
}

func GetUserPipelines(userID string) {
	req := &grpccrudpb.GetUserPipelinesRequest{
		UserId: userID,
	}

	resp, err := crudClient.GetUserPipelines(context.Background(), req)
	if err != nil {
		log.Fatalf("Error fetching pipelines: %v", err)
	}
	fmt.Printf("Pipeline fetched! : %s\n", resp.Pipelines)

}

func GetPipelineStages(pipelineID string) {
	req := &grpccrudpb.GetPipelineStagesRequest{
		PipelineId: pipelineID,
	}

	resp, err := crudClient.GetPipelineStages(context.Background(), req)
	if err != nil {
		log.Fatalf("Error fetching pipeline stages %v", err)
	}

	fmt.Printf("Pipeline stages : %s", resp.Stages)
}
