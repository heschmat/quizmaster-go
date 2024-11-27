package handlers

import (
	"encoding/json"
	"net/http"
	"quiz-api/db"
)

func GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	// Fetch top scores from the database
	leaderboard, err := db.FetchLeaderboard()
	if err != nil {
		http.Error(w, "Error fetching leaderboard", http.StatusInternalServerError)
		return
	}

	// Respond with the leaderboard
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(leaderboard)
}
