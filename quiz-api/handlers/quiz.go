package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"quiz-api/db"
)

func StartQuiz(w http.ResponseWriter, r *http.Request) {
	// Fetch random questions from the database
	questions, err := db.FetchQuestions(5) // Assume 5 questions
	if err != nil {
		http.Error(w, "Error fetching questions", http.StatusInternalServerError)
		return
	}

	// Create a new quiz session (you can generate unique session IDs here)
	sessionID := rand.Int() // Simplified session ID generation

	// Respond with session ID and questions
	response := map[string]interface{}{
		"session_id": sessionID,
		"questions":  questions,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
