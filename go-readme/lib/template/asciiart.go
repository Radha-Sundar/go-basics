// main.go
package main

import (
	"html/template"
	"net/http"
)

// Data struct represents the dynamic data to be rendered in the template.
type AData struct {
	Name     string
	ASCIIArt string
}

func main() {
	// Handle requests with the generateArtHandler function
	http.HandleFunc("/", generateArtHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

func generateArtHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user's name from the query parameter
	name := r.URL.Query().Get("name")

	// If the name is empty, set a default value
	if name == "" {
		name = "Guest"
	}

	// Generate ASCII art based on the user's name
	asciiArt := generateASCIIArt(name)

	// Prepare data
	data := AData{
		Name:     name,
		ASCIIArt: asciiArt,
	}

	// Define the template content
	templateContent := `
	<html>
		<head>
			<title>ASCII Art Generator</title>
			<style>
				body {
					font-family: 'Courier New', monospace;
					text-align: center;
					padding: 50px;
					background-color: #f0f0f0;
				}
				.ascii-art {
					font-size: 20px;
					color: #333;
					white-space: pre;
				}
			</style>
		</head>
		<body>
			<h1>ASCII Art Generator</h1>
			<h2>Hello, {{.Name}}!</h2>
			<div class="ascii-art">{{.ASCIIArt}}</div>
			<p>Change your name in the URL for personalized ASCII art, e.g., ?name=John</p>
		</body>
	</html>
	`

	// Create a new template
	tmpl, err := template.New("asciiArt").Parse(templateContent)
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

// generateASCIIArt generates ASCII art based on the user's name.
func generateASCIIArt(name string) string {
	// A simple example - you can replace this with a more sophisticated ASCII art generation logic
	return `
   ____  _              _______ _     _             
  / __ \| |            |__   __| |   (_)            
 | |  | | |_   _  ___     | |  | |__  _ _ __   __ _ 
 | |  | | | | | |/ _ \    | |  | '_ \| | '_ \ / _' |
 | |__| | | |_| |  __/    | |  | | | | | | | | (_| |
  \___\_\_|\__,_|\___|    |_|  |_| |_|_|_| |_|\__, |
                                                  __/ |
                                                 |___/ 
` + "\nHello, " + name + "!\n"
}
