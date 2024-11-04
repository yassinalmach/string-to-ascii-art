package utils

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var tSize int

// this function calculates the measurements of the Align Options
func Justify(alignOption string) {
	TerminalSize()
	switch alignOption {
	case "right":
		res := (tSize - artLength)
		result += strings.Repeat(" ", res)
	case "center":
		res := (tSize - artLength) / 2
		result += strings.Repeat(" ", res)
	case "justify":
		res := (tSize - artLength) / numSpace
		result += strings.Repeat(" ", res)
	}
}

// this function extract the current terminal size
func TerminalSize() {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	field := strings.Fields(string(out))
	tSize, err = strconv.Atoi(field[1])
	if err != nil {
		log.Fatal(err)
	}
}
