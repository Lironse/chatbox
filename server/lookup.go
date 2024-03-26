package main

import (
	"encoding/json"
	"net/http"
)

type LookupRequest struct {
	Username string `json:"username"`
}

type LookupResponse struct {
	Status string `json:"status"`
	Key    string `json:"key,omitempty"`
}

func handleLookupPeer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		handlePreflight(w)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	// Decode the request body into a RegistrationRequest struct
	var reqBody LookupRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userExists, key := routingTable.doesUserExist(reqBody.Username)

	var response LookupResponse

	if userExists {
		response = LookupResponse{
			Status: "success",
			Key:    key,
		}
	} else {
		response = LookupResponse{
			Status: "user does not exist",
			Key:    "",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
