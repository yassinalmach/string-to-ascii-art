package utils

var (
	result    string
	artLength int
	numSpace  int
)

// GenerateArt converts the input string into ASCII art using the provided ASCII art map.
func GenerateArt(asciiMap map[rune][]string, splittedStr []string, flags Flags) string {
	// Iterate through the input lines
	for _, line := range splittedStr {
		if line == "" {
			result += "\n"
			continue
		}

		asciiArt := [][]string{}
		// Iterate through each character in the line
		for index, char := range line {
			if char < 32 || char > 126 {
				PrintError("printable")
			}
			if flags.AlignOption == "justify" && char == 32 {
				// For justify alignment, mark spaces with a special character
				indice := []string{"+", "+", "+", "+", "+", "+", "+", "+"}
				asciiArt = append(asciiArt, indice)
				numSpace++
			} else if flags.Color != "" {
				coloredArt := colorize(asciiMap[char], line, flags.Color, flags.ColoredStr, index)
				asciiArt = append(asciiArt, coloredArt)
			} else {
				asciiArt = append(asciiArt, asciiMap[char])
				artLength += len(asciiMap[char][0])
			}
		}

		// make the ascii art
		for i := 0; i < 8; i++ {
			if flags.AlignOption == "right" || flags.AlignOption == "center" {
				Justify(flags.AlignOption)
			}
			for _, char := range asciiArt {
				if flags.AlignOption == "justify" && char[i] == "+" {
					Justify(flags.AlignOption)
				} else {
					result += char[i]
				}
			}
			result += "\n"
		}

		numSpace = 0
		artLength = 0
	}
	return result
}
