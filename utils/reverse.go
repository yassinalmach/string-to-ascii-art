package utils

var rev string

func Reverse(inputFile []string, asciiMap map[rune][]string) string {
    for i := 0; i < len(inputFile); i++ {
        if inputFile[i] == "" && i != len(inputFile)-1{
            rev += "\n"
            continue
        }

        if len(inputFile[i:]) >= 8 {
            for inputFile[i] != "" {
                for char := ' '; char <= '~'; char++ {
                    charWidth := len(asciiMap[char][0])
                    if isMatched(asciiMap[char], inputFile[i:], charWidth) {
                        rev += string(char)
                        inputFile = removeChar(charWidth, inputFile, i)
                    }
                }
            }
            i += 7
            rev += "\n"
        }
    }
    return rev[:len(rev)-1]
}

func isMatched(asciiChar, inputFile []string, charWidth int) bool {
    for i := 0; i < 8; i++ {
        if len(inputFile[i]) < len(asciiChar[i]) || asciiChar[i] != inputFile[i][:charWidth] {
            return false
        }
    }
    return true
}

func removeChar(width int, inputFile []string, idx int) []string {
    newInputFile := []string{}
    newInputFile = append(newInputFile, inputFile[:idx]...)
    for i := idx; i < idx+8; i++ {
        newInputFile = append(newInputFile, inputFile[i][width:])
    }
    newInputFile = append(newInputFile, inputFile[idx+8:]...)
    return newInputFile
}