package main

import (
	"fmt"
	"net/http"
)

func main() {
	go input()
	logInfo("Server started. Waiting for connections...")
	http.HandleFunc("/wss", handleWebSocket)

	err := http.ListenAndServe(":27357", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func input() {
	var command string
	client.init()
	for {
		fmt.Scanln(&command)
		switch command {
		case "list":
			logClients()
		}
	}
}
