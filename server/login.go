package main

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type LoginResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
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

	// registrationStatus = checkUsernameAvailability
	loginStatus := "successful"

	response := LoginResponse{}

	if loginStatus == "successful" {
		// Send a response indicating successful registration
		response = LoginResponse{
			Status:  "success",
			Message: "Client registered successfully",
		}
	}

	if loginStatus == "usernameTaken" {
		response = LoginResponse{
			Status:  "failure",
			Message: "Username was taken",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
