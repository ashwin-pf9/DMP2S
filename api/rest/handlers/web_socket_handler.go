package handlers

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/websocket"
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
