package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", handleWebSocket)

	err := http.ListenAndServe(":27357", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
