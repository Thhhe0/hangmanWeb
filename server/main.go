package main

import (
	"fmt"
	"hangman-web/server/db"
	Router "hangman-web/server/router"
	"log"
	"net/http"
)

func main() {
	db.InitDb()
	fileServer := http.FileServer(http.Dir("./server/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", Router.HandleHome)
	http.HandleFunc("/restart", Router.HandleRestart)
	http.HandleFunc("/score", Router.HandleScore)
	http.HandleFunc("/hangman", Router.HandleGame)
	http.HandleFunc("/punition", Router.HandlePunish)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
