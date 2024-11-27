package main

import (
	"log"
	"net/http"
	"quiz-api/db"
	"quiz-api/handlers"
)

func main() {
	// Initialize the database connection
	db.Init()

	// Register HTTP routes
	http.HandleFunc("/start-quiz", handlers.StartQuiz)
	http.HandleFunc("/submit-answer", handlers.SubmitAnswer)
	http.HandleFunc("/leaderboard", handlers.GetLeaderboard)
	http.HandleFunc("/submit-score", handlers.SubmitScore)

	// Start the server
	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
