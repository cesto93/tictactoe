package net

import (
	"html/template"
	"net/http"
	"strconv"

	"tictactoe/pkg/tictactoe"

	"github.com/sirupsen/logrus"
)

type Page struct {
	Title   string
	Content string
}

type Game struct {
	Player   string
	Board    [3][3]string
	Winner   string
	Terminal bool
}

var globalGame Game

func StartWeb() {
	logrus.SetReportCaller(true)
	
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/select", selectHandler)
	http.HandleFunc("/play", playHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Errorf("err %v", err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Title:   "TicTacToe",
		Content: "Choice X or O to play, X is first",
	}

	tmpl, err := template.New("index").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>{{.Title}}</title>
			<script src="/static/htmx.min.js"></script>
			<link rel="stylesheet" href="/static/style.css">
		</head>
		<body>

			<h1>{{.Title}}</h1>
			<div id="result">
				<p>{{.Content}}</p>
				<button hx-post="/select" hx-trigger="click" hx-swap="innerHTML" hx-target="#result" hx-vals='{"symbol" : "X"}'>X</button>
				<button hx-post="/select" hx-trigger="click" hx-swap="innerHTML" hx-target="#result" hx-vals='{"symbol" :"O"}'>O</button>
			</div>

		</body>
		</html>
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
		logrus.Errorf("err %+v", err)
		return
	}

	globalGame = Game{
		Board:  tictactoe.Init_state(),
		Player: symbol,
	}

	if globalGame.Player != tictactoe.X {
		err := cpuPlay(&globalGame)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logrus.Errorf("err %+v", err)
			return
		}
	}

	err = tmpl.Execute(w, globalGame)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	x, err := strconv.Atoi(r.FormValue("x"))
	if err != nil {
		logrus.Errorf("err %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	y, err := strconv.Atoi(r.FormValue("y"))
	if err != nil {
		logrus.Errorf("err %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	tmpl, err := template.New("index").Parse(`
	<p>You are the {{.Player}}</p>
	<div class="grid-container">
		{{range $x, $els := .Board}}
			{{range $y, $el := $els}}
				<div class="grid-item">
				{{ if eq . "" }}
					<button hx-post="/play" hx-trigger="click" hx-swap="innerHTML" hx-target="#result" hx-vals='{"x" : "{{$x}}", "y" : "{{$y}}"}'> {{.}} </button>
				{{ else }}
					<button hx-post="/play" disabled hx-trigger="click" hx-swap="innerHTML" hx-target="#result" hx-vals='{"x" : "{{$x}}", "y" : "{{$y}}"}'> {{.}} </button>
				{{ end }}
				</div>
			{{end}}
		{{end}}

	</div>

	{{ if .Terminal }}
		{{ if eq .Winner "" }}
			<p> Draw </p>
		{{ else }}
			<p> The Winner is {{.Winner}} </p>
		{{end}}
	{{end}}
	`)
	if err != nil {
		logrus.Errorf("err %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = playerPlay(&globalGame, [2]int{x, y})
	if err != nil {
		logrus.Errorf("err %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = cpuPlay(&globalGame)
	if err != nil {
		logrus.Errorf("err %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, globalGame)
	if err != nil {
		logrus.Errorf("err %+v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
