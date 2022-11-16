package router

import (
	"fmt"
	"hangman-web/hangman/src"
	"hangman-web/server/db"
	"net/http"
)

func HandleRestart(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Print(err)
			return
		}
		if r.Form.Get("difficulty") != "" {
			db.Game = src.NewGame(r.Form.Get("difficulty"))
			http.Redirect(w, r, "/hangman", 301)
		} else {
			http.Redirect(w, r, "/punition", 301)
		}
	} else {
		http.Redirect(w, r, "/punition", 301)
	}
}
