package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
)

var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

var natsConn *nats.Conn

// Initialize NATS and subscribe to updates
func InitNATSSubscriber() {
	var err error
	natsConn, err = nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	log.Println("Connected to NATS for subscription")

	_, err = natsConn.Subscribe("stage.updates", func(msg *nats.Msg) {
		log.Println("Received stage update:", string(msg.Data))
		// Broadcast to all WebSocket clients
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, msg.Data)
			if err != nil {
				log.Println("WebSocket write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
	})
	if err != nil {
		log.Fatalf("Failed to subscribe to stage.updates: %v", err)
	}
}

func StatusUpdatesHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()
	log.Println("New WebSocket client connected")

	clients[conn] = true

	// Keep connection open
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			delete(clients, conn)
			break
		}
	}
}

// package handlers

// import (
// 	"log"
// 	"net/http"
// )

// var clients = make(map[*websocket.Conn]bool) // Connected clients
// var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

// func StatusUpdatesHandler(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("WebSocket upgrade failed:", err)
// 		return
// 	}
// 	log.Println("New WebSocket client connected")
// 	defer conn.Close()

// 	clients[conn] = true

// 	// Continuously listen for messages from the broadcast channel
// 	for msg := range events.Broadcast {
// 		log.Println("Broadcasting message:", msg)
// 		for client := range clients {
// 			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
// 			if err != nil {
// 				log.Println("Write error:", err)
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 	}
// }
