package utils

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	currOs := runtime.GOOS
	if currOs == "windows" {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
