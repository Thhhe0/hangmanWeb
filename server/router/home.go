package router

import (
	"hangman-web/server/db"
	"hangman-web/server/utils"
	"html/template"
	"net/http"
)

type homePage struct {
	Title      string
	Css        string
	Js         string
	Scoreboard []utils.Score
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	home := homePage{"Accueil", "index.css", "index.js", db.Scoreboard}

	t := template.Must(template.ParseFiles("server/pages/index.html", "server/pages/templates/head.tmpl", "server/pages/templates/base.tmpl"))
	t.ExecuteTemplate(w, "base", home)
}
