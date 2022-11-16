package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Difficulty int

const (
	Easy   Difficulty = 1
	Normal            = 2
	Hard              = 3
)

type Score struct {
	Name         string
	Difficulty   string
	DifficultyId int
	Miss         int
}

func NewScore(name string, diff Difficulty, attemptLeft int) Score {
	if name == "" {
		name = "anonyme"
	}

	return Score{
		Name:         name,
		Difficulty:   diff.String(),
		DifficultyId: int(diff),
		Miss:         10 - attemptLeft,
	}
}

func SaveScoreBoard(sb []Score) {
	save, err := json.Marshal(sb)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("server/db/scoreboard.json", save, 0644)
	if err != nil {
		panic(err)
	}
}

func (diff Difficulty) String() string {
	switch diff {
	case Easy:
		return "easy"
	case Normal:
		return "normal"
	case Hard:
		return "Hard"
	}
	return "unknown"
}
