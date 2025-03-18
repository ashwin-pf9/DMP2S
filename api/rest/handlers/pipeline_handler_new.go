package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "github.com/ashwin-pf9/DMP2S/internal/protobuffs/pipeline"
	"google.golang.org/grpc"
)

var pipelineClient pb.PipelineServiceClient
var pipelineServer pb.PipelineOrchestratorServiceClient

func InitPipelineClient() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	pipelineClient = pb.NewPipelineServiceClient(conn)
}

// func CreatePipelineHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var req struct {
// 		Name   string `json:"name"`
// 		UserID string `json:"user_id"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Invalid request body", http.StatusBadRequest)
// 		return
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
// 	defer cancel()

// 	InitPipelineClient()

// 	resp, err := pipelineClient.CreatePipeline(ctx, &pb.CreatePipelineRequest{
// 		Name:   req.Name,
// 		UserId: req.UserID,
// 	})
// 	if err != nil {
// 		error := fmt.Sprintf("Pipeline creation failed %v", err.Error())
// 		http.Error(w, error, http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(resp)
// }

func InitPipelineServer() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	pipelineServer = pb.NewPipelineOrchestratorServiceClient(conn)
}

func ExecutePipelineHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		PipelineID string `json:"pipeline_id"`
	}
	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate pipeline_id
	if req.PipelineID == "" {
		http.Error(w, "pipeline_id is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	InitPipelineServer()

	log.Printf("from new_handler - calling pipelineServer.ExecutePipeline\n") //--for debugging

	resp, err := pipelineServer.ExecutePipeline(ctx, &pb.ExecutePipelineRequest{
		PipelineId: req.PipelineID,
	})
	if err != nil {
		error := fmt.Sprintf("Pipeline creation failed : %v", err.Error())
		http.Error(w, error, http.StatusInternalServerError)
		return
	}

	log.Printf("from new_handler - control came back from pipelineServer.ExecutePipeline\n") //--for debugging

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func AddStageHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		Name       string `json:"name"`
		PipelineID string `json:"pipeline_id"`
	}
	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate name and pipeline_id
	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}
	if req.PipelineID == "" {
		http.Error(w, "pipeline_id is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	InitPipelineServer()

	log.Printf("from addstage_handler - calling pipelineServer.AddStageToPipeline\n") //--for debugging

	_, err = pipelineServer.AddStageToPipeline(ctx, &pb.AddStageRequest{
		Stage: &pb.Stage{
			PipelineId: req.PipelineID,
			Name:       req.Name,
		},
	})
	if err != nil {
		error := fmt.Sprintf("Add stage failed: %v", err.Error())
		http.Error(w, error, http.StatusInternalServerError)
		return
	}

	log.Printf("from addstage_handler - control came back from pipelineServer.AddStageToPipeline\n") //--for debugging

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Stage added successfully"})
}

func DeletePipelineHandler(w http.ResponseWriter, r *http.Request) {

	var req struct {
		PipelineID string `json:"pipeline_id"`
	}
	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate pipeline_id
	if req.PipelineID == "" {
		http.Error(w, "pipeline_id is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	InitPipelineServer()

	log.Printf("from delete_handler - calling pipelineServer.DeletePipeline\n") //--for debugging

	_, err = pipelineServer.DeletePipeline(ctx, &pb.PipelineIDRequest{
		PipelineId: req.PipelineID,
	})
	if err != nil {
		error := fmt.Sprintf("Delete pipeline failed: %v", err.Error())
		http.Error(w, error, http.StatusInternalServerError)
		return
	}

	log.Printf("from delete_handler - control came back from pipelineServer.DeletePipeline\n") //--for debugging

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Pipeline deleted successfully"})
}
