package rest

import (
	"fmt"
	"net/http"

	"DMP2S/api/rest/handlers" // Import handlers

	"github.com/gorilla/mux"
)

// StartServer initializes and starts the HTTP server
func StartServer() error {
	//Better than mux.NewRouter(), gives more control, supports dynamic routes
	router := mux.NewRouter()

	// authentication routes - Response of these requests will be a jsonified USER object which will be sent back to the client,
	//Using that json response user will send a for e.g "/pipelines" request with that json object.
	//That request will be handled by the handlers.GetPipelineHandler and it will check that json body, and fetch required information,
	//such as - JWT token, User id, etc...

	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST") //maps /register endpoint to path "handlers.RegisterHandler"
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")       //maps /login endpoint to path "handlers.LoginHandler"

	/*
		To make following request the Client must have a valid json object of type user "*supabase.User"
	*/
	router.HandleFunc("/pipelines", handlers.GetPipelinesHandler).Methods("GET")           //maps /pipelines endpoint to path handlers.GetPipelinesHandler
	router.HandleFunc("/pipelines/create", handlers.CreatePipelineHandler).Methods("POST") //maps /pipelines/create endpoint to path handlers.CreatePipelinesHandler

	/*
		To make user first need to login and create pipeline, then from that pipeline(card on UI) option for creating state will be provided
	*/

	router.HandleFunc("/pipelines/{pipeline_id}/stages", handlers.GetStagesHandler).Methods("GET")
	router.HandleFunc("/pipelines/{pipeline_id}/stages/add", handlers.AddStageHandler).Methods("POST")

	/*
		Now that stages are added to the pipeline : user can call on these endpoints for further operation
	*/
	router.HandleFunc("/pipelines/{pipeline_id}/execute", handlers.ExecutePipelineHandler).Methods("POST")

	// Start the server
	port := ":8080"
	fmt.Println("Server started on", port)
	return http.ListenAndServe(port, router)
}
