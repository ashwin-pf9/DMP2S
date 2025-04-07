package auth

import (
	grpcauthpb "grpc/auth/proto"
	"grpc/auth/service"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var CONN *grpc.ClientConn

func AuthConnect(url string) *service.AuthServiceServer {
	/*-- CONNECTING TO AUTH SERVICE  --*/
	CONN, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to Connect to Auth Service %v", err)
	}
	// defer conn.Close()
	/*-- END --*/

	/*-- Create AuthServiceClient using the connection --*/
	authClient := grpcauthpb.NewAuthServiceClient(CONN)
	/*-- END --*/

	/*-- Initialize AuthServiceServer with the client --*/
	return &service.AuthServiceServer{
		AuthClient: authClient,
	}
	/*-- END --*/
}
