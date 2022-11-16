package anim

import (
	"fmt"
	"hangman-web/hangman/src"
	"hangman-web/hangman/src/assets"
	"hangman-web/hangman/src/utils"
	"io/ioutil"
	"runtime"
	"strings"
	"time"
)

func initPullAnim() []string {
	rawDB, err := ioutil.ReadFile("pull_anim.txt")
	if err != nil {
		panic(err)
	}

	currOs := runtime.GOOS
	if currOs == "windows" {
		return strings.Split(string(rawDB), ",\r\n")
	} else {
		return strings.Split(string(rawDB), ",\n")
	}
}

func BuildAnimPull(str string) (anim [][]string) {
	for _, frame := range initPullAnim() {
		anim = append(anim, strings.Split(frame, "\n"))
	}
	pull := anim[len(anim)-1]
	word := utils.BuildAsciiWord(utils.BuilsAsciiWordArr(str))
	wordLen := len(word[0])
	for i := 0; i < wordLen; i++ {
		anim = append(anim, []string{
			word[0][wordLen-i-1:] + pull[0],
			word[1][wordLen-i-1:] + pull[1],
			word[2][wordLen-i-1:] + pull[2],
			word[3][wordLen-i-1:] + pull[3],
			word[4][wordLen-i-1:] + pull[4],
			word[5][wordLen-i-1:] + pull[5],
			word[6][wordLen-i-1:] + pull[6],
			word[7][wordLen-i-1:] + pull[7],
		})
	}

	return
}

func RunAnimPull(anim [][]string, g src.Game) {
	for _, frame := range anim {
		time.Sleep((time.Second / 60) * 5)
		utils.Clear()
		g.PrintLife()
		for _, line := range frame {
			fmt.Println(line)
		}
		g.PrintWord(false, assets.Green)
	}
}
