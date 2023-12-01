package main

import (
	"html/template"
	"net/http"
	"sync"
	"time"
)

// Report represents a monthly report.
type Report struct {
	Month    time.Month
	Year     int
	Content  string
	Approved bool
}

// Data struct represents the dynamic data to be rendered in the template.
type RData struct {
	Reports []Report
}

var (
	mu      sync.Mutex
	reports []Report
)

func main() {
	// Serve static files (CSS, JS, etc.) from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Handle requests with the reportsHandler function
	http.HandleFunc("/reports", reportsHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}

func reportsHandler(w http.ResponseWriter, r *http.Request) {
	// Lock the mutex to ensure thread safety when accessing shared data
	mu.Lock()
	defer mu.Unlock()

	// Create or update a sample monthly report
	currentMonth := time.Now().Month()
	currentYear := time.Now().Year()

	// Check if a report for the current month already exists
	existingReportIndex := findReportIndex(currentMonth, currentYear)

	// If not, create a new report
	if existingReportIndex == -1 {
		reports = append(reports, Report{
			Month:    currentMonth,
			Year:     currentYear,
			Content:  "This is the monthly report content.",
			Approved: false,
		})
	} else {
		// If exists, update the content
		reports[existingReportIndex].Content = "Updated monthly report content."
	}

	// Prepare data
	data := RData{
		Reports: reports,
	}

	// Define the template content
	templateContent := `
	<html>
		<head>
			<title>Monthly Reports</title>
			<link rel="stylesheet" type="text/css" href="/static/style.css">
		</head>
		<body>
			<h1>Monthly Reports</h1>
			{{range .Reports}}
				<div class="report">
					<h2>{{.Month.String}} {{.Year}}</h2>
					<p>{{.Content}}</p>
					{{if .Approved}}
						<p class="approved">Approved</p>
					{{else}}
						<p class="not-approved">Not Approved</p>
					{{end}}
				</div>
			{{end}}
		</body>
	</html>
	`

	// Create a new template
	tmpl, err := template.New("reports").Parse(templateContent)
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

// findReportIndex returns the index of a report in the reports slice, based on month and year.
func findReportIndex(month time.Month, year int) int {
	for i, r := range reports {
		if r.Month == month && r.Year == year {
			return i
		}
	}
	return -1
}
