package router

import (
	"html/template"
	"net/http"
)

type punishPage struct {
	Title string
	Css   string
	Js    string
}

func HandlePunish(w http.ResponseWriter, r *http.Request) {
	punish := punishPage{"Punition", "punish.css", "punish.js"}
	t := template.Must(template.ParseFiles("server/pages/punish.html", "server/pages/templates/head.tmpl", "server/pages/templates/base.tmpl"))
	t.ExecuteTemplate(w, "base", punish)
}
