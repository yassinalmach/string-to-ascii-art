package reverse

import (
	"AsciiArt/utils"
	"fmt"
	"os"
	"strings"
)

// reverseAscii is the main point for reverse functionality
// it reads reverse file, and pick the right banner, then
// convert the art to a string
func ReverseAscii(flags utils.Flags) (string, error) {
	// read reverse file
	fileContent, err := os.ReadFile(flags.ReverseFile)
	if err != nil {
		return "", err
	}

	if string(fileContent) == "" {
		return "", fmt.Errorf("empty file")
	}

	// pick the right banner
	flags.BannerFile = pickBanner(string(fileContent))
	
	// split the revese file
	inputFile := strings.Split(string(fileContent), "\n")
	
	asciiMap, err := utils.MakeMap(flags.BannerFile)
	if err != nil {
		utils.PrintError("banner")
	}

	// reverse the art to a string
	result, err := reverse(inputFile, asciiMap)
	return result, err
}
