package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	crudpipelinepb "github.com/ashwin-pf9/DMP2S/internal/protobuffs/crud"
	"google.golang.org/grpc"
)

// gRPC client for the pipeline service
var crudClient crudpipelinepb.PipelineServiceClient

// Initialize gRPC client (Call this in main.go)
func InitCRUDClient() {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure()) // Adjust port if needed
	if err != nil {
		log.Fatalf("Failed to connect to CRUD gRPC server: %v", err)
	}
	crudClient = crudpipelinepb.NewPipelineServiceClient(conn)
}

// ----------- CreatePipelineHandler -----------

func CreatePipelineHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		UserID string `json:"user_id"`
		Name   string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	InitCRUDClient()

	resp, err := crudClient.CreatePipeline(ctx, &crudpipelinepb.CreatePipelineRequest{
		UserId: req.UserID,
		Name:   req.Name,
	})
	if err != nil {
		http.Error(w, "Failed to create pipeline: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp.Pipeline); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// ----------- GetUserPipelinesHandler -----------

func GetPipelinesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		UserID string `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	InitCRUDClient()

	resp, err := crudClient.GetUserPipelines(ctx, &crudpipelinepb.GetUserPipelinesRequest{
		UserId: req.UserID,
	})
	if err != nil {
		http.Error(w, "Failed to fetch pipelines: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp.Pipelines); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// ----------- GetPipelineStagesHandler -----------

func GetStagesHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		PipelineID string `json:"pipeline_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.PipelineID == "" {
		http.Error(w, "pipeline_id is required", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	InitCRUDClient()

	resp, err := crudClient.GetPipelineStages(ctx, &crudpipelinepb.GetPipelineStagesRequest{
		PipelineId: req.PipelineID,
	})
	if err != nil {
		http.Error(w, "Failed to fetch pipeline stages: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched %d stages for pipeline %s", len(resp.Stages), req.PipelineID)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp.Stages); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
