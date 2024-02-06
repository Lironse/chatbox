package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Packet struct {
	Type    string `json:"type"`
	Action  string `json:"action"`
	Payload string `json:"payload"`
	From    string `json:"from"`
	To      string `json:"to"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections for this example
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var clientsMutex sync.Mutex

func broadcastMessage(message []byte, sender *websocket.Conn) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	for client := range clients {
		// Send the message to all clients except the sender
		if client != sender {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				fmt.Println("Error writing message:", err)
				removeClient(client)
			}
		}
	}
}

func removeClient(client *websocket.Conn) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()

	delete(clients, client)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	fmt.Println("Connection from: ", r.RemoteAddr)

	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	for {
		var packet Packet
		// Read SDP message from the browser
		err = conn.ReadJSON(&packet)
		if err != nil {
			fmt.Println("Error reading message:", err)
			removeClient(conn)
			return
		}

		// Print the received SDP message
		fmt.Printf("Received SDP from: %s\n", r.RemoteAddr)
		// Broadcast the SDP message to all other clients
		packetString, _ := json.Marshal(packet)
		broadcastMessage([]byte(packetString), conn)
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	// Start the WebSocket server on port 8080
	err := http.ListenAndServe(":27357", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
