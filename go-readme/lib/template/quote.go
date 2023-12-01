// main.go
package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

// Data struct represents the dynamic data to be rendered in the template.
type QData struct {
	Quote string
}

var quotes = []string{
	"Believe you can and you're halfway there. -Theodore Roosevelt",
	"The only limit to our realization of tomorrow will be our doubts of today. -Franklin D. Roosevelt",
	"Your time is limited, don't waste it living someone else's life. -Steve Jobs",
	"The only way to do great work is to love what you do. -Steve Jobs",
	"Success is not final, failure is not fatal: It is the courage to continue that counts. -Winston Churchill",
}

func main() {
	// Handle requests with the quoteHandler function
	http.HandleFunc("/", quoteHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {
	// Prepare data
	data := QData{
		Quote: getRandomQuote(),
	}

	// Define the template content
	templateContent := `
	<html>
		<head>
			<title>Motivational Quote Generator</title>
			<style>
				body {
					font-family: 'Arial', sans-serif;
					text-align: center;
					padding: 50px;
					background-color: #f0f0f0;
				}
				.quote {
					font-size: 20px;
					font-style: italic;
					color: #333;
				}
			</style>
		</head>
		<body>
			<h1>Motivational Quote Generator</h1>
			<div class="quote">{{.Quote}}</div>
		</body>
	</html>
	`

	// Create a new template
	tmpl, err := template.New("motivationalQuote").Parse(templateContent)
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

// getRandomQuote returns a randomly selected quote from the quotes slice.
func getRandomQuote() string {
	rand.Seed(time.Now().UnixNano())
	return quotes[rand.Intn(len(quotes))]
}
