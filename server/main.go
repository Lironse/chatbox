package main

import (
	"net/http"
)

func main() {
	routingTable.addNode(localNode)

	routingTable.RegisterToServerList()

	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/register", handleRegister)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/lookup", handleLookupPeer)

	logInfo("Server starting. Listening on port 27357...")
	go input_commands()

	err := http.ListenAndServe(":27357", nil)
	if err != nil {
		logError(ErrServerOpenFailure)
	}
}
