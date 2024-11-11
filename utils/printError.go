package utils

import (
	"fmt"
	"os"
)

// PrintError prints the errors
func PrintError(err string) {
	switch err {
	case "color":
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
	case "output":
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
	case "justify":
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	case "reverse":
		fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
	case "banner":
		fmt.Println("Usage: please enter a valid BANNER name")
	case "printable":
		fmt.Println("Usage: One of the characters is not a Printable ASCII!")
	}
	os.Exit(0)
}
