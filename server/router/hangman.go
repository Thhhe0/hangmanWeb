package router

import (
	"fmt"
	"hangman-web/hangman/src"
	"hangman-web/server/db"
	"hangman-web/server/utils"
	"html/template"
	"math/rand"
	"net/http"
)

type hangmanPage struct {
	Title       string
	Css         string
	Js          string
	RequestSbId int
	Game        src.Game
}

func HandleGame(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Print(err)
			return
		}
		fmt.Println(r.Form)

		db.Game.AttemptUpdate(r.Form.Get("letter"))

		if db.Game.Word.Word == db.Game.Word.HiddenWordString || db.Game.Attempt < 1 {
			db.Game.InGame = false
		}
	}

	hangman := hangmanPage{"Hangman", "hangman.css", "hangman.js", rand.Intn(100000000), db.Game}
	hangman.Game.Word.HiddenWord = utils.MapSpace(hangman.Game.Word.HiddenWord)
	hangman.Game.Word.HiddenWordString = string(hangman.Game.Word.HiddenWord)

	t := template.Must(template.ParseFiles("server/pages/hangman.html", "server/pages/templates/head.tmpl", "server/pages/templates/base.tmpl"))
	fmt.Println(db.Game.Word.Word)
	fmt.Println(db.Game.Word.HiddenWordString)
	t.ExecuteTemplate(w, "base", hangman)
}
