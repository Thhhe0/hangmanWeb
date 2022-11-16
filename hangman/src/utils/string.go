package utils

import (
	"hangman-web/hangman/src/assets"
	"strings"
)

func MapSpace(word []rune) string {
	ret := []rune{}
	for _, char := range word {
		ret = append(ret, char, ' ')
	}

	return string(ret[:len(ret)-1])
}

func Normalize(str string) string {
	ret := strings.ToLower(str)
	ret = strings.ReplaceAll(ret, "é", "e")
	ret = strings.ReplaceAll(ret, "è", "e")
	ret = strings.ReplaceAll(ret, "ê", "e")
	ret = strings.ReplaceAll(ret, "ë", "e")
	ret = strings.ReplaceAll(ret, "à", "a")
	ret = strings.ReplaceAll(ret, "î", "i")
	ret = strings.ReplaceAll(ret, "ï", "i")
	ret = strings.ReplaceAll(ret, "û", "u")
	ret = strings.ReplaceAll(ret, "ü", "u")

	return ret
}

func ContainStr(arr []string, toFind string) bool {
	for _, d := range arr {
		if d == toFind {
			return true
		}
	}
	return false
}

func ColorString(color, str string) string {
	return color + str + assets.Reset
}
