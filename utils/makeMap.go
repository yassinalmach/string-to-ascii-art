package utils

import (
	"os"
	"strings"
)

// make MAP, to store every ascii art character (8 lines) with their decimal ascii number
func MakeMap(bannerFile string) (map[rune][]string, error) {
	file, err := os.ReadFile(bannerFile)
	if err != nil {
		return nil, err
	}

	fileContent := strings.Split(string(file), "\r\n")
	asciiMap := make(map[rune][]string)
	
	dec := 31
	for _, line := range fileContent {
		if line == "" {
			dec++
		} else {
			asciiMap[rune(dec)] = append(asciiMap[rune(dec)], line)
		}
	}
	return asciiMap, nil
}
