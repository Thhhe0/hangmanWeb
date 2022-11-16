package db

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func InitDb() {
	initScoreboard()
}

func initScoreboard() {
	if _, err := os.Stat("server/db/scoreboard.json"); err != nil {
		os.WriteFile("server/db/scoreboard.json", []byte("[]"), 0777)
	}

	saveJson, err := ioutil.ReadFile("server/db/scoreboard.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(saveJson, &Scoreboard); err != nil {
		panic(err)
	}
}
