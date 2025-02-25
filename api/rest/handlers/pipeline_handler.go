package handlers

import (
	"DMP2S/internal/core/domain"
	"DMP2S/internal/core/services"
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/nedpals/supabase-go"
)

//HANDLERS SHOULD FOLLOW GIVEN SIGNATURE
//func FunctionName(w http.ResponseWriter, r *http.Request)

/* CODE WRITTEN BY ME, MIGHT BE WRONG */

// Purpose: Should extract data from request data such as:
// jwt token to check if user is authenticated to access these pipelines or not
// user id
// user name?
// ...
func GetPipelinesHandler(w http.ResponseWriter, r *http.Request) { //Will be called on clients request
	// Extract Token from Authorization Header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("ANON_KEY") //anon_key because it is a client side request

	// Validate Token with Supabase
	client := supabase.CreateClient(url, key)
	user, err := client.Auth.User(context.TODO(), token)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	// Call actual function to get user pipelines
	pipelines := services.GetUsersPipelines(user.ID)

	// Send Response
	json.NewEncoder(w).Encode(pipelines)
}

func CreatePipelineHandler(w http.ResponseWriter, r *http.Request) {
	// Extract token from Authorization Header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("ANON_KEY") //anon_key because it is a client side request

	// Validate token with Supabase
	client := supabase.CreateClient(url, key) // Supabase client
	user, err := client.Auth.User(context.TODO(), token)

	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return
	}

	//..User authentication end...//

	// Parse request body to get pipeline name
	var req struct {
		Name string `json:"name"`
	}
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call service function to create pipeline
	pipeline, err := services.CreatePipeline(user.ID, req.Name)
	if err != nil {
		http.Error(w, "Failed to create pipeline", http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pipeline)
}

// -------------------------------------------------------------
// API handler

// CREATING IMPLEMENTATION OBJECT WHICH IMPLEMENTS ALL INTERFACE METHODS
var impl = services.NewPipelineOrchestratorImpl()

// CREATING CONNECTION WITH THE IMPLEMENTATION THROUGH INTERFACE
var orchestrator = services.NewPipelineOrchestratorService(impl) //PLUGGING IMPLEMENTATION TO THE ORCHESTRATOR

func AddStageHandler(w http.ResponseWriter, r *http.Request) {
	var stage domain.Stage
	if err := json.NewDecoder(r.Body).Decode(&stage); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := orchestrator.AddStageToPipeline(stage)
	if err != nil {
		http.Error(w, "Failed to add stage", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stage added successfully"))
}

// API: Execute Pipeline
func ExecutePipelineHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	vars := mux.Vars(r)
	pipelineID, err := uuid.Parse(vars["pipeline_id"])

	result, err := orchestrator.ExecutePipeline(ctx, pipelineID)
	if err != nil {
		http.Error(w, "Execution failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
