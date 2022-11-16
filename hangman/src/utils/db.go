package utils

import (
	"io/ioutil"
	"runtime"
	"strings"
)

func GetAsciiDb() []string {
	rawAscii, err := ioutil.ReadFile("hangman/ascii.txt")
	if err != nil {
		panic(err)
	}

	currOs := runtime.GOOS
	if currOs == "windows" {
		return strings.Split(string(rawAscii), "\r\n,\r\n")
	} else {
		return strings.Split(string(rawAscii), "\n,\n")
	}
}

func GetWordDb(fn string) []string {
	rawDB, err := ioutil.ReadFile(fn)
	if err != nil {
		panic(err)
	}

	currOs := runtime.GOOS
	if currOs == "windows" {
		return strings.Split(string(rawDB), "\r\n")
	} else {
		return strings.Split(string(rawDB), "\n")
	}
}

func GetJoseDb() []string {
	rawDB, err := ioutil.ReadFile("hangman/hangman.txt")
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
