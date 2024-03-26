package main

import (
	"hash/fnv"
	"net/http"
)

func handlePreflight(w http.ResponseWriter) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST") // Add other allowed methods if needed
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Respond with no content and 200 status code
	w.WriteHeader(http.StatusOK)
}

func hashUsername(username string) int {
	h := fnv.New32a()
	h.Write([]byte(username))
	return int(h.Sum32()) % keyspaceSize
}
