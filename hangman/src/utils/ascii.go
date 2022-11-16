package utils

import (
	"fmt"
	"runtime"
	"strings"
)

func GetAsciiLetter(asciiCode int) (ret [8]string) {
	asciiDb := GetAsciiDb()
	currOS := runtime.GOOS
	if currOS == "windows" {
		copy(ret[:], strings.Split(asciiDb[asciiCode-32], "\r\n")[:8])
	} else {
		copy(ret[:], strings.Split(asciiDb[asciiCode-32], "\n")[:8])
	}

	return
}

func BuilsAsciiWordArr(word string) (ret [][8]string) {
	for _, char := range word {
		ret = append(ret, GetAsciiLetter(int(char)))
	}

	return
}

func BuildAsciiWord(word [][8]string) (ret [8]string) {
	letterHeight := len(word[0])
	for i := 0; i < letterHeight; i++ {
		for _, letter := range word {
			ret[i] += letter[i]
		}
	}

	return
}

func PrintAsciiWord(word [][8]string, color string) {
	if color != "" {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				fmt.Print(ColorString(color, letter[i]))
			}
			fmt.Println()
		}
	} else {
		for i := 0; i < 8; i++ {
			for _, letter := range word {
				fmt.Print(letter[i])
			}
			fmt.Println()
		}
	}
}
