package net

import (
	"html/template"
	"net/http"
)

// Page struct represents the data to be passed to the HTML template
type Page struct {
	Title   string
	Content string
	SelectedSymbol string `json:"selectedSymbol"`
}

func StartWeb(){
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/select", selectHandler)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data to be passed to the HTML template
	page := Page{
		Title:   "TicTacToe",
		Content: "Choice X or O to play",
		SelectedSymbol: "",
	}

	// Parse the HTML template
	tmpl, err := template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>{{.Title}}</title>
			<link rel="stylesheet" href="https://unpkg.com/htmx.org@1.9.10/dist/htmx.css">
			<script src="https://unpkg.com/htmx.org@1.9.10/dist/htmx.js"></script>
			<style>
				body {
					font-family: Arial, sans-serif;
					text-align: center;
					margin: 50px;
				}

				button {
					font-size: 18px;
					padding: 10px 20px;
					margin: 10px;
					cursor: pointer;
				}

				#result {
					font-size: 24px;
					margin-top: 20px;
				}
			</style>
		</head>
		<body>

			<h1>{{.Title}}</h1>
			<p>{{.Content}}</p>
			<button hx-post="/select" hx-trigger="click" hx-swap="innerHTML" hx-target="#result" hx-vals='{"symbol" : "X"}'>X</button>
			<button hx-post="/select" hx-trigger="click" hx-swap="innerHTML" hx-target="#result" hx-vals='{"symbol" :"O"}'>O</button>
			<div id="result"></div>

		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute the template and write the result to the response writer
	err = tmpl.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func selectHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the selected symbol from the request
	symbol := r.FormValue("symbol")

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write([]byte("You selected: " + symbol))
}
