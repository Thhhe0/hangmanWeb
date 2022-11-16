package router

import (
	"fmt"
	"hangman-web/server/db"
	"hangman-web/server/utils"
	"net/http"
)

func HandleScore(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Print(err)
			return
		}
		if r.Form.Get("name") != "" {
			fmt.Println(r.Form)
			name := r.Form.Get("name")

			if !db.ScoreSaved {
				db.Scoreboard = append(db.Scoreboard, utils.NewScore(name, utils.Difficulty(db.Game.Difficulty), db.Game.Attempt))
				utils.SaveScoreBoard(db.Scoreboard)
				db.ScoreSaved = true

				http.Redirect(w, r, "/", 301)
			} else {
				http.Redirect(w, r, "/punition", 301)
			}

		} else {
			http.Redirect(w, r, "/punition", 301)
		}
	} else {
		http.Redirect(w, r, "/punition", 301)
	}
}
