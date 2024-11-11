package utils

import "fmt"

var colorMap = map[string]string{
	"white":  "\033[38;5;15m",
	"black":  "\033[38;5;30m",
	"red":    "\033[38;5;196m",
	"green":  "\033[38;5;46m",
	"blue":   "\033[38;5;27m",
	"yellow": "\033[38;5;226m",
	"orange": "\033[38;5;202m",
	"pink":   "\033[38;5;201m",
	"brown":  "\033[38;5;172m",
	"purple": "\033[38;5;93m",
}

// Colorize applies color to the ASCII art characters based on the specified indices.
func colorize(asciiArt []string, line string, flags Flags, index int) []string {
	resetCode := "\033[0m"
	colorCode := colorMap[flags.Color]
	coloredArt := []string{}
	var indices []int

	// This block finds all occurrences of a substring in a string and store their indices.
	if flags.ColoredStr != "" {
		for i := 0; i <= len(line)-len(flags.ColoredStr); i++ {
			if flags.ColoredStr == line[i:i+len(flags.ColoredStr)] {
				for index := i; index < i+len(flags.ColoredStr); index++ {
					indices = append(indices, index)
				}
			}
		}
	}

	for _, line := range asciiArt {
		if flags.ColoredStr == "" || contain(index, indices) {
			if flags.RgbFlag {
				coloredLine := fmt.Sprintf("\033[38;2;%s;%s;%sm", flags.RGB.R, flags.RGB.G, flags.RGB.B) + line + resetCode
				coloredArt = append(coloredArt, coloredLine)
			} else {
				coloredLine := colorCode + line + resetCode
				coloredArt = append(coloredArt, coloredLine)
			}
		} else {
			coloredArt = append(coloredArt, line)
		}
	}

	return coloredArt
}

// contains checks if a given index is present in the indices slice.
func contain(index int, indices []int) bool {
	for _, i := range indices {
		if i == index {
			return true
		}
	}
	return false
}
