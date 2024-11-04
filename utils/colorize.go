package utils

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
func colorize(asciiArt []string, line, color, subStr string, index int) []string {
	resetCode := "\033[0m"
	colorCode := colorMap[color]
	coloredArt := make([]string, 8)
	var indices []int

	// This block finds all occurrences of a substring in a string and store their indices. 
	if subStr != "" {
		for i := 0; i <= len(line)-len(subStr); i++ {
			if subStr == line[i:i+len(subStr)] {
				for index := i; index < i+len(subStr); index++ {
					indices = append(indices, index)
				}
			}
		}
	}
	
	// Apply color if subStr is empty or if the index is in the indices slice
	if subStr == "" || contain(index, indices) {
		for i, line := range asciiArt {
			coloredArt[i] = colorCode + line + resetCode
		}
	} else {
		return asciiArt
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
