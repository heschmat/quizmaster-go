package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quiz-api/db"
)

// AnswerSubmission represents the JSON payload for submitting an answer.
type AnswerSubmission struct {
	SessionID  int `json:"session_id"`
	QuestionID int `json:"question_id"`
	Answer     int `json:"answer"`
}

// SubmitAnswer handles the submission of an answer to a question.
func SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON payload from the request body.
	var submission AnswerSubmission
	if err := json.NewDecoder(r.Body).Decode(&submission); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Fetch the correct answer for the question from the database.
	correctAnswer, err := db.GetCorrectAnswer(submission.QuestionID)
	if err != nil {
		http.Error(w, "Error retrieving correct answer", http.StatusInternalServerError)
		return
	}

	// Compare the submitted answer with the correct answer.
	correct := submission.Answer == correctAnswer

	// Update the score for the session (this part would depend on how you're tracking sessions).
	// For simplicity, we'll print whether the answer was correct.
	if correct {
		fmt.Println("Correct answer for session:", submission.SessionID)
	} else {
		fmt.Println("Wrong answer for session:", submission.SessionID)
	}

	// Respond with the result.
	response := map[string]interface{}{
		"correct": correct,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
