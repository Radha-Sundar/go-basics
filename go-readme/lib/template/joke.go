// main.go
package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

// Joke represents a joke structure from the API response.
type Joke struct {
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
}

// Data struct represents the dynamic data to be rendered in the template.
type JData struct {
	Joke string
}

func main() {
	// Handle requests with the jokeHandler function
	http.HandleFunc("/", jokeHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch a random joke from an external API
	joke, err := getRandomJoke()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Prepare data
	data := JData{
		Joke: joke,
	}

	// Define the template content
	templateContent := `
	<html>
		<head>
			<title>Random Joke Generator</title>
			<style>
				body {
					font-family: 'Arial', sans-serif;
					text-align: center;
					padding: 50px;
					background-color: #f0f0f0;
				}
				.joke-container {
					border: 2px solid #333;
					border-radius: 10px;
					padding: 20px;
					margin: 20px auto;
					max-width: 600px;
					background-color: #fff;
				}
				.joke {
					font-size: 18px;
					font-weight: bold;
					color: #333;
				}
			</style>
		</head>
		<body>
			<h1>Random Joke Generator</h1>
			<div class="joke-container">
				<p class="joke">{{.Joke}}</p>
			</div>
			<p>Refresh the page for another joke!</p>
		</body>
	</html>
	`

	// Create a new template
	tmpl, err := template.New("randomJoke").Parse(templateContent)
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

// getRandomJoke fetches a random joke from an external API.
func getRandomJoke() (string, error) {
	apiURL := "https://v2.jokeapi.dev/joke/Any"

	// Make a GET request to the API
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the JSON response
	var joke Joke
	err = json.Unmarshal(body, &joke)
	if err != nil {
		return "", err
	}

	// Combine the setup and delivery for a complete joke
	return joke.Setup + " " + joke.Delivery, nil
}
