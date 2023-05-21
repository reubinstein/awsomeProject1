package modules

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Challenge represents a challenge submitted by a user
type Challenge struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Level       string `json:"level"`
}

var challenges []Challenge

func main() {
	// Initialize the challenges slice
	challenges = []Challenge{
		{1, "Lack of clean water", "The community lacks access to clean drinking water", "National"},
		{2, "Poor road conditions", "The roads in the district are in bad shape and need repair", "District"},
		{3, "Lack of electricity", "The community has no access to electricity", "Ward"},
	}

	// Define the HTTP routes
	http.HandleFunc("/challenges", getChallengesHandler)
	http.HandleFunc("/challenges/add", addChallengeHandler)

	// Start the HTTP server
	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

// getChallengesHandler returns all challenges
func getChallengesHandler(w http.ResponseWriter, r *http.Request) {
	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Convert the challenges slice to JSON
	jsonData, err := json.Marshal(challenges)
	if err != nil {
		// Handle the error if the JSON conversion fails
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON data to the response writer
	_, err = w.Write(jsonData)
	if err != nil {
		return
	}
}

// addChallengeHandler adds a new challenge to the challenges slice

func addChallengeHandler(w http.ResponseWriter, r *http.Request) {
	var challenge Challenge
	err := json.NewDecoder(r.Body).Decode(&challenge)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}
	challenge.ID = len(challenges) + 1
	challenges = append(challenges, challenge)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(challenge)
	if err != nil {
		return
	}
}

func getAllChallengesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(challenges)
}

func getChallengeByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/challenges/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid challenge ID")
		return
	}
	for _, challenge := range challenges {
		if challenge.ID == id {
			json.NewEncoder(w).Encode(challenge)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Challenge not found")
}

func updateChallengeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/challenges/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid challenge ID")
		return
	}
	var updatedChallenge Challenge
	err = json.NewDecoder(r.Body).Decode(&updatedChallenge)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request payload")
		return
	}
	for i, challenge := range challenges {
		if challenge.ID == id {
			updatedChallenge.ID = id
			challenges[i] = updatedChallenge
			json.NewEncoder(w).Encode(updatedChallenge)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Challenge not found")
}

func deleteChallengeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/challenges/"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid challenge ID")
		return
	}
	for i, challenge := range challenges {
		if challenge.ID == id {
			challenges = append(challenges[:i], challenges[i+1:]...)
			json.NewEncoder(w).Encode(challenge)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Challenge not found")
}
