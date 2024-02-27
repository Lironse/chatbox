package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	logInfo("Connection from: ", r.RemoteAddr)

	// upgrade the http connection to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logError(ErrUpgradeFailure)
		return // close server
	}

	// start reading and writing to clients
	for {
		// read the packet into a Packet struct
		var packet Packet
		err = conn.ReadJSON(&packet)
		if err != nil {
			logError(ErrReadPacketFailure)
			return // close server
		}

		// log the full incoming packet
		logPacket(packet)
		err = handlePacket(packet, conn)
		if err != nil {
			logError(err)
			continue // dismiss packet on error
		}
	}
}

func main() {
	logInfo("Server started. Waiting for connections...")
	http.HandleFunc("/ws", handleWebSocket)

	err := http.ListenAndServe(":27357", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
