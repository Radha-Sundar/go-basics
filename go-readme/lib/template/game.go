// main.go
package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Data struct represents the dynamic data to be rendered in the template.
type GData struct {
	Message string
}

var (
	targetNumber int
)

func main() {
	// Generate a random number between 1 and 100
	rand.Seed(time.Now().UnixNano())
	targetNumber = rand.Intn(100) + 1

	// Handle requests with the gameHandler function
	http.HandleFunc("/", gameHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	// Parse user input
	guessStr := r.URL.Query().Get("guess")
	guess, err := strconv.Atoi(guessStr)
	if err != nil {
		guess = -1 // Invalid guess
	}

	// Process the user's guess
	message := processGuess(guess)

	// Prepare data
	data := GData{
		Message: message,
	}

	// Define the template content
	templateContent := `
	<html>
		<head>
			<title>Guess the Number Game</title>
			<style>
				body {
					font-family: 'Arial', sans-serif;
					text-align: center;
					padding: 50px;
					background-color: #f0f0f0;
				}
				.game-container {
					border: 2px solid #333;
					border-radius: 10px;
					padding: 20px;
					margin: 20px auto;
					max-width: 600px;
					background-color: #fff;
				}
				.message {
					font-size: 18px;
					font-weight: bold;
					color: #333;
				}
			</style>
		</head>
		<body>
			<h1>Guess the Number Game</h1>
			<div class="game-container">
				<p>Guess a number between 1 and 100:</p>
				<form action="/" method="get">
					<input type="text" name="guess" />
					<input type="submit" value="Submit" />
				</form>
				<p class="message">{{.Message}}</p>
			</div>
		</body>
	</html>
	`

	// Create a new template
	tmpl, err := template.New("guessGame").Parse(templateContent)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the template, passing data to it
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func processGuess(guess int) string {
	if guess == -1 {
		return "Invalid guess. Please enter a number."
	}

	if guess < targetNumber {
		return "Too low. Try again!"
	} else if guess > targetNumber {
		return "Too high. Try again!"
	}

	return "Congratulations! You guessed the correct number."
}
