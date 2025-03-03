package handlers

import (
	"DMP2S/internal/core/domain"
	"DMP2S/internal/core/services"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
func authenticateRequest(w http.ResponseWriter, r *http.Request) (*supabase.User, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
		return nil, errors.New("missing authorization token")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("ANON_KEY") //anon_key because it is a client side request

	// Validate Token with Supabase
	client := supabase.CreateClient(url, key)
	user, err := client.Auth.User(context.TODO(), token)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		return nil, errors.New("Invalid authorization token")
	}
	return user, nil
}

func GetPipelinesHandler(w http.ResponseWriter, r *http.Request) { //Will be called on clients request
	//REQUEST AUTHENTICATION
	user, err := authenticateRequest(w, r)
	if err != nil {
		log.Printf("Error while authenticating user : %v", err)
	}
	//REQUEST AUTHENTICATION DONE

	// Call actual function to get user pipelines
	pipelines := services.GetUsersPipelines(user.ID)

	// Send Response
	json.NewEncoder(w).Encode(pipelines)
}

func CreatePipelineHandler(w http.ResponseWriter, r *http.Request) {
	//REQUEST AUTHENTICATION
	user, err := authenticateRequest(w, r)
	if err != nil {
		log.Printf("Error while authenticating user : %v", err)
		return
	}
	//REQUEST AUTHENTICATION DONE

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

func GetStagesHandler(w http.ResponseWriter, r *http.Request) {
	//REQUEST AUTHENTICATION
	_, err := authenticateRequest(w, r)
	if err != nil {
		log.Printf("Error while authenticating user : %v", err)
		return
	}
	//REQUEST AUTHENTICATION DONE

	//Fetching "pipeline_id" from request URL
	vars := mux.Vars(r)
	pipelineID := uuid.MustParse(vars["pipeline_id"])

	stages := services.GetPipelineStages(pipelineID)

	//Now i want to write this stages array to the response socket - for that i need to conver to json type
	json.NewEncoder(w).Encode(stages)
}

// ------------------------------------------------------------- //
// API handler

// CREATING IMPLEMENTATION OBJECT WHICH IMPLEMENTS ALL INTERFACE METHODS
var impl = services.NewPipelineOrchestratorImpl()

// CREATING CONNECTION WITH THE IMPLEMENTATION THROUGH INTERFACE
var orchestrator = services.NewPipelineOrchestratorService(impl) //PLUGGING IMPLEMENTATION TO THE ORCHESTRATOR

func AddStageHandler(w http.ResponseWriter, r *http.Request) {
	//REQUEST AUTHENTICATION
	_, err := authenticateRequest(w, r)
	if err != nil {
		log.Printf("Error while authenticating user : %v", err)
		return
	}
	//REQUEST AUTHENTICATION DONE

	var stage domain.Stage
	if err := json.NewDecoder(r.Body).Decode(&stage); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err = orchestrator.AddStageToPipeline(stage)
	if err != nil {
		http.Error(w, "Failed to add stage", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stage added successfully"))
}

// API: Execute Pipeline
func ExecutePipelineHandler(w http.ResponseWriter, r *http.Request) {
	//REQUEST AUTHENTICATION
	_, err := authenticateRequest(w, r)
	if err != nil {
		log.Printf("Error while authenticating user : %v", err)
		return
	}
	//REQUEST AUTHENTICATION DONE

	ctx := context.Background()
	//Extracting "pipeline_id" from request URL
	vars := mux.Vars(r)
	pipelineID, err := uuid.Parse(vars["pipeline_id"])

	result, err := orchestrator.ExecutePipeline(ctx, pipelineID)
	if err != nil {
		errString := fmt.Sprintf("Execution failed : %v", err)
		http.Error(w, errString, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
