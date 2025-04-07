package crud

import (
	grpccrudpb "grpc/crud/proto"
	"grpc/crud/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var CONN *grpc.ClientConn

func CrudConnect(url string) *service.CrudServiceServer {
	/*-- CONNECTING TO AUTH SERVICE  --*/
	CONN, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to Connect to Crud Service %v", err)
	}
	// defer conn.Close()
	/*-- END --*/

	/*-- Create CrudServiceClient using the connection --*/
	crudClient := grpccrudpb.NewPipelineServiceClient(CONN)
	/*-- END --*/

	/*-- Initialize AuthServiceServer with the client --*/
	return &service.CrudServiceServer{
		CrudClient: crudClient,
	}
	/*-- END --*/
}
