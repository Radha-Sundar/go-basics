// main.go
package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"net/url"
)

// WeatherInfo represents the weather information structure from the API response.
type WeatherInfo struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

// Data struct represents the dynamic data to be rendered in the template.
type Data struct {
	Location string
	Weather  WeatherInfo
	ASCIIArt string
}

func main() {
	// Handle requests with the weatherReportHandler function
	http.HandleFunc("/", weatherReportHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

func weatherReportHandler(w http.ResponseWriter, r *http.Request) {
	// Get user input from query parameter
	location := r.URL.Query().Get("location")

	// If location is empty, set a default value
	if location == "" {
		location = "City"
	}

	// Fetch the current weather information for the given location
	weather, err := getCurrentWeather(location)
	if err != nil {
		http.Error(w, "Failed to fetch weather information", http.StatusInternalServerError)
		return
	}

	// Generate colorful ASCII art based on the weather condition
	asciiArt, err := generateWeatherASCIIArt(weather.Main)
	if err != nil {
		http.Error(w, "Failed to generate ASCII art", http.StatusInternalServerError)
		return
	}

	// Prepare data
	data := Data{
		Location: location,
		Weather:  weather,
		ASCIIArt: asciiArt,
	}

	// Define the template content
	templateContent := `
	<html>
		<head>
			<title>Dynamic Weather Report</title>
			<style>
				body {
					font-family: 'Arial', sans-serif;
					text-align: center;
					padding: 50px;
					background-color: #f0f0f0;
				}
				.weather-report {
					border: 2px solid #333;
					border-radius: 10px;
					padding: 20px;
					margin: 20px auto;
					max-width: 600px;
					background-color: #fff;
				}
				.ascii-art {
					font-size: 20px;
					color: #333;
					white-space: pre;
				}
				.weather-info {
					font-size: 18px;
					font-weight: bold;
					color: #333;
				}
			</style>
		</head>
		<body>
			<h1>Dynamic Weather Report</h1>
			<div class="weather-report">
				<p>Current weather in {{.Location}}:</p>
				<p class="weather-info">{{.Weather.Description}}</p>
				<div class="ascii-art">{{.ASCIIArt}}</div>
			</div>
			<p>Change the location in the URL for weather updates, e.g., ?location=New+York</p>
		</body>
	</html>
	`

	// Create a new template
	tmpl, err := template.New("weatherReport").Parse(templateContent)
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

// getCurrentWeather fetches the current weather information for the given location from an external API.
func getCurrentWeather(location string) (WeatherInfo, error) {
	apiURL := "https://api.openweathermap.org/data/2.5/weather"
	apiKey := "YOUR_OPENWEATHERMAP_API_KEY" // Replace with your OpenWeatherMap API key

	// Build the URL for the API request
	u, err := url.Parse(apiURL)
	if err != nil {
		return WeatherInfo{}, err
	}

	q := u.Query()
	q.Set("q", location)
	q.Set("appid", apiKey)
	u.RawQuery = q.Encode()

	// Make a GET request to the OpenWeatherMap API
	resp, err := http.Get(u.String())
	if err != nil {
		return WeatherInfo{}, err
	}
	defer resp.Body.Close()

	// Decode the JSON response
	var weatherData struct {
		Weather []WeatherInfo `json:"weather"`
	}

	err = json.NewDecoder(resp.Body).Decode(&weatherData)
	if err != nil {
		return WeatherInfo{}, err
	}

	if len(weatherData.Weather) > 0 {
		return weatherData.Weather[0], nil
	}

	return WeatherInfo{}, nil
}

// generateWeatherASCIIArt generates colorful ASCII art based on the weather condition.
func generateWeatherASCIIArt(weatherCondition string) (string, error) {
	// A simple example - you can replace this with a more sophisticated ASCII art generation logic
	// You can customize ASCII art based on different weather conditions
	switch weatherCondition {
	case "Clear":
		return `
		\    /\
		 )  ( ')
		(  /  )
		 \(__)|`, nil
	case "Clouds":
		return `
		\ \ \
		 \_\_\`, nil
	case "Rain":
		return `
		\\   //
		 )\-/(
		/(_•_)\`, nil
	default:
		return `
		¯\_(ツ)_/¯`, nil
	}
}
