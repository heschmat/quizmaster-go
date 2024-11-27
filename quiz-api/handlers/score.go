package handlers

import (
	"encoding/json"
	"net/http"
	"quiz-api/db"
)

// SubmitScoreRequest represents the payload to submit a player's score.
type SubmitScoreRequest struct {
	Username string `json:"username"`
	Score    int    `json:"score"`
}

// SubmitScore handles score submission and leaderboard updates.
func SubmitScore(w http.ResponseWriter, r *http.Request) {
	var request SubmitScoreRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := db.UpdateLeaderboard(request.Username, request.Score); err != nil {
		http.Error(w, "Error updating leaderboard", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Score submitted successfully"))
}
