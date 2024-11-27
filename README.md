# quizmaster-go

# Go Quiz API

A lightweight and modular quiz API built with Go and SQLite. This project allows users to participate in quizzes, submit answers, track scores, and view the leaderboard.

## Features

- **Start a Quiz**: Fetch random questions for a quiz session.
- **Submit Answers**: Validate answers against the correct ones stored in the database.
- **Leaderboard**: View top scores submitted by players.
- **Score Submission**: Save your score and username to the leaderboard.

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: SQLite
- **API Format**: RESTful, with JSON payloads.

---

## Getting Started

### Prerequisites

- Go 1.19 or higher installed on your system.
- SQLite installed (optional, as the app creates the database if it doesn't exist).

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/go-quiz-api.git
   cd go-quiz-api
   ```
