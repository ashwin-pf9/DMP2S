package events

import (
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

type StageStatusUpdate struct {
	StageID string `json:"stage_id"`
	Status  string `json:"status"`
}

var natsConn *nats.Conn

// Initialize NATS connection
func InitNATS() {
	var err error
	natsConn, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	log.Println("Connected to NATS")
}

// Send update via NATS
func SendUpdate(stageID string, status string) {
	if natsConn == nil {
		log.Println("NATS connection not initialized")
		return
	}

	update := StageStatusUpdate{
		StageID: stageID,
		Status:  status,
	}
	message, err := json.Marshal(update)
	if err != nil {
		log.Println("Failed to marshal update:", err)
		return
	}

	err = natsConn.Publish("stage.updates", message)
	if err != nil {
		log.Println("Failed to publish to NATS:", err)
	}
}

// package events

// import (
// 	"encoding/json"
// 	"log"
// )

// var Broadcast = make(chan string, 100) // Make sure it matches `broadcast` in handlers

// type StageStatusUpdate struct {
// 	StageID string `json:"stage_id"`
// 	Status  string `json:"status"`
// }

// // Function to send updates
// func SendUpdate(stageID string, status string) {
// 	message, err := json.Marshal(StageStatusUpdate{StageID: stageID, Status: status})
// 	if err != nil {
// 		log.Println("Error marshalling JSON:", err)
// 		return
// 	}
// 	Broadcast <- string(message) // Send as JSON string
// }
