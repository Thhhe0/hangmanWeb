package src

import (
	"hangman-web/hangman/src/utils"
	"strings"
)

type Word struct {
	Word             string
	Normalize        string
	HiddenWord       []rune
	HiddenWordString string
	HiddenWordAscii  [][8]string
}

func NewWord(str string) Word {
	tempWord := Word{}

	tempWord.Word = str
	tempWord.Normalize = utils.Normalize(tempWord.Word)

	for _, char := range strings.ToUpper(tempWord.Word) {
		if char == ' ' || char == '-' || (char >= '0' && char <= '9') {
			tempWord.HiddenWord = append(tempWord.HiddenWord, char)
		} else {
			tempWord.HiddenWord = append(tempWord.HiddenWord, '_')
		}
	}
	UncoverLetter(&tempWord)

	tempWord.HiddenWordString = string(tempWord.HiddenWord)
	upperNorm := strings.ToUpper(utils.Normalize(string(tempWord.HiddenWord)))
	tempWord.HiddenWordAscii = utils.BuilsAsciiWordArr(upperNorm)

	return tempWord
}

func (w *Word) CheckStr(str string) bool {
	if len(str) > 1 {
		if str == w.Normalize {
			for i := range w.Word {
				w.displayLetter(i)
			}
			return true
		} else {
			return false
		}
	} else {
		if strings.Contains(w.Normalize, str) {
			indexes := []int{}
			for i, s := range w.Normalize {
				if s == rune(str[0]) {
					indexes = append(indexes, i)
				}
			}
			for _, i := range indexes {
				w.displayLetter(i)
			}
			return true
		} else {
			return false
		}
	}
}

func (w *Word) displayLetter(i int) {
	upperNorm := strings.ToUpper(w.Normalize)

	w.HiddenWord[i] = rune(w.Word[i])
	w.HiddenWordString = string(w.HiddenWord)
	w.HiddenWordAscii[i] = utils.GetAsciiLetter(int(upperNorm[i]))
}
