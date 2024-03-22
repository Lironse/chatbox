package main

import (
	"encoding/json"
	"net/http"
)

type RegistrationRequest struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type RegistrationResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func handlePreflight(w http.ResponseWriter) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST") // Add other allowed methods if needed
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Respond with no content and 200 status code
	w.WriteHeader(http.StatusOK)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
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

	response := RegistrationResponse{}

	registrationStatus := routingTable.isUsernameAvailable(reqBody.Name)

	if registrationStatus {
		// Send a response indicating successful registration
		response = RegistrationResponse{
			Status:  "success",
			Message: "Client registered successfully",
		}

		routingTable.addNode(Node{Id: len(reqBody.Name), LocalId: calculateLocalId(len(reqBody.Name)), Ip: reqBody.Key})
	}

	if !registrationStatus {
		response = RegistrationResponse{
			Status:  "failure",
			Message: "Username was taken",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
