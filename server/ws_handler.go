package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections for this example
		return true
	},
}
var clients = make(map[string]*websocket.Conn)

// var clientsMutex sync.Mutex

func passPacket(message []byte, sender string, to string) {
	_, ok := clients[to]
	if ok {
		err := clients[to].WriteMessage(websocket.TextMessage, message)
		if err != nil {
			fmt.Println("Error writing message:", err)
			removeClient(sender)
		}
	} else {
		fmt.Println("User", to, "is not connected!")
	}
}

func addClient(name string, conn *websocket.Conn) {
	clients[name] = conn
}

func removeClient(name string) {
	delete(clients, name)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connection from: ", r.RemoteAddr)

	// upgrade the http connection to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	// start reading and writing to clients
	for {
		// read the packet into a Packet struct
		var packet Packet
		err = conn.ReadJSON(&packet)
		// close connection on error
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		// log the full incoming packet
		logPacket(packet)

		// marshal the json struct into a string
		packetString, err := json.Marshal(packet)
		// dismiss packet on error
		if err != nil {
			fmt.Println("Faulty packet:", err)
			continue
		}

		// Act according to packet.Action
		// Accepted actions: register, passOffer, passAnswer
		switch packet.Action {
		case "open", "register":
			addClient(packet.From, conn)

		case "passOffer", "passAnswer":
			passPacket([]byte(packetString), packet.From, packet.To)
		}
	}
}
