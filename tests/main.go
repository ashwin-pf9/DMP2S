package main

import (
	"DMP2S/api/rest"
	"DMP2S/internal/core/services"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	// Create the actual implementation
	orchestrator := &services.PipelineOrchestrator{
		ID:   uuid.New(),
		Name: "Main Pipeline",
	}

	// Initialize service layer with implementation
	pipelineService := services.NewPipelineOrchestratorService(orchestrator)

	// API handler
	pipelineHandler := rest.NewPipelineHandler(pipelineService)

	// HTTP Routes
	http.HandleFunc("/add-stage", pipelineHandler.AddStageHandler)
	http.HandleFunc("/execute-pipeline", pipelineHandler.ExecutePipelineHandler)

	// Start Server
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
