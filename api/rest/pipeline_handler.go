package rest

import (
	"DMP2S/internal/core/domain"
	"DMP2S/internal/core/services"
	"context"
	"encoding/json"
	"net/http"
)

// API handler
type PipelineHandler struct {
	service *services.PipelineOrchestratorService
}

// Constructor
func NewPipelineHandler(service *services.PipelineOrchestratorService) *PipelineHandler {
	return &PipelineHandler{service: service}
}

// API: Add Stage
func (h *PipelineHandler) AddStageHandler(w http.ResponseWriter, r *http.Request) {
	var stage domain.Stage
	if err := json.NewDecoder(r.Body).Decode(&stage); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.service.AddStageToPipeline(stage)
	if err != nil {
		http.Error(w, "Failed to add stage", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stage added successfully"))
}

// API: Execute Pipeline
func (h *PipelineHandler) ExecutePipelineHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	input := "Sample Input"

	result, err := h.service.ExecutePipeline(ctx, input)
	if err != nil {
		http.Error(w, "Execution failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
