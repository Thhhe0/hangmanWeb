package src

import (
	"hangman-web/hangman/src/utils"
	"math/rand"
	"time"
)

func Roll(winRate int) bool {
	rand.Seed(time.Now().UnixNano())
	rdm := rand.Intn(100)

	if rdm < winRate {
		return true
	} else {
		return false
	}
}

func GetWord(level string) string {
	// fn := HandleWordDb()
	tab := utils.GetWordDb("hangman/words" + level + ".txt")

	rand.Seed(time.Now().UnixNano())
	randWord := tab[rand.Intn(len(tab)-1)]

	return randWord
	// return "congrats"
}

func UncoverLetter(word *Word) {
	a := len(word.Word)/2 - 1
	i := 0
	tempWord := []rune(word.Word)
	for a > 0 {
		if word.HiddenWord[i] == '_' && a > 0 && Roll(15) {
			word.HiddenWord[i] = tempWord[i]
			a--
		}
		if i == len(word.HiddenWord)-1 {
			i = 0
		} else {
			i++
		}
	}
}
