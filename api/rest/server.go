package main

import (
	"fmt"
	"net/http"

	"github.com/ashwin-pf9/DMP2S/api/rest/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	StartRESTServer()
}

// StartServer initializes and starts the HTTP server
// func StartRESTServer() error {
func StartRESTServer() error {
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
	router.HandleFunc("/pipelines", handlers.GetPipelinesHandler).Methods("POST")          //maps /pipelines endpoint to path handlers.GetPipelinesHandler
	router.HandleFunc("/pipelines/create", handlers.CreatePipelineHandler).Methods("POST") //maps /pipelines/create endpoint to path handlers.CreatePipelinesHandler
	router.HandleFunc("/pipelines/{pipeline_id}/stages", handlers.GetStagesHandler).Methods("GET")

	/*
		To make user first need to login and create pipeline, then from that pipeline(card on UI) option for creating state will be provided
		Now that stages are added to the pipeline : user can call on these endpoints for further operation
	*/

	router.HandleFunc("/pipelines/{pipeline_id}/stages/add", handlers.AddStageHandler).Methods("POST")
	router.HandleFunc("/pipelines/{pipeline_id}/start", handlers.ExecutePipelineHandler).Methods("POST")
	router.HandleFunc("/pipelines/{pipeline_id}/delete", handlers.DeletePipelineHandler).Methods("POST")

	/*--  Endpoint for WEB SOCKET --*/
	handlers.InitNATSSubscriber() // For subscribing
	router.HandleFunc("/ws/status-updates", handlers.StatusUpdatesHandler)

	// // Serve React frontend (static files)
	// staticDir := "../../web/frontend/build"
	// router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir+"/static"))))
	// router.PathPrefix("/").Handler(http.FileServer(http.Dir(staticDir)))

	// Enable CORS using `rs/cors` package
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:30001", "http://localhost:3000"}, // Allow frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Start the server
	port := ":8080"
	fmt.Println("Server started on", port)
	return http.ListenAndServe(port, corsHandler.Handler(router))
}
