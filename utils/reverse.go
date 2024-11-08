package utils

import "log"

var rev string

func Reverse(inputFile []string, asciiMap map[rune][]string, idx int) string {
	for i := idx; i < len(inputFile); i++ {
		if inputFile[idx] == "" {
			rev += "\n"
			i++
			Reverse(inputFile, asciiMap, i)
		}

		if len(inputFile[i:]) >= 8 {
			for char := ' '; char <= '~'; char++ {
				charWidth := len(asciiMap[char][0])
				if isMatched(asciiMap[char], inputFile[i:i+9], charWidth) {
					newInputFile := removeChar(charWidth, inputFile[i:i+9])
					rev += string(char)
					Reverse(newInputFile, asciiMap, i)
				}
			}
			if len(inputFile) > 8 {
				i += 8
				rev += "\n"
				Reverse(inputFile, asciiMap, i)
			}
		} else {
			log.Fatalln("error file format incorrect")
		}
		i += 7
	}
	return rev
}

func isMatched(asciiChar, inputFile []string, charWidth int) bool {
	for i := 0; i < 8; i++ {
		if len(inputFile[i]) < len(asciiChar[i]) || asciiChar[i] != inputFile[i][:charWidth] {
			return false
		}
	}
	return true
}

func removeChar(width int, inputFile []string) []string {
	newInputFile := []string{}
	for i := 0; i < 8; i++ {
		newInputFile = append(newInputFile, inputFile[i][width:])
	}
	return newInputFile
}
