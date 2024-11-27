package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// Initialize the database connection
func Init() {
	var err error
	db, err = sql.Open("sqlite3", "./data/quiz.db")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
}

func FetchQuestions(limit int) ([]map[string]interface{}, error) {
	rows, err := db.Query("SELECT id, prompt, option1, option2, option3, option4, answer FROM questions ORDER BY RANDOM() LIMIT ?", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []map[string]interface{}
	for rows.Next() {
		var id, answer int
		var prompt, opt1, opt2, opt3, opt4 string
		if err := rows.Scan(&id, &prompt, &opt1, &opt2, &opt3, &opt4, &answer); err != nil {
			return nil, err
		}
		questions = append(questions, map[string]interface{}{
			"id":      id,
			"prompt":  prompt,
			"options": []string{opt1, opt2, opt3, opt4},
			"answer":  answer,
		})
	}
	return questions, nil
}


// FetchLeaderboard retrieves the top scores from the leaderboard table.
func FetchLeaderboard() ([]map[string]interface{}, error) {
    rows, err := db.Query("SELECT username, score, created_at FROM leaderboard ORDER BY score DESC LIMIT 10")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var leaderboard []map[string]interface{}

    // Loop through each row in the result set.
    for rows.Next() {
        var username string
        var score int
        var createdAt string

        // Scan the values from the current row into variables.
        if err := rows.Scan(&username, &score, &createdAt); err != nil {
            return nil, err
        }

        // Add the entry as a map to the leaderboard slice.
        leaderboard = append(leaderboard, map[string]interface{}{
            "username":   username,
            "score":      score,
            "created_at": createdAt,
        })
    }

    return leaderboard, nil
}


// GetCorrectAnswer fetches the correct answer for a given question ID.
func GetCorrectAnswer(questionID int) (int, error) {
	var correctAnswer int
	err := db.QueryRow("SELECT answer FROM questions WHERE id = ?", questionID).Scan(&correctAnswer)
	if err != nil {
		return 0, err
	}
	return correctAnswer, nil
}


// UpdateLeaderboard adds a new score to the leaderboard.
func UpdateLeaderboard(username string, score int) error {
	_, err := db.Exec("INSERT INTO leaderboard (username, score) VALUES (?, ?)", username, score)
	return err
}
