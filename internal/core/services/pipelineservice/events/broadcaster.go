package events

import (
	"encoding/json"
	"log"
)

var Broadcast = make(chan string, 100) // Make sure it matches `broadcast` in handlers

type StageStatusUpdate struct {
	StageID string `json:"stage_id"`
	Status  string `json:"status"`
}

// Function to send updates
func SendUpdate(stageID string, status string) {
	message, err := json.Marshal(StageStatusUpdate{StageID: stageID, Status: status})
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		return
	}
	Broadcast <- string(message) // Send as JSON string
}
