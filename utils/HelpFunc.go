package utils

import (
	"fmt"
	"os"
)

// this function print a guide message to the user to help them use the program
func helpFunc() {
		helpMessage := `Usage: go run . [OPTION] [STRING] [BANNER]

Options:
	--output=<fileName.txt>    Save the output to a file instead of printing to terminal.
	--align=<alignment>        Align the output. Valid alignments: left, center, right, justify.
	--color=<color>            Color a specific parts or all of the output.
	--reverse=<filename.txt>   convert ascii art in a file to a string.

Color Flag:
	- Available colors:
		white, black, red, green, blue, yellow, orange, pink, brown, purple.
	- To color a specific charactes of the text, use the following syntax:
		go run . --color=<color> <letters to be colored> "your text here"
	- you can input RGB number to color a string
		go run . --color=rgb{<R>,<G>,<B>} <letters to be colored> "your text here"
		instead of R, G, B use numbers between 0 and 255

Banners:
	Available banners: standard, shadow, thinkertoy

Examples:
	go run . --output=output.txt "Hello, World!" standard
	go run . --align=center "Hello" shadow
	go run . --color=red He "Hello, World!"
	go run . --color=rbg{255,255,0} He "Hello, World!"
	go run . --reverse=output.txt`
	fmt.Println(helpMessage)
	os.Exit(0)
}
