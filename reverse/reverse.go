package reverse

import "fmt"

// reverse it takes the input file splitted and banner as a map, and returns
// the art reversed to a string.
func reverse(inputFile []string, asciiMap map[rune][]string) (string, error) {
	var result string
	for i := 0; i < len(inputFile); i++ {
		if (inputFile[i] == "" || inputFile[i] == "$") && i != len(inputFile)-1 {
			result += "\n"
			continue
		}

		if len(inputFile[i:]) >= 8 {
			for char := ' '; char <= '~'; char++ {
				charWidth := len(asciiMap[char][0])
				if isMatched(asciiMap[char], inputFile[i:], charWidth) {
					result += string(char)
					inputFile = removeChar(charWidth, inputFile, i)
					char = 31
				}
			}
			if isNotValidFormat(inputFile[i:]) {
				return "", fmt.Errorf("file format incorrect")
			}

			i += 7
			result += "\n"
		}
	}
	return result[:len(result)-1], nil
}

// isNotValidFormat checks input file format if it is correct
// it returns true if not correct, otherwise it retuns false
func isNotValidFormat(inputFile []string) bool {
	for i := 0; i < 8; i++ {
		if inputFile[i] != "" && inputFile[i] != "$" {
			return true
		}
	}
	return false
}

// isMatched it matches the ascii characters of the input file with the art in banner
// return true if the art matched
func isMatched(asciiChar, inputFile []string, charWidth int) bool {
	for i := 0; i < 8; i++ {
		if len(inputFile[i]) < len(asciiChar[i]) || asciiChar[i] != inputFile[i][:charWidth] {
			return false
		}
	}
	return true
}

// removeChar removes the matched character from the input file
func removeChar(width int, inputFile []string, idx int) []string {
	newInputFile := []string{}
	newInputFile = append(newInputFile, inputFile[:idx]...)
	for i := idx; i < idx+8; i++ {
		newInputFile = append(newInputFile, inputFile[i][width:])
	}
	newInputFile = append(newInputFile, inputFile[idx+8:]...)
	return newInputFile
}
