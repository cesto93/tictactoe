package net

import (
	"html/template"
	"net/http"
	"strconv"
	"tictactoe/pkg/tictactoe"
)

// Page struct represents the data to be passed to the HTML template
type Page struct {
	Title   string
	Content string
}

type Game struct {
	Player string
	Board [3][3] string
}

var globalGame Game

func StartWeb(){
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/select", selectHandler)
	http.HandleFunc("/play", playHandler)
	http.ListenAndServe(":8080", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data to be passed to the HTML template
	page := Page{
		Title:   "TicTacToe",
		Content: "Choice X or O to play",
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

				.grid-container {
			            display: grid;
			            grid-template-columns: repeat(3, 1fr);
			            gap: 10px;
			    }

			    .grid-item {
					border: 1px solid #ccc;
					padding: 20px;
					text-align: center;
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
	symbol := r.FormValue("symbol")

	w.Header().Set("Content-Type", "application/json")

	tmpl, err := template.New("index").Parse(`
	<p>You are the player {{.Player}}</p>
	<div class="grid-container">
		{{range $x, $els := .Board}}
			{{range $y, $el := $els}}
				<div class="grid-item">
				<button hx-post="/play" hx-trigger="click" hx-swap="innerHTML" hx-target="#result" hx-vals='{"x" : "{{$x}}", "y" : "{{$y}}"}'> {{.}} </button>
				</div>
			{{end}}
		{{end}}
	</div>
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	globalGame = Game {
		Board: tictactoe.Init_state(),	
		Player: symbol,
	}

	err = tmpl.Execute(w, globalGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	x, err := strconv.Atoi(r.FormValue("x"))

	y, err := strconv.Atoi(r.FormValue("y"))

	w.Header().Set("Content-Type", "application/json")

	tmpl, err := template.New("index").Parse(`
	<p>You are the player {{.Player}}</p>
	<div class="grid-container">
		{{range $x, $els := .Board}}
			{{range $y, $el := $els}}
				<div class="grid-item">
				<button hx-post="/play" hx-trigger="click" hx-swap="innerHTML" hx-target="#result" hx-vals='{"x" : "{{$x}}", "y" : "{{$y}}"}'> {{.}} </button>
				</div>
			{{end}}
		{{end}}
	</div>
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	globalGame.Board[x][y] = globalGame.Player

	err = tmpl.Execute(w, globalGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
