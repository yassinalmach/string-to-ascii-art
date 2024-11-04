package utils

import (
	"strings"
)

// this function split the input string into lines
func SplitString(str, alignOption string) []string {
	splitStr := strings.Split(str, "\\n")
	if alignOption == "justify" {
		for i, line := range splitStr {
			splitStr[i] = strings.Join(strings.Fields(line), " ")
		}
	}

	// Remove one new line if all the characters are "\n"
	if allNewLines(splitStr) {
		splitStr = splitStr[1:]
	}
	return splitStr
}

// this Function checks if all elements in the input string are new lines "\n"
func allNewLines(splitStr []string) bool {
	for _, line := range splitStr {
		if line != "" {
			return false
		}
	}
	return true
}
