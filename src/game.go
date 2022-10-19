package src

import (
	"bufio"
	"fmt"
	"hangman-web/hangman/src/assets"
	"hangman-web/hangman/src/utils"
	"os"
	"strconv"
)

type Game struct {
	Word        Word
	Attempt     int
	TriedLetter []string
	Difficulty  int
	SaveName    string
	InGame      bool
}

func NewGame(level string) Game {
	tempGame := Game{}

	tempGame.Word = NewWord(GetWord(level))
	tempGame.Attempt = 10
	tempGame.SaveName = ""
	tempGame.InGame = true
	tempGame.Difficulty, _ = strconv.Atoi(level)

	return tempGame
}

func (g Game) PrintWord(lowGraph bool, color string) {
	if color != "" {
		if !lowGraph {
			utils.PrintAsciiWord(g.Word.HiddenWordAscii, color)
		} else {
			fmt.Println(color + utils.MapSpace(g.Word.HiddenWord) + assets.Reset)
		}
	} else {
		if !lowGraph {
			utils.PrintAsciiWord(g.Word.HiddenWordAscii, "")
		} else {
			fmt.Println(utils.MapSpace(g.Word.HiddenWord))
		}
	}
}

func (g Game) PrintLife() {
	for i := 0; i < g.Attempt; i++ {
		fmt.Print(utils.ColorString(assets.Red, "♥") + assets.Reset + " ")
	}
	fmt.Println()
}

func (g *Game) AttemptUpdate(res string) {
	// jose := utils.GetJoseDb()
	res = utils.Normalize(res)
	length := len(res)

	if length > 0 {
		found := g.Word.CheckStr(res)
		contain := utils.ContainStr(g.TriedLetter, res)

		if length == 1 {
			if contain {
				// g.PrintLife()
				// fmt.Println(assets.Red + "You already tried this letter" + assets.Reset)
			} else {
				g.TriedLetter = append(g.TriedLetter, res)
				if !found {
					g.Attempt--
					// g.PrintLife()
					// fmt.Println(assets.Red + "Not present in the word" + assets.Reset)
					// fmt.Println(jose[10-(g.Attempt+1)] + "\n")
				} else {
					// g.PrintLife()
				}
			}
		} else {
			if !found {
				g.Attempt -= 2
				// g.PrintLife()
				// temp := g.Attempt
				if g.Attempt < 1 {
					g.Attempt = 0
					// temp = 0
				}

				// fmt.Println(assets.Red + "Wrong word" + assets.Reset)
				// fmt.Println(jose[10-(temp+1)] + "\n")
			} else {
				// g.PrintLife()
			}
		}
	} else {
		g.Attempt--
		// g.PrintLife()
		// fmt.Println(assets.Red + "Empty is not a valid character" + assets.Reset)
		// fmt.Println(jose[10-(g.Attempt+1)] + "\n")
	}
}

func (g *Game) Update(lowGraph bool) bool {
	utils.Clear()
	var res string
	scanner := bufio.NewScanner(os.Stdin)

	// init
	// fmt.Printf("Good luck, you have "+utils.ColorString(assets.Green, "%v")+" attempts.\n", g.Attempt)
	// fmt.Println("(if you're on windows, CHEH, you will not have text color)")
	// g.PrintWord(lowGraph, "")

	for {
		if g.Attempt < 1 {
			// lose := utils.GetJoseDb()[9]
			utils.Clear()

			// fmt.Println(assets.Red + "Dead" + assets.Reset)
			// fmt.Println(assets.Red + lose + assets.Reset)
			// g.PrintWord(lowGraph, assets.Red)

			return false
		}
		if string(g.Word.HiddenWord) == g.Word.Word {
			// win := utils.GetJoseDb()[10]

			// utils.Clear()
			// g.PrintLife()
			// fmt.Println(assets.Green + win + assets.Reset)
			// g.PrintWord(lowGraph, assets.Green)

			return true
		}

		fmt.Print("Choose : ")
		scanner.Scan()
		res = scanner.Text()
		res = utils.Normalize(res)

		if len(res) > 3 {
			// HandleStop(res, g)
		}
		utils.Clear()

		if res == "help" {
			// g.PrintLife()
			// HandleHelp(true)
		} else {
			g.AttemptUpdate(res)
		}
		// g.PrintWord(lowGraph, "")
	}
}
