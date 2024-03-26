package main

import (
	"encoding/json"
	"net/http"
)

type RegistrationRequest struct {
	Username string `json:"username"`
	Key      string `json:"key"`
}

type RegistrationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	// Check method
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
	var reqBody RegistrationRequest
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userExists, _ := routingTable.doesUserExist(reqBody.Username)

	var response RegistrationResponse

	if !userExists {
		response = RegistrationResponse{
			Status:  "success",
			Message: "Client registered successfully",
		}
		routingTable.addNode(Node{Id: hashUsername(reqBody.Username), LocalId: calculateLocalId(hashUsername(reqBody.Username)), Ip: reqBody.Key})
	} else {
		response = RegistrationResponse{
			Status:  "failure",
			Message: "Username is taken",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
